// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	"context"
	_ "embed"
	// TODO 	"encoding/base64"
	// TODO 	"errors"
	"fmt"
	"os"
	"os/exec"
	// TODO 	"path/filepath"
	// TODO 	"runtime"
	// TODO 	"strconv"
	"strings"
	// TODO 	"sync"
	// TODO 	"sync/atomic"
	"time"
	// TODO
	// TODO 	"github.com/evilsocket/islazy/zip"
	// TODO 	libtcl "modernc.org/libtcl9.0"
	// TODO 	libtk "modernc.org/libtk9.0"
	// TODO 	tklib "modernc.org/libtk9.0/library"
	// TODO 	tcl "modernc.org/tcl9.0"
)

const (
	gnuplotTimeout = time.Minute //TODO do not let the UI freeze
)

// App is the main/root application window.
var App *Window

// TODO // CollectErrors selects the behaviour on errors for certain functions that do
// TODO // not return error.
// TODO var CollectErrors bool
// TODO
// TODO // Error records errors when [CollectErrors] is true.
// TODO var Error error
// TODO
var (
	//TODO 	_ Widget = (*Window)(nil)
	//TODO
	//go:embed embed/gotk.png
	icon []byte

	cleanupDirs []string

	// TODO 	exitHandler Opt
	// TODO 	finished    atomic.Int32
	// TODO 	handlers    = map[int32]*eventHandler{}
	// TODO 	id          atomic.Int32
	// TODO 	interp      *tcl.Interp
	isBuilder = os.Getenv("MODERNC_BUILDER") != ""

// TODO 	tclDir      string
// TODO 	tkDir       string
// TODO
// TODO 	// https://pdos.csail.mit.edu/archive/rover/RoverDoc/escape_shell_table.html
// TODO 	//
// TODO 	// The following characters are dissallowed or have special meanings in Tcl and
// TODO 	// so are escaped:
// TODO 	//
// TODO 	//	&;`'"|*?~<>^()[]{}$\
// TODO 	badChars = [...]bool{
// TODO 		' ':  true,
// TODO 		'"':  true,
// TODO 		'$':  true,
// TODO 		'&':  true,
// TODO 		'(':  true,
// TODO 		')':  true,
// TODO 		'*':  true,
// TODO 		';':  true,
// TODO 		'<':  true,
// TODO 		'>':  true,
// TODO 		'?':  true,
// TODO 		'[':  true,
// TODO 		'\”: true,
// TODO 		'\\': true,
// TODO 		'\n': true,
// TODO 		'\r': true,
// TODO 		'\t': true,
// TODO 		']':  true,
// TODO 		'^':  true,
// TODO 		'`':  true,
// TODO 		'{':  true,
// TODO 		'|':  true,
// TODO 		'}':  true,
// TODO 		'~':  true,
// TODO 	}
// TODO
// TODO 	//TODO remove the associated tcl var on window destroy event both from the
// TODO 	//interp and this map.
// TODO 	textVariables = map[*Window]string{} // : tclName
)

//TODO
//TODO // Window represents a Tk window/widget. It implements common widget methods.
//TODO //
//TODO // Window implements Opt. When a Window instance is used as an Opt, it provides
//TODO // its path name.
//TODO type Window struct {
//TODO 	fpath string
//TODO }
//TODO
//TODO func (w *Window) isWidget() {}
//TODO
//TODO // Widget is implemented by every *Window
//TODO type Widget interface {
//TODO 	isWidget()
//TODO 	path() string
//TODO }
//TODO
//TODO func (w *Window) path() (r string) {
//TODO 	return w.String()
//TODO }
//TODO
//TODO // String implements fmt.Stringer.
//TODO func (w *Window) String() (r string) {
//TODO 	if r = w.fpath; r == "" {
//TODO 		r = "."
//TODO 	}
//TODO 	return r
//TODO }
//TODO
//TODO func (w *Window) optionString(_ *Window) string {
//TODO 	return w.String()
//TODO }
//TODO
//TODO func (w *Window) split(options []Opt) (opts []Opt, tvs []textVarOpt) {
//TODO 	for _, v := range options {
//TODO 		switch x := v.(type) {
//TODO 		case textVarOpt:
//TODO 			tvs = append(tvs, x)
//TODO 		default:
//TODO 			opts = append(opts, x)
//TODO 		}
//TODO 	}
//TODO 	return opts, tvs
//TODO }
//TODO
//TODO func (w *Window) newChild(nm string, options ...Opt) (rw *Window) {
//TODO 	class := strings.Replace(nm, "ttk_", "ttk::", 1)
//TODO 	nm = strings.Replace(nm, "ttk::", "t", 1)
//TODO 	if c := nm[len(nm)-1]; c >= '0' && c <= '9' {
//TODO 		nm += "_"
//TODO 	}
//TODO 	path := fmt.Sprintf("%s.%s%v", w, nm, id.Add(1))
//TODO 	options, tvs := w.split(options)
//TODO 	code := fmt.Sprintf("%s %s %s", class, path, winCollect(w, options...))
//TODO 	r, err := eval(code)
//TODO 	rw = &Window{fpath: r}
//TODO 	if err != nil {
//TODO 		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
//TODO 	}
//TODO 	if len(tvs) != 0 {
//TODO 		rw.Configure(tvs[len(tvs)-1])
//TODO 	}
//TODO 	return rw
//TODO }
//TODO
//TODO func eval(code string) (r string, err error) {
//TODO 	if dmesgs {
//TODO 		defer func() {
//TODO 			dmesg("code=%s -> r=%v err=%v", code, r, err)
//TODO 		}()
//TODO 	}
//TODO 	return interp.Eval(code, tcl.EvalGlobal)
//TODO }
//TODO
//TODO func evalErr(code string) (r string) {
//TODO 	r, err := eval(code)
//TODO 	if err != nil {
//TODO 		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
//TODO 	}
//TODO 	return r
//TODO }
//TODO
//TODO func fail(err error) {
//TODO 	if !CollectErrors {
//TODO 		if dmesgs {
//TODO 			dmesg("PANIC %v", err)
//TODO 		}
//TODO 		panic(err)
//TODO 	}
//TODO
//TODO 	Error = errors.Join(Error, err)
//TODO }
//TODO
//TODO func winCollect(w *Window, options ...Opt) string {
//TODO 	var a []string
//TODO 	for _, v := range options {
//TODO 		a = append(a, v.optionString(w))
//TODO 	}
//TODO 	return strings.Join(a, " ")
//TODO }
//TODO
//TODO func collect(options ...Opt) string {
//TODO 	var a []string
//TODO 	for _, v := range options {
//TODO 		a = append(a, v.optionString(nil))
//TODO 	}
//TODO 	return strings.Join(a, " ")
//TODO }
//TODO
//TODO // Opts is a list of options. It implements Opt.
//TODO type Opts []Opt
//TODO
//TODO func (o Opts) optionString(w *Window) string {
//TODO 	return winCollect(w, []Opt(o)...)
//TODO }
//TODO
//TODO // Opt represents an optional argument.
//TODO type Opt interface {
//TODO 	optionString(w *Window) string
//TODO }
//TODO
//TODO type rawOption string
//TODO
//TODO func (s rawOption) optionString(w *Window) string {
//TODO 	return string(s)
//TODO }
//TODO
//TODO type stringOption string
//TODO
//TODO func (s stringOption) optionString(w *Window) string {
//TODO 	return tclSafeString(string(s))
//TODO }
//TODO
//TODO // EventHandler is invoked when its associated event fires. The 'data' argument
//TODO // is the additional value passed when the handler was registered
//TODO type EventHandler func(w *Window, data any) (any, error)
//TODO
//TODO // EventDetacher is invoked when the handler is detached. The 'data' argument
//TODO // is the additional value passed when the handler was registered.
//TODO type EventDetacher func(w *Window, data any)
//TODO
//TODO type eventHandler struct {
//TODO 	data     any
//TODO 	detacher EventDetacher
//TODO 	handler  EventHandler
//TODO 	id       int32
//TODO 	tcl      string
//TODO 	w        *Window
//TODO }
//TODO
//TODO func newEventHandler(option string, args ...any) (r *eventHandler) {
//TODO 	if len(args) == 0 {
//TODO 		fail(fmt.Errorf("registering event handler: need at least one argument"))
//TODO 		return nil
//TODO 	}
//TODO
//TODO 	var handler EventHandler
//TODO 	var detacher EventDetacher
//TODO 	var data any
//TODO 	for _, v := range args {
//TODO 		switch x := v.(type) {
//TODO 		case EventHandler:
//TODO 			if handler != nil {
//TODO 				fail(fmt.Errorf("registering event handler: multiple handling functions"))
//TODO 				return nil
//TODO 			}
//TODO
//TODO 			handler = x
//TODO 		case func(*Window, any) (any, error):
//TODO 			if handler != nil {
//TODO 				fail(fmt.Errorf("registering event handler: multiple handling functions"))
//TODO 				return nil
//TODO 			}
//TODO
//TODO 			handler = x
//TODO 		case func():
//TODO 			if handler != nil {
//TODO 				if detacher != nil {
//TODO 					fail(fmt.Errorf("registering event handler: multiple detaching functions"))
//TODO 					return nil
//TODO 				}
//TODO
//TODO 				detacher = func(*Window, any) { x() }
//TODO 				break
//TODO 			}
//TODO
//TODO 			handler = func(*Window, any) (any, error) { x(); return nil, nil }
//TODO 		case EventDetacher:
//TODO 			if detacher != nil {
//TODO 				fail(fmt.Errorf("registering event handler: multiple detaching functions"))
//TODO 				return nil
//TODO 			}
//TODO
//TODO 			detacher = x
//TODO 		case func(*Window, any):
//TODO 			if detacher != nil {
//TODO 				fail(fmt.Errorf("registering event handler: multiple detaching functions"))
//TODO 				return nil
//TODO 			}
//TODO
//TODO 			detacher = x
//TODO 		default:
//TODO 			if data != nil {
//TODO 				fail(fmt.Errorf("registering event handler: multiple data values"))
//TODO 				return nil
//TODO 			}
//TODO
//TODO 			data = x
//TODO 		}
//TODO 	}
//TODO 	if handler == nil {
//TODO 		fail(fmt.Errorf("registering event handler: no event handler argument"))
//TODO 		return nil
//TODO 	}
//TODO
//TODO 	r = &eventHandler{
//TODO 		handler:  handler,
//TODO 		detacher: detacher,
//TODO 		data:     data,
//TODO 		id:       id.Add(1),
//TODO 		tcl:      option,
//TODO 	}
//TODO 	switch {
//TODO 	case r.detacher == nil:
//TODO 		r.detacher = func(w *Window, v any) { delete(handlers, r.id) }
//TODO 	default:
//TODO 		r.detacher = func(w *Window, v any) {
//TODO 			detacher(w, v)
//TODO 			delete(handlers, r.id)
//TODO 		}
//TODO 	}
//TODO 	handlers[r.id] = r
//TODO 	return r
//TODO }
//TODO
//TODO func (e *eventHandler) optionString(w *Window) string {
//TODO 	if e == nil {
//TODO 		return ""
//TODO 	}
//TODO
//TODO 	e.w = w
//TODO 	return fmt.Sprintf("%s {eventDispatcher %v}", e.tcl, e.id)
//TODO }
//TODO
//TODO func optionString(v any) string {
//TODO 	switch x := v.(type) {
//TODO 	case time.Duration:
//TODO 		return fmt.Sprint(int64((x + time.Millisecond/2) / time.Millisecond))
//TODO 	case []byte:
//TODO 		return base64.StdEncoding.EncodeToString(x)
//TODO 	case []FileType:
//TODO 		var a []string
//TODO 		for _, v := range x {
//TODO 			a = append(a, fmt.Sprintf("{%s {%s} %s}", tclSafeString(v.TypeName), tclSafeStrings(v.Extensions...), v.MacType))
//TODO 		}
//TODO 		return fmt.Sprintf("{%s}", strings.Join(a, " "))
//TODO 	default:
//TODO 		return tclSafeString(fmt.Sprint(v))
//TODO 	}
//TODO }
//TODO
//TODO func tclSafeStringBind(s string) string {
//TODO 	if s == "" {
//TODO 		return "{}"
//TODO 	}
//TODO
//TODO 	const badString = "&;`'\"|*?~^()[]{}$\\\n\r\t "
//TODO 	if strings.ContainsAny(s, badString) {
//TODO 		var b strings.Builder
//TODO 		for _, c := range s {
//TODO 			switch {
//TODO 			case int(c) < len(badChars) && badChars[c]:
//TODO 				fmt.Fprintf(&b, "\\x%02x", c)
//TODO 			default:
//TODO 				b.WriteRune(c)
//TODO 			}
//TODO 		}
//TODO 		s = b.String()
//TODO 	}
//TODO 	return s
//TODO }
//TODO
//TODO func tclSafeStrings(s ...string) string {
//TODO 	var a []string
//TODO 	for _, v := range s {
//TODO 		a = append(a, tclSafeString(v))
//TODO 	}
//TODO 	return strings.Join(a, " ")
//TODO }

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

