// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	"flag"
	"fmt"
	"os"
	"testing"

	_ "github.com/adrg/xdg"   // generator.go
	_ "golang.org/x/net/html" // generator.go
	_ "modernc.org/ngrab/lib" // generator.go
)

func TestMain(m *testing.M) {
	if isBuilder {
		os.Exit(0)
	}

	if Error != nil {
		fmt.Fprintln(os.Stderr, Error)
		os.Exit(1)
	}

	flag.Parse()
	rc := m.Run()
	Finalize()
	os.Exit(rc)
}

// func Test(t *testing.T) {
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
// }

// func Test2(t *testing.T) {
// 	b := TButton(Txt("Save"), Underline(0), State("disabled"))
// 	Pack(
// 		Label(Image(ImageCreatePhoto(File("testdata/gimp.png")))),
// 		b,
// 		Label(Txt(fmt.Sprintf("button above .Txt=%s", b.Cget(Txt)))),
// 		TButton(Txt("Close"), Underline(0), Command(func() { Destroy(App) })),
// 		Ipadx(30),
// 	)
// 	App.Configure(Pady(30))
// 	fmt.Println(App.Configure())
// 	eval(`
// ttk::style theme use clam
// tk::PlaceWindow . center
// wm title . "modernc.org/tk9.0 example"
// wm geometry . 300x[winfo height .]
// tkwait window .
// `)
// }

func Test3(t *testing.T) {
	if _, err := eval(`

# ttk::style theme use clam
# . configure -pady 30
# image create photo gopher -file testdata/go.png
# label .l -image gopher
# ttk::button .save -text Save -state disabled -underline 0
# label .l2 -text "button above .Txt=[.save cget -text]"
# ttk::button .close -text Close -command { destroy . } -underline 0
# pack .l .save .l2 .close -ipadx 30
# tk::PlaceWindow . center
# wm title . "modernc.org/tk9.0 example"
# wm geometry . 500x[winfo height .]
# tkwait window .

button .b -text hello -command { destroy . }
pack .b
tkwait window .

`,
	); err != nil {
		t.Fatal(err)
	}
}
