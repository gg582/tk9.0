// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tk9.0 is an idiomatic Go wrapper for [libtk9.0]. It is similar to
// Python's tkinter.
//
// # The options pattern
//
// Zero or more options can be specified when creating a widget. For example
//
//	b := Button(Text("Hello"), OverRelief("flat"))
//
// or
//
//	lbl := myFrame.Label(State("disabled"), Width(200))
//
// # Widget path names
//
// Tcl/Tk uses widget pathnames explicitly set by user code. This package
// generates pathnames automatically.
//
// # OS thread
//
// This package should be used from the same goroutine that initialized the
// package. Package initialization performs a runtime.LockOSThread.
//
// # Event handlers
//
// The Command() and similar options expect arguments that must be one of:
//
// - EventHandler or a function literal of signature func(*Window, any) (any, error).
//
// - EventDetacher or a function literal of signature func(*Window, any).
//
// - Any other type, used as the additonal 'data' argument when invoking the event handler/detacher.
//
// Each of the three types must be present at most once and only the event handler is mandatory.
// The event detacher and additional data are both optional.
//
// For convenience, additionally type 'func()' literals are accepted as well.
// They are converted internally to EventHandler and or EventDetacher. The
// first occurrence is used as an EventHandler if it was not yet specified or
// as an EventDetacher if that was not yet specified. The second instance is
// used as an EventDetacher if that was not yet specified.
//
// # time.Duration
//
// When passing an argument of type time.Durarion to a function accepting
// 'any', the duration is converted to an integer number of milliseconds.
//
// # Tcl/Tk code
//
// At least some minimal knowledge of Tcl/Tk is required for using this
// package. However you will not need to write any Tcl code and you do not need
// to care about the grammar of Tcl words/string literals and how it differs
// from Go.
//
// There are several Tcl/Tk tutorials available, for example at
// [tutorialspoint].
//
// # Notes
//
// Most of the documentation comments are generated directly from the Tcl/Tk
// documentation and may not be entirely correct for the Go package. They serve
// more as a quick/offline Tcl/Tk reference.
//
// # Additional copyrights
//
// Parts of the documentation are copied and/or modified from the [tcl.tk
// site], see the LICENSE-TCLTK file for details.
//
// Parts of the documentation are copied and/or modified from [TkDocs],
// see the LICENSE-TKDOCS file for details.
//
// [tutorialspoint]: https://www.tutorialspoint.com/tcl-tk/tk_overview.htm
// [TkDocs]: https://tkdocs.com/about.html
// [libtk9.0]: https://pkg.go.dev/modernc.org/libtk9.0
// [tcl.tk site]: https://www.tcl.tk/man/tcl9.0/TkCmd/index.html
package tk9_0 // import "modernc.org/tk9.0"

//TODO Use Tcl 'uknown' for prefixed commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/evilsocket/islazy/zip"
	libtcl "modernc.org/libtcl9.0"
	libtk "modernc.org/libtk9.0"
	tklib "modernc.org/libtk9.0/library"
	tcl "modernc.org/tcl9.0"
)

