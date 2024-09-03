// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package tk9_0 // import "modernc.org/tk9.0"

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/evilsocket/islazy/zip"
	libtcl "modernc.org/libtcl9.0"
	libtk "modernc.org/libtk9.0"
	tklib "modernc.org/libtk9.0/library"
	tcl "modernc.org/tcl9.0"
)

var (
	interp *tcl.Interp
	tclDir string
	tkDir  string
)

func init() {
	if isBuilder {
		return
	}

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

	if Error = interp.RegisterCommand("eventDispatcher", eventDispatcher, nil, nil); Error == nil {
		CollectErrors = true

		defer func() { CollectErrors = false }()

		App = &Window{}
		exitHandler = Command(func() { Destroy(App) })
		// Set some defaults.
		evalErr("option add *tearOff 0") // https://tkdocs.com/tutorial/menus.html
		App.Center()
		App.IconPhoto(NewPhoto(Data(icon)))
		App.WmTitle(filepath.Base(os.Args[0]))
	}
}

func eval(code string) (r string, err error) {
	if dmesgs {
		defer func() {
			dmesg("code=%s -> r=%v err=%v", code, r, err)
		}()
	}
	return interp.Eval(code, tcl.EvalDirect)
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
	for _, v := range append([]string{tclDir, tkDir}, cleanupDirs...) {
		err = errors.Join(err, os.RemoveAll(v))
	}
	return err
}
