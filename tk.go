// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tk9.0 is an idiomatic Go wrapper for [libtk9.0]. It is similar to
// Python's tkinter.
//
// # Event handlers
//
// The various command options, like Command() expect arguments that must be one of:
//
// - EventHandler or a function literal of signature func(*Window, any) (any, error), ie. the same as EventHandler.
//
// - EventDetacher or a function literal of signature func(*Window, any), ie. the same as EventDetacher.
//
// - Any other type, used as the additonal 'data' argument when invoking the event handler/detacher.
//
// Each of the three types must be present at most once and only the event handler is mandatory.
// The event detacher and additional data are both optional.
//
// Note: Parts of the documentation are copied and/or modified from [TkDocs],
// see the LICENSE-TKDOCS file for details.
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
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/evilsocket/islazy/zip"
	libtcl "modernc.org/libtcl9.0"
	lib "modernc.org/libtk9.0"
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

		if rc := lib.XTk_Init(in.TLS(), in.Handle()); rc != lib.TCL_OK {
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

func (w *Window) newChild(nm string, options ...Option) *Window {
	cls := strings.Replace(nm, "ttk_", "ttk::", 1)
	nm = strings.Replace(nm, "ttk_", "t", 1)
	if c := nm[len(nm)-1]; c >= '0' && c <= '9' {
		nm += "_"
	}
	path := fmt.Sprintf("%s.%s%v", w.path(), nm, id.Add(1))
	var a []string
	for _, v := range options {
		a = append(a, v.option(w))
	}
	code := fmt.Sprintf("%s %s %s", cls, path, strings.Join(a, " "))
	r, err := Inter.eval(code)
	if err != nil {
		Inter.fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
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

func bool2int(b bool) int {
	if b {
		return 1
	}

	return 0
}

// Option represents an optional argument.
type Option interface {
	option(w *Window) string
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
	Inter.handlers[r.id] = r
	return r
}

func (e *eventHandler) option(w *Window) string {
	if e == nil {
		return ""
	}

	e.w = w
	return fmt.Sprintf("%s {eventDispatcher %v}", e.tcl, e.id)
}

//TODO type xscrollcommandOption string
//TODO
//TODO func (o xscrollcommandOption) option(w *Window) string {
//TODO 	return fmt.Sprintf(`-xscrollcommand %s`, tclSafeString(string(o)))
//TODO }
//TODO
//TODO // Specifies the prefix for a command used to communicate with horizontal
//TODO // scrollbars.
//TODO // When the view in the widget's window changes (or
//TODO // whenever anything else occurs that could change the display in a
//TODO // scrollbar, such as a change in the total size of the widget's
//TODO // contents), the widget will
//TODO // generate a Tcl command by concatenating the scroll command and
//TODO // two numbers.
//TODO // Each of the numbers is a fraction between 0 and 1, which indicates
//TODO // a position in the document.  0 indicates the beginning of the document,
//TODO // 1 indicates the end, .333 indicates a position one third the way through
//TODO // the document, and so on.
//TODO // The first fraction indicates the first information in the document
//TODO // that is visible in the window, and the second fraction indicates
//TODO // the information just after the last portion that is visible.
//TODO // The command is
//TODO // then passed to the Tcl interpreter for execution.  Typically the
//TODO // '-xscrollcommand' option consists of the path name of a scrollbar
//TODO // widget followed by 'set', e.g. '.x.scrollbar set': this will cause
//TODO // the scrollbar to be updated whenever the view in the window changes. If this
//TODO // option is not specified, then no command will be executed.
//TODO //
//TODO // More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//TODO //
//TODO // Note: This option applies to all windows/widgets.
//TODO //
//TODO // [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-xscrollcommand
//TODO func Xscrollcommand(value string) Option {
//TODO 	return xscrollcommandOption(value)
//TODO }
//TODO
//TODO type yscrollcommandOption string
//TODO
//TODO func (o yscrollcommandOption) option(w *Window) string {
//TODO 	return fmt.Sprintf(`-yscrollcommand %s`, tclSafeString(string(o)))
//TODO }
//TODO
//TODO // Specifies the prefix for a command used to communicate with vertical
//TODO // scrollbars.  This option is treated in the same way as the
//TODO // '-xscrollcommand' option, except that it is used for vertical
//TODO // scrollbars and is provided by widgets that support vertical scrolling.
//TODO // See the description of '-xscrollcommand' for details
//TODO // on how this option is used.
//TODO //
//TODO // More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//TODO //
//TODO // Note: This option applies to all windows/widgets.
//TODO //
//TODO // [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-yscrollcommand
//TODO func Yscrollcommand(value string) Option {
//TODO 	return yscrollcommandOption(value)
//TODO }

// Specifies a Tcl command to associate with the button.  This command
// is typically invoked when mouse button 1 is released over the button
// window.
//
// More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//
// Note: This option applies to Button, Checkbutton, Radiobutton, Scale, Scrollbar, Spinbox, TButton, TCheckbutton, TRadiobutton, TScale, TScrollbar, TSpinbox.
//
// [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-command
func Command(args ...any) Option {
	return newEventHandler("-command", args...)
}

// If this option is specified then it provides a Tcl command to execute
// each time the menu is posted.  The command is invoked by the 'post'
// widget command before posting the menu. Note that in Tk 8.0 on Macintosh
// and Windows, all post-commands in a system of menus are executed before any
// of those menus are posted.
// This is due to the limitations in the individual platforms' menu managers.
//
// More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//
// Note: This option applies to Menu, TCombobox.
//
// [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-postcommand
func Postcommand(args ...any) Option {
	return newEventHandler("-postcommand", args...)
}

// If this option has a non-empty value, then it specifies a Tcl command
// to invoke whenever the menu is torn off.  The actual command will
// consist of the value of this option, followed by a space, followed
// by the name of the menu window, followed by a space, followed by
// the name of the name of the torn off menu window.  For example, if
// the option's value is
//
// More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//
// Note: This option applies to Menu.
//
// [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-tearoffcommand
func Tearoffcommand(args ...any) Option {
	return newEventHandler("-tearoffcommand", args...)
}

// Specifies a script to eval when '-validatecommand' returns 0.
// Setting it to {} disables this feature (the default).  The best use
// of this option is to set it to 'bell'.  See 'VALIDATION'
// below for more information.
//
// More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//
// Note: This option applies to Entry, Spinbox, TEntry.
//
// [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-invalidcommand
func Invalidcommand(args ...any) Option {
	return newEventHandler("-invalidcommand", args...)
}

// Specifies a script to eval when '-validatecommand' returns 0.
// Setting it to {} disables this feature (the default).  The best use
// of this option is to set it to 'bell'.  See 'VALIDATION'
// below for more information.
//
// More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//
// Note: This option applies to Entry, Spinbox.
//
// [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-invcmd
func Invcmd(args ...any) Option {
	return newEventHandler("-invcmd", args...)
}

// Specifies a script to eval when you want to validate the input into
// the entry widget.  Setting it to {} disables this feature (the default).
// This command must return a valid Tcl boolean value.  If it returns 0 (or
// the valid Tcl boolean equivalent) then it means you reject the new edition
// and it will not occur and the '-invalidcommand' will be evaluated if it
// is set. If it returns 1, then the new edition occurs.
// See 'VALIDATION' below for more information.
//
// More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//
// Note: This option applies to Entry, Spinbox, TEntry.
//
// [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-validatecommand
func Validatecommand(args ...any) Option {
	return newEventHandler("-validatecommand", args...)
}

// Specifies a script to eval when you want to validate the input into
// the entry widget.  Setting it to {} disables this feature (the default).
// This command must return a valid Tcl boolean value.  If it returns 0 (or
// the valid Tcl boolean equivalent) then it means you reject the new edition
// and it will not occur and the '-invalidcommand' will be evaluated if it
// is set. If it returns 1, then the new edition occurs.
// See 'VALIDATION' below for more information.
//
// More details about the option and the values it accepts can be possibly found at the [Tcl/Tk documentation].
//
// Note: This option applies to Entry, Spinbox.
//
// [Tcl/Tk documentation]: https://www.tcl.tk/man/tcl9.0/TkCmd/options.html#M-vcmd
func Vcmd(args ...any) Option {
	return newEventHandler("-vcmd", args...)
}
