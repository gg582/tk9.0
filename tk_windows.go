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
	"modernc.org/memory"
)

var (
	//go:embed embed_windows/tk_library.zip
	tkLibrary []byte

	// No mutex, the package must be used by a single goroutine only.
	allocator        memory.Allocator
	evalExProc       *windows.Proc
	getObjResultProc *windows.Proc
	getStringProc    *windows.Proc
	newStringObjProc *windows.Proc
	setObjResultProc *windows.Proc
	tclDll           *windows.DLL
	tkDll            *windows.DLL
)

func init() {
	if isBuilder {
		return
	}

	runtime.LockOSThread()
	var cacheDir string
	if cacheDir, Error = getCacheDir(); Error != nil {
		return
	}

	var wd string
	if wd, Error = os.Getwd(); Error != nil {
		return
	}

	func() {
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

		if setObjResultProc, Error = tclDll.FindProc("Tcl_SetObjResult"); Error != nil {
			return
		}

		if getObjResultProc, Error = tclDll.FindProc("Tcl_GetObjResult"); Error != nil {
			return
		}

		if getStringProc, Error = tclDll.FindProc("Tcl_GetString"); Error != nil {
			return
		}

		if newStringObjProc, Error = tclDll.FindProc("Tcl_NewStringObj"); Error != nil {
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

		if _, Error := eval("zipfs mount tk_library.zip /lib/tk"); Error != nil {
			return
		}

		if r, _, _ := tkInit.Call(interp); r != tcl_ok {
			Error = fmt.Errorf("failed to initialize Tk")
			return
		}
	}()

	s, err := eval(`
image create photo img -file _examples/gopher.png
label .l -image img
ttk::button .b -text Exit -command { destroy . }
pack .l .b -padx 1m -pady 2m -ipadx 1m -ipady 1m
. configure -padx 4m -pady 3m
tkwait window .`,
	)
	trc("s=%s, err=%v", s, err)

	// if Error = interp.RegisterCommand("eventDispatcher", eventDispatcher, nil, nil); Error == nil {
	// 	CollectErrors = true

	// 	defer func() { CollectErrors = false }()

	// 	panic(todo(""))
	// 	//TODO App = &Window{}
	// 	//TODO exitHandler = Command(func() { Destroy(App) })
	// 	//TODO // Set some defaults.
	// 	//TODO evalErr("option add *tearOff 0") // https://tkdocs.com/tutorial/menus.html
	// 	//TODO App.Center()
	// 	//TODO App.IconPhoto(NewPhoto(Data(icon)))
	// 	//TODO App.WmTitle(filepath.Base(os.Args[0]))
	// }
}

func getCacheDir() (r string, err error) {
	defer func() { trc("168:->(r=%s err=%v)", r, err) }()
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
	zf = filepath.Join(tmp, "tk_library.zip")
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

func eval(code string) (r string, err error) {
	if dmesgs {
		defer func() {
			dmesg("code=%s -> r=%v err=%v", code, r, err)
		}()
	}
	cs, err := cString(code)
	if err != nil {
		return "", err
	}

	defer allocator.UintptrFree(cs)

	if r0, _, _ := evalExProc.Call(interp, cs, uintptr(len(code)), tcl_eval_direct); r0 == tcl_ok {
		return tclResult(), nil
	}

	return "", fmt.Errorf("%s", tclResult())
}

func cString(s string) (r uintptr, err error) {
	if s == "" {
		return 0, nil
	}

	if r, err = allocator.UintptrMalloc(len(s) + 1); err != nil {
		return 0, err
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(r)), len(s)), s)
	*(*byte)(unsafe.Add(unsafe.Pointer(r), len(s))) = 0
	return r, nil
}

func setResult(s string) (err error) {
	cs, err := cString(s)
	if err != nil {
		return err
	}

	defer allocator.UintptrFree(cs)

	obj, _, _ := newStringObjProc.Call(cs, uintptr(len(s)))
	if obj == 0 {
		return fmt.Errorf("OOM")
	}

	setObjResultProc.Call(interp, obj)
	return nil
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
