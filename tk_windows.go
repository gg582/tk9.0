// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"unsafe"

	"github.com/evilsocket/islazy/zip"
	"golang.org/x/sys/windows"
)

var (
	//go:embed embed_windows/tk_library.zip
	tkLibrary []byte

	evalExProc       *windows.Proc
	getObjResultProc *windows.Proc
	getStringProc    *windows.Proc
	tclDll           *windows.DLL
	tkDll            *windows.DLL
)

func init() {
	if isBuilder {
		return
	}

	initOnce.Do(func() {
		runtime.LockOSThread()
		var cacheDir string
		if cacheDir, Error = getCacheDir(); Error != nil {
			return
		}

		var wd string
		if wd, Error = os.Getwd(); Error != nil {
			return
		}

		defer func() {
			Error = errors.Join(Error, os.Chdir(wd))
		}()

		if Error = os.Chdir(cacheDir); Error != nil {
			return
		}

		if tclDll, Error = windows.LoadDLL("tcl90.dll"); Error != nil {
			return
		}

		if tkDll, Error = windows.LoadDLL("tcl9tk90.dll"); Error != nil {
			return
		}

		var tclCreateInterp, tclInit, tkInit *windows.Proc
		if tclCreateInterp, Error = tclDll.FindProc("Tcl_CreateInterp"); Error != nil {
			return
		}

		if tclInit, Error = tclDll.FindProc("Tcl_Init"); Error != nil {
			return
		}

		if evalExProc, Error = tclDll.FindProc("Tcl_EvalEx"); Error != nil {
			return
		}

		if getObjResultProc, Error = tclDll.FindProc("Tcl_GetObjResult"); Error != nil {
			return
		}

		if getStringProc, Error = tclDll.FindProc("Tcl_GetString"); Error != nil {
			return
		}

		if tkInit, Error = tkDll.FindProc("Tk_Init"); Error != nil {
			return
		}

		if interp, _, _ = tclCreateInterp.Call(); interp == 0 {
			Error = fmt.Errorf("failed to create a Tcl interpreter")
			return
		}

		if r, _, _ := tclInit.Call(interp); r != tcl_ok {
			Error = fmt.Errorf("failed to initialize the Tcl interpreter")
			return
		}

		s, err := tclEvalEx(fmt.Sprintf("zipfs mount %s /lib/tk", tclSafeString(filepath.Join(cacheDir, "lib.zip"))))
		trc("---- MOUNT err=%v\n%s\n---", err, s)
		s, err = tclEvalEx("zipfs list")
		trc("---- LIST err=%v\n%s\n---", err, s)

		if r, r2, err := tkInit.Call(interp); r != tcl_ok {
			trc("r=%0x r2=%#0x err=%v res=%q", r, r2, err, tclResult())
			Error = fmt.Errorf("failed to initialize Tk")
			return
		}

		// r0, r1, err := tclInit.Call(uintptr(unsafe.Pointer(&interp)))
		// trc("interp=%#0x r0=%v r1=%v err=%v", interp, r0, r1, err)

		// tclDll = windows.NewLazyDLL(filepath.Join(cacheDir, "tcl90.dll"))
		// trc("A %v", tclDll.Load())

		// tkDll = windows.NewLazyDLL(filepath.Join(cacheDir, "tcl9tk90.dll"))
		// trc("B %v", tkDll.Load())

		// if tclDir, Error = tcl.Stdlib(); Error != nil {
		// 	return
		// }

		// if tkDir, Error = stdlib(); Error != nil {
		// 	return
		// }

		// if interp, Error = tcl.NewInterp(map[string]string{
		// 	"tcl_library": tclDir,
		// 	"tk_library":  tkDir,
		// }); Error != nil {
		// 	return
		// }

		// if rc := libtk.XTk_Init(interp.TLS(), interp.Handle()); rc != libtk.TCL_OK {
		// 	interp.Close()
		// 	Error = fmt.Errorf("failed to initialize the Tk subsystem")
		// 	return
		// }

		// Error = interp.RegisterCommand("eventDispatcher", eventDispatcher, nil, nil)
	})
	trc("Error=%v", Error)
	s, err := tclEvalEx(`
image create photo img -file _examples/gopher.png
label .l -image img
ttk::button .b -text Exit -command { destroy . }
pack .l .b -padx 1m -pady 2m -ipadx 1m -ipady 1m
configure . -padx 4m -pady 3m
tkwait window .
		`)
	trc("---- WAIT err=%v\n%s\n---", err, s)
	//TODO if Error == nil {
	//TODO 	CollectErrors = true

	//TODO 	defer func() { CollectErrors = false }()

	//TODO 	App = &Window{}
	//TODO 	exitHandler = Command(func() { Destroy(App) })
	//TODO 	// Set some defaults.
	//TODO 	evalErr("option add *tearOff 0") // https://tkdocs.com/tutorial/menus.html
	//TODO 	App.Center()
	//TODO 	App.IconPhoto(NewPhoto(Data(icon)))
	//TODO 	App.WmTitle(filepath.Base(os.Args[0]))
	//TODO }
}

func getCacheDir() (r string, err error) {
	if r, err = os.UserCacheDir(); err != nil {
		return "", err
	}

	r0 := filepath.Join(r, "modernc.org")
	r = filepath.Join(r0, "tk9.0")
	fi, err := os.Stat(r)
	if err == nil && fi.IsDir() {
		return r, nil
	}

	err = os.MkdirAll(r0, 0700)
	trc("%s -> %v", r, err)
	tmp, err := os.MkdirTemp("", "tk9.0-")
	if err != nil {
		return "", err
	}

	zf := filepath.Join(tmp, "dll.zip")
	if err = os.WriteFile(zf, dlls, 0660); err != nil {
		return "", err
	}

	if _, err = zip.Unzip(zf, tmp); err != nil {
		os.Remove(zf)
		return "", err
	}

	os.Remove(zf)
	zf = filepath.Join(tmp, "lib.zip")
	if err = os.WriteFile(zf, tkLibrary, 0660); err != nil {
		return "", err
	}

	if err = os.Rename(tmp, r); err == nil {
		trc("OK %s -> %s", tmp, r)
		return r, nil
	}

	trc("NOT OK: %s %v", tmp, err)
	cleanupDirs = append(cleanupDirs, tmp)
	return tmp, nil
}

func tclResult() string {
	r, _, _ := getObjResultProc.Call(interp)
	if r == 0 {
		return ""
	}

	if r, _, _ = getStringProc.Call(r); r != 0 {
		r0 := r
		var n int
		for ; *(*byte)(unsafe.Pointer(r)) != 0; n++ {
			r++
		}
		if n != 0 {
			return unsafe.String((*byte)(unsafe.Pointer(r0)), n)
		}
	}

	return ""
}

func tclEvalEx(s string) (r string, err error) {
	trc("EVAL `%s`", s)
	if r0, _, _ := evalExProc.Call(interp, uintptr(unsafe.Pointer(unsafe.StringData(s))), uintptr(len(s)), tcl_eval_direct); r0 == tcl_ok {
		return tclResult(), nil
	}

	return "", fmt.Errorf("%s", tclResult())
}
