// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	"flag"
	"fmt"
	"os"
	"testing"

	tcl "modernc.org/tcl9.0"
)

var (
	isBuilder = os.Getenv("MODERNC_BUILDER") != ""
)

func TestMain(m *testing.M) {
	if isBuilder {
		os.Exit(0)
	}

	flag.Parse()
	os.Exit(m.Run())
}

// Configuring a new Tcl interpreter to use the Tcl/Tk standard libraries.
// Error handling omitted.
func Example() {
	in, _ := NewInterp(map[string]string{
		"tcl_library": tcl.MustStdlib(),
		"tk_library":  MustStdlib(),
	})
	fmt.Println(in.Eval("package require http;", tcl.EvalGlobal))
	fmt.Println(in.Eval("package require Tk;", tcl.EvalGlobal))
	// Output:
	// 2.10b1 <nil>
	// 9.0b1 <nil>
}
