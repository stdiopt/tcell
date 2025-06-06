// Copyright 2024 The Tcell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package views

import (
	"sync"

	"github.com/stdiopt/tcell"
)

// Application represents an event-driven application running on a screen.
type Application struct {
	widget   Widget
	screen   tcell.Screen
	style    tcell.Style
	err      error
	wg       sync.WaitGroup
	paste    bool
	stopQ    chan struct{}
	eventQ   chan tcell.Event
	stopOnce sync.Once
}

// SetRootWidget sets the primary (root, main) Widget to be displayed.
func (app *Application) SetRootWidget(widget Widget) {
	app.widget = widget
}

// EnablePaste enables or disables pasting support in the application.
// It must be called before Start or Run or else it will have no effect.
func (app *Application) EnablePaste(on bool) {
	app.paste = on
}

// initialize initializes the application.  It will normally attempt to
// allocate a default screen if one is not already established.
func (app *Application) initialize() error {
	if app.screen == nil {
		if app.screen, app.err = tcell.NewScreen(); app.err != nil {
			return app.err
		}
		app.screen.SetStyle(app.style)
	}
	return nil
}

// Quit causes the application to shutdown gracefully.  It does not wait
// for the application to exit, but returns immediately.
func (app *Application) Quit() {
	app.stopOnce.Do(func() { close(app.stopQ) })
}

// Refresh causes the application forcibly redraw everything.  Use this
// to clear up screen corruption, etc.
func (app *Application) Refresh() {
	if scr := app.screen; scr != nil {
		scr.Sync()
	}
}

// Update asks the application to draw any screen updates that have not
// been drawn yet.  It is not necessary to call this from inside an
// event handler, as a draw will be done implicitly after handling events.
func (app *Application) Update() {
	if scr := app.screen; scr != nil {
		scr.Show()
	}
}

// PostFunc posts a function to be executed in the context of the
// application's event loop.  Functions that need to update displayed
// state, etc. can do this to avoid holding locks.
func (app *Application) PostFunc(fn func()) {
	ev := &eventAppFunc{fn: fn}
	ev.SetEventNow()
	if scr := app.screen; scr != nil {
		go func() { scr.PostEventWait(ev) }()
	}
}

// SetScreen sets the screen to use for the application.  This must be
// done before the application starts to run or is initialized.
func (app *Application) SetScreen(scr tcell.Screen) {
	if app.screen == nil {
		app.screen = scr
		app.err = nil
	}
}

// SetStyle sets the default style (background) to be used for Widgets
// that have not specified any other style.
func (app *Application) SetStyle(style tcell.Style) {
	app.style = style
	if app.screen != nil {
		app.screen.SetStyle(style)
	}
}

func (app *Application) run() {

	screen := app.screen
	widget := app.widget

	if widget == nil {
		app.wg.Done()
		return
	}
	if screen == nil {
		if e := app.initialize(); e != nil {
			app.wg.Done()
			return
		}
		screen = app.screen
	}
	defer func() {
		screen.Fini()
		app.wg.Done()
	}()
	screen.Init()
	screen.EnableMouse()
	if app.paste {
		screen.EnablePaste()
	}
	screen.Clear()
	widget.SetView(screen)

	app.eventQ = make(chan tcell.Event, 16)
	app.stopQ = make(chan struct{})
	go screen.ChannelEvents(app.eventQ, app.stopQ)

loop:
	for {
		if widget = app.widget; widget == nil {
			break
		}
		widget.Draw()
		screen.Show()

		ev := <-app.eventQ
		if ev == nil {
			screen.Fini()
			break loop
		}
		switch nev := ev.(type) {
		case *eventAppFunc:
			nev.fn()
		case *tcell.EventResize:
			screen.Sync()
			widget.Resize()
		default:
			widget.HandleEvent(ev)
		}
	}
}

// Start starts the application loop, initializing the screen
// and starting the Event loop.  The application will run in a goroutine
// until Quit is called.
func (app *Application) Start() {
	app.wg.Add(1)
	go app.run()
}

// Wait waits until the application finishes.
func (app *Application) Wait() error {
	app.wg.Wait()
	return app.err
}

// Run runs the application, waiting until the application loop exits.
// It is equivalent to app.Start() followed by app.Wait()
func (app *Application) Run() error {
	app.Start()
	return app.Wait()
}

type eventAppFunc struct {
	tcell.EventTime
	fn func()
}
