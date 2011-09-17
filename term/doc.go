// Package term provides a basic terminal emulation framework.
//
//   The TTY class abstracts the logic of providing line editing and line
// buffering, as well as escape sequence recognition.  When creating a TTY with
// the NewTTY function, interactive echo will automatically be enabled if the
// provided io.Reader's underlying object also implements io.Writer.
//
// Line editing capabilities
//
//   The line editing facilities are very basic; you can type, and you can
// backspace out characters up to the beginning of the line.  Note that for all
// internal purposes, typing a control character (e.g. ^D or ^C) starts a new
// line, including for line history below.
//
// Line history
//
//   Currently the TTY only has a single-line history.  Pressing the return key
// will save the current line in that history, and pressing the "up" arrow at
// any time will restore the previous line.
package term