var (
	isBuilder = os.Getenv("MODERNC_BUILDER") != ""

	id atomic.Int32

	// Inter is the singleton Tk instance created on package initialization.
	Inter *Tk

	interErr  error
	interOnce sync.Once

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

	// https://pdos.csail.mit.edu/archive/rover/RoverDoc/escape_shell_table.html
	//
	// The following characters are dissallowed or have special meanings in Tcl and
	// so are escaped:
	//
	//	&;`'"|*?~<>^()[]{}$\
	badChars = [...]bool{
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
)

func init() {
	if isBuilder {
		return
	}

	initialize()
	if interErr != nil {
		panic(interErr)
	}
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

func (w *Window) newChild(nm string, options ...option) *Window {
	cls := strings.Replace(nm, "ttk_", "ttk::", 1)
	nm = strings.Replace(nm, "ttk_", "t", 1)
	if c := nm[len(nm)-1]; c >= '0' && c <= '9' {
		nm += "_"
	}
	path := fmt.Sprintf("%s.%s%v", w.path(), nm, id.Add(1))
	var a []string
	for _, v := range options {
		a = append(a, v.optionString(w))
	}
	code := fmt.Sprintf("%s %s %s", cls, path, strings.Join(a, " "))
	r, err := Inter.eval(code)
	if err != nil {
		Inter.fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
	}
	return &Window{fpath: r}
}

// option represents an optional argument.
type option interface {
	optionString(w *Window) string
}

// Tk represents the main window of an application. It has an associated Tcl
// interpreter.
//
// Note: Tk has all *Window methods promoted.
type Tk struct {
	*Window
	handlers map[int32]*eventHandler
	in       *tcl.Interp

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

func (tk *Tk) fail(err error) {
	if !CollectErrors {
		if dmesgs {
			dmesg("PANIC %v", err)
		}
		panic(err)
	}

	Error = errors.Join(Error, err)
}

// EventHandler is invoked when its associated event fires. The 'data' argument
// is the additional value passed when the handler was registered
type EventHandler func(w *Window, data any) (any, error)

// EventDetacher is invoked when the handler is detached. The 'data' argument
// is the additional value passed when the handler was registered.
type EventDetacher func(w *Window, data any)

type eventHandler struct {
	data     any
	detacher EventDetacher
	handler  EventHandler
	id       int32
	tcl      string
	w        *Window
}

func newEventHandler(option string, args ...any) (r *eventHandler) {
	if len(args) == 0 {
		Inter.fail(fmt.Errorf("registering event handler: need at least one argument"))
		return nil
	}

	var handler EventHandler
	var detacher EventDetacher
	var data any
	for _, v := range args {
		switch x := v.(type) {
		case EventHandler:
			if handler != nil {
				Inter.fail(fmt.Errorf("registering event handler: multiple handling functions"))
				return nil
			}

			handler = x
		case func(*Window, any) (any, error):
			if handler != nil {
				Inter.fail(fmt.Errorf("registering event handler: multiple handling functions"))
				return nil
			}

			handler = x
		case func():
			if handler != nil {
				if detacher != nil {
					Inter.fail(fmt.Errorf("registering event handler: multiple detaching functions"))
					return nil
				}

				detacher = func(*Window, any) { x() }
				break
			}

			handler = func(*Window, any) (any, error) { x(); return nil, nil }
		case EventDetacher:
			if detacher != nil {
				Inter.fail(fmt.Errorf("registering event handler: multiple detaching functions"))
				return nil
			}

			detacher = x
		case func(*Window, any):
			if detacher != nil {
				Inter.fail(fmt.Errorf("registering event handler: multiple detaching functions"))
				return nil
			}

			detacher = x
		default:
			if data != nil {
				Inter.fail(fmt.Errorf("registering event handler: multiple data values"))
				return nil
			}

			data = x
		}
	}
	if handler == nil {
		Inter.fail(fmt.Errorf("registering event handler: no event handler argument"))
		return nil
	}

	r = &eventHandler{
		handler:  handler,
		detacher: detacher,
		data:     data,
		id:       id.Add(1),
		tcl:      option,
	}
	switch {
	case r.detacher == nil:
		r.detacher = func(w *Window, v any) { delete(Inter.handlers, r.id) }
	default:
		r.detacher = func(w *Window, v any) {
			detacher(w, v)
			delete(Inter.handlers, r.id)
		}
	}
	Inter.handlers[r.id] = r
	return r
}

func (e *eventHandler) optionString(w *Window) string {
	if e == nil {
		return ""
	}

	e.w = w
	return fmt.Sprintf("%s {eventDispatcher %v}", e.tcl, e.id)
}

func optionString(v any) string {
	switch x := v.(type) {
	case time.Duration:
		return fmt.Sprint(int64((x + time.Millisecond/2) / time.Millisecond))
	default:
		return tclSafeString(fmt.Sprint(v))
	}
}

func tclSafeString(s string) string {
	if s == "" {
		return "{}"
	}

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

// initialize performs package initialization and returns a *Tk or error, if
// any.
//
// The returned value is a singleton. Calls to initialize() are idempotent and
// all return the same (instance, error) tuple.
//
// initialize will perform runtime.LockOSThread. All further uses of this
// package should be done using the same goroutine that first called
// initialize.
func initialize() (r *Tk, err error) {
	interOnce.Do(func() {
		runtime.LockOSThread()
		if tclDir, interErr = tcl.Stdlib(); err != nil {
			return
		}

		if tkDir, interErr = stdlib(); interErr != nil {
			return
		}

		var in *tcl.Interp
		if in, interErr = tcl.NewInterp(map[string]string{
			"tcl_library": tclDir,
			"tk_library":  tkDir,
		}); interErr != nil {
			return
		}

		if rc := libtk.XTk_Init(in.TLS(), in.Handle()); rc != libtk.TCL_OK {
			in.Close()
			interErr = fmt.Errorf("failed to initialize the Tk subsystem")
			return
		}

		Inter = &Tk{
			Window:   &Window{},
			handlers: map[int32]*eventHandler{},
			in:       in,
		}
		interErr = Inter.in.RegisterCommand("eventDispatcher", eventDispatcher, nil, nil)
	})
	return Inter, interErr
}

func eventDispatcher(data any, in *tcl.Interp, args []string) int {
	id, err := strconv.Atoi(args[1])
	if err != nil {
		panic(todo("event dispatcher internal error: %q", args))
	}

	h := Inter.handlers[int32(id)]
	r, err := h.handler(h.w, h.data)
	Inter.in.SetResult(tclSafeString(fmt.Sprint(r)))
	if err != nil {
		return libtcl.TCL_ERROR
	}

	return libtcl.TCL_OK
}

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
	if Inter != nil {
		err = Inter.in.Close()
		Inter = nil
	}
	for _, v := range []string{tclDir, tkDir} {
		err = errors.Join(err, os.RemoveAll(v))
	}
	return err
}

// bell — Ring a display's bell
//
// bell ?-displayof window? ?-nice?
// func Bell(opts ...option) {
// }

// Displayof

//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/bell.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/bind.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/bindtags.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/bitmap.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/busy.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/button.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/canvas.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/clipboard.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/colors.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/console.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/cursors.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/destroy.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/dialog.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/entry.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/event.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/focus.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/focusNext.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/fontchooser.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/font.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/frame.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/getOpenFile.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/grab.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/grid.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/checkbutton.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/chooseColor.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/chooseDirectory.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/image.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/keysyms.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/labelframe.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/label.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/listbox.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/loadTk.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/lower.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/menubutton.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/menu.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/messageBox.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/message.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/nsimage.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/optionMenu.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/option.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/options.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/pack.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/palette.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/panedwindow.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/photo.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/place.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/popup.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/print.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/radiobutton.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/raise.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/scale.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/scrollbar.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/selection.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/send.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/spinbox.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/sysnotify.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/systray.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/text.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/tkerror.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/tk_mac.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/tk.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/tkvars.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/tkwait.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/toplevel.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_button.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_combobox.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_entry.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_frame.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_checkbutton.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_image.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_intro.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_labelframe.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_label.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_menubutton.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_notebook.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_panedwindow.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_progressbar.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_radiobutton.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_scale.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_scrollbar.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_separator.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_sizegrip.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_spinbox.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_style.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_treeview.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_vsapi.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/ttk_widget.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/winfo.n
//TODO ~/.config/ccgo/v4/libtk9.0/linux/amd64/tk9.0b2/doc/wm.n