//TODO func eventDispatcher(data any, interp *tcl.Interp, args []string) int {
//TODO 	id, err := strconv.Atoi(args[1])
//TODO 	if err != nil {
//TODO 		panic(todo("event dispatcher internal error: %q", args))
//TODO 	}
//TODO
//TODO 	h := handlers[int32(id)]
//TODO 	r, err := h.handler(h.w, h.data)
//TODO 	interp.SetResult(tclSafeString(fmt.Sprint(r)))
//TODO 	if err != nil {
//TODO 		return libtcl.TCL_ERROR
//TODO 	}
//TODO
//TODO 	return libtcl.TCL_OK
//TODO }
//TODO
//TODO func stdlib() (dir string, err error) {
//TODO 	if dir, err = os.MkdirTemp("", "tk-library-"); err != nil {
//TODO 		return "", err
//TODO 	}
//TODO
//TODO 	fn := filepath.Join(dir, "library.zip")
//TODO 	if err = os.WriteFile(fn, []byte(tklib.Zip), 0600); err != nil {
//TODO 		return
//TODO 	}
//TODO
//TODO 	if _, err = zip.Unzip(fn, dir); err != nil {
//TODO 		return
//TODO 	}
//TODO
//TODO 	return filepath.Join(dir, "library"), nil
//TODO }
//TODO
//TODO // Finalize releases all resources held, if any. This may include temporary
//TODO // files. Finalize is intended to be called on process shutdown only.
//TODO func Finalize() (err error) {
//TODO 	if finished.Swap(1) != 0 {
//TODO 		return
//TODO 	}
//TODO
//TODO 	defer runtime.UnlockOSThread()
//TODO
//TODO 	if interp != nil {
//TODO 		err = interp.Close()
//TODO 		interp = nil
//TODO 	}
//TODO 	for _, v := range []string{tclDir, tkDir} {
//TODO 		err = errors.Join(err, os.RemoveAll(v))
//TODO 	}
//TODO 	return err
//TODO }
//TODO
//TODO // bind — Arrange for X events to invoke functions
//TODO //
//TODO // # Description
//TODO //
//TODO // Bind tag options...
//TODO //
//TODO // The bind command associates commands with X events. If all three
//TODO // arguments are specified, bind will arrange for script (a Tcl script called
//TODO // the “binding script”) to be evaluated whenever the event(s) given by
//TODO // sequence occur in the window(s) identified by tag. If script is prefixed
//TODO // with a “+”, then it is appended to any existing binding for sequence;
//TODO // otherwise script replaces any existing binding. If script is an empty string
//TODO // then the current binding for sequence is destroyed, leaving sequence
//TODO // unbound. In all of the cases where a script argument is provided, bind
//TODO // returns an empty string.
//TODO //
//TODO // If sequence is specified without a script, then the script currently bound
//TODO // to sequence is returned, or an empty string is returned if there is no
//TODO // binding for sequence. If neither sequence nor script is specified, then the
//TODO // return value is a list whose elements are all the sequences for which there
//TODO // exist bindings for tag.
//TODO //
//TODO // The tag argument determines which window(s) the binding applies to. If tag
//TODO // begins with a dot, as in .a.b.c, then it must be the path name for a window;
//TODO // otherwise it may be an arbitrary string. Each window has an associated list
//TODO // of tags, and a binding applies to a particular window if its tag is among
//TODO // those specified for the window. Although the bindtags command may be used to
//TODO // assign an arbitrary set of binding tags to a window, the default binding
//TODO // tags provide the following behavior:
//TODO //
//TODO //   - If a tag is the name of an internal window the binding applies to that window.
//TODO //   - If the tag is the name of a class of widgets, such as Button, the binding applies to all widgets in that class.
//TODO //   - If the tag is the name of a toplevel window the binding applies to the toplevel window and all its internal windows.
//TODO //   - If tag has the value all, the binding applies to all windows in the application.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk bind] page.
//TODO //
//TODO // [Tcl/Tk bind]: https://www.tcl.tk/man/tcl9.0/TkCmd/bind.htm
//TODO func Bind(options ...any) {
//TODO 	a := []string{"bind"}
//TODO 	for _, v := range options {
//TODO 		switch x := v.(type) {
//TODO 		case *Window:
//TODO 			a = append(a, x.String())
//TODO 		case *eventHandler:
//TODO 			x.tcl = ""
//TODO 			a = append(a, x.optionString(nil))
//TODO 		default:
//TODO 			a = append(a, tclSafeStringBind(fmt.Sprint(x)))
//TODO 		}
//TODO 	}
//TODO 	evalErr(strings.Join(a, " "))
//TODO }
//TODO
//TODO // Img represents a Tk image.
//TODO type Img struct {
//TODO 	name string
//TODO }
//TODO
//TODO // String implements fmt.Stringer.
//TODO func (m *Img) String() string {
//TODO 	return m.optionString(nil)
//TODO }
//TODO
//TODO func (m *Img) optionString(_ *Window) string {
//TODO 	if m != nil {
//TODO 		return m.name
//TODO 	}
//TODO
//TODO 	return "img0" // does not exist
//TODO }
//TODO
//TODO // Bitmap — Images that display two colors
//TODO //
//TODO // # Description
//TODO //
//TODO // A bitmap is an image whose pixels can display either of two colors or be
//TODO // transparent. A bitmap image is defined by four things: a background color, a
//TODO // foreground color, and two bitmaps, called the source and the mask. Each of
//TODO // the bitmaps specifies 0/1 values for a rectangular array of pixels, and the
//TODO // two bitmaps must have the same dimensions. For pixels where the mask is
//TODO // zero, the image displays nothing, producing a transparent effect. For other
//TODO // pixels, the image displays the foreground color if the source data is one
//TODO // and the background color if the source data is zero.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk bitmap] page.
//TODO //
//TODO //   - [Background] color
//TODO //
//TODO // Specifies a background color for the image in any of the standard ways
//TODO // accepted by Tk. If this option is set to an empty string then the background
//TODO // pixels will be transparent. This effect is achieved by using the source
//TODO // bitmap as the mask bitmap, ignoring any -maskdata or -maskfile options.
//TODO //
//TODO //   - [Data] string
//TODO //
//TODO // Specifies the contents of the source bitmap as a string. The string must
//TODO // adhere to X11 bitmap format (e.g., as generated by the bitmap program). If
//TODO // both the -data and -file options are specified, the -data option takes
//TODO // precedence.
//TODO //
//TODO //   - [File] name
//TODO //
//TODO // name gives the name of a file whose contents define the source bitmap. The
//TODO // file must adhere to X11 bitmap format (e.g., as generated by the bitmap
//TODO // program).
//TODO //
//TODO //   - [Foreground] color
//TODO //
//TODO // Specifies a foreground color for the image in any of the standard ways
//TODO // accepted by Tk.
//TODO //
//TODO //   - [Maskdata] string
//TODO //
//TODO // Specifies the contents of the mask as a string. The string must adhere to
//TODO // X11 bitmap format (e.g., as generated by the bitmap program). If both the
//TODO // -maskdata and -maskfile options are specified, the -maskdata option takes
//TODO // precedence.
//TODO //
//TODO //   - [Maskfile] name
//TODO //
//TODO // name gives the name of a file whose contents define the mask. The file must
//TODO // adhere to X11 bitmap format (e.g., as generated by the bitmap program).
//TODO //
//TODO // [Tcl/Tk bitmap]: https://www.tcl.tk/man/tcl9.0/TkCmd/bitmap.htm
//TODO func NewBitmap(options ...Opt) *Img {
//TODO 	nm := fmt.Sprintf("bmp%v", id.Add(1))
//TODO 	code := fmt.Sprintf("image create bitmap %s %s", nm, collect(options...))
//TODO 	r, err := eval(code)
//TODO 	if err != nil {
//TODO 		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
//TODO 		return nil
//TODO 	}
//TODO
//TODO 	return &Img{name: nm}
//TODO }
//TODO
//TODO // Photo — Full-color images
//TODO //
//TODO // A photo is an image whose pixels can display any color with a varying degree
//TODO // of transparency (the alpha channel). A photo image is stored internally in
//TODO // full color (32 bits per pixel), and is displayed using dithering if
//TODO // necessary. Image data for a photo image can be obtained from a file or a
//TODO // string, or it can be supplied from C code through a procedural interface. At
//TODO // present, only PNG, GIF, PPM/PGM, and (read-only) SVG formats are supported,
//TODO // but an interface exists to allow additional image file formats to be added
//TODO // easily. A photo image is (semi)transparent if the image data it was obtained
//TODO // from had transparency information. In regions where no image data has been
//TODO // supplied, it is fully transparent. Transparency may also be modified with
//TODO // the transparency set subcommand.
//TODO //
//TODO //   - [Data] string
//TODO //
//TODO // Specifies the contents of the image as a string. The string should contain
//TODO // binary data or, for some formats, base64-encoded data (this is currently
//TODO // guaranteed to be supported for PNG and GIF images). The format of the string
//TODO // must be one of those for which there is an image file format handler that
//TODO // will accept string data. If both the -data and -file options are specified,
//TODO // the -file option takes precedence.
//TODO //
//TODO //   - [Format] format-name
//TODO //
//TODO // Specifies the name of the file format for the data specified with the -data
//TODO // or -file option.
//TODO //
//TODO //   - [File] name
//TODO //
//TODO // name gives the name of a file that is to be read to supply data for the
//TODO // photo image. The file format must be one of those for which there is an
//TODO // image file format handler that can read data.
//TODO //
//TODO //   - [Gamma] value
//TODO //
//TODO // Specifies that the colors allocated for displaying this image in a window
//TODO // should be corrected for a non-linear display with the specified gamma
//TODO // exponent value. (The intensity produced by most CRT displays is a power
//TODO // function of the input value, to a good approximation; gamma is the exponent
//TODO // and is typically around 2). The value specified must be greater than zero.
//TODO // The default value is one (no correction). In general, values greater than
//TODO // one will make the image lighter, and values less than one will make it
//TODO // darker.
//TODO //
//TODO //   - [Height] number
//TODO //
//TODO // Specifies the height of the image, in pixels. This option is useful
//TODO // primarily in situations where the user wishes to build up the contents of
//TODO // the image piece by piece. A value of zero (the default) allows the image to
//TODO // expand or shrink vertically to fit the data stored in it.
//TODO //
//TODO //   - [Palette] palette-spec
//TODO //
//TODO // Specifies the resolution of the color cube to be allocated for displaying
//TODO // this image, and thus the number of colors used from the colormaps of the
//TODO // windows where it is displayed. The palette-spec string may be either a
//TODO // single decimal number, specifying the number of shades of gray to use, or
//TODO // three decimal numbers separated by slashes (/), specifying the number of
//TODO // shades of red, green and blue to use, respectively. If the first form (a
//TODO // single number) is used, the image will be displayed in monochrome (i.e.,
//TODO // grayscale).
//TODO //
//TODO //   - [Width] number
//TODO //
//TODO // Specifies the width of the image, in pixels. This option is useful primarily
//TODO // in situations where the user wishes to build up the contents of the image
//TODO // piece by piece. A value of zero (the default) allows the image to expand or
//TODO // shrink horizontally to fit the data stored in it.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk photo] page.
//TODO //
//TODO // [Tcl/Tk photo]: https://www.tcl.tk/man/tcl9.0/TkCmd/photo.htm
//TODO func NewPhoto(options ...Opt) *Img {
//TODO 	nm := fmt.Sprintf("img%v", id.Add(1))
//TODO 	code := fmt.Sprintf("image create photo %s %s", nm, collect(options...))
//TODO 	r, err := eval(code)
//TODO 	if err != nil {
//TODO 		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
//TODO 		return nil
//TODO 	}
//TODO
//TODO 	return &Img{name: nm}
//TODO }
//TODO
//TODO // Width — Get the configured option value.
//TODO func (m *Img) Width() string {
//TODO 	return evalErr(fmt.Sprintf(`%s cget -width`, m))
//TODO }
//TODO
//TODO // Height — Get the configured option value.
//TODO func (m *Img) Height() string {
//TODO 	return evalErr(fmt.Sprintf(`%s cget -height`, m))
//TODO }
//TODO
//TODO // Graph — use gnuplot to draw on a photo. Graph returns 'm'
//TODO //
//TODO // The 'script' argument is passed to a gnuplot executable, which must be
//TODO // installed on the machine.  See the [gnuplot site] for documentation about
//TODO // producing graphs. The script must not use the 'set term <device>' command.
//TODO //
//TODO // The content of 'm' is replaced, including its internal name.
//TODO //
//TODO // [gnuplot site]: http://www.gnuplot.info/
//TODO func (m *Img) Graph(script string) *Img {
//TODO 	switch {
//TODO 	case strings.HasPrefix(m.name, "img"):
//TODO 		w, h := m.Width(), m.Height()
//TODO 		script = fmt.Sprintf("set terminal pngcairo size %s, %s\n%s", w, h, script)
//TODO 		out, err := gnuplot(script)
//TODO 		if err != nil {
//TODO 			fail(fmt.Errorf("plot: executing script: %s", err))
//TODO 			break
//TODO 		}
//TODO
//TODO 		*m = *NewPhoto(Width(w), Height(h), Data(out))
//TODO 	default:
//TODO 		fail(fmt.Errorf("plot: %s is not a photo", m))
//TODO 	}
//TODO 	return m
//TODO }
//TODO
//TODO // Destroy — Destroy one or more windows
//TODO //
//TODO // # Description
//TODO //
//TODO // This command deletes the windows given by the window arguments, plus all of
//TODO // their descendants. If a window “.” (App) is deleted then all windows will be
//TODO // destroyed and the application will (normally) exit. The windows are
//TODO // destroyed in order, and if an error occurs in destroying a window the
//TODO // command aborts without destroying the remaining windows. No error is
//TODO // returned if window does not exist.
//TODO func Destroy(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("destroy %s", collect(options...)))
//TODO }
//TODO
//TODO // Pack — Geometry manager that packs around edges of cavity
//TODO //
//TODO // # Description
//TODO //
//TODO // The options consist of one or more content windows followed
//TODO // by options that specify how to manage the content. See THE PACKER
//TODO // ALGORITHM for details on how the options are used by the packer.
//TODO //
//TODO // The first argument must be a *Window.
//TODO //
//TODO // The following options are supported:
//TODO //
//TODO //   - [After] other
//TODO //
//TODO // Other must the name of another window. Use its container as the container
//TODO // for the content, and insert the content just after other in the packing
//TODO // order.
//TODO //
//TODO //   - [Anchor] anchor
//TODO //
//TODO // Anchor must be a valid anchor position such as n or sw; it specifies where
//TODO // to position each content in its parcel. Defaults to center.
//TODO //
//TODO //   - [Before] other
//TODO //
//TODO // Other must the name of another window. Use its container as the container
//TODO // for the content, and insert the content just before other in the packing
//TODO // order.
//TODO //
//TODO //   - [Expand] boolean
//TODO //
//TODO // Specifies whether the content should be expanded to consume extra space in
//TODO // their container. Boolean may have any proper boolean value, such as 1 or no.
//TODO // Defaults to 0.
//TODO //
//TODO //   - [Fill] style
//TODO //
//TODO // If a content's parcel is larger than its requested dimensions, this option
//TODO // may be used to stretch the content. Style must have one of the following
//TODO // values:
//TODO //
//TODO //   - "none" - Give the content its requested dimensions plus any internal
//TODO //     padding requested with -ipadx or -ipady. This is the default.
//TODO //   - "x" - Stretch the content horizontally to fill the entire width of its
//TODO //     parcel (except leave external padding as specified by -padx).
//TODO //   - "y" - Stretch the content vertically to fill the entire height of its parcel
//TODO //     (except leave external padding as specified by -pady).
//TODO //   - "both": Stretch the content both horizontally and vertically.
//TODO //
//TODO // .
//TODO //
//TODO //   - [In] container
//TODO //
//TODO // Insert the window at the end of the packing order for the container window
//TODO // given by container.
//TODO //
//TODO //   - [Ipadx] amount
//TODO //
//TODO // Amount specifies how much horizontal internal padding to leave on each side
//TODO // of the content. Amount must be a valid screen distance, such as 2 or .5c. It
//TODO // defaults to 0.
//TODO //
//TODO //   - [Ipady] amount
//TODO //
//TODO // Amount specifies how much vertical internal padding to leave on each side of
//TODO // the content. Amount defaults to 0.
//TODO //
//TODO //   - [Padx] amount
//TODO //
//TODO // Amount specifies how much horizontal external padding to leave on each side
//TODO // of the content. Amount may be a list of two values to specify padding for
//TODO // left and right separately. Amount defaults to 0.
//TODO //
//TODO //   - [Pady] amount
//TODO //
//TODO // Amount specifies how much vertical external padding to leave on each side of
//TODO // the content. Amount may be a list of two values to specify padding for top
//TODO // and bottom separately. Amount defaults to 0.
//TODO //
//TODO //   - [Side] side
//TODO //
//TODO // Specifies which side of the container the content will be packed against.
//TODO // Must be "left", "right", "top", or "bottom". Defaults to top.
//TODO //
//TODO // If no -in, -after or -before option is specified then each of the content
//TODO // will be inserted at the end of the packing list for its parent unless it is
//TODO // already managed by the packer (in which case it will be left where it is).
//TODO // If one of these options is specified then all the content will be inserted
//TODO // at the specified point. If any of the content are already managed by the
//TODO // geometry manager then any unspecified options for them retain their previous
//TODO // values rather than receiving default values.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk pack] page.
//TODO //
//TODO // [Tcl/Tk pack]: https://www.tcl.tk/man/tcl9.0/TkCmd/pack.htm
//TODO func Pack(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("pack %s", collect(options...)))
//TODO }
//TODO
//TODO // Wait — Wait for a window to be destroyed
//TODO //
//TODO // # Description
//TODO //
//TODO // Wait command waits for 'w' to be destroyed. This is typically used to wait
//TODO // for a user to finish interacting with a dialog box before using the result
//TODO // of that interaction.
//TODO //
//TODO // While the Wwait command is waiting it processes events in the normal
//TODO // fashion, so the application will continue to respond to user interactions.
//TODO // If an event handler invokes Wait again, the nested call to Wait must
//TODO // complete before the outer call can complete.
//TODO func (w *Window) Wait() {
//TODO 	evalErr(fmt.Sprintf("tkwait window %s", w))
//TODO }
//TODO
//TODO // WaitVisibility — Wait for a window to change visibility
//TODO //
//TODO // # Description
//TODO //
//TODO // WaitVisibility command waits for a change in w's visibility state (as
//TODO // indicated by the arrival of a VisibilityNotify event). This form is
//TODO // typically used to wait for a newly-created window to appear on the screen
//TODO // before taking some action.
//TODO //
//TODO // While the Wwait command is waiting it processes events in the normal
//TODO // fashion, so the application will continue to respond to user interactions.
//TODO // If an event handler invokes Wait again, the nested call to Wait must
//TODO // complete before the outer call can complete.
//TODO func (w *Window) WaitVisibility() {
//TODO 	evalErr(fmt.Sprintf("tkwait visibility %s", w))
//TODO }
//TODO
//TODO // IconPhoto — change window icon
//TODO //
//TODO // # Description
//TODO //
//TODO // IconPhoto sets the titlebar icon for window based on the named photo images.
//TODO // If -default is specified, this is applied to all future created toplevels as
//TODO // well. The data in the images is taken as a snapshot at the time of
//TODO // invocation. If the images are later changed, this is not reflected to the
//TODO // titlebar icons. Multiple images are accepted to allow different images sizes
//TODO // (e.g., 16x16 and 32x32) to be provided. The window manager may scale
//TODO // provided icons to an appropriate size.
//TODO //
//TODO // On Windows, the images are packed into a Windows icon structure. This will
//TODO // override an ico specified to wm iconbitmap, and vice versa. This command
//TODO // sets the taskbar icon for the window.
//TODO //
//TODO // On X, the images are arranged into the _NET_WM_ICON X property, which most
//TODO // modern window managers support. A wm iconbitmap may exist simultaneously. It is
//TODO // recommended to use not more than 2 icons, placing the larger icon first. This
//TODO // command also sets the panel icon for the application if the window manager or
//TODO // desktop environment supports it.
//TODO //
//TODO // On Macintosh, the first image called is loaded into an OSX-native icon
//TODO // format, and becomes the application icon in dialogs, the Dock, and other
//TODO // contexts. At the script level the command will accept only the first image
//TODO // passed in the parameters as support for multiple sizes/resolutions on macOS
//TODO // is outside Tk's scope. Developers should use the largest icon they can
//TODO // support (preferably 512 pixels) to ensure smooth rendering on the Mac.
//TODO //
//TODO // More information might be available at the [Tcl/Tk wm] page.
//TODO //
//TODO // [Tcl/Tk wm]: https://www.tcl.tk/man/tcl9.0/TkCmd/wm.htm#M50
//TODO func (w *Window) IconPhoto(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("wm iconphoto %s %s", w, collect(options...)))
//TODO }
//TODO
//TODO // WmTitle — change the window manager title
//TODO //
//TODO // # Description
//TODO //
//TODO // If string is specified, then it will be passed to the window manager for use
//TODO // as the title for window (the window manager should display this string in
//TODO // window's title bar). In this case the command returns an empty string. If
//TODO // string is not specified then the command returns the current title for the
//TODO // window. The title for a window defaults to its name.
//TODO //
//TODO // More information might be available at the [Tcl/Tk wm] page.
//TODO //
//TODO // [Tcl/Tk wm]: https://www.tcl.tk/man/tcl9.0/TkCmd/wm.htm#M63
//TODO func (w *Window) WmTitle(s string) string {
//TODO 	if s != "" {
//TODO 		s = tclSafeString(s)
//TODO 	}
//TODO 	return evalErr(fmt.Sprintf("wm title %s %s", w, s))
//TODO }
//TODO
//TODO // Center centers 'w' and returns 'w'.
//TODO func (w *Window) Center() *Window {
//TODO 	evalErr(fmt.Sprintf("tk::PlaceWindow %s center", w))
//TODO 	return w
//TODO }
//TODO
//TODO // Grid — Geometry manager that arranges widgets in a grid
//TODO //
//TODO // # Description
//TODO //
//TODO // The arguments consist of the names of one or more content windows followed
//TODO // by pairs of arguments that specify how to manage the content. The characters
//TODO // -, x and ^, can be specified instead of a window name to alter the default
//TODO // location of a window, as described in the RELATIVE PLACEMENT section, below.
//TODO //
//TODO // The following options are supported:
//TODO //
//TODO //   - [Column] n
//TODO //
//TODO // Insert the window so that it occupies the nth column in the grid. Column
//TODO // numbers start with 0. If this option is not supplied, then the window is
//TODO // arranged just to the right of previous window specified on this call to
//TODO // grid, or column “0” if it is the first window. For each x that immediately
//TODO // precedes the window, the column position is incremented by one. Thus the x
//TODO // represents a blank column for this row in the grid.
//TODO //
//TODO //   - [Columnspan] n
//TODO //
//TODO // Insert the window so that it occupies n columns in the grid. The default is
//TODO // one column, unless the window name is followed by a -, in which case the
//TODO // columnspan is incremented once for each immediately following -.
//TODO //
//TODO //   - [In] container
//TODO //
//TODO // Insert the window(s) in the container window given by container. The default
//TODO // is the first window's parent window.
//TODO //
//TODO //   - [Ipadx] amount
//TODO //
//TODO // The amount specifies how much horizontal internal padding to leave on each
//TODO // side of the content. This is space is added inside the content border. The
//TODO // amount must be a valid screen distance, such as 2 or .5c. It defaults to 0.
//TODO //
//TODO //   - [Ipady] amount
//TODO //
//TODO // The amount specifies how much vertical internal padding to leave on the top
//TODO // and bottom of the content. This space is added inside the content border.
//TODO // The amount defaults to 0.
//TODO //
//TODO //   - [Padx] amount
//TODO //
//TODO // The amount specifies how much horizontal external padding to leave on each
//TODO // side of the content, in screen units. Amount may be a list of two values to
//TODO // specify padding for left and right separately. The amount defaults to 0.
//TODO // This space is added outside the content border.
//TODO //
//TODO //   - [Pady] amount
//TODO //
//TODO // The amount specifies how much vertical external padding to leave on the top
//TODO // and bottom of the content, in screen units. Amount may be a list of two
//TODO // values to specify padding for top and bottom separately. The amount defaults
//TODO // to 0. This space is added outside the content border.
//TODO //
//TODO //   - [Row] n
//TODO //
//TODO // Insert the content so that it occupies the nth row in the grid. Row numbers
//TODO // start with 0. If this option is not supplied, then the content is arranged
//TODO // on the same row as the previous content specified on this call to grid, or
//TODO // the next row after the highest occupied row if this is the first content.
//TODO //
//TODO //   - [Rowspan] n
//TODO //
//TODO // Insert the content so that it occupies n rows in the grid. The default is
//TODO // one row. If the next grid command contains ^ characters instead of content
//TODO // that line up with the columns of this content, then the rowspan of this
//TODO // content is extended by one.
//TODO //
//TODO //   - [Sticky] style
//TODO //
//TODO // If a content's cell is larger than its requested dimensions, this option may
//TODO // be used to position (or stretch) the content within its cell. Style is a
//TODO // string that contains zero or more of the characters n, s, e or w. The string
//TODO // can optionally contain spaces or commas, but they are ignored. Each letter
//TODO // refers to a side (north, south, east, or west) that the content will “stick”
//TODO // to. If both n and s (or e and w) are specified, the content will be
//TODO // stretched to fill the entire height (or width) of its cavity. The -sticky
//TODO // option subsumes the combination of -anchor and -fill that is used by pack.
//TODO // The default is “”, which causes the content to be centered in its cavity, at
//TODO // its requested size.
//TODO //
//TODO // If any of the content is already managed by the geometry manager then any
//TODO // unspecified options for them retain their previous values rather than
//TODO // receiving default values.
//TODO //
//TODO // More information might be available at the [Tcl/Tk grid] page.
//TODO //
//TODO // [Tcl/Tk grid]: https://www.tcl.tk/man/tcl9.0/TkCmd/grid.html#M9
//TODO func Grid(w Widget, options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("grid configure %s %s", w, collect(options...)))
//TODO }
//TODO
//TODO // Configure alters the configuration of 'w' and returns 'w'.
//TODO func (w *Window) Configure(options ...Opt) *Window {
//TODO 	options, tvs := w.split(options)
//TODO 	if len(options) != 0 {
//TODO 		evalErr(fmt.Sprintf("%s configure %s", w, collect(options...)))
//TODO 	}
//TODO 	if len(tvs) != 0 {
//TODO 		tvo := tvs[len(tvs)-1]
//TODO 		tclVar := textVariables[w]
//TODO 		if tclVar == "" {
//TODO 			tclVar = fmt.Sprintf("textVar%d", id.Add(1))
//TODO 			textVariables[w] = tclVar
//TODO 			evalErr(fmt.Sprintf("%s configure -textvariable %s", w, tclVar))
//TODO 		}
//TODO 		evalErr(fmt.Sprintf("set %s %s", tclVar, tclSafeString(string(tvo))))
//TODO 	}
//TODO 	return w
//TODO }
//TODO
//TODO // tk_messageBox — pops up a message window and waits for user response.
//TODO //
//TODO // # Description
//TODO //
//TODO // This procedure creates and displays a message window with an
//TODO // application-specified message, an icon and a set of buttons. Each of the
//TODO // buttons in the message window is identified by a unique symbolic name (see
//TODO // the -type options). After the message window is popped up, tk_messageBox
//TODO // waits for the user to select one of the buttons. Then it returns the
//TODO // symbolic name of the selected button.  The following optins are
//TODO // supported:
//TODO //
//TODO //   - [Command] handler
//TODO //
//TODO // Specifies the handler to invoke when the user closes the
//TODO // dialog. The actual command consists of string followed by a space and the
//TODO // name of the button clicked by the user to close the dialog. This is only
//TODO // available on Mac OS X.
//TODO //
//TODO //   - [Default] name
//TODO //
//TODO // Name gives the symbolic name of the default button for this message window (
//TODO // “ok”, “cancel”, and so on). See -type for a list of the symbolic names. If
//TODO // this option is not specified, the first button in the dialog will be made
//TODO // the default.
//TODO //
//TODO //   - [Detail] string
//TODO //
//TODO // Specifies an auxiliary message to the main message given by the -message
//TODO // option. The message detail will be presented beneath the main message and,
//TODO // where supported by the OS, in a less emphasized font than the main message.
//TODO //
//TODO //   - [Icon] iconImage
//TODO //
//TODO // Specifies an icon to display. IconImage must be one of the following: error,
//TODO // info, question or warning. If this option is not specified, then the info
//TODO // icon will be displayed.
//TODO //
//TODO //   - [Message] string
//TODO //
//TODO // Specifies the message to display in this message box. The default value is
//TODO // an empty string.
//TODO //
//TODO //   - [Parent] window
//TODO //
//TODO // Makes window the logical parent of the message box. The message box is
//TODO // displayed on top of its parent window.
//TODO //
//TODO //   - [Title] titleString
//TODO //
//TODO // Specifies a string to display as the title of the message box. The default
//TODO // value is an empty string.
//TODO //
//TODO //   - [Type] predefinedType
//TODO //
//TODO // Arranges for a predefined set of buttons to be displayed. The following
//TODO // values are possible for predefinedType:
//TODO //
//TODO //   - abortretryignore - Displays three buttons whose symbolic names are abort, retry and ignore.
//TODO //   - ok - Displays one button whose symbolic name is ok.
//TODO //   - okcancel - Displays two buttons whose symbolic names are ok and cancel.
//TODO //   - retrycancel - Displays two buttons whose symbolic names are retry and cancel.
//TODO //   - yesno - Displays two buttons whose symbolic names are yes and no.
//TODO //   - yesnocancel - Displays three buttons whose symbolic names are yes, no and cancel.
//TODO //
//TODO // More information might be available at the [Tcl/Tk messageBox] page.
//TODO //
//TODO // [Tcl/Tk messageBox]: https://www.tcl.tk/man/tcl9.0/TkCmd/messageBox.html
//TODO func MessageBox(options ...Opt) string {
//TODO 	return evalErr(fmt.Sprintf("tk_messageBox %s", collect(options...)))
//TODO }
//TODO
//TODO // Bell — Ring a display's bell
//TODO //
//TODO // # Description
//TODO //
//TODO // This command rings the bell on the display for window and returns an empty
//TODO // string. If the -displayof option is omitted, the display of the
//TODO // application's main window is used by default. The command uses the current
//TODO // bell-related settings for the display, which may be modified with programs
//TODO // such as xset.
//TODO //
//TODO // If -nice is not specified, this command also resets the screen saver for the
//TODO // screen. Some screen savers will ignore this, but others will reset so that
//TODO // the screen becomes visible again.
//TODO //
//TODO //   - [Displayof] window
//TODO //   - [Nice]
//TODO //
//TODO // More information might be available at the [Tcl/Tk bell] page.
//TODO //
//TODO // [Tcl/Tk bell]: https://www.tcl.tk/man/tcl9.0/TclCmd/bell.htm
//TODO func Bell(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("bell %s", collect(options...)))
//TODO }
//TODO
//TODO // ChooseColor — pops up a dialog box for the user to select a color.
//TODO //
//TODO // # Description
//TODO //
//TODO // ChooseColor pops up a dialog box for the user to select a color. The
//TODO // following option-value pairs are possible as command line arguments:
//TODO //
//TODO //   - [Initialcolor] color
//TODO //
//TODO // Specifies the color to display in the color dialog when it pops up. color
//TODO // must be in a form acceptable to the Tk_GetColor function.
//TODO //
//TODO //   - [Parent] window
//TODO //
//TODO // Makes window the logical parent of the color dialog. The color dialog is
//TODO // displayed on top of its parent window.
//TODO //
//TODO //   - [Title] titleString
//TODO //
//TODO // Specifies a string to display as the title of the dialog box. If this option
//TODO // is not specified, then a default title will be displayed.
//TODO //
//TODO // If the user selects a color, ChooseColor will return the name of the
//TODO // color in a form acceptable to Tk_GetColor. If the user cancels the
//TODO // operation, both commands will return the empty string.
//TODO //
//TODO // More information might be available at the [Tcl/Tk choosecolor] page.
//TODO //
//TODO // [Tcl/Tk choosecolor]: https://www.tcl.tk/man/tcl9.0/TclCmd/chooseColor.htm
//TODO func ChooseColor(options ...Opt) string {
//TODO 	return evalErr(fmt.Sprintf("tk_chooseColor %s", collect(options...)))
//TODO }
//TODO
//TODO // Busy — confine pointer events to a window sub-tree
//TODO //
//TODO // # Description
//TODO //
//TODO // The Busy command provides a simple means to block pointer events from Tk
//TODO // widgets, while overriding the widget's cursor with a configurable busy
//TODO // cursor. Note this command does not prevent keyboard events from being sent
//TODO // to the widgets made busy.
//TODO //
//TODO //   - [Cursor] cursorName
//TODO //
//TODO // Specifies the cursor to be displayed when the widget is made busy.
//TODO // CursorName can be in any form accepted by Tk_GetCursor. The default cursor
//TODO // is wait on Windows and watch on other platforms.
//TODO //
//TODO // More information might be available at the [Tcl/Tk busy] page.
//TODO //
//TODO // [Tcl/Tk update]: https://www.tcl.tk/man/tcl9.0/TclCmd/busy.htm
//TODO func (w *Window) Busy(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("tk busy %s %s", w, collect(options...)))
//TODO }
//TODO
//TODO // BusyForget — undo Busy
//TODO //
//TODO // # Description
//TODO //
//TODO // Releases resources allocated by the [Window.Busy] command for window, including
//TODO // the transparent window. User events will again be received by window.
//TODO // Resources are also released when window is destroyed. Window must be the
//TODO // name of a widget specified in the hold operation, otherwise an error is
//TODO // reported.
//TODO //
//TODO // More information might be available at the [Tcl/Tk busy] page.
//TODO //
//TODO // [Tcl/Tk update]: https://www.tcl.tk/man/tcl9.0/TclCmd/busy.htm
//TODO func (w *Window) BusyForget(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("tk busy forget %s %s", w, collect(options...)))
//TODO }
//TODO
//TODO // Update — Process pending events and idle callbacks
//TODO //
//TODO // More information might be available at the [Tcl/Tk update] page.
//TODO //
//TODO // [Tcl/Tk update]: https://www.tcl.tk/man/tcl9.0/TclCmd/update.htm
//TODO func Update() {
//TODO 	evalErr("update")
//TODO }
//TODO
//TODO //TODO?
//TODO //  - [Command] string
//TODO //
//TODO // Specifies the prefix of a Tcl command to invoke when the user closes the
//TODO // dialog after having selected an item. This callback is not called if the
//TODO // user cancelled the dialog. The actual command consists of string followed by
//TODO // a space and the value selected by the user in the dialog. This is only
//TODO // available on Mac OS X.
//TODO
//TODO // ChooseDirectory — pops up a dialog box for the user to select a directory.
//TODO //
//TODO // # Description
//TODO //
//TODO // The procedure tk_chooseDirectory pops up a dialog box for the user to select
//TODO // a directory. The following option-value pairs are possible as command line
//TODO // arguments:
//TODO //
//TODO //   - [Initialdir] dirname
//TODO //
//TODO // Specifies that the directories in directory should be displayed when the
//TODO // dialog pops up. If this parameter is not specified, the initial directory
//TODO // defaults to the current working directory on non-Windows systems and on
//TODO // Windows systems prior to Vista. On Vista and later systems, the initial
//TODO // directory defaults to the last user-selected directory for the application.
//TODO // If the parameter specifies a relative path, the return value will convert
//TODO // the relative path to an absolute path.
//TODO //
//TODO //   - [Message] string
//TODO //
//TODO // Specifies a message to include in the client area of the dialog. This is
//TODO // only available on Mac OS X.
//TODO //
//TODO //   - [Mustexist] boolean
//TODO //
//TODO // Specifies whether the user may specify non-existent directories. If this
//TODO // parameter is true, then the user may only select directories that already
//TODO // exist. The default value is false.
//TODO //
//TODO //   - [Parent] window
//TODO //
//TODO // Makes window the logical parent of the dialog. The dialog is displayed on
//TODO // top of its parent window. On Mac OS X, this turns the file dialog into a
//TODO // sheet attached to the parent window.
//TODO //
//TODO //   - [Title] titleString
//TODO //
//TODO // Specifies a string to display as the title of the dialog box. If this option
//TODO // is not specified, then a default title will be displayed.
//TODO //
//TODO // More information might be available at the [Tcl/Tk chooseDirectory] page.
//TODO //
//TODO // [Tcl/Tk chooseDirectory]: https://www.tcl.tk/man/tcl9.0/TkCmd/chooseDirectory.htm
//TODO func ChooseDirectory(options ...Opt) string {
//TODO 	return evalErr(fmt.Sprintf("tk_chooseDirectory %s", collect(options...)))
//TODO }
//TODO
//TODO // ClipboardAppend — Manipulate Tk clipboard
//TODO //
//TODO // # Description
//TODO //
//TODO // This command provides a Tcl interface to the Tk clipboard, which stores data
//TODO // for later retrieval using the selection mechanism (via the -selection
//TODO // CLIPBOARD option). In order to copy data into the clipboard, clipboard clear
//TODO // must be called, followed by a sequence of one or more calls to clipboard
//TODO // append. To ensure that the clipboard is updated atomically, all appends
//TODO // should be completed before returning to the event loop.
//TODO //
//TODO // ClipboardAppend appends 'data' to the clipboard on window's display in the
//TODO // form given by type with the representation given by format and claims
//TODO // ownership of the clipboard on window's display.
//TODO //
//TODO //   - [Displayof] window
//TODO //
//TODO //   - [Format] format
//TODO //
//TODO // The format argument specifies the representation that should be used to
//TODO // transmit the selection to the requester (the second column of Table 2 of the
//TODO // ICCCM), and defaults to STRING. If format is STRING, the selection is
//TODO // transmitted as 8-bit ASCII characters. If format is ATOM, then the data is
//TODO // divided into fields separated by white space; each field is converted to its
//TODO // atom value, and the 32-bit atom value is transmitted instead of the atom
//TODO // name. For any other format, data is divided into fields separated by white
//TODO // space and each field is converted to a 32-bit integer; an array of integers
//TODO // is transmitted to the selection requester. Note that strings passed to
//TODO // clipboard append are concatenated before conversion, so the caller must take
//TODO // care to ensure appropriate spacing across string boundaries. All items
//TODO // appended to the clipboard with the same type must have the same format.
//TODO //
//TODO // The format argument is needed only for compatibility with clipboard
//TODO // requesters that do not use Tk. If the Tk toolkit is being used to retrieve
//TODO // the CLIPBOARD selection then the value is converted back to a string at the
//TODO // requesting end, so format is irrelevant.
//TODO //
//TODO //   - [Type] type
//TODO //
//TODO // Type specifies the form in which the selection is to be returned (the
//TODO // desired “target” for conversion, in ICCCM terminology), and should be an
//TODO // atom name such as STRING or FILE_NAME; see the Inter-Client Communication
//TODO // Conventions Manual for complete details. Type defaults to STRING.
//TODO //
//TODO // More information might be available at the [Tcl/Tk clipboard] page.
//TODO //
//TODO // [Tcl/Tk clipboard]: https://www.tcl.tk/man/tcl9.0/TkCmd/clipboard.htm
//TODO func ClipboardAppend(data string, options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("clipboard append %s -- %s", collect(options...), tclSafeString(data)))
//TODO }
//TODO
//TODO // ClipboardClear — Manipulate Tk clipboard
//TODO //
//TODO // # Description
//TODO //
//TODO // Claims ownership of the clipboard on window's display and removes any
//TODO // previous contents. Window defaults to App. Returns an empty string.
//TODO //
//TODO // More information might be available at the [Tcl/Tk clipboard] page.
//TODO //
//TODO // [Tcl/Tk clipboard]: https://www.tcl.tk/man/tcl9.0/TkCmd/clipboard.htm
//TODO func ClipboardClear(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("clipboard clear %s", collect(options...)))
//TODO }
//TODO
//TODO // ClipboardGet — Manipulate Tk clipboard
//TODO //
//TODO // # Description
//TODO //
//TODO // Retrieve data from the clipboard on window's display. Window defaults to App.
//TODO //
//TODO //   - [Displayof] window
//TODO //
//TODO //   - [Type] type
//TODO //
//TODO // Type specifies the form in which the data is to be returned and should be an
//TODO // atom name such as STRING or FILE_NAME. Type defaults to STRING. This command
//TODO // is equivalent to [SelectionGet](Selection("CLIPBOARD").
//TODO //
//TODO // Note that on modern X11 systems, the most useful type to retrieve for
//TODO // transferred strings is not STRING, but rather UTF8_STRING.
//TODO //
//TODO // More information might be available at the [Tcl/Tk clipboard] page.
//TODO //
//TODO // [Tcl/Tk clipboard]: https://www.tcl.tk/man/tcl9.0/TkCmd/clipboard.htm
//TODO func ClipboardGet(options ...Opt) string {
//TODO 	return evalErr(fmt.Sprintf("clipboard get %s", collect(options...)))
//TODO }
//TODO
//TODO // ExitHandler returns a canned [Command] that destroys the [App].
//TODO func ExitHandler() Opt {
//TODO 	return exitHandler
//TODO }
//TODO
//TODO // Exit provides a canned [Button] with default [Txt] "Exit", bound to the
//TODO // [ExitHandler].
//TODO //
//TODO // Use [Window.Exit] to create a Exit with a particular parent.
//TODO func Exit(options ...Opt) *ButtonWidget {
//TODO 	return App.Exit(options...)
//TODO }
//TODO
//TODO // Exit provides a canned [Button] with default [Txt] "Exit", bound to the
//TODO // [ExitHandler].
//TODO //
//TODO // The resulting [Window] is a child of 'w'
//TODO func (w *Window) Exit(options ...Opt) *ButtonWidget {
//TODO 	return w.Button(append([]Opt{Txt("Exit"), ExitHandler()}, options...)...)
//TODO }
//TODO
//TODO // TExit provides a canned [TButton] with default [Txt] "Exit", bound to the
//TODO // [ExitHandler].
//TODO //
//TODO // Use [Window.TExit] to create a TExit with a particular parent.
//TODO func TExit(options ...Opt) *TButtonWidget {
//TODO 	return App.TExit(options...)
//TODO }
//TODO
//TODO // TExit provides a canned [TButton] with default [Txt] "Exit", bound to the
//TODO // [ExitHandler].
//TODO //
//TODO // The resulting [Window] is a child of 'w'
//TODO func (w *Window) TExit(options ...Opt) *TButtonWidget {
//TODO 	return w.TButton(append([]Opt{Txt("Exit"), ExitHandler()}, options...)...)
//TODO }
//TODO
//TODO type textVarOpt string
//TODO
//TODO func (textVarOpt) optionString(*Window) string {
//TODO 	panic("internal error") // Not supposed to be invoked.
//TODO }
//TODO
//TODO // Textvariable option.
//TODO //
//TODO // Specifies the value to be displayed inside the widget.
//TODO // The way in which the string is displayed in the widget depends on the
//TODO // particular widget and may be determined by other options, such as
//TODO // -anchor or -justify.
//TODO //
//TODO // Known uses:
//TODO //   - [Button]
//TODO //   - [Checkbutton]
//TODO //   - [Entry]
//TODO //   - [Label]
//TODO //   - [Menubutton]
//TODO //   - [Message]
//TODO //   - [Radiobutton]
//TODO //   - [Spinbox]
//TODO //   - [TButton]
//TODO //   - [TCheckbutton]
//TODO //   - [TCombobox] (widget specific)
//TODO //   - [TEntry] (widget specific)
//TODO //   - [TLabel]
//TODO //   - [TMenubutton]
//TODO //   - [TRadiobutton]
//TODO func Textvariable(s string) Opt {
//TODO 	return textVarOpt(s)
//TODO }
//TODO
//TODO // Textvariable — Get the configured option value.
//TODO //
//TODO // Known uses:
//TODO //   - [Button]
//TODO //   - [Checkbutton]
//TODO //   - [Entry]
//TODO //   - [Label]
//TODO //   - [Menubutton]
//TODO //   - [Message]
//TODO //   - [Radiobutton]
//TODO //   - [Spinbox]
//TODO //   - [TButton]
//TODO //   - [TCheckbutton]
//TODO //   - [TCombobox] (widget specific)
//TODO //   - [TEntry] (widget specific)
//TODO //   - [TLabel]
//TODO //   - [TMenubutton]
//TODO //   - [TRadiobutton]
//TODO func (w *Window) Textvariable() (r string) {
//TODO 	if tclVar := textVariables[w]; tclVar != "" {
//TODO 		return evalErr(fmt.Sprintf("set %s", tclVar))
//TODO 	}
//TODO
//TODO 	return ""
//TODO }
//TODO
//TODO // Focus — Manage the input focus
//TODO //
//TODO // # Description
//TODO //
//TODO // The focus command is used to manage the Tk input focus. At any given time,
//TODO // one window on each display is designated as the focus window; any key press
//TODO // or key release events for the display are sent to that window. It is
//TODO // normally up to the window manager to redirect the focus among the top-level
//TODO // windows of a display. For example, some window managers automatically set
//TODO // the input focus to a top-level window whenever the mouse enters it; others
//TODO // redirect the input focus only when the user clicks on a window. Usually the
//TODO // window manager will set the focus only to top-level windows, leaving it up
//TODO // to the application to redirect the focus among the children of the
//TODO // top-level.
//TODO //
//TODO // Tk remembers one focus window for each top-level (the most recent descendant
//TODO // of that top-level to receive the focus); when the window manager gives the
//TODO // focus to a top-level, Tk automatically redirects it to the remembered
//TODO // window. Within a top-level Tk uses an explicit focus model by default.
//TODO // Moving the mouse within a top-level does not normally change the focus; the
//TODO // focus changes only when a widget decides explicitly to claim the focus
//TODO // (e.g., because of a button click), or when the user types a key such as Tab
//TODO // that moves the focus.
//TODO //
//TODO // The Tcl procedure tk_focusFollowsMouse may be invoked to create an implicit
//TODO // focus model: it reconfigures Tk so that the focus is set to a window
//TODO // whenever the mouse enters it. The Tcl procedures tk_focusNext and
//TODO // tk_focusPrev implement a focus order among the windows of a top-level; they
//TODO // are used in the default bindings for Tab and Shift-Tab, among other things.
//TODO //
//TODO // The focus command can take any of the following forms:
//TODO //
//TODO //	Focus()
//TODO //
//TODO // Returns the path name of the focus window on the display containing the
//TODO // application's main window, or an empty string if no window in this
//TODO // application has the focus on that display. Note: it is better to specify the
//TODO // display explicitly using -displayof (see below) so that the code will work
//TODO // in applications using multiple displays.
//TODO //
//TODO //	Focus(window)
//TODO //
//TODO // If the application currently has the input focus on window's display, this
//TODO // command resets the input focus for window's display to window and returns an
//TODO // empty string. If the application does not currently have the input focus on
//TODO // window's display, window will be remembered as the focus for its top-level;
//TODO // the next time the focus arrives at the top-level, Tk will redirect it to
//TODO // window. If window is an empty string then the command does nothing.
//TODO //
//TODO //	Focus(Displayof(window))
//TODO //
//TODO // Returns the name of the focus window on the display containing window. If
//TODO // the focus window for window's display is not in this application, the return
//TODO // value is an empty string.
//TODO //
//TODO //	Focus(Force(window))
//TODO //
//TODO // Sets the focus of window's display to window, even if the application does
//TODO // not currently have the input focus for the display. This command should be
//TODO // used sparingly, if at all. In normal usage, an application should not claim
//TODO // the focus for itself; instead, it should wait for the window manager to give
//TODO // it the focus. If window is an empty string then the command does nothing.
//TODO //
//TODO //	Focus(Lastfor(window))
//TODO //
//TODO // Returns the name of the most recent window to have the input focus among all
//TODO // the windows in the same top-level as window. If no window in that top-level
//TODO // has ever had the input focus, or if the most recent focus window has been
//TODO // deleted, then the name of the top-level is returned. The return value is the
//TODO // window that will receive the input focus the next time the window manager
//TODO // gives the focus to the top-level.
//TODO //
//TODO // # Quirks
//TODO //
//TODO // When an internal window receives the input focus, Tk does not actually set
//TODO // the X focus to that window; as far as X is concerned, the focus will stay on
//TODO // the top-level window containing the window with the focus. However, Tk
//TODO // generates FocusIn and FocusOut events just as if the X focus were on the
//TODO // internal window. This approach gets around a number of problems that would
//TODO // occur if the X focus were actually moved; the fact that the X focus is on
//TODO // the top-level is invisible unless you use C code to query the X server
//TODO // directly.
//TODO //
//TODO // More information might be available at the [Tcl/Tk focus] page.
//TODO //
//TODO // [Tcl/Tk focus]: https://www.tcl.tk/man/tcl9.0/TkCmd/focus.htm
//TODO func Focus(options ...Opt) string {
//TODO 	return evalErr(fmt.Sprintf("focus %s", collect(options...)))
//TODO }
//TODO
//TODO // Font represents a Tk font.
//TODO type Font struct {
//TODO 	name string
//TODO }
//TODO
//TODO func (f *Font) optionString(_ *Window) (r string) {
//TODO 	if f != nil {
//TODO 		return f.name
//TODO 	}
//TODO
//TODO 	return "font0" // does not exist
//TODO }
//TODO
//TODO // String implements fmt.Stringer.
//TODO func (f *Font) String() string {
//TODO 	return f.optionString(nil)
//TODO }
//TODO
//TODO // NewFont — Create and inspect fonts.
//TODO //
//TODO // # Description
//TODO //
//TODO // Creates a new font.
//TODO //
//TODO // The following options are supported on all platforms, and are used when
//TODO // creating/specifying a font:
//TODO //
//TODO //   - [Family] name
//TODO //
//TODO // The case-insensitive font family name. Tk guarantees to support the font
//TODO // families named Courier (a monospaced “typewriter” font), Times (a serifed
//TODO // “newspaper” font), and Helvetica (a sans-serif “European” font). The most
//TODO // closely matching native font family will automatically be substituted when
//TODO // one of the above font families is used. The name may also be the name of a
//TODO // native, platform-specific font family; in that case it will work as desired
//TODO // on one platform but may not display correctly on other platforms. If the
//TODO // family is unspecified or unrecognized, a platform-specific default font will
//TODO // be chosen.
//TODO //
//TODO //   - [Size] size
//TODO //
//TODO // The desired size of the font. If the size argument is a positive number, it
//TODO // is interpreted as a size in points. If size is a negative number, its
//TODO // absolute value is interpreted as a size in pixels. If a font cannot be
//TODO // displayed at the specified size, a nearby size will be chosen. If size is
//TODO // unspecified or zero, a platform-dependent default size will be chosen.
//TODO //
//TODO // Sizes should normally be specified in points so the application will remain
//TODO // the same ruler size on the screen, even when changing screen resolutions or
//TODO // moving scripts across platforms. However, specifying pixels is useful in
//TODO // certain circumstances such as when a piece of text must line up with respect
//TODO // to a fixed-size bitmap. The mapping between points and pixels is set when
//TODO // the application starts, based on properties of the installed monitor, but it
//TODO // can be overridden by calling the tk scaling command.
//TODO //
//TODO //   - [Weight] weight
//TODO //
//TODO // The nominal thickness of the characters in the font. The value normal
//TODO // specifies a normal weight font, while bold specifies a bold font. The
//TODO // closest available weight to the one specified will be chosen. The default
//TODO // weight is normal.
//TODO //
//TODO //   - [Slant] slant
//TODO //
//TODO // The amount the characters in the font are slanted away from the vertical.
//TODO // Valid values for slant are roman and italic. A roman font is the normal,
//TODO // upright appearance of a font, while an italic font is one that is tilted
//TODO // some number of degrees from upright. The closest available slant to the one
//TODO // specified will be chosen. The default slant is roman.
//TODO //
//TODO //   - [Underline] boolean
//TODO //
//TODO // The value is a boolean flag that specifies whether characters in this font
//TODO // should be underlined. The default value for underline is false.
//TODO //
//TODO //   - [Overstrike] boolean
//TODO //
//TODO // The value is a boolean flag that specifies whether a horizontal line should
//TODO // be drawn through the middle of characters in this font. The default value
//TODO // for overstrike is false.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk font] page.
//TODO //
//TODO // [Tcl/Tk font]: https://www.tcl.tk/man/tcl9.0/TkCmd/font.htm
//TODO func NewFont(options ...Opt) *Font {
//TODO 	nm := fmt.Sprintf("font%v", id.Add(1))
//TODO 	code := fmt.Sprintf("font create %s %s", nm, collect(options...))
//TODO 	r, err := eval(code)
//TODO 	if err != nil {
//TODO 		fail(fmt.Errorf("code=%s -> r=%s err=%v", code, r, err))
//TODO 		return nil
//TODO 	}
//TODO
//TODO 	return &Font{name: nm}
//TODO }
//TODO
//TODO // FontFamilies — Create and inspect fonts.
//TODO //
//TODO // # Description
//TODO //
//TODO // The return value is a list of the case-insensitive names of all font
//TODO // families that exist on window's display. If the Displayof argument is
//TODO // omitted, it defaults to the main window.
//TODO //
//TODO //   - [Displayof] window
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk font] page.
//TODO //
//TODO // [Tcl/Tk font]: https://www.tcl.tk/man/tcl9.0/TkCmd/font.htm
//TODO func FontFamilies(options ...Opt) []string {
//TODO 	return parseList(evalErr(fmt.Sprintf("font families %s", collect(options...))))
//TODO }
//TODO
//TODO // Can handle "{abc def} ghi" -> ["abc def", "ghi"].
//TODO // Can not handle "{abc {def ghi}} jkl", too many levels.
//TODO func parseList(s string) (r []string) {
//TODO 	r = strings.Fields(s)
//TODO 	w := 0
//TODO 	in := false
//TODO 	var a []string
//TODO 	for _, v := range r {
//TODO 		switch {
//TODO 		case in:
//TODO 			switch {
//TODO 			case strings.HasSuffix(v, "}"):
//TODO 				a = append(a, v[:len(v)-1])
//TODO 				r[w] = strings.Join(a, " ")
//TODO 				w++
//TODO 				in = false
//TODO 			default:
//TODO 				a = append(a, v)
//TODO 			}
//TODO 		default:
//TODO 			switch {
//TODO 			case strings.HasPrefix(v, "{"):
//TODO 				a = append(a[:0], v[1:])
//TODO 				in = true
//TODO 			default:
//TODO 				r[w] = v
//TODO 				w++
//TODO 			}
//TODO 		}
//TODO 	}
//TODO 	return r[:w]
//TODO }
//TODO
//TODO // Delete — Manipulate fonts.
//TODO //
//TODO // # Description
//TODO //
//TODO // Delete the font. If there are widgets using the named font, the named font
//TODO // will not actually be deleted until all the instances are released. Those
//TODO // widgets will continue to display using the last known values for the named
//TODO // font. If a deleted named font is subsequently recreated with another call to
//TODO // font create, the widgets will use the new named font and redisplay
//TODO // themselves using the new attributes of that font.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk font] page.
//TODO //
//TODO // [Tcl/Tk font]: https://www.tcl.tk/man/tcl9.0/TkCmd/font.htm
//TODO func (f *Font) Delete() {
//TODO 	evalErr(fmt.Sprintf("font delete %s", f))
//TODO }
//TODO
//TODO // Text — Create and manipulate 'text' hypertext editing widgets
//TODO //
//TODO // # Description
//TODO //
//TODO // Inserts all of the chars arguments just before the character at index. If
//TODO // index refers to the end of the text (the character after the last newline)
//TODO // then the new text is inserted just before the last newline instead. If there
//TODO // is a single chars argument and no tagList, then the new text will receive
//TODO // any tags that are present on both the character before and the character
//TODO // after the insertion point; if a tag is present on only one of these
//TODO // characters then it will not be applied to the new text. If tagList is
//TODO // specified then it consists of a list of tag names; the new characters will
//TODO // receive all of the tags in this list and no others, regardless of the tags
//TODO // present around the insertion point. If multiple chars-tagList argument pairs
//TODO // are present, they produce the same effect as if a separate pathName insert
//TODO // widget command had been issued for each pair, in order. The last tagList
//TODO // argument may be omitted.
//TODO //
//TODO // The value that is passed to Tcl/Tk for the 'index' argument is obtained by
//TODO // fmt.Sprint(index), enabling any custom index encoding via implementing
//TODO // fmt.Stringer.
//TODO //
//TODO // # Indices
//TODO //
//TODO // Many of the widget commands for texts take one or more indices as arguments.
//TODO // An index is a string used to indicate a particular place within a text, such
//TODO // as a place to insert characters or one endpoint of a range of characters to
//TODO // delete. Indices have the syntax
//TODO //
//TODO //	base modifier modifier modifier ...
//TODO //
//TODO // Where base gives a starting point and the modifiers adjust the index from
//TODO // the starting point (e.g. move forward or backward one character). Every
//TODO // index must contain a base, but the modifiers are optional. Most modifiers
//TODO // (as documented below) allow an optional submodifier. Valid submodifiers are
//TODO // any and display. If the submodifier is abbreviated, then it must be followed
//TODO // by whitespace, but otherwise there need be no space between the submodifier
//TODO // and the following modifier. Typically the display submodifier adjusts the
//TODO // meaning of the following modifier to make it refer to visual or non-elided
//TODO // units rather than logical units, but this is explained for each relevant
//TODO // case below. Lastly, where count is used as part of a modifier, it can be
//TODO // positive or negative, so “base - -3 lines” is perfectly valid (and
//TODO // equivalent to “base +3lines”).
//TODO //
//TODO // The base for an index must have one of the following forms:
//TODO //
//TODO //	"line.char"
//TODO //
//TODO // Indicates char'th character on line line. Lines are numbered from 1 for
//TODO // consistency with other UNIX programs that use this numbering scheme. Within
//TODO // a line, characters are numbered from 0. If char is end then it refers to the
//TODO // newline character that ends the line.
//TODO //
//TODO // This form of index can be passed as [LC]{line, char}.
//TODO //
//TODO //	"@x,y"
//TODO //
//TODO // Indicates the character that covers the pixel whose x and y coordinates
//TODO // within the text's window are x and y.
//TODO //
//TODO //	"end"
//TODO //
//TODO // Indicates the end of the text (the character just after the last newline).
//TODO //
//TODO //	"mark"
//TODO //
//TODO // Indicates the character just after the mark whose name is mark (see MARKS
//TODO // for details).
//TODO //
//TODO //	"tag.first"
//TODO //
//TODO // Indicates the first character in the text that has been tagged with tag.
//TODO // This form generates an error if no characters are currently tagged with tag.
//TODO //
//TODO //	"tag.last"
//TODO //
//TODO // Indicates the character just after the last one in the text that has been
//TODO // tagged with tag. This form generates an error if no characters are currently
//TODO // tagged with tag.
//TODO //
//TODO //	"pathName"
//TODO //
//TODO // Indicates the position of the embedded window whose name is pathName. This
//TODO // form generates an error if there is no embedded window by the given name.
//TODO //
//TODO //	"imageName"
//TODO //
//TODO // Indicates the position of the embedded image whose name is imageName. This
//TODO // form generates an error if there is no embedded image by the given name.
//TODO //
//TODO // If the base could match more than one of the above forms, such as a mark and
//TODO // imageName both having the same value, then the form earlier in the above
//TODO // list takes precedence. If modifiers follow the base index, each one of them
//TODO // must have one of the forms listed below. Keywords such as chars and wordend
//TODO // may be abbreviated as long as the abbreviation is unambiguous.
//TODO //
//TODO //	"+ count ?submodifier? chars"
//TODO //
//TODO // Adjust the index forward by count characters, moving to later lines in the
//TODO // text if necessary. If there are fewer than count characters in the text
//TODO // after the current index, then set the index to the last index in the text.
//TODO // Spaces on either side of count are optional. If the display submodifier is
//TODO // given, elided characters are skipped over without being counted. If any is
//TODO // given, then all characters are counted. For historical reasons, if neither
//TODO // modifier is given then the count actually takes place in units of index
//TODO // positions (see INDICES for details). This behaviour may be changed in a
//TODO // future major release, so if you need an index count, you are encouraged to
//TODO // use indices instead wherever possible.
//TODO //
//TODO //	"- count ?submodifier? chars"
//TODO //
//TODO // Adjust the index backward by count characters, moving to earlier lines in
//TODO // the text if necessary. If there are fewer than count characters in the text
//TODO // before the current index, then set the index to the first index in the text
//TODO // (1.0). Spaces on either side of count are optional. If the display
//TODO // submodifier is given, elided characters are skipped over without being
//TODO // counted. If any is given, then all characters are counted. For historical
//TODO // reasons, if neither modifier is given then the count actually takes place in
//TODO // units of index positions (see INDICES for details). This behavior may be
//TODO // changed in a future major release, so if you need an index count, you are
//TODO // encouraged to use indices instead wherever possible.
//TODO //
//TODO //	"+ count ?submodifier? indices"
//TODO //
//TODO // Adjust the index forward by count index positions, moving to later lines in
//TODO // the text if necessary. If there are fewer than count index positions in the
//TODO // text after the current index, then set the index to the last index position
//TODO // in the text. Spaces on either side of count are optional. Note that an index
//TODO // position is either a single character or a single embedded image or embedded
//TODO // window. If the display submodifier is given, elided indices are skipped over
//TODO // without being counted. If any is given, then all indices are counted; this
//TODO // is also the default behaviour if no modifier is given.
//TODO //
//TODO //	"- count ?submodifier? indices"
//TODO //
//TODO // Adjust the index backward by count index positions, moving to earlier lines
//TODO // in the text if necessary. If there are fewer than count index positions in
//TODO // the text before the current index, then set the index to the first index
//TODO // position (1.0) in the text. Spaces on either side of count are optional. If
//TODO // the display submodifier is given, elided indices are skipped over without
//TODO // being counted. If any is given, then all indices are counted; this is also
//TODO // the default behaviour if no modifier is given.
//TODO //
//TODO //	"+ count ?submodifier? lines"
//TODO //
//TODO // Adjust the index forward by count lines, retaining the same character
//TODO // position within the line. If there are fewer than count lines after the line
//TODO // containing the current index, then set the index to refer to the same
//TODO // character position on the last line of the text. Then, if the line is not
//TODO // long enough to contain a character at the indicated character position,
//TODO // adjust the character position to refer to the last character of the line
//TODO // (the newline). Spaces on either side of count are optional. If the display
//TODO // submodifier is given, then each visual display line is counted separately.
//TODO // Otherwise, if any (or no modifier) is given, then each logical line (no
//TODO // matter how many times it is visually wrapped) counts just once. If the
//TODO // relevant lines are not wrapped, then these two methods of counting are
//TODO // equivalent.
//TODO //
//TODO //	"- count ?submodifier? lines"
//TODO //
//TODO // Adjust the index backward by count logical lines, retaining the same
//TODO // character position within the line. If there are fewer than count lines
//TODO // before the line containing the current index, then set the index to refer to
//TODO // the same character position on the first line of the text. Then, if the line
//TODO // is not long enough to contain a character at the indicated character
//TODO // position, adjust the character position to refer to the last character of
//TODO // the line (the newline). Spaces on either side of count are optional. If the
//TODO // display submodifier is given, then each visual display line is counted
//TODO // separately. Otherwise, if any (or no modifier) is given, then each logical
//TODO // line (no matter how many times it is visually wrapped) counts just once. If
//TODO // the relevant lines are not wrapped, then these two methods of counting are
//TODO // equivalent.
//TODO //
//TODO //	"?submodifier? linestart"
//TODO //
//TODO // Adjust the index to refer to the first index on the line. If the display
//TODO // submodifier is given, this is the first index on the display line, otherwise
//TODO // on the logical line.
//TODO //
//TODO //	"?submodifier? lineend"
//TODO //
//TODO // Adjust the index to refer to the last index on the line (the newline). If
//TODO // the display submodifier is given, this is the last index on the display
//TODO // line, otherwise on the logical line.
//TODO //
//TODO //	"?submodifier? wordstart"
//TODO //
//TODO // Adjust the index to refer to the first character of the word containing the
//TODO // current index. A word consists of any number of adjacent characters that are
//TODO // letters, digits, or underscores, or a single character that is not one of
//TODO // these. If the display submodifier is given, this only examines non-elided
//TODO // characters, otherwise all characters (elided or not) are examined.
//TODO //
//TODO //	"?submodifier? wordend"
//TODO //
//TODO // Adjust the index to refer to the character just after the last one of the
//TODO // word containing the current index. If the current index refers to the last
//TODO // character of the text then it is not modified. If the display submodifier is
//TODO // given, this only examines non-elided characters, otherwise all characters
//TODO // (elided or not) are examined.
//TODO //
//TODO // If more than one modifier is present then they are applied in left-to-right
//TODO // order. For example, the index “end - 1 chars” refers to the next-to-last
//TODO // character in the text and “insert wordstart - 1 c” refers to the character
//TODO // just before the first one in the word containing the insertion cursor.
//TODO // Modifiers are applied one by one in this left to right order, and after each
//TODO // step the resulting index is constrained to be a valid index in the text
//TODO // widget. So, for example, the index “1.0 -1c +1c” refers to the index “2.0”.
//TODO //
//TODO // Where modifiers result in index changes by display lines, display chars or
//TODO // display indices, and the base refers to an index inside an elided tag, that
//TODO // base index is considered to be equivalent to the first following non-elided
//TODO // index.
//TODO //
//TODO // Insert returns its index argument.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk text] page.
//TODO //
//TODO // [Tcl/Tk text]: https://www.tcl.tk/man/tcl9.0/TkCmd/text.htm
//TODO func (w *TextWidget) Insert(index any, chars string, options ...string) any {
//TODO 	idx := fmt.Sprint(index)
//TODO 	evalErr(fmt.Sprintf("%s insert %s %s %s", w, tclSafeString(idx), tclSafeString(chars), tclSafeStrings(options...)))
//TODO 	return index
//TODO }
//TODO
//TODO // LC encodes a text index consisting of a line and char number.
//TODO type LC struct {
//TODO 	Line int // 1-based line number within the text content.
//TODO 	Char int // 0-based char number within the line.
//TODO }
//TODO
//TODO // String implements fmt.Stringer.
//TODO func (lc LC) String() string {
//TODO 	return fmt.Sprintf("%d.%d", lc.Line, lc.Char)
//TODO }
//TODO
//TODO // Text — Create and manipulate 'text' hypertext editing widgets
//TODO //
//TODO // # Description
//TODO //
//TODO // This command is similar to the pathName configure widget command except that
//TODO // it modifies options associated with the tag given by tagName instead of
//TODO // modifying options for the overall text widget. If no option is specified,
//TODO // the command returns a list describing all of the available options for
//TODO // tagName (see Tk_ConfigureInfo for information on the format of this list).
//TODO // If option is specified with no value, then the command returns a list
//TODO // describing the one named option (this list will be identical to the
//TODO // corresponding sublist of the value returned if no option is specified). If
//TODO // one or more option-value pairs are specified, then the command modifies the
//TODO // given option(s) to have the given value(s) in tagName; in this case the
//TODO // command returns an empty string. See TAGS above for details on the options
//TODO // available for tags.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk text] page.
//TODO //
//TODO // [Tcl/Tk text]: https://www.tcl.tk/man/tcl9.0/TkCmd/text.htm
//TODO func (w *TextWidget) TagConfigure(name string, options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("%s tag configure %s %s", w, tclSafeString(name), collect(options...)))
//TODO }
//TODO
//TODO // Text — Create and manipulate 'text' hypertext editing widgets
//TODO //
//TODO // # Description
//TODO //
//TODO // Associate the tag name with all of the characters starting with index1
//TODO // and ending just before index2 (the character at index2 is not tagged). A
//TODO // single command may contain any number of index1-index2 pairs. If the last
//TODO // index2 is omitted then the single character at index1 is tagged. If there
//TODO // are no characters in the specified range (e.g. index1 is past the end of the
//TODO // file or index2 is less than or equal to index1) then the command has no
//TODO // effect.
//TODO //
//TODO // The function returns the new tag name.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk text] page.
//TODO //
//TODO // [Tcl/Tk text]: https://www.tcl.tk/man/tcl9.0/TkCmd/text.htm
//TODO func (w *TextWidget) TagAdd(options ...any) string {
//TODO 	tag := fmt.Sprintf("tag%d", id.Add(1))
//TODO 	var a []Opt
//TODO 	for _, v := range options {
//TODO 		switch x := v.(type) {
//TODO 		case Opt:
//TODO 			a = append(a, x)
//TODO 		case Opts:
//TODO 			a = append(a, x)
//TODO 		default:
//TODO 			a = append(a, stringOption(fmt.Sprint(x)))
//TODO 		}
//TODO 	}
//TODO 	evalErr(fmt.Sprintf("%s tag add %s %s", w, tag, collect(a...)))
//TODO 	return tag
//TODO }
//TODO
//TODO // Fontchooser — control font selection dialog
//TODO //
//TODO // # Description
//TODO //
//TODO // The tk fontchooser command controls the Tk font selection dialog. It uses
//TODO // the native platform font selection dialog where available, or a dialog
//TODO // implemented in Tcl otherwise.
//TODO //
//TODO // Unlike most of the other Tk dialog commands, tk fontchooser does not return
//TODO // an immediate result, as on some platforms (Mac OS X) the standard font
//TODO // dialog is modeless while on others (Windows) it is modal. To get the
//TODO // user-selected font use FontchooserFont() from a handler assigned via
//TODO // [Command].
//TODO //
//TODO // Set one or more of the configurations options below (analogous to Tk widget configuration).
//TODO //
//TODO //   - [Parent]
//TODO //
//TODO // Specifies/returns the logical parent window of the font selection dialog
//TODO // (similar to the -parent option to other dialogs). The font selection dialog
//TODO // is hidden if it is visible when the parent window is destroyed.
//TODO //
//TODO //   - [Title]
//TODO //
//TODO // Specifies/returns the title of the dialog. Has no effect on platforms where
//TODO // the font selection dialog does not support titles.
//TODO //
//TODO //   - [Font]
//TODO //
//TODO // Specifies/returns the font that is currently selected in the dialog if it is
//TODO // visible, or that will be initially selected when the dialog is shown (if
//TODO // supported by the platform). Can be set to the empty string to indicate that
//TODO // no font should be selected. Fonts can be specified in any form given by the
//TODO // "FONT DESCRIPTION" section in the font manual page.
//TODO //
//TODO //   - [Command]
//TODO //
//TODO // Specifies the command called when a font selection has been made by the
//TODO // user. To obtain the font description, call [FontchooserFont] from the
//TODO // handler.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk fontchooser] page.
//TODO //
//TODO // [Tcl/Tk fontchooser]: https://www.tcl.tk/man/tcl9.0/TkCmd/fontchooser.htm
//TODO func Fontchooser(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("tk fontchooser configure %s", collect(options...)))
//TODO }
//TODO
//TODO // FontchooserFont — control font selection dialog
//TODO //
//TODO // # Description
//TODO //
//TODO // Returns the selected font description in the form
//TODO //
//TODO //	family size style...
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk fontchooser] page.
//TODO //
//TODO // [Tcl/Tk fontchooser]: https://www.tcl.tk/man/tcl9.0/TkCmd/fontchooser.htm
//TODO func FontchooserFont() []string {
//TODO 	return parseList(evalErr("tk fontchooser config -font"))
//TODO }
//TODO
//TODO // FontchooserShow — control font selection dialog
//TODO //
//TODO // # Description
//TODO //
//TODO // Show the font selection dialog. Depending on the platform, may return
//TODO // immediately or only once the dialog has been withdrawn.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk fontchooser] page.
//TODO //
//TODO // [Tcl/Tk fontchooser]: https://www.tcl.tk/man/tcl9.0/TkCmd/fontchooser.htm
//TODO func FontchooserShow() {
//TODO 	evalErr("tk  fontchooser show")
//TODO }
//TODO
//TODO // FontchooserHide — control font selection dialog
//TODO //
//TODO // # Description
//TODO //
//TODO // Hide the font selection dialog if it is visible and cause any pending tk
//TODO // fontchooser show command to return.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk fontchooser] page.
//TODO //
//TODO // [Tcl/Tk fontchooser]: https://www.tcl.tk/man/tcl9.0/TkCmd/fontchooser.htm
//TODO func FontchooserHide() {
//TODO 	evalErr("tk fontchooser hide")
//TODO }
//TODO
//TODO // GetOpenFile — pop up a dialog box for the user to select a file to open.
//TODO //
//TODO // # Description
//TODO //
//TODO // GetOpenFile pops up a dialog box for the user to select a file to open. The
//TODO // function is usually associated with the Open command in the File menu. Its
//TODO // purpose is for the user to select an existing file only. If the user enters
//TODO // a non-existent file, the dialog box gives the user an error prompt and
//TODO // requires the user to give an alternative selection. If an application allows
//TODO // the user to create new files, it should do so by providing a separate New
//TODO // menu command.
//TODO //
//TODO //   - [Defaultextension] extension
//TODO //
//TODO // Specifies a string that will be appended to the filename if the user enters
//TODO // a filename without an extension. The default value is the empty string,
//TODO // which means no extension will be appended to the filename in any case. This
//TODO // option is ignored on Mac OS X, which does not require extensions to
//TODO // filenames, and the UNIX implementation guesses reasonable values for this
//TODO // from the -filetypes option when this is not supplied.
//TODO //
//TODO //   - [Filetypes] filePatternList ([][FileType])
//TODO //
//TODO // If a File types listbox exists in the file dialog on the particular
//TODO // platform, this option gives the filetypes in this listbox. When the user
//TODO // choose a filetype in the listbox, only the files of that type are listed. If
//TODO // this option is unspecified, or if it is set to the empty list, or if the
//TODO // File types listbox is not supported by the particular platform then all
//TODO // files are listed regardless of their types. See the section SPECIFYING FILE
//TODO // PATTERNS below for a discussion on the contents of filePatternList.
//TODO //
//TODO //   - [Initialdir] directory
//TODO //
//TODO // Specifies that the files in directory should be displayed when the dialog
//TODO // pops up. If this parameter is not specified, the initial directory defaults
//TODO // to the current working directory on non-Windows systems and on Windows
//TODO // systems prior to Vista. On Vista and later systems, the initial directory
//TODO // defaults to the last user-selected directory for the application. If the
//TODO // parameter specifies a relative path, the return value will convert the
//TODO // relative path to an absolute path.
//TODO //
//TODO //   - [Initialfile] filename
//TODO //
//TODO // Specifies a filename to be displayed in the dialog when it pops up.
//TODO //
//TODO //   - [Multiple] boolean
//TODO //
//TODO // Allows the user to choose multiple files from the Open dialog.
//TODO //
//TODO //   - [Parent] window
//TODO //
//TODO // Makes window the logical parent of the file dialog. The file dialog is
//TODO // displayed on top of its parent window. On Mac OS X, this turns the file
//TODO // dialog into a sheet attached to the parent window.
//TODO //
//TODO //   - [Title] titleString
//TODO //
//TODO // Specifies a string to display as the title of the dialog box. If this option
//TODO // is not specified, then a default title is displayed.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk getopenfile] page.
//TODO //
//TODO // [Tcl/Tk getopenfile]: https://www.tcl.tk/man/tcl9.0/TkCmd/getOpenFile.html
//TODO func GetOpenFile(options ...Opt) []string {
//TODO 	return parseList(evalErr(fmt.Sprintf("tk_getOpenFile %s", collect(options...))))
//TODO }
//TODO
//TODO // FileType specifies a single file type for the [Filetypes] option.
//TODO type FileType struct {
//TODO 	TypeName   string   // Eg. "Go files"
//TODO 	Extensions []string // Eg. []string{".go"}
//TODO 	MacType    string   // Eg. "TEXT"
//TODO }
//TODO
//TODO // GetSaveFile — pop up a dialog box for the user to select a file to save.
//TODO //
//TODO // # Description
//TODO //
//TODO // GetSaveFile pops up a dialog box for the user to select a file to save.
//TODO //
//TODO // The functio is usually associated with the Save as command in the File menu.
//TODO // If the user enters a file that already exists, the dialog box prompts the
//TODO // user for confirmation whether the existing file should be overwritten or
//TODO // not.
//TODO //
//TODO //   - [Confirmoverwrite] boolean
//TODO //
//TODO // Configures how the Save dialog reacts when the selected file already exists,
//TODO // and saving would overwrite it. A true value requests a confirmation dialog
//TODO // be presented to the user. A false value requests that the overwrite take
//TODO // place without confirmation. Default value is true.
//TODO //
//TODO //   - [Defaultextension] extension
//TODO //
//TODO // Specifies a string that will be appended to the filename if the user enters
//TODO // a filename without an extension. The default value is the empty string,
//TODO // which means no extension will be appended to the filename in any case. This
//TODO // option is ignored on Mac OS X, which does not require extensions to
//TODO // filenames, and the UNIX implementation guesses reasonable values for this
//TODO // from the -filetypes option when this is not supplied.
//TODO //
//TODO //   - [Filetypes] filePatternList ([][FileType])
//TODO //
//TODO // If a File types listbox exists in the file dialog on the particular
//TODO // platform, this option gives the filetypes in this listbox. When the user
//TODO // choose a filetype in the listbox, only the files of that type are listed. If
//TODO // this option is unspecified, or if it is set to the empty list, or if the
//TODO // File types listbox is not supported by the particular platform then all
//TODO // files are listed regardless of their types. See the section SPECIFYING FILE
//TODO // PATTERNS below for a discussion on the contents of filePatternList.
//TODO //
//TODO //   - [Initialdir] directory
//TODO //
//TODO // Specifies that the files in directory should be displayed when the dialog
//TODO // pops up. If this parameter is not specified, the initial directory defaults
//TODO // to the current working directory on non-Windows systems and on Windows
//TODO // systems prior to Vista. On Vista and later systems, the initial directory
//TODO // defaults to the last user-selected directory for the application. If the
//TODO // parameter specifies a relative path, the return value will convert the
//TODO // relative path to an absolute path.
//TODO //
//TODO //   - [Initialfile] filename
//TODO //
//TODO // Specifies a filename to be displayed in the dialog when it pops up.
//TODO //
//TODO //   - [Parent] window
//TODO //
//TODO // Makes window the logical parent of the file dialog. The file dialog is
//TODO // displayed on top of its parent window. On Mac OS X, this turns the file
//TODO // dialog into a sheet attached to the parent window.
//TODO //
//TODO //   - [Title] titleString
//TODO //
//TODO // Specifies a string to display as the title of the dialog box. If this option
//TODO // is not specified, then a default title is displayed.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk getopenfile] page.
//TODO //
//TODO // [Tcl/Tk getopenfile]: https://www.tcl.tk/man/tcl9.0/TkCmd/getOpenFile.html
//TODO func GetSaveFile(options ...Opt) string {
//TODO 	return evalErr(fmt.Sprintf("tk_getSaveFile %s", collect(options...)))
//TODO }
//TODO
//TODO // Place — Geometry manager for fixed or rubber-sheet placement
//TODO //
//TODO // # Description
//TODO //
//TODO // The placer is a geometry manager for Tk. It provides simple fixed placement
//TODO // of windows, where you specify the exact size and location of one window,
//TODO // called the content, within another window, called the container. The placer
//TODO // also provides rubber-sheet placement, where you specify the size and
//TODO // location of the content in terms of the dimensions of the container, so that
//TODO // the content changes size and location in response to changes in the size of
//TODO // the container. Lastly, the placer allows you to mix these styles of
//TODO // placement so that, for example, the content has a fixed width and height but
//TODO // is centered inside the container.
//TODO //
//TODO // The first argument must be a *Window.
//TODO //
//TODO // The following options are supported:
//TODO //
//TODO //   - [Anchor] where
//TODO //
//TODO // Where specifies which point of window is to be positioned at the (x,y)
//TODO // location selected by the -x, -y, -relx, and -rely options. The anchor point
//TODO // is in terms of the outer area of window including its border, if any. Thus
//TODO // if where is se then the lower-right corner of window's border will appear at
//TODO // the given (x,y) location in the container. The anchor position defaults to
//TODO // nw.
//TODO //
//TODO //   - [Bordermode] mode
//TODO //
//TODO // Mode determines the degree to which borders within the container are used in
//TODO // determining the placement of the content. The default and most common value
//TODO // is inside. In this case the placer considers the area of the container to be
//TODO // the innermost area of the container, inside any border: an option of -x 0
//TODO // corresponds to an x-coordinate just inside the border and an option of
//TODO // -relwidth 1.0 means window will fill the area inside the container's border.
//TODO //
//TODO // If mode is outside then the placer considers the area of the container to
//TODO // include its border; this mode is typically used when placing window outside
//TODO // its container, as with the options -x 0 -y 0 -anchor ne. Lastly, mode may be
//TODO // specified as ignore, in which case borders are ignored: the area of the
//TODO // container is considered to be its official X area, which includes any
//TODO // internal border but no external border. A bordermode of ignore is probably
//TODO // not very useful.
//TODO //
//TODO //   - [Height] size
//TODO //
//TODO // Size specifies the height for window in screen units (i.e. any of the forms
//TODO // accepted by Tk_GetPixels). The height will be the outer dimension of window
//TODO // including its border, if any. If size is an empty string, or if no -height
//TODO // or -relheight option is specified, then the height requested internally by
//TODO // the window will be used.
//TODO //
//TODO //   - [In] container
//TODO //
//TODO // Container specifies the path name of the window relative to which window is
//TODO // to be placed. Container must either be window's parent or a descendant of
//TODO // window's parent. In addition, container and window must both be descendants
//TODO // of the same top-level window. These restrictions are necessary to guarantee
//TODO // that window is visible whenever container is visible. If this option is not
//TODO // specified then the other window defaults to window's parent.
//TODO //
//TODO //   - [Relheight] size
//TODO //
//TODO // Size specifies the height for window. In this case the height is specified
//TODO // as a floating-point number relative to the height of the container: 0.5
//TODO // means window will be half as high as the container, 1.0 means window will
//TODO // have the same height as the container, and so on. If both -height and
//TODO // -relheight are specified for a content, their values are summed. For
//TODO // example, -relheight 1.0 -height -2 makes the content 2 pixels shorter than
//TODO // the container.
//TODO //
//TODO //   - [Relwidth] size
//TODO //
//TODO // Size specifies the width for window. In this case the width is specified as
//TODO // a floating-point number relative to the width of the container: 0.5 means
//TODO // window will be half as wide as the container, 1.0 means window will have the
//TODO // same width as the container, and so on. If both -width and -relwidth are
//TODO // specified for a content, their values are summed. For example, -relwidth 1.0
//TODO // -width 5 makes the content 5 pixels wider than the container.
//TODO //
//TODO //   - [Relx] location
//TODO //
//TODO // Location specifies the x-coordinate within the container window of the
//TODO // anchor point for window. In this case the location is specified in a
//TODO // relative fashion as a floating-point number: 0.0 corresponds to the left
//TODO // edge of the container and 1.0 corresponds to the right edge of the
//TODO // container. Location need not be in the range 0.0-1.0. If both -x and -relx
//TODO // are specified for a content then their values are summed. For example, -relx
//TODO // 0.5 -x -2 positions the left edge of the content 2 pixels to the left of the
//TODO // center of its container.
//TODO //
//TODO //   - [Rely] location
//TODO //
//TODO // Location specifies the y-coordinate within the container window of the
//TODO // anchor point for window. In this case the value is specified in a relative
//TODO // fashion as a floating-point number: 0.0 corresponds to the top edge of the
//TODO // container and 1.0 corresponds to the bottom edge of the container. Location
//TODO // need not be in the range 0.0-1.0. If both -y and -rely are specified for a
//TODO // content then their values are summed. For example, -rely 0.5 -x 3 positions
//TODO // the top edge of the content 3 pixels below the center of its container.
//TODO //
//TODO //   - [Width] size
//TODO //
//TODO // Size specifies the width for window in screen units (i.e. any of the forms
//TODO // accepted by Tk_GetPixels). The width will be the outer width of window
//TODO // including its border, if any. If size is an empty string, or if no -width or
//TODO // -relwidth option is specified, then the width requested internally by the
//TODO // window will be used.
//TODO //
//TODO //   - [X] location
//TODO //
//TODO // Location specifies the x-coordinate within the container window of the
//TODO // anchor point for window. The location is specified in screen units (i.e. any
//TODO // of the forms accepted by Tk_GetPixels) and need not lie within the bounds of
//TODO // the container window.
//TODO //
//TODO //   - [Y] location
//TODO //
//TODO // Location specifies the y-coordinate within the container window of the
//TODO // anchor point for window. The location is specified in screen units (i.e. any
//TODO // of the forms accepted by Tk_GetPixels) and need not lie within the bounds of
//TODO // the container window.
//TODO //
//TODO // If the same value is specified separately with two different options, such
//TODO // as -x and -relx, then the most recent option is used and the older one is
//TODO // ignored.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk place] page.
//TODO //
//TODO // [Tcl/Tk place]: https://www.tcl.tk/man/tcl9.0/TkCmd/place.htm
//TODO func Place(options ...Opt) {
//TODO 	evalErr(fmt.Sprintf("place %s", collect(options...)))
//TODO }
//TODO
//TODO // Lower — Change a window's position in the stacking order
//TODO //
//TODO // # Description
//TODO //
//TODO // If the belowThis argument is nil then the command lowers window so that
//TODO // it is below all of its siblings in the stacking order (it will be obscured
//TODO // by any siblings that overlap it and will not obscure any siblings). If
//TODO // belowThis is specified then it must be the path name of a window that is
//TODO // either a sibling of window or the descendant of a sibling of window. In this
//TODO // case the lower command will insert window into the stacking order just below
//TODO // belowThis (or the ancestor of belowThis that is a sibling of window); this
//TODO // could end up either raising or lowering window.
//TODO //
//TODO // All toplevel windows may be restacked with respect to each other, whatever
//TODO // their relative path names, but the window manager is not obligated to
//TODO // strictly honor requests to restack.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk lower] page.
//TODO //
//TODO // [Tcl/Tk lower]: https://www.tcl.tk/man/tcl9.0/TkCmd/lower.html
//TODO func (w *Window) Lower(belowThis Widget) {
//TODO 	b := ""
//TODO 	if belowThis != nil {
//TODO 		b = belowThis.path()
//TODO 	}
//TODO 	evalErr(fmt.Sprintf("lower %s %s", w, b))
//TODO }
//TODO
//TODO // Raise — Change a window's position in the stacking order
//TODO //
//TODO // # Description
//TODO //
//TODO // If the aboveThis argument is nil then the command raises window so that it
//TODO // is above all of its siblings in the stacking order (it will not be obscured
//TODO // by any siblings and will obscure any siblings that overlap it). If aboveThis
//TODO // is specified then it must be the path name of a window that is either a
//TODO // sibling of window or the descendant of a sibling of window. In this case the
//TODO // raise command will insert window into the stacking order just above
//TODO // aboveThis (or the ancestor of aboveThis that is a sibling of window); this
//TODO // could end up either raising or lowering window.
//TODO //
//TODO // All toplevel windows may be restacked with respect to each other, whatever
//TODO // their relative path names, but the window manager is not obligated to
//TODO // strictly honor requests to restack.
//TODO //
//TODO // On macOS raising an iconified toplevel window causes it to be deiconified.
//TODO //
//TODO // Additional information might be available at the [Tcl/Tk raise] page.
//TODO //
//TODO // [Tcl/Tk raise]: https://www.tcl.tk/man/tcl9.0/TkCmd/raise.html
//TODO func (w *Window) Raise(aboveThis Widget) {
//TODO 	b := ""
//TODO 	if aboveThis != nil {
//TODO 		b = aboveThis.path()
//TODO 	}
//TODO 	evalErr(fmt.Sprintf("raise %s %s", w, b))
//TODO }

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

