// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tk9.0 is an idiomatic Go wrapper for [libtk9.0].
//
// [libtk9.0]: https://pkg.go.dev/modernc.org/libtk9.0
package tk9_0 // import "modernc.org/tk9.0"

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/evilsocket/islazy/zip"
	lib "modernc.org/libtk9.0"
	tklib "modernc.org/libtk9.0/library"
	tcl "modernc.org/tcl9.0"
)

var (
	onceStdlib    sync.Once
	onceStdlibErr error
	stdlib        string
)

// Stdlib returns the path to the Tk standard library or an error, if any. It
// once creates a temporary directory where the standard library is written.
// Subsequent calls to Stdlib share the same temporary directory.
//
// Stdlib is safe for concurrent access by multiple goroutines.
func Stdlib() (string, error) {
	onceStdlib.Do(func() {
		dir, err := os.MkdirTemp("", "tk-library-")
		defer func() { onceStdlibErr = err }()
		if err != nil {
			return
		}

		fn := filepath.Join(dir, "library.zip")
		if err = os.WriteFile(fn, []byte(tklib.Zip), 0600); err != nil {
			return
		}

		if _, err = zip.Unzip(fn, dir); err != nil {
			return
		}

		stdlib = filepath.Join(dir, "library")
	})
	return stdlib, onceStdlibErr
}

// MustStdlib is like Stdlib but panics on error.
func MustStdlib() (r string) {
	r, err := Stdlib()
	if err != nil {
		panic(err)
	}

	return r
}

// NewInterp is like [tcl.NewInterp] but additionally initializes the Tk
// subsystem.
//
// [tcl.NewInterp]: https://pkg.go.dev/modernc.org/tcl9.0#NewInterp
func NewInterp(tclvars map[string]string) (r *tcl.Interp, err error) {
	if r, err = tcl.NewInterp(tclvars); err != nil {
		return nil, err
	}

	if rc := lib.XTk_Init(r.TLS(), r.Handle()); rc != lib.TCL_OK {
		r.Close()
		return nil, fmt.Errorf("failed to initialize the Tk subsystem")
	}

	return r, nil
}
