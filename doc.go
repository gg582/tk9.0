// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tk9.0 is a CGo-free cross platform GUI for Go. It is similar to
// [tkinter] for Python.
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
//   - [Img.Graph] and [CanvasWidget.Graph] require the gnuplot 5.4+ executable
//     available in $PATH.
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
// package exports the widget as type 'TextWidget', its constructor as function
// 'Text' and the option as function 'Txt'. The complete list is:
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
// [tkinter]: https://en.wikipedia.org/wiki/Tkinter
// [tutorialspoint]: https://www.tutorialspoint.com/tcl-tk/tk_overview.htm
package tk9_0 // import "modernc.org/tk9.0"
