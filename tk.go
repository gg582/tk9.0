// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tk9.0 is a CGo-free cross platform GUI for Go. It is similar to
// [tkinter] for Python.
//
// # Tk8
//
// An older version of Tcl/Tk exists at [tk8].
//
// # Hello world
//
// Also available in _examples/hello.go
//
//	package main
//
//	import . "modernc.org/tk9.0"
//
//	func main() {
//		Pack(Button(Txt("Hello"), Command(func() { Destroy(App) })))
//		App.Wait()
//	}
//
// To execute the above program on any supported target issue something like
//
//	$ CGO_ENABLED=0 go run hello.go
//
// The CGO_ENABLED=0 is optional and here it only demonstrates the program can
// be built without CGo.
//
// # Debugging
//
// Consider this program in _examples/debugging.go:
//
//	// Build this program using -tags=tk.dmesg
//	package main
//
//	import . "modernc.org/tk9.0"
//
//	func main() {
//		Pack(
//			TButton(Txt("Hello"), Command(func() { Destroy(App) })),
//			Ipadx(10), Ipady(5), Padx(20), Pady(10),
//		)
//		App.Wait()
//	}
//
// Execute the program using the tags as indicated, then close the window or
// click the Hello button. With the tk.dmesg tag the package initialization
// prints the debug messages path. So we can view it, for example, like this:
//
//	$ cat $(go run -tags=tk.dmesg ./_examples/debugging.go)
//	[27190 debugging] enter [dmesgon.go:32:0 proc.go:7176:doInit1 proc.go:7143:doInit]
//	[27190 debugging] code=ttk::button ..ttk::button2 -text Hello -command {eventDispatcher 1} -> r=.ttk::button2 err=<nil> [tk.go:304:eval tk.go:304:eval tk.go:291:newChild]
//	[27190 debugging] code=pack .ttk::button2 -ipadx 10 -ipady 5 -padx 20 -pady 10 -> r= err=<nil> [tk.go:304:eval tk.go:304:eval tk.go:800:Pack]
//	[27190 debugging] code=destroy . -> r= err=<nil> [tk.go:304:eval tk.go:304:eval tk.go:704:Destroy]
//	[27190 debugging] code=tkwait window . -> r= err=<nil> [tk.go:304:eval tk.go:304:eval tk.go:809:Wait]
//	$
//
// 27190 was the process PID in this particular run. Using the tags allows to
// inspect the Tcl/Tk code executed during the lifetime of the process.
//
// # Supported targets
//
// These combinations of GOOS and GOARCH are currently supported
//
//	OS      Arch
//	-------------
//	darwin  arm64
//	linux   386
//	linux   amd64
//	linux   arm
//	linux   arm64
//	linux   loong64
//	linux   ppc64le
//	linux   riscv64
//	linux   s390x
//
// # Runtime dependencies
//
//   - [Img.Graph] and [CanvasWidget.Graph] require the gnuplot 5.4+
//     executable available in $PATH.
//   - darwin (macOS) requires [XQuartz].
//
// # Completeness
//
// At the moment the package does not cover all of the functionality C Tcl/Tk
// provides, but only a minimal subset considered to be usable. Please report
// needed, but missing and not yet requested Tcl/Tk features at the [issue
// tracker], thanks.
//
// # Error handling
//
// The [CollectErrors] variable selects the behaviour on errors for certain
// functions that do not return error.
//
// When CollectErrors is false, errors will panic, providing a stack trace.
//
// When CollectErrors is true, errors will be recorded using [errors.Join] in
// the [Error] variable.  Even if a function does not return error, it is still
// possible to handle errors in the usual way when needed, except that Error is
// now a static variable. That's a problem in the general case, but less so in
// this package that must be used from a single goroutine only, as
// documented elsewhere.
//
//	// Explicit error handling.
//	CollectErrors = true
//	if SomeFunnction(someArgument); Error != nil {
//		... error handling goes here
//	}
//
// This is obviously a compromise enabling to have a way to check for errors
// and, at the same time, the ability to write concise code like:
//
//	// Deferred error handling.
//	if b = Button(Txt("Foo"), Padx(5), Pady(10)); Error != nil {
//		...
//	}
//
// There are altogether four different places where the call to the Button
// function can produce errors as additionally to the call itself, every of its
// three arguments can independently fail as well.  Checking each and one of
// them separately is not always necessary in GUI code. But the explicit option
// in the first example is still available when needed.
//
// # Package initialization
//
// User code should check if the package variable 'Error' is nil before using
// this package.  If package initialization failed, the 'Error' variable will
// be non nil.
//
// Initialization will fail if a Unix process starts on a machine with no
// X server or the process is started in a way that it has no access to the X
// server. On the other hand, this package will work on Unix machines with no X
// server if the process is started remotely using '$ ssh -X foo@bar'.
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
// # Widget path names, image and font names
//
// Tcl/Tk uses widget pathnames, image and font names explicitly set by user
// code.  This package generates those names automatically and they are not
// directly needed in code that uses this package.
//
// # Renamed options
//
// There is, for a example, a Tcl/tk 'text' widget and a '-text' option. This
// package exports the widget as 'Text' and the option as 'Txt'. The complete
// list is:
//
//   - [Button]  option is renamed to [Btn]
//   - [Font]    option is renamed to [Fnt]
//   - [Label]   option is renamed to [Lbl]
//   - [Menu]    option is renamed to [Mnu]
//   - [Message] option is renamed to [Msg]
//   - [Text]    option is renamed to [Txt]
//
// # OS thread
//
// This package should be used from the same goroutine that initialized the
// package. Package initialization performs a runtime.LockOSThread, meaning
// func main() will start execuing locked on the same OS thread.
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
// # Specially handled types
//
//   - [time.Duration]
//
// When passing an argument of type [time.Durarion] to a function accepting
// 'any', the duration is converted to an integer number of milliseconds.
//
//   - []byte
//
// When passing an argument of type []byte to a function accepting 'any', the
// byte slice is converted to a [encoding/base64] encoded string.
//
//   - []FileType
//
// When passing an argument of type []FileType to a function accepting 'any',
// the slice is converted to the representation the Tcl/Tk -filetypes option
// expects.
//
// # Tcl/Tk code
//
// At least some minimal knowledge of reading Tcl/Tk code is probably required
// for using this package and/or using the related documentation. However you
// will not need to write any Tcl code and you do not need to care about the
// grammar of Tcl words/string literals and how it differs from Go.
//
// There are several Tcl/Tk tutorials available, for example at
// [tutorialspoint].
//
// # Notes
//
// Most of the documentation is generated directly from the Tcl/Tk
// documentation and may not be entirely correct for the Go package. Those
// parts hopefully still serve as a quick/offline Tcl/Tk reference.
//
// # Additional copyrights
//
// Parts of the documentation are copied and/or modified from the [tcl.tk
// site], see the LICENSE-TCLTK file for details.
//
// [XQuartz]: https://en.wikipedia.org/wiki/XQuartz
// [issue tracker]: https://gitlab.com/cznic/tk9.0/-/issues
// [tcl.tk site]: https://www.tcl.tk/man/tcl9.0/TkCmd/index.html
// [tk8]: https://pkg.go.dev/modernc.org/tk8.6
// [tkinter]: https://en.wikipedia.org/wiki/Tkinter
// [tutorialspoint]: https://www.tutorialspoint.com/tcl-tk/tk_overview.htm
package tk9_0 // import "modernc.org/tk9.0"

import (
	"context"
	_ "embed"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"os/exec"
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

const (
	gnuplotTimeout = time.Minute //TODO do not let the UI freeze
)

// App is the main/root application window.
var App *Window

// CollectErrors selects the behaviour on errors for certain functions that do
// not return error.
var CollectErrors bool

// Error records errors when [CollectErrors] is true.
var Error error

var (
	_ Widget = (*Window)(nil)

	//go:embed embed/gotk.png
	icon []byte

	exitHandler Opt
	finished    atomic.Int32
	handlers    = map[int32]*eventHandler{}
	id          atomic.Int32
	initOnce    sync.Once
	interp      *tcl.Interp
	isBuilder   = os.Getenv("MODERNC_BUILDER") != ""
	tclDir      string
	tkDir       string

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

	//TODO remove the associated tcl var on window destroy event both from the
	//interp and this map.
	textVariables = map[*Window]string{} // : tclName
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
		exitHandler = Command(func() { Destroy(App) })
		// Set some defaults.
		App.Center()
		App.IconPhoto(NewPhoto(Data(icon)))
		App.WmTitle(filepath.Base(os.Args[0]))
	}
}

// Window represents a Tk window/widget. It implements common widget methods.
//
// Window implements Opt. When a Window instance is used as an Opt, it provides
// its path name.
type Window struct {
	fpath string
}

func (w *Window) isWidget() {}

// Widget is implemented by every *Window
type Widget interface {
	isWidget()
	path() string
}

func (w *Window) path() (r string) {
	return w.String()
}

// String implements fmt.Stringer.
func (w *Window) String() (r string) {
	if r = w.fpath; r == "" {
		r = "."
	}
	return r
}

func (w *Window) optionString(_ *Window) string {
	return w.String()
}

func (w *Window) split(options []Opt) (opts []Opt, tvs []textVarOpt) {
	for _, v := range options {
		switch x := v.(type) {
		case textVarOpt:
			tvs = append(tvs, x)
		default:
			opts = append(opts, x)
		}
	}
	return opts, tvs
}

func (w *Window) newChild(nm string, options ...Opt) (rw *Window) {
	class := strings.Replace(nm, "ttk_", "ttk::", 1)
	nm = strings.Replace(nm, "ttk::", "t", 1)
	if c := nm[len(nm)-1]; c >= '0' && c <= '9' {
		nm += "_"
	}
	path := fmt.Sprintf("%s.%s%v", w, nm, id.Add(1))
	options, tvs := w.split(options)
	code := fmt.Sprintf("%s %s %s", class, path, winCollect(w, options...))
	r, err := eval(code)
	rw = &Window{fpath: r}
	if err != nil {
		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
	}
	if len(tvs) != 0 {
		rw.Configure(tvs[len(tvs)-1])
	}
	return rw
}

