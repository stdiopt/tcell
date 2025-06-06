// Generated automatically.  DO NOT HAND-EDIT.

package kterm

import "github.com/stdiopt/tcell/terminfo"

func init() {

	// kterm kanji terminal emulator (X window system)
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:              "kterm",
		Columns:           80,
		Lines:             24,
		Colors:            8,
		Bell:              "\a",
		Clear:             "\x1b[H\x1b[2J",
		EnterCA:           "\x1b7\x1b[?47h",
		ExitCA:            "\x1b[2J\x1b[?47l\x1b8",
		AttrOff:           "\x1b[m\x1b(B",
		Underline:         "\x1b[4m",
		Bold:              "\x1b[1m",
		Reverse:           "\x1b[7m",
		EnterKeypad:       "\x1b[?1h\x1b=",
		ExitKeypad:        "\x1b[?1l\x1b>",
		SetFg:             "\x1b[3%p1%dm",
		SetBg:             "\x1b[4%p1%dm",
		SetFgBg:           "\x1b[3%p1%d;4%p2%dm",
		ResetFgBg:         "\x1b[39;49m",
		PadChar:           "\x00",
		AltChars:          "``aajjkkllmmnnooppqqrrssttuuvvwwxx~~",
		EnterAcs:          "\x1b(0",
		ExitAcs:           "\x1b(B",
		EnableAutoMargin:  "\x1b[?7h",
		DisableAutoMargin: "\x1b[?7l",
		Mouse:             "\x1b[M",
		SetCursor:         "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:       "\b",
		CursorUp1:         "\x1b[A",
		KeyUp:             "\x1bOA",
		KeyDown:           "\x1bOB",
		KeyRight:          "\x1bOC",
		KeyLeft:           "\x1bOD",
		KeyInsert:         "\x1b[2~",
		KeyDelete:         "\x1b[3~",
		KeyBackspace:      "\x7f",
		KeyPgUp:           "\x1b[5~",
		KeyPgDn:           "\x1b[6~",
		KeyF1:             "\x1b[11~",
		KeyF2:             "\x1b[12~",
		KeyF3:             "\x1b[13~",
		KeyF4:             "\x1b[14~",
		KeyF5:             "\x1b[15~",
		KeyF6:             "\x1b[17~",
		KeyF7:             "\x1b[18~",
		KeyF8:             "\x1b[19~",
		KeyF9:             "\x1b[20~",
		KeyF10:            "\x1b[21~",
		KeyF11:            "\x1b[23~",
		KeyF12:            "\x1b[24~",
		KeyF13:            "\x1b[25~",
		KeyF14:            "\x1b[26~",
		KeyF15:            "\x1b[28~",
		KeyF16:            "\x1b[29~",
		KeyF17:            "\x1b[31~",
		KeyF18:            "\x1b[32~",
		KeyF19:            "\x1b[33~",
		KeyF20:            "\x1b[34~",
		AutoMargin:        true,
		XTermLike:         true,
	})
}
