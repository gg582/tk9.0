// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	"flag"
	"fmt"
	"os"
	"testing"
	// tcl "modernc.org/tcl9.0"
)

var (
	isBuilder = os.Getenv("MODERNC_BUILDER") != ""
)

func TestMain(m *testing.M) {
	if isBuilder {
		os.Exit(0)
	}

	flag.Parse()
	rc := m.Run()
	Finalize()
	os.Exit(rc)
}

func Test(t *testing.T) {
	//	in, _ := Initialize()
	//	if _, err := in.in.Eval(`
	//
	// ttk::style theme use clam
	// . configure -pady 10
	// ttk::label .l -text "Hello, World!" -padding 3
	// ttk::spinbox .s -from 1 -to 10
	// ttk::button .b -text "Close dialog and exit" -command { destroy . }
	// pack .l .s .b
	// tk::PlaceWindow . center
	// wm title . "modernc.org/tk9.0 example"
	// wm geometry . 300x[winfo height .]
	// tkwait window .
	//
	// `,
	//
	//		tcl.EvalDirect,
	//	); err != nil {
	//		t.Fatal(err)
	//	}
}

func Test2(t *testing.T) {
	tk, err := Initialize()
	if err != nil {
		t.Fatal(err)
	}

	tk.trace = true
	b := TButton(Txt("{[Close\nMe]}"))
	tk.eval(fmt.Sprintf(`
ttk::style theme use clam
. configure -pady 10
pack %s
tk::PlaceWindow . center
wm title . "modernc.org/tk9.0 example"
wm geometry . 300x[winfo height .]
tkwait window .

`, b.path()))
	trc("", b)
}
