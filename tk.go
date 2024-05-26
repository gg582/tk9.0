// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tk9.0 is an idiomatic Go wrapper for [libtk9.0].
//
// Parts of the documentation are copied and/or modified from [TkDocs], see the
// LICENSE-TKDOCS file for details.
//
// [TkDocs]: https://tkdocs.com/about.html
// [libtk9.0]: https://pkg.go.dev/modernc.org/libtk9.0
package tk9_0 // import "modernc.org/tk9.0"

//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/bell.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/bind.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/bindtags.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/bitmap.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/busy.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/button.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/canvas.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/clipboard.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/colors.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/console.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/cursors.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/destroy.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/dialog.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/entry.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/event.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/focus.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/focusNext.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/fontchooser.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/font.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/frame.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/getOpenFile.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/grab.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/grid.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/checkbutton.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/chooseColor.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/chooseDirectory.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/image.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/keysyms.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/labelframe.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/label.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/listbox.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/loadTk.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/lower.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/menubutton.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/menu.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/messageBox.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/message.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/nsimage.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/optionMenu.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/option.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/options.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/pack.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/palette.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/panedwindow.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/photo.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/place.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/popup.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/print.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/radiobutton.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/raise.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/scale.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/scrollbar.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/selection.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/send.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/spinbox.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/sysnotify.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/systray.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/text.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/tkerror.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/tk_mac.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/tk.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/tkvars.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/tkwait.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/toplevel.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_button.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_combobox.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_entry.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_frame.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_checkbutton.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_image.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_intro.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_labelframe.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_label.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_menubutton.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_notebook.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_panedwindow.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_progressbar.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_radiobutton.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_scale.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_scrollbar.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_separator.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_sizegrip.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_spinbox.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_style.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_treeview.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_vsapi.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/ttk_widget.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/winfo.n
//	/home/jnml/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b1/doc/wm.n

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/evilsocket/islazy/zip"
	lib "modernc.org/libtk9.0"
	tklib "modernc.org/libtk9.0/library"
	tcl "modernc.org/tcl9.0"
)

var (
	id atomic.Int32

	tk     *Tk
	tkErr  error
	tkOnce sync.Once

	tclDir string
	tkDir  string

	finished atomic.Int32

	// CollectErrors selects the behaviour on errors for certain functions
	// that do not return error.
	//
	// When false, errors will panic, providing a stack trace.
	//
	// When true, errors will be recorded using errors.Join in the Error
	// variable.
	CollectErrors bool
	// Error records errors when ErrModeCollect is true.
	Error error
)

// stdlib returns the path to the Tk standard library or an error, if any. It
// once creates a temporary directory where the standard library is written.
// Subsequent calls to stdlib share the same temporary directory.
//
// stdlib is safe for concurrent access by multiple goroutines.
func stdlib() (dir string, err error) {
	if dir, err = os.MkdirTemp("", "tk-library-"); err != nil {
		return "", err
	}

	fn := filepath.Join(dir, "library.zip")
	if err = os.WriteFile(fn, []byte(tklib.Zip), 0600); err != nil {
		return
	}

	if _, err = zip.Unzip(fn, dir); err != nil {
		return
	}

	return filepath.Join(dir, "library"), nil
}

// Finalize releases all resources held, if any. Finalize is intended to be
// called on process shutdown only.
func Finalize() (err error) {
	if finished.Swap(1) != 0 {
		return
	}

	runtime.UnlockOSThread()
	if tk != nil {
		err = tk.in.Close()
		tk = nil
	}
	for _, v := range []string{tclDir, tkDir} {
		err = errors.Join(err, os.RemoveAll(v))
	}
	return err
}

// Tk represents the main window of an application. It has an associated Tcl
// interpreter.
//
// Note: Tk has all *Window methods promoted.
type Tk struct {
	*Window
	in *tcl.Interp

	trace bool
}

func (tk *Tk) eval(code string) (r string, err error) {
	if dmesgs {
		defer func() {
			dmesg("code=%s -> r=%v err=%v", code, r, err)
		}()
	}
	return tk.in.Eval(code, tcl.EvalGlobal)
}

// Initialize performs package initialization and returns a *Tk or error, if
// any.
//
// The returned value is a singleton. Calls to Initialize() are idempotent and
// all return the same (instance, error) tuple.
//
// Initialize will perform runtime.LockOSThread. All further uses of this
// package should be done using the same goroutine that first called
// Initialize.
func Initialize() (r *Tk, err error) {
	tkOnce.Do(func() {
		runtime.LockOSThread()
		if tclDir, tkErr = tcl.Stdlib(); err != nil {
			return
		}

		if tkDir, tkErr = stdlib(); tkErr != nil {
			return
		}

		var in *tcl.Interp
		if in, tkErr = tcl.NewInterp(map[string]string{
			"tcl_library": tclDir,
			"tk_library":  tkDir,
		}); tkErr != nil {
			return
		}

		if rc := lib.XTk_Init(in.TLS(), in.Handle()); rc != lib.TCL_OK {
			in.Close()
			tkErr = fmt.Errorf("failed to initialize the Tk subsystem")
			return
		}

		tk = &Tk{
			Window: &Window{},
			in:     in,
		}
	})
	return tk, tkErr
}

// Window represents a Tk window/widget.
type Window struct {
	fpath string
}

func (w *Window) path() (r string) {
	if r = w.fpath; r == "" {
		r = "."
	}
	return r
}

func (w *Window) newChild(nm string, opts ...Opt) *Window {
	cls := strings.Replace(nm, "ttk_", "ttk::", 1)
	nm = strings.Replace(nm, "ttk_", "t", 1)
	if c := nm[len(nm)-1]; c >= '0' && c <= '9' {
		nm += "_"
	}
	path := fmt.Sprintf("%s.%s%v", w.path(), nm, id.Add(1))
	var a []string
	for _, v := range opts {
		a = append(a, v.opt())
	}
	r, err := tk.eval(fmt.Sprintf("%s %s %s", cls, path, strings.Join(a, " ")))
	if err != nil {
		if !CollectErrors {
			panic(err)
		}

		Error = errors.Join(Error, err)
	}
	return &Window{fpath: r}
}

// https://pdos.csail.mit.edu/archive/rover/RoverDoc/escape_shell_table.html
//
// The following characters are dissallowed or have special meanings in Tcl and
// so are escaped:
//
//	&;`'"|*?~<>^()[]{}$\
var badChars = [...]bool{
	' ':  true,
	'"':  true,
	'$':  true,
	'&':  true,
	'(':  true,
	')':  true,
	'*':  true,
	';':  true,
	'<':  true,
	'>':  true,
	'?':  true,
	'[':  true,
	'\'': true,
	'\\': true,
	'\n': true,
	'\r': true,
	'\t': true,
	']':  true,
	'^':  true,
	'`':  true,
	'{':  true,
	'|':  true,
	'}':  true,
	'~':  true,
}

func tclSafeString(s string) string {
	const badString = "&;`'\"|*?~<>^()[]{}$\\\n\r\t "
	if strings.ContainsAny(s, badString) {
		var b strings.Builder
		for _, c := range s {
			switch {
			case int(c) < len(badChars) && badChars[c]:
				fmt.Fprintf(&b, "\\x%02x", c)
			default:
				b.WriteRune(c)
			}
		}
		s = b.String()
	}
	return s
}

// Opt represents an optional argument.
type Opt interface {
	opt() string
}

type text string

func (t text) opt() string {
	return fmt.Sprintf("-text %s", tclSafeString(string(t)))
}

// Txt produces the '-text s' configuration option.
func Txt(s string) Opt {
	return text(s)
}