// Menu — Create and manipulate 'menu' widgets and menubars
//
// # Description
//
// Add a new command entry to the bottom of the menu.
//
// Additional information might be available at the [Tcl/Tk menu] page.
//
// [Tcl/Tk menu]: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.htm
func (w *MenuWidget) AddCommand(options ...Opt) {
	evalErr(fmt.Sprintf("%s add command %s", w, winCollect(w.Window, options...)))
}

// Menu — Create and manipulate 'menu' widgets and menubars
//
// # Description
//
// Add a new cascade entry to the end of the menu.
//
// Additional information might be available at the [Tcl/Tk menu] page.
//
// [Tcl/Tk menu]: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.htm
func (w *MenuWidget) AddCascade(options ...Opt) {
	evalErr(fmt.Sprintf("%s add cascade %s", w, winCollect(w.Window, options...)))
}

// Menu — Create and manipulate 'menu' widgets and menubars
//
// # Description
//
// Add a new separator entry to the bottom of the menu.
//
// Additional information might be available at the [Tcl/Tk menu] page.
//
// [Tcl/Tk menu]: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.htm
func (w *MenuWidget) AddSeparator(options ...Opt) {
	evalErr(fmt.Sprintf("%s add separator %s", w, winCollect(w.Window, options...)))
}

// Menu — Create and manipulate 'menu' widgets and menubars
//
// # Description
//
// Invoke the action of the menu entry. See the sections on the individual
// entries above for details on what happens. If the menu entry is disabled
// then nothing happens. If the entry has a command associated with it then the
// result of that command is returned as the result of the invoke widget
// command. Otherwise the result is an empty string. Note: invoking a menu
// entry does not automatically unpost the menu; the default bindings normally
// take care of this before invoking the invoke widget command.
//
// Additional information might be available at the [Tcl/Tk menu] page.
//
// [Tcl/Tk menu]: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.htm
func (w *MenuWidget) Invoke(index uint) {
	evalErr(fmt.Sprintf("%s invoke %d", w, index))
}
