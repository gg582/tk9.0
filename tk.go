// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tk9.0 is an idiomatic Go wrapper for [libtk9.0]. It is similar to
// Python's tkinter.
//
// # Supported platforms and architectures
//
// These combinations of GOOS and GOARCH are currently supported
//
//	OS      Arch
//	-------------
//	linux	amd64
//
// # Package initialization
//
// User code should check if the package variable 'Error' is nil before using
// this package.  If package initialization failed the 'Error' variable will be
// non nil.
//
// # The options pattern
//
// Zero or more options can be specified when creating a widget. For example
//
//	b := Button(Txt("Hello"), OverRelief("flat"))
//
// or
//
//	lbl := myFrame.Label(State("disabled"), Width(200))
//
// # Widget path names, image names
//
// Tcl/Tk uses widget pathnames and image names explicitly set by user code.
// This package generates those names automatically.
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

//TODO Use Tcl 'unknown' for prefixed commands?

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
	// App is the main/root application window.
	App *Window

	// Error records errors when ErrModeCollect is true.
	Error error

	// CollectErrors selects the behaviour on errors for certain functions that do
	// not return error.
	//
	// When false, errors will panic, providing a stack trace.
	//
	// When true, errors will be recorded using errors.Join in the Error variable.
	CollectErrors bool

	finished  atomic.Int32
	handlers  = map[int32]*eventHandler{}
	id        atomic.Int32
	initOnce  sync.Once
	interp    *tcl.Interp
	isBuilder = os.Getenv("MODERNC_BUILDER") != ""
	tclDir    string
	tkDir     string

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

	initOnce.Do(func() {
		runtime.LockOSThread()
		if tclDir, Error = tcl.Stdlib(); Error != nil {
			return
		}

		if tkDir, Error = stdlib(); Error != nil {
			return
		}

		if interp, Error = tcl.NewInterp(map[string]string{
			"tcl_library": tclDir,
			"tk_library":  tkDir,
		}); Error != nil {
			return
		}

		if rc := libtk.XTk_Init(interp.TLS(), interp.Handle()); rc != libtk.TCL_OK {
			interp.Close()
			Error = fmt.Errorf("failed to initialize the Tk subsystem")
			return
		}

		Error = interp.RegisterCommand("eventDispatcher", eventDispatcher, nil, nil)
	})
	if Error == nil {
		App = &Window{}
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

func (w *Window) optionString(_ *Window) string {
	return w.path()
}

func (w *Window) newChild(nm string, options ...option) *Window {
	class := strings.Replace(nm, "ttk_", "ttk::", 1)
	nm = strings.Replace(nm, "ttk_", "t", 1)
	if c := nm[len(nm)-1]; c >= '0' && c <= '9' {
		nm += "_"
	}
	path := fmt.Sprintf("%s.%s%v", w.path(), nm, id.Add(1))
	code := fmt.Sprintf("%s %s %s", class, path, winCollect(w, options...))
	r, err := eval(code)
	if err != nil {
		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
	}
	return &Window{fpath: r}
}

func eval(code string) (r string, err error) {
	if dmesgs {
		defer func() {
			dmesg("code=%s -> r=%v err=%v", code, r, err)
		}()
	}
	return interp.Eval(code, tcl.EvalGlobal)
}

func evalAny(code string) (r any) {
	s := evalErr(code)
	if s == "" {
		return nil
	}

	if n, err := strconv.ParseInt(s, 0, 64); err == nil {
		return n
	}

	if n, err := strconv.ParseFloat(s, 64); err == nil {
		return n
	}

	return s
}

func evalErr(code string) (r string) {
	r, err := eval(code)
	if err != nil {
		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
	}
	return r
}

func fail(err error) {
	if !CollectErrors {
		if dmesgs {
			dmesg("PANIC %v", err)
		}
		panic(err)
	}

	Error = errors.Join(Error, err)
}

func winCollect(w *Window, options ...option) string {
	var a []string
	for _, v := range options {
		a = append(a, v.optionString(w))
	}
	return strings.Join(a, " ")
}

func collect(options ...option) string {
	var a []string
	for _, v := range options {
		a = append(a, v.optionString(nil))
	}
	return strings.Join(a, " ")
}

// option represents an optional argument.
type option interface {
	optionString(w *Window) string
}

type stringOption string

func (s stringOption) optionString(w *Window) string {
	return tclSafeString(string(s))
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
		fail(fmt.Errorf("registering event handler: need at least one argument"))
		return nil
	}

	var handler EventHandler
	var detacher EventDetacher
	var data any
	for _, v := range args {
		switch x := v.(type) {
		case EventHandler:
			if handler != nil {
				fail(fmt.Errorf("registering event handler: multiple handling functions"))
				return nil
			}

			handler = x
		case func(*Window, any) (any, error):
			if handler != nil {
				fail(fmt.Errorf("registering event handler: multiple handling functions"))
				return nil
			}

			handler = x
		case func():
			if handler != nil {
				if detacher != nil {
					fail(fmt.Errorf("registering event handler: multiple detaching functions"))
					return nil
				}

				detacher = func(*Window, any) { x() }
				break
			}

			handler = func(*Window, any) (any, error) { x(); return nil, nil }
		case EventDetacher:
			if detacher != nil {
				fail(fmt.Errorf("registering event handler: multiple detaching functions"))
				return nil
			}

			detacher = x
		case func(*Window, any):
			if detacher != nil {
				fail(fmt.Errorf("registering event handler: multiple detaching functions"))
				return nil
			}

			detacher = x
		default:
			if data != nil {
				fail(fmt.Errorf("registering event handler: multiple data values"))
				return nil
			}

			data = x
		}
	}
	if handler == nil {
		fail(fmt.Errorf("registering event handler: no event handler argument"))
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
		r.detacher = func(w *Window, v any) { delete(handlers, r.id) }
	default:
		r.detacher = func(w *Window, v any) {
			detacher(w, v)
			delete(handlers, r.id)
		}
	}
	handlers[r.id] = r
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
	case *Img:
		return x.name
	default:
		return tclSafeString(fmt.Sprint(v))
	}
}

