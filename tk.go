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

// dmesgs cc/v4 style
// error modes: panic (default), collect, ignore

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
type Tk struct {
	*Window
	in *tcl.Interp

	trace bool
}

func (tk *Tk) eval(code string) (r string, err error) {
	if tk.trace {
		defer func() {
			fmt.Fprintf(os.Stderr, "code=%s -> (r=%v err=%v)\n", code, r, err)
		}()
	}
	return tk.in.Eval(code, tcl.EvalGlobal)
}

// Initialize performs package initialization and returns a *Tk or error, if
// any.
//
// The returned value is a singleton. Calls to Initialize() are idempotent and
// all return the same (instance, error) tuple.
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

func (w *Window) newChild(nm string, opts ...Opt) (*Window, error) {
	cls := strings.Replace(nm, "ttk_", "ttk::", 1)
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
		return nil, fmt.Errorf("%v %v", err, r)
	}

	return &Window{fpath: r}, nil
}

// https://pdos.csail.mit.edu/archive/rover/RoverDoc/escape_shell_table.html
//
// The following characters are dissallowed or have special meanings in Tcl and
// so are escaped:
//
//	&;`'"|*?~<>^()[]{}$\
func tclSafeString(s string) string {
	//TODO const chars = "\\&;`'\"|*?~<>^()[]{}$"
	//TODO if strings.ContainsAny(s, chars) {
	//TODO 	for i := 0; i < len(chars); i++ {
	//TODO 		c := chars[i : i+1]
	//TODO 		s = strings.ReplaceAll(s, c, `\`+c)
	//TODO 	}
	//TODO }
	//TODO if strings.ContainsAny(s, "\n\r\t ") {
	//TODO 	s = "{" + s + "}"
	//TODO }
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

// Txt produces the '-text s' option. It is not named 'Text' because of the
// Text() function.
func Txt(s string) Opt {
	return text(s)
}