func eval(code string) (r string, err error) {
	if dmesgs {
		defer func() {
			dmesg("code=%s -> r=%v err=%v", code, r, err)
		}()
	}
	return interp.Eval(code, tcl.EvalGlobal)
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

func winCollect(w *Window, options ...Opt) string {
	var a []string
	for _, v := range options {
		a = append(a, v.optionString(w))
	}
	return strings.Join(a, " ")
}

func collect(options ...Opt) string {
	var a []string
	for _, v := range options {
		a = append(a, v.optionString(nil))
	}
	return strings.Join(a, " ")
}

// Opts is a list of options. It implements Opt.
type Opts []Opt

func (o Opts) optionString(w *Window) string {
	return winCollect(w, []Opt(o)...)
}

// Opt represents an optional argument.
type Opt interface {
	optionString(w *Window) string
}

type rawOption string

func (s rawOption) optionString(w *Window) string {
	return string(s)
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
	case []byte:
		return base64.StdEncoding.EncodeToString(x)
	case []FileType:
		var a []string
		for _, v := range x {
			a = append(a, fmt.Sprintf("{%s {%s} %s}", tclSafeString(v.TypeName), tclSafeStrings(v.Extensions...), v.MacType))
		}
		return fmt.Sprintf("{%s}", strings.Join(a, " "))
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

func tclSafeStrings(s ...string) string {
	var a []string
	for _, v := range s {
		a = append(a, tclSafeString(v))
	}
	return strings.Join(a, " ")
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

// Finalize releases all resources held, if any. This may include temporary
// files. Finalize is intended to be called on process shutdown only.
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
// # Description
//
// Bind tag options...
//
// The bind command associates commands with X events. If all three
// arguments are specified, bind will arrange for script (a Tcl script called
// the “binding script”) to be evaluated whenever the event(s) given by
// sequence occur in the window(s) identified by tag. If script is prefixed
// with a “+”, then it is appended to any existing binding for sequence;
// otherwise script replaces any existing binding. If script is an empty string
// then the current binding for sequence is destroyed, leaving sequence
// unbound. In all of the cases where a script argument is provided, bind
// returns an empty string.
//
// If sequence is specified without a script, then the script currently bound
// to sequence is returned, or an empty string is returned if there is no
// binding for sequence. If neither sequence nor script is specified, then the
// return value is a list whose elements are all the sequences for which there
// exist bindings for tag.
//
// The tag argument determines which window(s) the binding applies to. If tag
// begins with a dot, as in .a.b.c, then it must be the path name for a window;
// otherwise it may be an arbitrary string. Each window has an associated list
// of tags, and a binding applies to a particular window if its tag is among
// those specified for the window. Although the bindtags command may be used to
// assign an arbitrary set of binding tags to a window, the default binding
// tags provide the following behavior:
//
//   - If a tag is the name of an internal window the binding applies to that window.
//   - If the tag is the name of a class of widgets, such as Button, the binding applies to all widgets in that class.
//   - If the tag is the name of a toplevel window the binding applies to the toplevel window and all its internal windows.
//   - If tag has the value all, the binding applies to all windows in the application.
//
// Additional information might be available at the [Tcl/Tk bind] page.
//
// [Tcl/Tk bind]: https://www.tcl.tk/man/tcl9.0/TkCmd/bind.htm
func Bind(options ...any) {
	a := []string{"bind"}
	for _, v := range options {
		switch x := v.(type) {
		case *Window:
			a = append(a, x.String())
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

// String implements fmt.Stringer.
func (m *Img) String() string {
	return m.optionString(nil)
}

func (m *Img) optionString(_ *Window) string {
	if m != nil {
		return m.name
	}

	return "img0" // does not exist
}

// Bitmap — Images that display two colors
//
// # Description
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
//   - [Background] color
//
// Specifies a background color for the image in any of the standard ways
// accepted by Tk. If this option is set to an empty string then the background
// pixels will be transparent. This effect is achieved by using the source
// bitmap as the mask bitmap, ignoring any -maskdata or -maskfile options.
//
//   - [Data] string
//
// Specifies the contents of the source bitmap as a string. The string must
// adhere to X11 bitmap format (e.g., as generated by the bitmap program). If
// both the -data and -file options are specified, the -data option takes
// precedence.
//
//   - [File] name
//
// name gives the name of a file whose contents define the source bitmap. The
// file must adhere to X11 bitmap format (e.g., as generated by the bitmap
// program).
//
//   - [Foreground] color
//
// Specifies a foreground color for the image in any of the standard ways
// accepted by Tk.
//
//   - [Maskdata] string
//
// Specifies the contents of the mask as a string. The string must adhere to
// X11 bitmap format (e.g., as generated by the bitmap program). If both the
// -maskdata and -maskfile options are specified, the -maskdata option takes
// precedence.
//
//   - [Maskfile] name
//
// name gives the name of a file whose contents define the mask. The file must
// adhere to X11 bitmap format (e.g., as generated by the bitmap program).
//
// [Tcl/Tk bitmap]: https://www.tcl.tk/man/tcl9.0/TkCmd/bitmap.htm
func NewBitmap(options ...Opt) *Img {
	nm := fmt.Sprintf("bmp%v", id.Add(1))
	code := fmt.Sprintf("image create bitmap %s %s", nm, collect(options...))
	r, err := eval(code)
	if err != nil {
		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
		return nil
	}

	return &Img{name: nm}
}

// Photo — Full-color images
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
//   - [Data] string
//
// Specifies the contents of the image as a string. The string should contain
// binary data or, for some formats, base64-encoded data (this is currently
// guaranteed to be supported for PNG and GIF images). The format of the string
// must be one of those for which there is an image file format handler that
// will accept string data. If both the -data and -file options are specified,
// the -file option takes precedence.
//
//   - [Format] format-name
//
// Specifies the name of the file format for the data specified with the -data
// or -file option.
//
//   - [File] name
//
// name gives the name of a file that is to be read to supply data for the
// photo image. The file format must be one of those for which there is an
// image file format handler that can read data.
//
//   - [Gamma] value
//
// Specifies that the colors allocated for displaying this image in a window
// should be corrected for a non-linear display with the specified gamma
// exponent value. (The intensity produced by most CRT displays is a power
// function of the input value, to a good approximation; gamma is the exponent
// and is typically around 2). The value specified must be greater than zero.
// The default value is one (no correction). In general, values greater than
// one will make the image lighter, and values less than one will make it
// darker.
//
//   - [Height] number
//
// Specifies the height of the image, in pixels. This option is useful
// primarily in situations where the user wishes to build up the contents of
// the image piece by piece. A value of zero (the default) allows the image to
// expand or shrink vertically to fit the data stored in it.
//
//   - [Palette] palette-spec
//
// Specifies the resolution of the color cube to be allocated for displaying
// this image, and thus the number of colors used from the colormaps of the
// windows where it is displayed. The palette-spec string may be either a
// single decimal number, specifying the number of shades of gray to use, or
// three decimal numbers separated by slashes (/), specifying the number of
// shades of red, green and blue to use, respectively. If the first form (a
// single number) is used, the image will be displayed in monochrome (i.e.,
// grayscale).
//
//   - [Width] number
//
// Specifies the width of the image, in pixels. This option is useful primarily
// in situations where the user wishes to build up the contents of the image
// piece by piece. A value of zero (the default) allows the image to expand or
// shrink horizontally to fit the data stored in it.
//
// Additional information might be available at the [Tcl/Tk photo] page.
//
// [Tcl/Tk photo]: https://www.tcl.tk/man/tcl9.0/TkCmd/photo.htm
func NewPhoto(options ...Opt) *Img {
	nm := fmt.Sprintf("img%v", id.Add(1))
	code := fmt.Sprintf("image create photo %s %s", nm, collect(options...))
	r, err := eval(code)
	if err != nil {
		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
		return nil
	}

	return &Img{name: nm}
}

// Width — Get the configured option value.
func (m *Img) Width() string {
	return evalErr(fmt.Sprintf(`%s cget -width`, m))
}

// Height — Get the configured option value.
func (m *Img) Height() string {
	return evalErr(fmt.Sprintf(`%s cget -height`, m))
}

// Graph — use gnuplot to draw on a photo. Graph returns 'm'
//
// The 'script' argument is passed to a gnuplot executable, which must be
// installed on the machine.  See the [gnuplot site] for documentation about
// producing graphs. The script must not use the 'set term <device>' command.
//
// The content of 'm' is replaced, including its internal name.
//
// [gnuplot site]: http://www.gnuplot.info/
func (m *Img) Graph(script string) *Img {
	switch {
	case strings.HasPrefix(m.name, "img"):
		w, h := m.Width(), m.Height()
		script = fmt.Sprintf("set terminal pngcairo size %s, %s\n%s", w, h, script)
		out, err := gnuplot(script)
		if err != nil {
			fail(fmt.Errorf("plot: executing script: %s", err))
			break
		}

		*m = *NewPhoto(Width(w), Height(h), Data(out))
	default:
		fail(fmt.Errorf("plot: %s is not a photo", m))
	}
	return m
}

// Destroy — Destroy one or more windows
//
// # Description
//
// This command deletes the windows given by the window arguments, plus all of
// their descendants. If a window “.” (App) is deleted then all windows will be
// destroyed and the application will (normally) exit. The windows are
// destroyed in order, and if an error occurs in destroying a window the
// command aborts without destroying the remaining windows. No error is
// returned if window does not exist.
func Destroy(options ...Opt) {
	evalErr(fmt.Sprintf("destroy %s", collect(options...)))
}

// Pack — Geometry manager that packs around edges of cavity
//
// # Description
//
// The options consist of one or more content windows followed
// by options that specify how to manage the content. See THE PACKER
// ALGORITHM for details on how the options are used by the packer.
//
// The first argument must be a *Window.
//
// The following options are supported:
//
//   - [After] other
//
// Other must the name of another window. Use its container as the container
// for the content, and insert the content just after other in the packing
// order.
//
//   - [Anchor] anchor
//
// Anchor must be a valid anchor position such as n or sw; it specifies where
// to position each content in its parcel. Defaults to center.
//
//   - [Before] other
//
// Other must the name of another window. Use its container as the container
// for the content, and insert the content just before other in the packing
// order.
//
//   - [Expand] boolean
//
// Specifies whether the content should be expanded to consume extra space in
// their container. Boolean may have any proper boolean value, such as 1 or no.
// Defaults to 0.
//
//   - [Fill] style
//
// If a content's parcel is larger than its requested dimensions, this option
// may be used to stretch the content. Style must have one of the following
// values:
//
//   - "none" - Give the content its requested dimensions plus any internal
//     padding requested with -ipadx or -ipady. This is the default.
//   - "x" - Stretch the content horizontally to fill the entire width of its
//     parcel (except leave external padding as specified by -padx).
//   - "y" - Stretch the content vertically to fill the entire height of its parcel
//     (except leave external padding as specified by -pady).
//   - "both": Stretch the content both horizontally and vertically.
//
// .
//
//   - [In] container
//
// Insert the window at the end of the packing order for the container window
// given by container.
//
//   - [Ipadx] amount
//
// Amount specifies how much horizontal internal padding to leave on each side
// of the content. Amount must be a valid screen distance, such as 2 or .5c. It
// defaults to 0.
//
//   - [Ipady] amount
//
// Amount specifies how much vertical internal padding to leave on each side of
// the content. Amount defaults to 0.
//
//   - [Padx] amount
//
// Amount specifies how much horizontal external padding to leave on each side
// of the content. Amount may be a list of two values to specify padding for
// left and right separately. Amount defaults to 0.
//
//   - [Pady] amount
//
// Amount specifies how much vertical external padding to leave on each side of
// the content. Amount may be a list of two values to specify padding for top
// and bottom separately. Amount defaults to 0.
//
//   - [Side] side
//
// Specifies which side of the container the content will be packed against.
// Must be "left", "right", "top", or "bottom". Defaults to top.
//
// If no -in, -after or -before option is specified then each of the content
// will be inserted at the end of the packing list for its parent unless it is
// already managed by the packer (in which case it will be left where it is).
// If one of these options is specified then all the content will be inserted
// at the specified point. If any of the content are already managed by the
// geometry manager then any unspecified options for them retain their previous
// values rather than receiving default values.
//
// Additional information might be available at the [Tcl/Tk pack] page.
//
// [Tcl/Tk pack]: https://www.tcl.tk/man/tcl9.0/TkCmd/pack.htm
func Pack(options ...Opt) {
	evalErr(fmt.Sprintf("pack %s", collect(options...)))
}

// Wait — Wait for a window to be destroyed
//
// # Description
//
// Wait command waits for 'w' to be destroyed. This is typically used to wait
// for a user to finish interacting with a dialog box before using the result
// of that interaction.
//
// While the Wwait command is waiting it processes events in the normal
// fashion, so the application will continue to respond to user interactions.
// If an event handler invokes Wait again, the nested call to Wait must
// complete before the outer call can complete.
func (w *Window) Wait() {
	evalErr(fmt.Sprintf("tkwait window %s", w))
}

// WaitVisibility — Wait for a window to change visibility
//
// # Description
//
// WaitVisibility command waits for a change in w's visibility state (as
// indicated by the arrival of a VisibilityNotify event). This form is
// typically used to wait for a newly-created window to appear on the screen
// before taking some action.
//
// While the Wwait command is waiting it processes events in the normal
// fashion, so the application will continue to respond to user interactions.
// If an event handler invokes Wait again, the nested call to Wait must
// complete before the outer call can complete.
func (w *Window) WaitVisibility() {
	evalErr(fmt.Sprintf("tkwait visibility %s", w))
}

// IconPhoto — change window icon
//
// # Description
//
// IconPhoto sets the titlebar icon for window based on the named photo images.
// If -default is specified, this is applied to all future created toplevels as
// well. The data in the images is taken as a snapshot at the time of
// invocation. If the images are later changed, this is not reflected to the
// titlebar icons. Multiple images are accepted to allow different images sizes
// (e.g., 16x16 and 32x32) to be provided. The window manager may scale
// provided icons to an appropriate size.
//
// On Windows, the images are packed into a Windows icon structure. This will
// override an ico specified to wm iconbitmap, and vice versa. This command
// sets the taskbar icon for the window.
//
// On X, the images are arranged into the _NET_WM_ICON X property, which most
// modern window managers support. A wm iconbitmap may exist simultaneously. It is
// recommended to use not more than 2 icons, placing the larger icon first. This
// command also sets the panel icon for the application if the window manager or
// desktop environment supports it.
//
// On Macintosh, the first image called is loaded into an OSX-native icon
// format, and becomes the application icon in dialogs, the Dock, and other
// contexts. At the script level the command will accept only the first image
// passed in the parameters as support for multiple sizes/resolutions on macOS
// is outside Tk's scope. Developers should use the largest icon they can
// support (preferably 512 pixels) to ensure smooth rendering on the Mac.
//
// More information might be available at the [Tcl/Tk wm] page.
//
// [Tcl/Tk wm]: https://www.tcl.tk/man/tcl9.0/TkCmd/wm.htm#M50
func (w *Window) IconPhoto(options ...Opt) {
	evalErr(fmt.Sprintf("wm iconphoto %s %s", w, collect(options...)))
}

// WmTitle — change the window manager title
//
// # Description
//
// If string is specified, then it will be passed to the window manager for use
// as the title for window (the window manager should display this string in
// window's title bar). In this case the command returns an empty string. If
// string is not specified then the command returns the current title for the
// window. The title for a window defaults to its name.
//
// More information might be available at the [Tcl/Tk wm] page.
//
// [Tcl/Tk wm]: https://www.tcl.tk/man/tcl9.0/TkCmd/wm.htm#M63
func (w *Window) WmTitle(s string) string {
	if s != "" {
		s = tclSafeString(s)
	}
	return evalErr(fmt.Sprintf("wm title %s %s", w, s))
}

// Center centers 'w' and returns 'w'.
func (w *Window) Center() *Window {
	evalErr(fmt.Sprintf("tk::PlaceWindow %s center", w))
	return w
}

// Grid — Geometry manager that arranges widgets in a grid
//
// # Description
//
// The arguments consist of the names of one or more content windows followed
// by pairs of arguments that specify how to manage the content. The characters
// -, x and ^, can be specified instead of a window name to alter the default
// location of a window, as described in the RELATIVE PLACEMENT section, below.
//
// The following options are supported:
//
//   - [Column] n
//
// Insert the window so that it occupies the nth column in the grid. Column
// numbers start with 0. If this option is not supplied, then the window is
// arranged just to the right of previous window specified on this call to
// grid, or column “0” if it is the first window. For each x that immediately
// precedes the window, the column position is incremented by one. Thus the x
// represents a blank column for this row in the grid.
//
//   - [Columnspan] n
//
// Insert the window so that it occupies n columns in the grid. The default is
// one column, unless the window name is followed by a -, in which case the
// columnspan is incremented once for each immediately following -.
//
//   - [In] container
//
// Insert the window(s) in the container window given by container. The default
// is the first window's parent window.
//
//   - [Ipadx] amount
//
// The amount specifies how much horizontal internal padding to leave on each
// side of the content. This is space is added inside the content border. The
// amount must be a valid screen distance, such as 2 or .5c. It defaults to 0.
//
//   - [Ipady] amount
//
// The amount specifies how much vertical internal padding to leave on the top
// and bottom of the content. This space is added inside the content border.
// The amount defaults to 0.
//
//   - [Padx] amount
//
// The amount specifies how much horizontal external padding to leave on each
// side of the content, in screen units. Amount may be a list of two values to
// specify padding for left and right separately. The amount defaults to 0.
// This space is added outside the content border.
//
//   - [Pady] amount
//
// The amount specifies how much vertical external padding to leave on the top
// and bottom of the content, in screen units. Amount may be a list of two
// values to specify padding for top and bottom separately. The amount defaults
// to 0. This space is added outside the content border.
//
//   - [Row] n
//
// Insert the content so that it occupies the nth row in the grid. Row numbers
// start with 0. If this option is not supplied, then the content is arranged
// on the same row as the previous content specified on this call to grid, or
// the next row after the highest occupied row if this is the first content.
//
//   - [Rowspan] n
//
// Insert the content so that it occupies n rows in the grid. The default is
// one row. If the next grid command contains ^ characters instead of content
// that line up with the columns of this content, then the rowspan of this
// content is extended by one.
//
//   - [Sticky] style
//
// If a content's cell is larger than its requested dimensions, this option may
// be used to position (or stretch) the content within its cell. Style is a
// string that contains zero or more of the characters n, s, e or w. The string
// can optionally contain spaces or commas, but they are ignored. Each letter
// refers to a side (north, south, east, or west) that the content will “stick”
// to. If both n and s (or e and w) are specified, the content will be
// stretched to fill the entire height (or width) of its cavity. The -sticky
// option subsumes the combination of -anchor and -fill that is used by pack.
// The default is “”, which causes the content to be centered in its cavity, at
// its requested size.
//
// If any of the content is already managed by the geometry manager then any
// unspecified options for them retain their previous values rather than
// receiving default values.
//
// More information might be available at the [Tcl/Tk grid] page.
//
// [Tcl/Tk grid]: https://www.tcl.tk/man/tcl9.0/TkCmd/grid.html#M9
func Grid(w Widget, options ...Opt) {
	evalErr(fmt.Sprintf("grid configure %s %s", w, collect(options...)))
}

// Configure alters the configuration of 'w' and returns 'w'.
func (w *Window) Configure(options ...Opt) *Window {
	options, tvs := w.split(options)
	if len(options) != 0 {
		evalErr(fmt.Sprintf("%s configure %s", w, collect(options...)))
	}
	if len(tvs) != 0 {
		tvo := tvs[len(tvs)-1]
		tclVar := textVariables[w]
		if tclVar == "" {
			tclVar = fmt.Sprintf("textVar%d", id.Add(1))
			textVariables[w] = tclVar
			evalErr(fmt.Sprintf("%s configure -textvariable %s", w, tclVar))
		}
		evalErr(fmt.Sprintf("set %s %s", tclVar, tclSafeString(string(tvo))))
	}
	return w
}

// tk_messageBox — pops up a message window and waits for user response.
//
// # Description
//
// This procedure creates and displays a message window with an
// application-specified message, an icon and a set of buttons. Each of the
// buttons in the message window is identified by a unique symbolic name (see
// the -type options). After the message window is popped up, tk_messageBox
// waits for the user to select one of the buttons. Then it returns the
// symbolic name of the selected button.  The following optins are
// supported:
//
//   - [Command] handler
//
// Specifies the handler to invoke when the user closes the
// dialog. The actual command consists of string followed by a space and the
// name of the button clicked by the user to close the dialog. This is only
// available on Mac OS X.
//
//   - [Default] name
//
// Name gives the symbolic name of the default button for this message window (
// “ok”, “cancel”, and so on). See -type for a list of the symbolic names. If
// this option is not specified, the first button in the dialog will be made
// the default.
//
//   - [Detail] string
//
// Specifies an auxiliary message to the main message given by the -message
// option. The message detail will be presented beneath the main message and,
// where supported by the OS, in a less emphasized font than the main message.
//
//   - [Icon] iconImage
//
// Specifies an icon to display. IconImage must be one of the following: error,
// info, question or warning. If this option is not specified, then the info
// icon will be displayed.
//
//   - [Message] string
//
// Specifies the message to display in this message box. The default value is
// an empty string.
//
//   - [Parent] window
//
// Makes window the logical parent of the message box. The message box is
// displayed on top of its parent window.
//
//   - [Title] titleString
//
// Specifies a string to display as the title of the message box. The default
// value is an empty string.
//
//   - [Type] predefinedType
//
// Arranges for a predefined set of buttons to be displayed. The following
// values are possible for predefinedType:
//
//   - abortretryignore - Displays three buttons whose symbolic names are abort, retry and ignore.
//   - ok - Displays one button whose symbolic name is ok.
//   - okcancel - Displays two buttons whose symbolic names are ok and cancel.
//   - retrycancel - Displays two buttons whose symbolic names are retry and cancel.
//   - yesno - Displays two buttons whose symbolic names are yes and no.
//   - yesnocancel - Displays three buttons whose symbolic names are yes, no and cancel.
//
// More information might be available at the [Tcl/Tk messageBox] page.
//
// [Tcl/Tk messageBox]: https://www.tcl.tk/man/tcl9.0/TkCmd/messageBox.html
func MessageBox(options ...Opt) string {
	return evalErr(fmt.Sprintf("tk_messageBox %s", collect(options...)))
}

// Bell — Ring a display's bell
//
// # Description
//
// This command rings the bell on the display for window and returns an empty
// string. If the -displayof option is omitted, the display of the
// application's main window is used by default. The command uses the current
// bell-related settings for the display, which may be modified with programs
// such as xset.
//
// If -nice is not specified, this command also resets the screen saver for the
// screen. Some screen savers will ignore this, but others will reset so that
// the screen becomes visible again.
//
//   - [Displayof] window
//   - [Nice]
//
// More information might be available at the [Tcl/Tk bell] page.
//
// [Tcl/Tk bell]: https://www.tcl.tk/man/tcl9.0/TclCmd/bell.htm
func Bell(options ...Opt) {
	evalErr(fmt.Sprintf("bell %s", collect(options...)))
}

// ChooseColor — pops up a dialog box for the user to select a color.
//
// # Description
//
// ChooseColor pops up a dialog box for the user to select a color. The
// following option-value pairs are possible as command line arguments:
//
//   - [Initialcolor] color
//
// Specifies the color to display in the color dialog when it pops up. color
// must be in a form acceptable to the Tk_GetColor function.
//
//   - [Parent] window
//
// Makes window the logical parent of the color dialog. The color dialog is
// displayed on top of its parent window.
//
//   - [Title] titleString
//
// Specifies a string to display as the title of the dialog box. If this option
// is not specified, then a default title will be displayed.
//
// If the user selects a color, ChooseColor will return the name of the
// color in a form acceptable to Tk_GetColor. If the user cancels the
// operation, both commands will return the empty string.
//
// More information might be available at the [Tcl/Tk choosecolor] page.
//
// [Tcl/Tk choosecolor]: https://www.tcl.tk/man/tcl9.0/TclCmd/chooseColor.htm
func ChooseColor(options ...Opt) string {
	return evalErr(fmt.Sprintf("tk_chooseColor %s", collect(options...)))
}

// Busy — confine pointer events to a window sub-tree
//
// # Description
//
// The Busy command provides a simple means to block pointer events from Tk
// widgets, while overriding the widget's cursor with a configurable busy
// cursor. Note this command does not prevent keyboard events from being sent
// to the widgets made busy.
//
//   - [Cursor] cursorName
//
// Specifies the cursor to be displayed when the widget is made busy.
// CursorName can be in any form accepted by Tk_GetCursor. The default cursor
// is wait on Windows and watch on other platforms.
//
// More information might be available at the [Tcl/Tk busy] page.
//
// [Tcl/Tk update]: https://www.tcl.tk/man/tcl9.0/TclCmd/busy.htm
func (w *Window) Busy(options ...Opt) {
	evalErr(fmt.Sprintf("tk busy %s %s", w, collect(options...)))
}

// BusyForget — undo Busy
//
// # Description
//
// Releases resources allocated by the [Window.Busy] command for window, including
// the transparent window. User events will again be received by window.
// Resources are also released when window is destroyed. Window must be the
// name of a widget specified in the hold operation, otherwise an error is
// reported.
//
// More information might be available at the [Tcl/Tk busy] page.
//
// [Tcl/Tk update]: https://www.tcl.tk/man/tcl9.0/TclCmd/busy.htm
func (w *Window) BusyForget(options ...Opt) {
	evalErr(fmt.Sprintf("tk busy forget %s %s", w, collect(options...)))
}

// Update — Process pending events and idle callbacks
//
// More information might be available at the [Tcl/Tk update] page.
//
// [Tcl/Tk update]: https://www.tcl.tk/man/tcl9.0/TclCmd/update.htm
func Update() {
	evalErr("update")
}

//TODO?
//  - [Command] string
//
// Specifies the prefix of a Tcl command to invoke when the user closes the
// dialog after having selected an item. This callback is not called if the
// user cancelled the dialog. The actual command consists of string followed by
// a space and the value selected by the user in the dialog. This is only
// available on Mac OS X.

// ChooseDirectory — pops up a dialog box for the user to select a directory.
//
// # Description
//
// The procedure tk_chooseDirectory pops up a dialog box for the user to select
// a directory. The following option-value pairs are possible as command line
// arguments:
//
//   - [Initialdir] dirname
//
// Specifies that the directories in directory should be displayed when the
// dialog pops up. If this parameter is not specified, the initial directory
// defaults to the current working directory on non-Windows systems and on
// Windows systems prior to Vista. On Vista and later systems, the initial
// directory defaults to the last user-selected directory for the application.
// If the parameter specifies a relative path, the return value will convert
// the relative path to an absolute path.
//
//   - [Message] string
//
// Specifies a message to include in the client area of the dialog. This is
// only available on Mac OS X.
//
//   - [Mustexist] boolean
//
// Specifies whether the user may specify non-existent directories. If this
// parameter is true, then the user may only select directories that already
// exist. The default value is false.
//
//   - [Parent] window
//
// Makes window the logical parent of the dialog. The dialog is displayed on
// top of its parent window. On Mac OS X, this turns the file dialog into a
// sheet attached to the parent window.
//
//   - [Title] titleString
//
// Specifies a string to display as the title of the dialog box. If this option
// is not specified, then a default title will be displayed.
//
// More information might be available at the [Tcl/Tk chooseDirectory] page.
//
// [Tcl/Tk chooseDirectory]: https://www.tcl.tk/man/tcl9.0/TkCmd/chooseDirectory.htm
func ChooseDirectory(options ...Opt) string {
	return evalErr(fmt.Sprintf("tk_chooseDirectory %s", collect(options...)))
}

// ClipboardAppend — Manipulate Tk clipboard
//
// # Description
//
// This command provides a Tcl interface to the Tk clipboard, which stores data
// for later retrieval using the selection mechanism (via the -selection
// CLIPBOARD option). In order to copy data into the clipboard, clipboard clear
// must be called, followed by a sequence of one or more calls to clipboard
// append. To ensure that the clipboard is updated atomically, all appends
// should be completed before returning to the event loop.
//
// ClipboardAppend appends 'data' to the clipboard on window's display in the
// form given by type with the representation given by format and claims
// ownership of the clipboard on window's display.
//
//   - [Displayof] window
//
//   - [Format] format
//
// The format argument specifies the representation that should be used to
// transmit the selection to the requester (the second column of Table 2 of the
// ICCCM), and defaults to STRING. If format is STRING, the selection is
// transmitted as 8-bit ASCII characters. If format is ATOM, then the data is
// divided into fields separated by white space; each field is converted to its
// atom value, and the 32-bit atom value is transmitted instead of the atom
// name. For any other format, data is divided into fields separated by white
// space and each field is converted to a 32-bit integer; an array of integers
// is transmitted to the selection requester. Note that strings passed to
// clipboard append are concatenated before conversion, so the caller must take
// care to ensure appropriate spacing across string boundaries. All items
// appended to the clipboard with the same type must have the same format.
//
// The format argument is needed only for compatibility with clipboard
// requesters that do not use Tk. If the Tk toolkit is being used to retrieve
// the CLIPBOARD selection then the value is converted back to a string at the
// requesting end, so format is irrelevant.
//
//   - [Type] type
//
// Type specifies the form in which the selection is to be returned (the
// desired “target” for conversion, in ICCCM terminology), and should be an
// atom name such as STRING or FILE_NAME; see the Inter-Client Communication
// Conventions Manual for complete details. Type defaults to STRING.
//
// More information might be available at the [Tcl/Tk clipboard] page.
//
// [Tcl/Tk clipboard]: https://www.tcl.tk/man/tcl9.0/TkCmd/clipboard.htm
func ClipboardAppend(data string, options ...Opt) {
	evalErr(fmt.Sprintf("clipboard append %s -- %s", collect(options...), tclSafeString(data)))
}

// ClipboardClear — Manipulate Tk clipboard
//
// # Description
//
// Claims ownership of the clipboard on window's display and removes any
// previous contents. Window defaults to App. Returns an empty string.
//
// More information might be available at the [Tcl/Tk clipboard] page.
//
// [Tcl/Tk clipboard]: https://www.tcl.tk/man/tcl9.0/TkCmd/clipboard.htm
func ClipboardClear(options ...Opt) {
	evalErr(fmt.Sprintf("clipboard clear %s", collect(options...)))
}

// ClipboardGet — Manipulate Tk clipboard
//
// # Description
//
// Retrieve data from the clipboard on window's display. Window defaults to App.
//
//   - [Displayof] window
//
//   - [Type] type
//
// Type specifies the form in which the data is to be returned and should be an
// atom name such as STRING or FILE_NAME. Type defaults to STRING. This command
// is equivalent to [SelectionGet](Selection("CLIPBOARD").
//
// Note that on modern X11 systems, the most useful type to retrieve for
// transferred strings is not STRING, but rather UTF8_STRING.
//
// More information might be available at the [Tcl/Tk clipboard] page.
//
// [Tcl/Tk clipboard]: https://www.tcl.tk/man/tcl9.0/TkCmd/clipboard.htm
func ClipboardGet(options ...Opt) string {
	return evalErr(fmt.Sprintf("clipboard get %s", collect(options...)))
}

// ExitHandler returns a canned [Command] that destroys the [App].
func ExitHandler() Opt {
	return exitHandler
}

// Exit provides a canned [Button] with default [Txt] "Exit", bound to the
// [ExitHandler].
//
// Use [Window.Exit] to create a Exit with a particular parent.
func Exit(options ...Opt) *ButtonWidget {
	return App.Exit(options...)
}

// Exit provides a canned [Button] with default [Txt] "Exit", bound to the
// [ExitHandler].
//
// The resulting [Window] is a child of 'w'
func (w *Window) Exit(options ...Opt) *ButtonWidget {
	return w.Button(append([]Opt{Txt("Exit"), ExitHandler()}, options...)...)
}

// TExit provides a canned [TButton] with default [Txt] "Exit", bound to the
// [ExitHandler].
//
// Use [Window.TExit] to create a TExit with a particular parent.
func TExit(options ...Opt) *TButtonWidget {
	return App.TExit(options...)
}

// TExit provides a canned [TButton] with default [Txt] "Exit", bound to the
// [ExitHandler].
//
// The resulting [Window] is a child of 'w'
func (w *Window) TExit(options ...Opt) *TButtonWidget {
	return w.TButton(append([]Opt{Txt("Exit"), ExitHandler()}, options...)...)
}

type textVarOpt string

func (textVarOpt) optionString(*Window) string {
	panic("internal error") // Not supposed to be invoked.
}

// Textvariable option.
//
// Specifies the value to be displayed inside the widget.
// The way in which the string is displayed in the widget depends on the
// particular widget and may be determined by other options, such as
// -anchor or -justify.
//
// Known uses:
//   - [Button]
//   - [Checkbutton]
//   - [Entry]
//   - [Label]
//   - [Menubutton]
//   - [Message]
//   - [Radiobutton]
//   - [Spinbox]
//   - [TButton]
//   - [TCheckbutton]
//   - [TCombobox] (widget specific)
//   - [TEntry] (widget specific)
//   - [TLabel]
//   - [TMenubutton]
//   - [TRadiobutton]
func Textvariable(s string) Opt {
	return textVarOpt(s)
}

// Textvariable — Get the configured option value.
//
// Known uses:
//   - [Button]
//   - [Checkbutton]
//   - [Entry]
//   - [Label]
//   - [Menubutton]
//   - [Message]
//   - [Radiobutton]
//   - [Spinbox]
//   - [TButton]
//   - [TCheckbutton]
//   - [TCombobox] (widget specific)
//   - [TEntry] (widget specific)
//   - [TLabel]
//   - [TMenubutton]
//   - [TRadiobutton]
func (w *Window) Textvariable() (r string) {
	if tclVar := textVariables[w]; tclVar != "" {
		return evalErr(fmt.Sprintf("set %s", tclVar))
	}

	return ""
}

// Focus — Manage the input focus
//
// # Description
//
// The focus command is used to manage the Tk input focus. At any given time,
// one window on each display is designated as the focus window; any key press
// or key release events for the display are sent to that window. It is
// normally up to the window manager to redirect the focus among the top-level
// windows of a display. For example, some window managers automatically set
// the input focus to a top-level window whenever the mouse enters it; others
// redirect the input focus only when the user clicks on a window. Usually the
// window manager will set the focus only to top-level windows, leaving it up
// to the application to redirect the focus among the children of the
// top-level.
//
// Tk remembers one focus window for each top-level (the most recent descendant
// of that top-level to receive the focus); when the window manager gives the
// focus to a top-level, Tk automatically redirects it to the remembered
// window. Within a top-level Tk uses an explicit focus model by default.
// Moving the mouse within a top-level does not normally change the focus; the
// focus changes only when a widget decides explicitly to claim the focus
// (e.g., because of a button click), or when the user types a key such as Tab
// that moves the focus.
//
// The Tcl procedure tk_focusFollowsMouse may be invoked to create an implicit
// focus model: it reconfigures Tk so that the focus is set to a window
// whenever the mouse enters it. The Tcl procedures tk_focusNext and
// tk_focusPrev implement a focus order among the windows of a top-level; they
// are used in the default bindings for Tab and Shift-Tab, among other things.
//
// The focus command can take any of the following forms:
//
//	Focus()
//
// Returns the path name of the focus window on the display containing the
// application's main window, or an empty string if no window in this
// application has the focus on that display. Note: it is better to specify the
// display explicitly using -displayof (see below) so that the code will work
// in applications using multiple displays.
//
//	Focus(window)
//
// If the application currently has the input focus on window's display, this
// command resets the input focus for window's display to window and returns an
// empty string. If the application does not currently have the input focus on
// window's display, window will be remembered as the focus for its top-level;
// the next time the focus arrives at the top-level, Tk will redirect it to
// window. If window is an empty string then the command does nothing.
//
//	Focus(Displayof(window))
//
// Returns the name of the focus window on the display containing window. If
// the focus window for window's display is not in this application, the return
// value is an empty string.
//
//	Focus(Force(window))
//
// Sets the focus of window's display to window, even if the application does
// not currently have the input focus for the display. This command should be
// used sparingly, if at all. In normal usage, an application should not claim
// the focus for itself; instead, it should wait for the window manager to give
// it the focus. If window is an empty string then the command does nothing.
//
//	Focus(Lastfor(window))
//
// Returns the name of the most recent window to have the input focus among all
// the windows in the same top-level as window. If no window in that top-level
// has ever had the input focus, or if the most recent focus window has been
// deleted, then the name of the top-level is returned. The return value is the
// window that will receive the input focus the next time the window manager
// gives the focus to the top-level.
//
// # Quirks
//
// When an internal window receives the input focus, Tk does not actually set
// the X focus to that window; as far as X is concerned, the focus will stay on
// the top-level window containing the window with the focus. However, Tk
// generates FocusIn and FocusOut events just as if the X focus were on the
// internal window. This approach gets around a number of problems that would
// occur if the X focus were actually moved; the fact that the X focus is on
// the top-level is invisible unless you use C code to query the X server
// directly.
//
// More information might be available at the [Tcl/Tk focus] page.
//
// [Tcl/Tk focus]: https://www.tcl.tk/man/tcl9.0/TkCmd/focus.htm
func Focus(options ...Opt) string {
	return evalErr(fmt.Sprintf("focus %s", collect(options...)))
}

// Font represents a Tk font.
type Font struct {
	name string
}

func (f *Font) optionString(_ *Window) (r string) {
	if f != nil {
		return f.name
	}

	return "font0" // does not exist
}

// String implements fmt.Stringer.
func (f *Font) String() string {
	return f.optionString(nil)
}

// NewFont — Create and inspect fonts.
//
// # Description
//
// Creates a new font.
//
// The following options are supported on all platforms, and are used when
// creating/specifying a font:
//
//   - [Family] name
//
// The case-insensitive font family name. Tk guarantees to support the font
// families named Courier (a monospaced “typewriter” font), Times (a serifed
// “newspaper” font), and Helvetica (a sans-serif “European” font). The most
// closely matching native font family will automatically be substituted when
// one of the above font families is used. The name may also be the name of a
// native, platform-specific font family; in that case it will work as desired
// on one platform but may not display correctly on other platforms. If the
// family is unspecified or unrecognized, a platform-specific default font will
// be chosen.
//
//   - [Size] size
//
// The desired size of the font. If the size argument is a positive number, it
// is interpreted as a size in points. If size is a negative number, its
// absolute value is interpreted as a size in pixels. If a font cannot be
// displayed at the specified size, a nearby size will be chosen. If size is
// unspecified or zero, a platform-dependent default size will be chosen.
//
// Sizes should normally be specified in points so the application will remain
// the same ruler size on the screen, even when changing screen resolutions or
// moving scripts across platforms. However, specifying pixels is useful in
// certain circumstances such as when a piece of text must line up with respect
// to a fixed-size bitmap. The mapping between points and pixels is set when
// the application starts, based on properties of the installed monitor, but it
// can be overridden by calling the tk scaling command.
//
//   - [Weight] weight
//
// The nominal thickness of the characters in the font. The value normal
// specifies a normal weight font, while bold specifies a bold font. The
// closest available weight to the one specified will be chosen. The default
// weight is normal.
//
//   - [Slant] slant
//
// The amount the characters in the font are slanted away from the vertical.
// Valid values for slant are roman and italic. A roman font is the normal,
// upright appearance of a font, while an italic font is one that is tilted
// some number of degrees from upright. The closest available slant to the one
// specified will be chosen. The default slant is roman.
//
//   - [Underline] boolean
//
// The value is a boolean flag that specifies whether characters in this font
// should be underlined. The default value for underline is false.
//
//   - [Overstrike] boolean
//
// The value is a boolean flag that specifies whether a horizontal line should
// be drawn through the middle of characters in this font. The default value
// for overstrike is false.
//
// Additional information might be available at the [Tcl/Tk font] page.
//
// [Tcl/Tk font]: https://www.tcl.tk/man/tcl9.0/TkCmd/font.htm
func NewFont(options ...Opt) *Font {
	nm := fmt.Sprintf("font%v", id.Add(1))
	code := fmt.Sprintf("font create %s %s", nm, collect(options...))
	r, err := eval(code)
	if err != nil {
		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
		return nil
	}

	return &Font{name: nm}
}

// FontFamilies — Create and inspect fonts.
//
// # Description
//
// The return value is a list of the case-insensitive names of all font
// families that exist on window's display. If the Displayof argument is
// omitted, it defaults to the main window.
//
//   - [Displayof] window
//
// Additional information might be available at the [Tcl/Tk font] page.
//
// [Tcl/Tk font]: https://www.tcl.tk/man/tcl9.0/TkCmd/font.htm
func FontFamilies(options ...Opt) []string {
	return parseList(evalErr(fmt.Sprintf("font families %s", collect(options...))))
}

// Can handle "{abc def} ghi" -> ["abc def", "ghi"].
// Can not handle "{abc {def ghi}} jkl", too many levels.
func parseList(s string) (r []string) {
	r = strings.Fields(s)
	w := 0
	in := false
	var a []string
	for _, v := range r {
		switch {
		case in:
			switch {
			case strings.HasSuffix(v, "}"):
				a = append(a, v[:len(v)-1])
				r[w] = strings.Join(a, " ")
				w++
				in = false
			default:
				a = append(a, v)
			}
		default:
			switch {
			case strings.HasPrefix(v, "{"):
				a = append(a[:0], v[1:])
				in = true
			default:
				r[w] = v
				w++
			}
		}
	}
	return r[:w]
}

// Delete — Manipulate fonts.
//
// # Description
//
// Delete the font. If there are widgets using the named font, the named font
// will not actually be deleted until all the instances are released. Those
// widgets will continue to display using the last known values for the named
// font. If a deleted named font is subsequently recreated with another call to
// font create, the widgets will use the new named font and redisplay
// themselves using the new attributes of that font.
//
// Additional information might be available at the [Tcl/Tk font] page.
//
// [Tcl/Tk font]: https://www.tcl.tk/man/tcl9.0/TkCmd/font.htm
func (f *Font) Delete() {
	evalErr(fmt.Sprintf("font delete %s", f))
}

// Text — Create and manipulate 'text' hypertext editing widgets
//
// # Description
//
// Inserts all of the chars arguments just before the character at index. If
// index refers to the end of the text (the character after the last newline)
// then the new text is inserted just before the last newline instead. If there
// is a single chars argument and no tagList, then the new text will receive
// any tags that are present on both the character before and the character
// after the insertion point; if a tag is present on only one of these
// characters then it will not be applied to the new text. If tagList is
// specified then it consists of a list of tag names; the new characters will
// receive all of the tags in this list and no others, regardless of the tags
// present around the insertion point. If multiple chars-tagList argument pairs
// are present, they produce the same effect as if a separate pathName insert
// widget command had been issued for each pair, in order. The last tagList
// argument may be omitted.
//
// The value that is passed to Tcl/Tk for the 'index' argument is obtained by
// fmt.Sprint(index), enabling any custom index encoding via implementing
// fmt.Stringer.
//
// # Indices
//
// Many of the widget commands for texts take one or more indices as arguments.
// An index is a string used to indicate a particular place within a text, such
// as a place to insert characters or one endpoint of a range of characters to
// delete. Indices have the syntax
//
//	base modifier modifier modifier ...
//
// Where base gives a starting point and the modifiers adjust the index from
// the starting point (e.g. move forward or backward one character). Every
// index must contain a base, but the modifiers are optional. Most modifiers
// (as documented below) allow an optional submodifier. Valid submodifiers are
// any and display. If the submodifier is abbreviated, then it must be followed
// by whitespace, but otherwise there need be no space between the submodifier
// and the following modifier. Typically the display submodifier adjusts the
// meaning of the following modifier to make it refer to visual or non-elided
// units rather than logical units, but this is explained for each relevant
// case below. Lastly, where count is used as part of a modifier, it can be
// positive or negative, so “base - -3 lines” is perfectly valid (and
// equivalent to “base +3lines”).
//
// The base for an index must have one of the following forms:
//
//	"line.char"
//
// Indicates char'th character on line line. Lines are numbered from 1 for
// consistency with other UNIX programs that use this numbering scheme. Within
// a line, characters are numbered from 0. If char is end then it refers to the
// newline character that ends the line.
//
// This form of index can be passed as [LC]{line, char}.
//
//	"@x,y"
//
// Indicates the character that covers the pixel whose x and y coordinates
// within the text's window are x and y.
//
//	"end"
//
// Indicates the end of the text (the character just after the last newline).
//
//	"mark"
//
// Indicates the character just after the mark whose name is mark (see MARKS
// for details).
//
//	"tag.first"
//
// Indicates the first character in the text that has been tagged with tag.
// This form generates an error if no characters are currently tagged with tag.
//
//	"tag.last"
//
// Indicates the character just after the last one in the text that has been
// tagged with tag. This form generates an error if no characters are currently
// tagged with tag.
//
//	"pathName"
//
// Indicates the position of the embedded window whose name is pathName. This
// form generates an error if there is no embedded window by the given name.
//
//	"imageName"
//
// Indicates the position of the embedded image whose name is imageName. This
// form generates an error if there is no embedded image by the given name.
//
// If the base could match more than one of the above forms, such as a mark and
// imageName both having the same value, then the form earlier in the above
// list takes precedence. If modifiers follow the base index, each one of them
// must have one of the forms listed below. Keywords such as chars and wordend
// may be abbreviated as long as the abbreviation is unambiguous.
//
//	"+ count ?submodifier? chars"
//
// Adjust the index forward by count characters, moving to later lines in the
// text if necessary. If there are fewer than count characters in the text
// after the current index, then set the index to the last index in the text.
// Spaces on either side of count are optional. If the display submodifier is
// given, elided characters are skipped over without being counted. If any is
// given, then all characters are counted. For historical reasons, if neither
// modifier is given then the count actually takes place in units of index
// positions (see INDICES for details). This behaviour may be changed in a
// future major release, so if you need an index count, you are encouraged to
// use indices instead wherever possible.
//
//	"- count ?submodifier? chars"
//
// Adjust the index backward by count characters, moving to earlier lines in
// the text if necessary. If there are fewer than count characters in the text
// before the current index, then set the index to the first index in the text
// (1.0). Spaces on either side of count are optional. If the display
// submodifier is given, elided characters are skipped over without being
// counted. If any is given, then all characters are counted. For historical
// reasons, if neither modifier is given then the count actually takes place in
// units of index positions (see INDICES for details). This behavior may be
// changed in a future major release, so if you need an index count, you are
// encouraged to use indices instead wherever possible.
//
//	"+ count ?submodifier? indices"
//
// Adjust the index forward by count index positions, moving to later lines in
// the text if necessary. If there are fewer than count index positions in the
// text after the current index, then set the index to the last index position
// in the text. Spaces on either side of count are optional. Note that an index
// position is either a single character or a single embedded image or embedded
// window. If the display submodifier is given, elided indices are skipped over
// without being counted. If any is given, then all indices are counted; this
// is also the default behaviour if no modifier is given.
//
//	"- count ?submodifier? indices"
//
// Adjust the index backward by count index positions, moving to earlier lines
// in the text if necessary. If there are fewer than count index positions in
// the text before the current index, then set the index to the first index
// position (1.0) in the text. Spaces on either side of count are optional. If
// the display submodifier is given, elided indices are skipped over without
// being counted. If any is given, then all indices are counted; this is also
// the default behaviour if no modifier is given.
//
//	"+ count ?submodifier? lines"
//
// Adjust the index forward by count lines, retaining the same character
// position within the line. If there are fewer than count lines after the line
// containing the current index, then set the index to refer to the same
// character position on the last line of the text. Then, if the line is not
// long enough to contain a character at the indicated character position,
// adjust the character position to refer to the last character of the line
// (the newline). Spaces on either side of count are optional. If the display
// submodifier is given, then each visual display line is counted separately.
// Otherwise, if any (or no modifier) is given, then each logical line (no
// matter how many times it is visually wrapped) counts just once. If the
// relevant lines are not wrapped, then these two methods of counting are
// equivalent.
//
//	"- count ?submodifier? lines"
//
// Adjust the index backward by count logical lines, retaining the same
// character position within the line. If there are fewer than count lines
// before the line containing the current index, then set the index to refer to
// the same character position on the first line of the text. Then, if the line
// is not long enough to contain a character at the indicated character
// position, adjust the character position to refer to the last character of
// the line (the newline). Spaces on either side of count are optional. If the
// display submodifier is given, then each visual display line is counted
// separately. Otherwise, if any (or no modifier) is given, then each logical
// line (no matter how many times it is visually wrapped) counts just once. If
// the relevant lines are not wrapped, then these two methods of counting are
// equivalent.
//
//	"?submodifier? linestart"
//
// Adjust the index to refer to the first index on the line. If the display
// submodifier is given, this is the first index on the display line, otherwise
// on the logical line.
//
//	"?submodifier? lineend"
//
// Adjust the index to refer to the last index on the line (the newline). If
// the display submodifier is given, this is the last index on the display
// line, otherwise on the logical line.
//
//	"?submodifier? wordstart"
//
// Adjust the index to refer to the first character of the word containing the
// current index. A word consists of any number of adjacent characters that are
// letters, digits, or underscores, or a single character that is not one of
// these. If the display submodifier is given, this only examines non-elided
// characters, otherwise all characters (elided or not) are examined.
//
//	"?submodifier? wordend"
//
// Adjust the index to refer to the character just after the last one of the
// word containing the current index. If the current index refers to the last
// character of the text then it is not modified. If the display submodifier is
// given, this only examines non-elided characters, otherwise all characters
// (elided or not) are examined.
//
// If more than one modifier is present then they are applied in left-to-right
// order. For example, the index “end - 1 chars” refers to the next-to-last
// character in the text and “insert wordstart - 1 c” refers to the character
// just before the first one in the word containing the insertion cursor.
// Modifiers are applied one by one in this left to right order, and after each
// step the resulting index is constrained to be a valid index in the text
// widget. So, for example, the index “1.0 -1c +1c” refers to the index “2.0”.
//
// Where modifiers result in index changes by display lines, display chars or
// display indices, and the base refers to an index inside an elided tag, that
// base index is considered to be equivalent to the first following non-elided
// index.
//
// Insert returns its index argument.
//
// Additional information might be available at the [Tcl/Tk text] page.
//
// [Tcl/Tk text]: https://www.tcl.tk/man/tcl9.0/TkCmd/text.htm
func (w *TextWidget) Insert(index any, chars string, options ...string) any {
	idx := fmt.Sprint(index)
	evalErr(fmt.Sprintf("%s insert %s %s %s", w, tclSafeString(idx), tclSafeString(chars), tclSafeStrings(options...)))
	return index
}

// LC encodes a text index consisting of a line and char number.
type LC struct {
	Line int // 1-based line number within the text content.
	Char int // 0-based char number within the line.
}

// String implements fmt.Stringer.
func (lc LC) String() string {
	return fmt.Sprintf("%d.%d", lc.Line, lc.Char)
}

// Text — Create and manipulate 'text' hypertext editing widgets
//
// # Description
//
// This command is similar to the pathName configure widget command except that
// it modifies options associated with the tag given by tagName instead of
// modifying options for the overall text widget. If no option is specified,
// the command returns a list describing all of the available options for
// tagName (see Tk_ConfigureInfo for information on the format of this list).
// If option is specified with no value, then the command returns a list
// describing the one named option (this list will be identical to the
// corresponding sublist of the value returned if no option is specified). If
// one or more option-value pairs are specified, then the command modifies the
// given option(s) to have the given value(s) in tagName; in this case the
// command returns an empty string. See TAGS above for details on the options
// available for tags.
//
// Additional information might be available at the [Tcl/Tk text] page.
//
// [Tcl/Tk text]: https://www.tcl.tk/man/tcl9.0/TkCmd/text.htm
func (w *TextWidget) TagConfigure(name string, options ...Opt) {
	evalErr(fmt.Sprintf("%s tag configure %s %s", w, tclSafeString(name), collect(options...)))
}

// Text — Create and manipulate 'text' hypertext editing widgets
//
// # Description
//
// Associate the tag name with all of the characters starting with index1
// and ending just before index2 (the character at index2 is not tagged). A
// single command may contain any number of index1-index2 pairs. If the last
// index2 is omitted then the single character at index1 is tagged. If there
// are no characters in the specified range (e.g. index1 is past the end of the
// file or index2 is less than or equal to index1) then the command has no
// effect.
//
// The function returns the new tag name.
//
// Additional information might be available at the [Tcl/Tk text] page.
//
// [Tcl/Tk text]: https://www.tcl.tk/man/tcl9.0/TkCmd/text.htm
func (w *TextWidget) TagAdd(options ...any) string {
	tag := fmt.Sprintf("tag%d", id.Add(1))
	var a []Opt
	for _, v := range options {
		switch x := v.(type) {
		case Opt:
			a = append(a, x)
		case Opts:
			a = append(a, x)
		default:
			a = append(a, stringOption(fmt.Sprint(x)))
		}
	}
	evalErr(fmt.Sprintf("%s tag add %s %s", w, tag, collect(a...)))
	return tag
}

// Fontchooser — control font selection dialog
//
// # Description
//
// The tk fontchooser command controls the Tk font selection dialog. It uses
// the native platform font selection dialog where available, or a dialog
// implemented in Tcl otherwise.
//
// Unlike most of the other Tk dialog commands, tk fontchooser does not return
// an immediate result, as on some platforms (Mac OS X) the standard font
// dialog is modeless while on others (Windows) it is modal. To get the
// user-selected font use FontchooserFont() from a handler assigned via
// [Command].
//
// Set one or more of the configurations options below (analogous to Tk widget configuration).
//
//   - [Parent]
//
// Specifies/returns the logical parent window of the font selection dialog
// (similar to the -parent option to other dialogs). The font selection dialog
// is hidden if it is visible when the parent window is destroyed.
//
//   - [Title]
//
// Specifies/returns the title of the dialog. Has no effect on platforms where
// the font selection dialog does not support titles.
//
//   - [Font]
//
// Specifies/returns the font that is currently selected in the dialog if it is
// visible, or that will be initially selected when the dialog is shown (if
// supported by the platform). Can be set to the empty string to indicate that
// no font should be selected. Fonts can be specified in any form given by the
// "FONT DESCRIPTION" section in the font manual page.
//
//   - [Command]
//
// Specifies the command called when a font selection has been made by the
// user. To obtain the font description, call [FontchooserFont] from the
// handler.
//
// Additional information might be available at the [Tcl/Tk fontchooser] page.
//
// [Tcl/Tk fontchooser]: https://www.tcl.tk/man/tcl9.0/TkCmd/fontchooser.htm
func Fontchooser(options ...Opt) {
	evalErr(fmt.Sprintf("tk fontchooser configure %s", collect(options...)))
}

// FontchooserFont — control font selection dialog
//
// # Description
//
// Returns the selected font description in the form
//
//	family size style...
//
// Additional information might be available at the [Tcl/Tk fontchooser] page.
//
// [Tcl/Tk fontchooser]: https://www.tcl.tk/man/tcl9.0/TkCmd/fontchooser.htm
func FontchooserFont() []string {
	return parseList(evalErr("tk fontchooser config -font"))
}

// FontchooserShow — control font selection dialog
//
// # Description
//
// Show the font selection dialog. Depending on the platform, may return
// immediately or only once the dialog has been withdrawn.
//
// Additional information might be available at the [Tcl/Tk fontchooser] page.
//
// [Tcl/Tk fontchooser]: https://www.tcl.tk/man/tcl9.0/TkCmd/fontchooser.htm
func FontchooserShow() {
	evalErr("tk  fontchooser show")
}

// FontchooserHide — control font selection dialog
//
// # Description
//
// Hide the font selection dialog if it is visible and cause any pending tk
// fontchooser show command to return.
//
// Additional information might be available at the [Tcl/Tk fontchooser] page.
//
// [Tcl/Tk fontchooser]: https://www.tcl.tk/man/tcl9.0/TkCmd/fontchooser.htm
func FontchooserHide() {
	evalErr("tk fontchooser hide")
}

// GetOpenFile — pop up a dialog box for the user to select a file to open.
//
// # Description
//
// GetOpenFile pops up a dialog box for the user to select a file to open. The
// function is usually associated with the Open command in the File menu. Its
// purpose is for the user to select an existing file only. If the user enters
// a non-existent file, the dialog box gives the user an error prompt and
// requires the user to give an alternative selection. If an application allows
// the user to create new files, it should do so by providing a separate New
// menu command.
//
//   - [Defaultextension] extension
//
// Specifies a string that will be appended to the filename if the user enters
// a filename without an extension. The default value is the empty string,
// which means no extension will be appended to the filename in any case. This
// option is ignored on Mac OS X, which does not require extensions to
// filenames, and the UNIX implementation guesses reasonable values for this
// from the -filetypes option when this is not supplied.
//
//   - [Filetypes] filePatternList ([][FileType])
//
// If a File types listbox exists in the file dialog on the particular
// platform, this option gives the filetypes in this listbox. When the user
// choose a filetype in the listbox, only the files of that type are listed. If
// this option is unspecified, or if it is set to the empty list, or if the
// File types listbox is not supported by the particular platform then all
// files are listed regardless of their types. See the section SPECIFYING FILE
// PATTERNS below for a discussion on the contents of filePatternList.
//
//   - [Initialdir] directory
//
// Specifies that the files in directory should be displayed when the dialog
// pops up. If this parameter is not specified, the initial directory defaults
// to the current working directory on non-Windows systems and on Windows
// systems prior to Vista. On Vista and later systems, the initial directory
// defaults to the last user-selected directory for the application. If the
// parameter specifies a relative path, the return value will convert the
// relative path to an absolute path.
//
//   - [Initialfile] filename
//
// Specifies a filename to be displayed in the dialog when it pops up.
//
//   - [Multiple] boolean
//
// Allows the user to choose multiple files from the Open dialog.
//
//   - [Parent] window
//
// Makes window the logical parent of the file dialog. The file dialog is
// displayed on top of its parent window. On Mac OS X, this turns the file
// dialog into a sheet attached to the parent window.
//
//   - [Title] titleString
//
// Specifies a string to display as the title of the dialog box. If this option
// is not specified, then a default title is displayed.
//
// Additional information might be available at the [Tcl/Tk getopenfile] page.
//
// [Tcl/Tk getopenfile]: https://www.tcl.tk/man/tcl9.0/TkCmd/getOpenFile.html
func GetOpenFile(options ...Opt) []string {
	return parseList(evalErr(fmt.Sprintf("tk_getOpenFile %s", collect(options...))))
}

// FileType specifies a single file type for the [Filetypes] option.
type FileType struct {
	TypeName   string   // Eg. "Go files"
	Extensions []string // Eg. []string{".go"}
	MacType    string   // Eg. "TEXT"
}

// GetSaveFile — pop up a dialog box for the user to select a file to save.
//
// # Description
//
// GetSaveFile pops up a dialog box for the user to select a file to save.
//
// The functio is usually associated with the Save as command in the File menu.
// If the user enters a file that already exists, the dialog box prompts the
// user for confirmation whether the existing file should be overwritten or
// not.
//
//   - [Confirmoverwrite] boolean
//
// Configures how the Save dialog reacts when the selected file already exists,
// and saving would overwrite it. A true value requests a confirmation dialog
// be presented to the user. A false value requests that the overwrite take
// place without confirmation. Default value is true.
//
//   - [Defaultextension] extension
//
// Specifies a string that will be appended to the filename if the user enters
// a filename without an extension. The default value is the empty string,
// which means no extension will be appended to the filename in any case. This
// option is ignored on Mac OS X, which does not require extensions to
// filenames, and the UNIX implementation guesses reasonable values for this
// from the -filetypes option when this is not supplied.
//
//   - [Filetypes] filePatternList ([][FileType])
//
// If a File types listbox exists in the file dialog on the particular
// platform, this option gives the filetypes in this listbox. When the user
// choose a filetype in the listbox, only the files of that type are listed. If
// this option is unspecified, or if it is set to the empty list, or if the
// File types listbox is not supported by the particular platform then all
// files are listed regardless of their types. See the section SPECIFYING FILE
// PATTERNS below for a discussion on the contents of filePatternList.
//
//   - [Initialdir] directory
//
// Specifies that the files in directory should be displayed when the dialog
// pops up. If this parameter is not specified, the initial directory defaults
// to the current working directory on non-Windows systems and on Windows
// systems prior to Vista. On Vista and later systems, the initial directory
// defaults to the last user-selected directory for the application. If the
// parameter specifies a relative path, the return value will convert the
// relative path to an absolute path.
//
//   - [Initialfile] filename
//
// Specifies a filename to be displayed in the dialog when it pops up.
//
//   - [Parent] window
//
// Makes window the logical parent of the file dialog. The file dialog is
// displayed on top of its parent window. On Mac OS X, this turns the file
// dialog into a sheet attached to the parent window.
//
//   - [Title] titleString
//
// Specifies a string to display as the title of the dialog box. If this option
// is not specified, then a default title is displayed.
//
// Additional information might be available at the [Tcl/Tk getopenfile] page.
//
// [Tcl/Tk getopenfile]: https://www.tcl.tk/man/tcl9.0/TkCmd/getOpenFile.html
func GetSaveFile(options ...Opt) string {
	return evalErr(fmt.Sprintf("tk_getSaveFile %s", collect(options...)))
}

// Place — Geometry manager for fixed or rubber-sheet placement
//
// # Description
//
// The placer is a geometry manager for Tk. It provides simple fixed placement
// of windows, where you specify the exact size and location of one window,
// called the content, within another window, called the container. The placer
// also provides rubber-sheet placement, where you specify the size and
// location of the content in terms of the dimensions of the container, so that
// the content changes size and location in response to changes in the size of
// the container. Lastly, the placer allows you to mix these styles of
// placement so that, for example, the content has a fixed width and height but
// is centered inside the container.
//
// The first argument must be a *Window.
//
// The following options are supported:
//
//   - [Anchor] where
//
// Where specifies which point of window is to be positioned at the (x,y)
// location selected by the -x, -y, -relx, and -rely options. The anchor point
// is in terms of the outer area of window including its border, if any. Thus
// if where is se then the lower-right corner of window's border will appear at
// the given (x,y) location in the container. The anchor position defaults to
// nw.
//
//   - [Bordermode] mode
//
// Mode determines the degree to which borders within the container are used in
// determining the placement of the content. The default and most common value
// is inside. In this case the placer considers the area of the container to be
// the innermost area of the container, inside any border: an option of -x 0
// corresponds to an x-coordinate just inside the border and an option of
// -relwidth 1.0 means window will fill the area inside the container's border.
//
// If mode is outside then the placer considers the area of the container to
// include its border; this mode is typically used when placing window outside
// its container, as with the options -x 0 -y 0 -anchor ne. Lastly, mode may be
// specified as ignore, in which case borders are ignored: the area of the
// container is considered to be its official X area, which includes any
// internal border but no external border. A bordermode of ignore is probably
// not very useful.
//
//   - [Height] size
//
// Size specifies the height for window in screen units (i.e. any of the forms
// accepted by Tk_GetPixels). The height will be the outer dimension of window
// including its border, if any. If size is an empty string, or if no -height
// or -relheight option is specified, then the height requested internally by
// the window will be used.
//
//   - [In] container
//
// Container specifies the path name of the window relative to which window is
// to be placed. Container must either be window's parent or a descendant of
// window's parent. In addition, container and window must both be descendants
// of the same top-level window. These restrictions are necessary to guarantee
// that window is visible whenever container is visible. If this option is not
// specified then the other window defaults to window's parent.
//
//   - [Relheight] size
//
// Size specifies the height for window. In this case the height is specified
// as a floating-point number relative to the height of the container: 0.5
// means window will be half as high as the container, 1.0 means window will
// have the same height as the container, and so on. If both -height and
// -relheight are specified for a content, their values are summed. For
// example, -relheight 1.0 -height -2 makes the content 2 pixels shorter than
// the container.
//
//   - [Relwidth] size
//
// Size specifies the width for window. In this case the width is specified as
// a floating-point number relative to the width of the container: 0.5 means
// window will be half as wide as the container, 1.0 means window will have the
// same width as the container, and so on. If both -width and -relwidth are
// specified for a content, their values are summed. For example, -relwidth 1.0
// -width 5 makes the content 5 pixels wider than the container.
//
//   - [Relx] location
//
// Location specifies the x-coordinate within the container window of the
// anchor point for window. In this case the location is specified in a
// relative fashion as a floating-point number: 0.0 corresponds to the left
// edge of the container and 1.0 corresponds to the right edge of the
// container. Location need not be in the range 0.0-1.0. If both -x and -relx
// are specified for a content then their values are summed. For example, -relx
// 0.5 -x -2 positions the left edge of the content 2 pixels to the left of the
// center of its container.
//
//   - [Rely] location
//
// Location specifies the y-coordinate within the container window of the
// anchor point for window. In this case the value is specified in a relative
// fashion as a floating-point number: 0.0 corresponds to the top edge of the
// container and 1.0 corresponds to the bottom edge of the container. Location
// need not be in the range 0.0-1.0. If both -y and -rely are specified for a
// content then their values are summed. For example, -rely 0.5 -x 3 positions
// the top edge of the content 3 pixels below the center of its container.
//
//   - [Width] size
//
// Size specifies the width for window in screen units (i.e. any of the forms
// accepted by Tk_GetPixels). The width will be the outer width of window
// including its border, if any. If size is an empty string, or if no -width or
// -relwidth option is specified, then the width requested internally by the
// window will be used.
//
//   - [X] location
//
// Location specifies the x-coordinate within the container window of the
// anchor point for window. The location is specified in screen units (i.e. any
// of the forms accepted by Tk_GetPixels) and need not lie within the bounds of
// the container window.
//
//   - [Y] location
//
// Location specifies the y-coordinate within the container window of the
// anchor point for window. The location is specified in screen units (i.e. any
// of the forms accepted by Tk_GetPixels) and need not lie within the bounds of
// the container window.
//
// If the same value is specified separately with two different options, such
// as -x and -relx, then the most recent option is used and the older one is
// ignored.
//
// Additional information might be available at the [Tcl/Tk place] page.
//
// [Tcl/Tk place]: https://www.tcl.tk/man/tcl9.0/TkCmd/place.htm
func Place(options ...Opt) {
	evalErr(fmt.Sprintf("place %s", collect(options...)))
}

// Lower — Change a window's position in the stacking order
//
// # Description
//
// If the belowThis argument is nil then the command lowers window so that
// it is below all of its siblings in the stacking order (it will be obscured
// by any siblings that overlap it and will not obscure any siblings). If
// belowThis is specified then it must be the path name of a window that is
// either a sibling of window or the descendant of a sibling of window. In this
// case the lower command will insert window into the stacking order just below
// belowThis (or the ancestor of belowThis that is a sibling of window); this
// could end up either raising or lowering window.
//
// All toplevel windows may be restacked with respect to each other, whatever
// their relative path names, but the window manager is not obligated to
// strictly honor requests to restack.
//
// Additional information might be available at the [Tcl/Tk lower] page.
//
// [Tcl/Tk lower]: https://www.tcl.tk/man/tcl9.0/TkCmd/lower.html
func (w *Window) Lower(belowThis Widget) {
	b := ""
	if belowThis != nil {
		b = belowThis.path()
	}
	evalErr(fmt.Sprintf("lower %s %s", w, b))
}

// Raise — Change a window's position in the stacking order
//
// # Description
//
// If the aboveThis argument is nil then the command raises window so that it
// is above all of its siblings in the stacking order (it will not be obscured
// by any siblings and will obscure any siblings that overlap it). If aboveThis
// is specified then it must be the path name of a window that is either a
// sibling of window or the descendant of a sibling of window. In this case the
// raise command will insert window into the stacking order just above
// aboveThis (or the ancestor of aboveThis that is a sibling of window); this
// could end up either raising or lowering window.
//
// All toplevel windows may be restacked with respect to each other, whatever
// their relative path names, but the window manager is not obligated to
// strictly honor requests to restack.
//
// On macOS raising an iconified toplevel window causes it to be deiconified.
//
// Additional information might be available at the [Tcl/Tk raise] page.
//
// [Tcl/Tk raise]: https://www.tcl.tk/man/tcl9.0/TkCmd/raise.html
func (w *Window) Raise(aboveThis Widget) {
	b := ""
	if aboveThis != nil {
		b = aboveThis.path()
	}
	evalErr(fmt.Sprintf("raise %s %s", w, b))
}

// Graph — use gnuplot to draw on a canvas. Graph returns 'w'.
//
// The 'script' argument is passed to a gnuplot executable, which must be
// installed on the machine.  See the [gnuplot site] for documentation about
// producing graphs. The script must not use the 'set term <device>' command.
//
// [gnuplot site]: http://www.gnuplot.info/
func (w *CanvasWidget) Graph(script string) *CanvasWidget {
	script = fmt.Sprintf("set terminal tkcanvas size %s, %s\n%s", w.Width(), w.Height(), script)
	out, err := gnuplot(script)
	if err != nil {
		fail(fmt.Errorf("plot: executing script: %s", err))
		return w
	}

	evalErr(fmt.Sprintf("%s\ngnuplot %s", out, w))
	return w
}

func gnuplot(script string) (out []byte, err error) {
	f, err := os.CreateTemp("", "tk8.6-")
	if err != nil {
		return nil, err
	}

	defer os.Remove(f.Name())

	if err := os.WriteFile(f.Name(), []byte(script), 0660); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), gnuplotTimeout)

	defer cancel()

	return exec.CommandContext(ctx, "gnuplot", f.Name()).Output()
}