func tclSafeStringBind(s string) string {
	if s == "" {
		return "{}"
	}

	const badString = "&;`'\"|*?~^()[]{}$\\\n\r\t "
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

func eventDispatcher(data any, interp *tcl.Interp, args []string) int {
	id, err := strconv.Atoi(args[1])
	if err != nil {
		panic(todo("event dispatcher internal error: %q", args))
	}

	h := handlers[int32(id)]
	r, err := h.handler(h.w, h.data)
	interp.SetResult(tclSafeString(fmt.Sprint(r)))
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

	defer runtime.UnlockOSThread()

	if interp != nil {
		err = interp.Close()
		interp = nil
	}
	for _, v := range []string{tclDir, tkDir} {
		err = errors.Join(err, os.RemoveAll(v))
	}
	return err
}

// bind — Arrange for X events to invoke functions
//
// Example: Bind all TLabel widgets mouse button 1 click to a function printing
// "Clicked".
//
//	Bind("TLabel", "<1>", Command(func() { fmt.Println("Clicked!") }))
//
// Additional information might be available at the [Tcl/Tk bind] page.
//
// [Tcl/Tk bind]: https://www.tcl.tk/man/tcl9.0/TkCmd/bind.html
func Bind(options ...any) {
	a := []string{"bind"}
	for _, v := range options {
		switch x := v.(type) {
		case *Window:
			a = append(a, x.path())
		case *eventHandler:
			x.tcl = ""
			a = append(a, x.optionString(nil))
		default:
			a = append(a, tclSafeStringBind(fmt.Sprint(x)))
		}
	}
	evalErr(strings.Join(a, " "))
}

// Img represents a Tk image.
type Img struct {
	name string
}

func (m *Img) optionString(_ *Window) string {
	return m.name
}

// bitmap — Images that display two colors
//
// A bitmap is an image whose pixels can display either of two colors or be
// transparent. A bitmap image is defined by four things: a background color, a
// foreground color, and two bitmaps, called the source and the mask. Each of
// the bitmaps specifies 0/1 values for a rectangular array of pixels, and the
// two bitmaps must have the same dimensions. For pixels where the mask is
// zero, the image displays nothing, producing a transparent effect. For other
// pixels, the image displays the foreground color if the source data is one
// and the background color if the source data is zero.
//
// Additional information might be available at the [Tcl/Tk bitmap] page.
//
// [Tcl/Tk bitmap]: https://www.tcl.tk/man/tcl9.0/TkCmd/bitmap.html
func ImageCreateBitmap(options ...option) *Img {
	nm := fmt.Sprintf("img%v", id.Add(1))
	code := fmt.Sprintf("image create bitmap %s %s", nm, collect(options...))
	r, err := eval(code)
	if err != nil {
		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
		return nil
	}

	return &Img{name: nm}
}

// photo — Full-color images
//
// A photo is an image whose pixels can display any color with a varying degree
// of transparency (the alpha channel). A photo image is stored internally in
// full color (32 bits per pixel), and is displayed using dithering if
// necessary. Image data for a photo image can be obtained from a file or a
// string, or it can be supplied from C code through a procedural interface. At
// present, only PNG, GIF, PPM/PGM, and (read-only) SVG formats are supported,
// but an interface exists to allow additional image file formats to be added
// easily. A photo image is (semi)transparent if the image data it was obtained
// from had transparency information. In regions where no image data has been
// supplied, it is fully transparent. Transparency may also be modified with
// the transparency set subcommand.
//
// Additional information might be available at the [Tcl/Tk photo] page.
//
// [Tcl/Tk photo]: https://www.tcl.tk/man/tcl9.0/TkCmd/photo.html
func ImageCreatePhoto(options ...option) *Img {
	nm := fmt.Sprintf("img%v", id.Add(1))
	code := fmt.Sprintf("image create photo %s %s", nm, collect(options...))
	r, err := eval(code)
	if err != nil {
		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
		return nil
	}

	return &Img{name: nm}
}

func collect0(f func(any) option) string {
	s := f("").optionString(nil)
	s = s[:len(s)-len("{}")-1]
	return s
}
