// Copyright 2024 The TCell Authors
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

// Package extended contains an extended set of terminal descriptions.
// Applications desiring to have a better chance of Just Working by
// default should include this package.  This will significantly increase
// the size of the program.
package extended

import (
	// The following imports just register themselves --
	// these are the terminal types we aggregate in this package.
	_ "github.com/stdiopt/tcell/terminfo/a/aixterm"
	_ "github.com/stdiopt/tcell/terminfo/a/alacritty"
	_ "github.com/stdiopt/tcell/terminfo/a/ansi"
	_ "github.com/stdiopt/tcell/terminfo/b/beterm"
	_ "github.com/stdiopt/tcell/terminfo/c/cygwin"
	_ "github.com/stdiopt/tcell/terminfo/d/dtterm"
	_ "github.com/stdiopt/tcell/terminfo/e/emacs"
	_ "github.com/stdiopt/tcell/terminfo/f/foot"
	_ "github.com/stdiopt/tcell/terminfo/g/gnome"
	_ "github.com/stdiopt/tcell/terminfo/h/hpterm"
	_ "github.com/stdiopt/tcell/terminfo/k/konsole"
	_ "github.com/stdiopt/tcell/terminfo/k/kterm"
	_ "github.com/stdiopt/tcell/terminfo/l/linux"
	_ "github.com/stdiopt/tcell/terminfo/p/pcansi"
	_ "github.com/stdiopt/tcell/terminfo/r/rxvt"
	_ "github.com/stdiopt/tcell/terminfo/s/screen"
	_ "github.com/stdiopt/tcell/terminfo/s/simpleterm"
	_ "github.com/stdiopt/tcell/terminfo/s/sun"
	_ "github.com/stdiopt/tcell/terminfo/t/tmux"
	_ "github.com/stdiopt/tcell/terminfo/v/vt100"
	_ "github.com/stdiopt/tcell/terminfo/v/vt102"
	_ "github.com/stdiopt/tcell/terminfo/v/vt220"
	_ "github.com/stdiopt/tcell/terminfo/v/vt320"
	_ "github.com/stdiopt/tcell/terminfo/v/vt400"
	_ "github.com/stdiopt/tcell/terminfo/v/vt420"
	_ "github.com/stdiopt/tcell/terminfo/v/vt52"
	_ "github.com/stdiopt/tcell/terminfo/w/wy50"
	_ "github.com/stdiopt/tcell/terminfo/w/wy60"
	_ "github.com/stdiopt/tcell/terminfo/w/wy99_ansi"
	_ "github.com/stdiopt/tcell/terminfo/x/xfce"
	_ "github.com/stdiopt/tcell/terminfo/x/xterm"
	_ "github.com/stdiopt/tcell/terminfo/x/xterm_ghostty"
	_ "github.com/stdiopt/tcell/terminfo/x/xterm_kitty"
)
