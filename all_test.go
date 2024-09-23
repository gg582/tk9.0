// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package tk9_0 // import "modernc.org/tk9.0"

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"testing"

	_ "github.com/adrg/xdg"       // generator.go
	_ "github.com/expr-lang/expr" // examples
	_ "golang.org/x/net/html"     // generator.go
	_ "modernc.org/ngrab/lib"     // generator.go
	_ "modernc.org/rec/lib"       // generator.go
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

func TestTokenizer(t *testing.T) {
	for i, test := range []struct {
		s    string
		ids  []int
		toks []string
	}{
		{},
		{"a", []int{0}, []string{"a"}},
		{"\\$", []int{0}, []string{"\\$"}},
		{"\\$\\$", []int{0}, []string{"\\$\\$"}},
		{"\\$\\$\\$", []int{0}, []string{"\\$\\$\\$"}},

		{"\\$\\$\\$\\$", []int{0}, []string{"\\$\\$\\$\\$"}},
		{"a\\$", []int{0}, []string{"a\\$"}},
		{"a\\$\\$", []int{0}, []string{"a\\$\\$"}},
		{"a\\$\\$\\$", []int{0}, []string{"a\\$\\$\\$"}},
		{"a\\$\\$\\$\\$", []int{0}, []string{"a\\$\\$\\$\\$"}},

		{"$a$", []int{1}, []string{"$a$"}},
		{"$$a$", []int{2}, []string{"$$a$"}},
		{"$$a$$", []int{2}, []string{"$$a$$"}},
		{"$a$$", []int{2}, []string{"$a$$"}},
		{"x$a$", []int{0, 1}, []string{"x", "$a$"}},

		{"x$$a$", []int{0, 2}, []string{"x", "$$a$"}},
		{"x$$a$$", []int{0, 2}, []string{"x", "$$a$$"}},
		{"x$a$$", []int{0, 2}, []string{"x", "$a$$"}},
		{"x$a$y", []int{0, 1, 0}, []string{"x", "$a$", "y"}},
		{"x$$a$y", []int{0, 2, 0}, []string{"x", "$$a$", "y"}},

		{"x$$a$$y", []int{0, 2, 0}, []string{"x", "$$a$$", "y"}},
		{"x$a$$y", []int{0, 2, 0}, []string{"x", "$a$$", "y"}},
		{"x\\$0$a\\$1b$$\\$y", []int{0, 2, 0}, []string{"x\\$0", "$a\\$1b$$", "\\$y"}},
	} {
		ids, toks := tokenize(test.s)
		if g, e := fmt.Sprintf("%v %q", ids, toks), fmt.Sprintf("%v %q", test.ids, test.toks); g != e {
			t.Errorf("#%3v: `%s`\ngot %s\nexp %s", i, test.s, g, e)
		}
	}
}

func TestTmp(t *testing.T) {
	// trc("", StyleLayout("TButton"))
	// 	Button.border -sticky nswe -border 1 -children {
	//		Button.focus -sticky nswe -children {
	//			Button.padding -sticky nswe -children {
	//				Button.label -sticky nswe}}}

	trc("", evalErr(`ttk::style element options TButton.border`))
	// TRC -background -borderwidth -relief

	trc("", evalErr(`ttk::style element options TButton.focus`))
	// TRC -focuscolor -focusthickness -focussolid

	trc("", evalErr(`ttk::style element options TButton.padding`))
	// TRC -padding -relief -shiftrelief

	trc("", evalErr(`ttk::style element options TButton.label`))
	// TRC -compound -space -text -font -foreground -underline -width -anchor -justify -wraplength -embossed -image -stipple -background

	trc("", evalErr(`ttk::style map TButton`))
	// TRC -relief {{!disabled pressed} sunken}

	trc("", evalErr(`ttk::style map TButton.border`))
	// TRC
	trc("", evalErr(`ttk::style map TButton.focus`))
	// TRC
	trc("", evalErr(`ttk::style map TButton.padding`))
	// TRC
	trc("", evalErr(`ttk::style map TButton.label`))
	// TRC

	// ttk::style configure primary.TButton -foreground #ffffff -background #4582ec -bordercolor #4582ec -darkcolor #4582ec -lightcolor #4582ec -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5} -anchor center
	// ttk::style map primary.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed} -lightcolor {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed}

	// 	const svg = `
	// <svg width="22cm" height="17cm">
	// 	<rect x="100" y="100" width="200" height="150" fill="#fff" stroke-width="10" stroke="#c2dbfe" rx="10" />
	// 	<rect x="400" y="100" width="200" height="150" fill="#fff" stroke-width="10" stroke="#c2dbfe" rx="20" />
	// 	<rect x="100" y="350" width="200" height="150" fill="#fff" stroke-width="10" stroke="#c2dbfe" rx="30" />
	// 	<rect x="400" y="350" width="200" height="150" fill="#fff" stroke-width="10" stroke="#c2dbfe" rx="40" />
	// </svg>`

	// Pack(
	// 	Label(Background("#fff"), Image(NewPhoto(Data(svg)))),
	// 	Button(Txt("foo")),
	// 	TButton(Txt("bar")),
	// 	TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
	// )
	// StyleConfigure("TButton", Borderwidth(10))
	// StyleConfigure("TButton", FocusColor(Blue))
	// StyleConfigure("TButton", FocusThickness(10))
	// StyleConfigure("TButton", FocusSolid(true))
	// App.Configure(Background("#fff")).Center().Wait()
}

/*

image create photo pyimage1 -data iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAFxEAABcRAcom8z8AAAT/SURBVFhHzZd9TNR1HMff/O64BwKOp+NRnkIScJI2n7ZqFmqjVhkrHExrrDFZLnPkqmEyW2WtlQ1SQawWPSAYiJLNGvK0VhJOaDgVBDPCw0PwEOUQfsfd/fp+v/cFbt0Bd6yxXvfP5+F79/l+P9+H+3w8wMnO3r40I+P5vOCQ4I1ymVzrQeCu/wSJYDabB/sHbp45XlnzQUnJgcvcBRQUHEjX6/vH6aCFQK/XjxcWHnyBxpbRle/Ysf3n0NAQJTWMmiRUnB/FNy1GNHaNYdwsIS7IE4JdPqxWKxoam3D0aAVq6+rR39+P2NhYKBQKPsJGx3Abyq8dxGldBbrvXkSoVyR8Pf3g7e0tj46O2mQ0Gqs96urqv1u/PmUL/cLfQ2a8/O0t9A2b2Q9MsjpGicOZQfBSeEAURbyZtxstLee410awVovCgk8RGxPD9NLuT1B6dT+TJ/EUFMhL/gwpYZuYXl/fUCbQPaeKVQJyqwwOwSnnekR8VDvM5JIjnzsEpwwMDmL32/mwWCxoHqhzCE6ZsJrw4YWduHGvh+k0tkAPHFXadSZc1k8whzNOtt+DcWwCJ2p+4BZH/urpwfnWNpzs/YpbHJmwivjxehmTaWxh8rTT9M+GSM7CpZ5BjI2NcYtzdH069PEVzsSkn8YWmETQek+JTpER9/3h/pDL5dzinKDAQAQqg7nmHHv/VNRV5KCFaWRcc2RdvBpajRopjz/GLY4EBARgzerVeCKc3TCneJDPxohp/9QEFDIPfJwWwE76v1nkL8fep/yYnLvzNURFRTLZHqVSib35e6BSqfDkokysC32ae6ahwbPidyFRs4JbiK2jo1NKSFjCVaDHYMaRX0fQ3mcikwIeiVMh+2EfaNTTW2Q0jqKsvBxnzzaza5mYmICXXtw6dQUpVslK7n85ztyogkEcQIRXLNKis7BWu4GPADo7rzhOYCGZ9wQmxi249NN19P5xCxaTFdo4XyQ/Gw3fEDUf4RrzmoA4MoFT77RiqNfILTY8VTKk5i1HWJI/t8wNncDsd88JzV93OQSn0Kw0FF5kGXEHtyZgJj/+5283uebI6JAI3QUD11zDrQmIIyZYzLOvcNQgcsk13JqASqOAXDnzY0XxcfMgujUBmVzAkpRwrjmiCfNCxLIArrmG24dwzZbFTk+6mmRnw+vLIJAX1R3m9Q5IpHjo/kWP3jYDOZgW9g4sTY2EyseTj3CN/8VL6PYWMO70A1W7gHeTgfzFwBcZpBpp4U7OuAl4jxQm8ZvJ/pB/0IeygC9P0fKYD7DhfgaGeoFDzwAjA9zAEcjtyDwEPEjqPZFUVqm5QFMbd9rxShpQ9AYT55eBmj2OwSlWC1D9Fln5XaC42nlwSvEJoLGVK+7eApE8wVcauOKEsTtAVxPwfT03zMCxab9AGwUuz834iG2lszF6GzCQLMzGkM1PYwu0XWKaK/iQAlqt4coMhDwAJE0XJk7hfhpboL0a01xBIAXpo9u44oRIUmrFriV1G7kV9q2UPb73AdueYyKNLfNSazpXrlyRTdslZp2LmFXAbR2gn+4tGSHkJmWVkgz5AtGhpOsgT3ItaWBIGzeFnw9QuQ9YHk/bObG4qGgrs9NGkTaMdE9c5trvknR6nyTV5EtS23HS+pq4w46rOkl6v1SSXt0vSQXHJGlwmJltzWkhK42n8pSTk5OUnr5598K055WkPS8hKQT+AVyRrtzM5URAAAAAAElFTkSuQmCC
ttk::style theme names
ttk::style theme create litera -parent clam -settings {}
ttk::style theme names
ttk::style theme use litera
ttk::style configure . -background #ffffff -darkcolor #bfbfbf -foreground #343a40 -troughcolor #ffffff -selectbg #adb5bd -selectfg #ffffff -selectforeground #ffffff -selectbackground #adb5bd -fieldbg white -borderwidth 1 -focuscolor {}
ttk::style configure Link.TButton -foreground #343a40 -background #ffffff -bordercolor #ffffff -darkcolor #ffffff -lightcolor #ffffff -relief raised -focusthickness 0 -focuscolor #343a40 -anchor center -padding {10 5}
ttk::style map Link.TButton -shiftrelief {{pressed !disabled} -1} -foreground {disabled #c2c3c5 {pressed !disabled} #17a2b8 {hover !disabled} #17a2b8} -focuscolor {{pressed !disabled} #17a2b8 {hover !disabled} #17a2b8} -background {disabled #ffffff {pressed !disabled} #ffffff {hover !disabled} #ffffff} -bordercolor {disabled #ffffff {pressed !disabled} #ffffff {hover !disabled} #ffffff} -darkcolor {disabled #ffffff {pressed !disabled} #ffffff {hover !disabled} #ffffff} -lightcolor {disabled #ffffff {pressed !disabled} #ffffff {hover !disabled} #ffffff}
ttk::style configure symbol.Link.TButton -font {-size 16}
ttk::style configure TFrame -background #ffffff
ttk::style configure TLabel -foreground #343a40 -background #ffffff
ttk::style element create Combobox.downarrow from default
ttk::style element create Combobox.padding from clam
ttk::style element create Combobox.textarea from clam
ttk::style configure TCombobox -bordercolor #bfbfbf -darkcolor #fff -lightcolor #fff -arrowcolor #343a40 -foreground #343a40 -fieldbackground #fff -background #fff -insertcolor #343a40 -relief flat -padding 5 -arrowsize 12
ttk::style map TCombobox -background {readonly #F8F9FA} -fieldbackground {readonly #F8F9FA} -foreground {disabled #bfbfbf} -bordercolor {invalid #d9534f {focus !disabled} #4582ec {hover !disabled} #4582ec} -lightcolor {{focus invalid} #d9534f {focus !disabled} #4582ec {pressed !disabled} #4582ec readonly #F8F9FA} -darkcolor {{focus invalid} #d9534f {focus !disabled} #4582ec {pressed !disabled} #4582ec readonly #F8F9FA} -arrowcolor {disabled #bfbfbf {pressed !disabled} #4582ec {focus !disabled} #4582ec {hover !disabled} #4582ec}
ttk::style layout TCombobox {combo.Spinbox.field -side top -sticky ew -children {
  Combobox.downarrow -side right -sticky ns
  Combobox.padding -expand 1 -sticky nsew -children {
    Combobox.textarea -sticky nsew
  }
}}
image create photo pyimage2 -width 40 -height 1
ttk::style element create Horizontal.Separator.separator image {pyimage2 }
ttk::style layout Horizontal.TSeparator {Horizontal.Separator.separator -sticky ew}
image create photo pyimage3 -width 1 -height 40
ttk::style element create Vertical.Separator.separator image {pyimage3 }
ttk::style layout Vertical.TSeparator {Vertical.Separator.separator -sticky ns}
ttk::style configure TLabelframe.Label -foreground #343a40 -background #ffffff
ttk::style configure TLabelframe -relief raised -borderwidth 1 -bordercolor #bfbfbf -lightcolor #ffffff -darkcolor #ffffff -background #ffffff
ttk::style configure primary.TButton -foreground #ffffff -background #4582ec -bordercolor #4582ec -darkcolor #4582ec -lightcolor #4582ec -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5} -anchor center
ttk::style map primary.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed} -lightcolor {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed}
ttk::style configure secondary.TButton -foreground #ffffff -background #adb5bd -bordercolor #adb5bd -darkcolor #adb5bd -lightcolor #adb5bd -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5} -anchor center
ttk::style map secondary.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #bdc3ca {hover !disabled} #b5bcc3} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #bdc3ca {hover !disabled} #b5bcc3} -lightcolor {disabled #eaebeb {pressed !disabled} #bdc3ca {hover !disabled} #b5bcc3}
ttk::style configure success.TButton -foreground #ffffff -background #02b875 -bordercolor #02b875 -darkcolor #02b875 -lightcolor #02b875 -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5} -anchor center
ttk::style map success.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #34c690 {hover !disabled} #1bbf82} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #34c690 {hover !disabled} #1bbf82} -lightcolor {disabled #eaebeb {pressed !disabled} #34c690 {hover !disabled} #1bbf82}
ttk::style configure info.TButton -foreground #ffffff -background #17a2b8 -bordercolor #17a2b8 -darkcolor #17a2b8 -lightcolor #17a2b8 -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5} -anchor center
ttk::style map info.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #45b4c6 {hover !disabled} #2eabbf} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #45b4c6 {hover !disabled} #2eabbf} -lightcolor {disabled #eaebeb {pressed !disabled} #45b4c6 {hover !disabled} #2eabbf}
ttk::style configure warning.TButton -foreground #ffffff -background #f0ad4e -bordercolor #f0ad4e -darkcolor #f0ad4e -lightcolor #f0ad4e -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5} -anchor center
ttk::style map warning.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #f3bd71 {hover !disabled} #f1b55f} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #f3bd71 {hover !disabled} #f1b55f} -lightcolor {disabled #eaebeb {pressed !disabled} #f3bd71 {hover !disabled} #f1b55f}
ttk::style configure danger.TButton -foreground #ffffff -background #d9534f -bordercolor #d9534f -darkcolor #d9534f -lightcolor #d9534f -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5} -anchor center
ttk::style map danger.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #e07572 {hover !disabled} #dc6460} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #e07572 {hover !disabled} #dc6460} -lightcolor {disabled #eaebeb {pressed !disabled} #e07572 {hover !disabled} #dc6460}
ttk::style configure light.TButton -foreground #343A40 -background #F8F9FA -bordercolor #F8F9FA -darkcolor #F8F9FA -lightcolor #F8F9FA -relief raised -focusthickness 0 -focuscolor #343A40 -padding {10 5} -anchor center
ttk::style map light.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #f9fafb {hover !disabled} #f8f9fa} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #f9fafb {hover !disabled} #f8f9fa} -lightcolor {disabled #eaebeb {pressed !disabled} #f9fafb {hover !disabled} #f8f9fa}
ttk::style configure dark.TButton -foreground #F8F9FA -background #343A40 -bordercolor #343A40 -darkcolor #343A40 -lightcolor #343A40 -relief raised -focusthickness 0 -focuscolor #F8F9FA -padding {10 5} -anchor center
ttk::style map dark.TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #5c6166 {hover !disabled} #484d53} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #5c6166 {hover !disabled} #484d53} -lightcolor {disabled #eaebeb {pressed !disabled} #5c6166 {hover !disabled} #484d53}
image create photo pyimage4 -width 14 -height 14
image create photo pyimage5 -width 14 -height 14
image create photo pyimage6 -width 14 -height 14
image create photo pyimage7 -width 14 -height 14
image create photo pyimage8 -width 14 -height 14
image create photo pyimage9 -width 14 -height 14
ttk::style element create TCheckbutton.indicator image {pyimage5 {disabled selected} pyimage6 {disabled alternate} pyimage8 disabled pyimage9 alternate pyimage7 !selected pyimage4} -width 20 -border 4 -sticky w
ttk::style configure TCheckbutton -foreground #343a40
ttk::style map TCheckbutton -foreground {disabled #c2c3c5}
ttk::style layout TCheckbutton {Checkbutton.padding -sticky nsew -children {
  TCheckbutton.indicator -side left -sticky {}
  Checkbutton.focus -side left -sticky {} -children {
    Checkbutton.label -sticky nsew
  }
}}
image create photo pyimage10 -width 14 -height 14
image create photo pyimage11 -width 14 -height 14
image create photo pyimage12 -width 14 -height 14
image create photo pyimage13 -width 14 -height 14
ttk::style element create TRadiobutton.indicator image {pyimage11 {disabled selected} pyimage12 disabled pyimage13 !selected pyimage10} -width 20 -border 4 -sticky w
ttk::style map TRadiobutton -foreground {disabled #c2c3c5}
ttk::style configure TRadiobutton
ttk::style layout TRadiobutton {Radiobutton.padding -sticky nsew -children {
  TRadiobutton.indicator -side left -sticky {}
  Radiobutton.focus -side left -sticky {} -children {
    Radiobutton.label -sticky nsew
  }
}}
ttk::style configure Treeview.Heading -background #fff -foreground #343a40 -relief flat -padding 5
ttk::style map Treeview.Heading -foreground {disabled #cccccc} -bordercolor {{focus !disabled} #fff}
ttk::style configure Treeview -background #fff -fieldbackground #fff -foreground #343a40 -bordercolor #bfbfbf -lightcolor #fff -darkcolor #fff -borderwidth 2 -padding 0 -rowheight 23 -relief raised
ttk::style map Treeview -background {selected #adb5bd} -foreground {disabled #cccccc selected #ffffff} -bordercolor {disabled #bfbfbf focus #4582ec pressed #4582ec hover #4582ec} -lightcolor {focus #4582ec} -darkcolor {focus #4582ec}
ttk::style layout Treeview {Button.border -sticky nsew -border 1 -children {
  Treeview.padding -sticky nsew -children {
    Treeview.treearea -sticky nsew
  }
}}
ttk::style element create Treeitem.indicator from alt
ttk::style configure TNotebook -background #ffffff -bordercolor #bfbfbf -lightcolor #ffffff -darkcolor #ffffff -tabmargins {0 1 1 0}
ttk::style configure TNotebook.Tab -focuscolor {} -foreground #343a40 -padding {6 5}
ttk::style map TNotebook.Tab -background {selected #ffffff !selected #fff} -lightcolor {selected #ffffff !selected #fff} -bordercolor {selected #bfbfbf !selected #bfbfbf} -padding {selected {6 5} !selected {6 5}} -foreground {selected #343a40 !selected #343a40}
image create photo pyimage14 -width 28 -height 9
image create photo pyimage15 -width 28 -height 9
image create photo pyimage16 -width 28 -height 9
image create photo pyimage17 -width 9 -height 28
image create photo pyimage18 -width 9 -height 28
image create photo pyimage19 -width 9 -height 28
ttk::style configure Horizontal.TScrollbar -troughcolor #F8F9FA -darkcolor #F8F9FA -bordercolor #F8F9FA -lightcolor #F8F9FA -arrowcolor #bfbfbf -arrowsize 11 -background #F8F9FA -relief flat -borderwidth 0
ttk::style element create Horizontal.TScrollbar.thumb image {pyimage14 pressed pyimage15 active pyimage16} -border {3 0} -sticky nsew
ttk::style layout Horizontal.TScrollbar {Horizontal.Scrollbar.trough -sticky we -children {
  Horizontal.Scrollbar.leftarrow -side left -sticky {}
  Horizontal.Scrollbar.rightarrow -side right -sticky {}
  Horizontal.TScrollbar.thumb -expand 1 -sticky nswe
}}
ttk::style configure Horizontal.TScrollbar -arrowcolor #bfbfbf
ttk::style map Horizontal.TScrollbar -arrowcolor {pressed #b5b5b5 active #c8c8c8}
ttk::style configure Vertical.TScrollbar -troughcolor #F8F9FA -darkcolor #F8F9FA -bordercolor #F8F9FA -lightcolor #F8F9FA -arrowcolor #bfbfbf -arrowsize 11 -background #F8F9FA -relief flat -borderwidth 0
ttk::style element create Vertical.TScrollbar.thumb image {pyimage17 pressed pyimage18 active pyimage19} -border {0 3} -sticky nsew
ttk::style layout Vertical.TScrollbar {Vertical.Scrollbar.trough -sticky ns -children {
  Vertical.Scrollbar.uparrow -side top -sticky {}
  Vertical.Scrollbar.downarrow -side bottom -sticky {}
  Vertical.TScrollbar.thumb -expand 1 -sticky nswe
}}
ttk::style configure Vertical.TScrollbar -arrowcolor #bfbfbf
ttk::style map Vertical.TScrollbar -arrowcolor {pressed #b5b5b5 active #c8c8c8}
image create photo pyimage20 -width 14 -height 14
image create photo pyimage21 -width 14 -height 14
image create photo pyimage22 -width 14 -height 14
image create photo pyimage23 -width 14 -height 14
image create photo pyimage24 -width 40 -height 5
image create photo pyimage25 -width 5 -height 40
ttk::style element create Horizontal.Scale.slider image {pyimage20 disabled pyimage23 pressed pyimage21 hover pyimage22}
ttk::style element create Horizontal.Scale.track image {pyimage24 }
ttk::style layout Horizontal.TScale {Horizontal.Scale.focus -expand 1 -sticky nsew -children {
  Horizontal.Scale.track -sticky ew
  Horizontal.Scale.slider -side left -sticky {}
}}
ttk::style element create Vertical.Scale.slider image {pyimage20 disabled pyimage23 pressed pyimage21 hover pyimage22}
ttk::style element create Vertical.Scale.track image {pyimage25 }
ttk::style layout Vertical.TScale {Vertical.Scale.focus -expand 1 -sticky nsew -children {
  Vertical.Scale.track -sticky ns
  Vertical.Scale.slider -side top -sticky {}
}}
ttk::style configure Horizontal.TProgressbar -thickness 10 -borderwidth 1 -bordercolor #F8F9FA -lightcolor #bfbfbf -pbarrelief flat -troughcolor #F8F9FA
ttk::style element names
ttk::style configure Vertical.TProgressbar -thickness 10 -borderwidth 1 -bordercolor #F8F9FA -lightcolor #bfbfbf -pbarrelief flat -troughcolor #F8F9FA
ttk::style element names
ttk::style element create Horizontal.Progressbar.trough from clam
ttk::style element create Horizontal.Progressbar.pbar from default
ttk::style layout Horizontal.TProgressbar {Horizontal.Progressbar.trough -sticky nswe -children {
  Horizontal.Progressbar.pbar -side left -sticky ns
}}
ttk::style configure Horizontal.TProgressbar -background #4582ec
ttk::style element create Vertical.Progressbar.trough from clam
ttk::style element create Vertical.Progressbar.pbar from default
ttk::style configure Vertical.TProgressbar -background #4582ec
ttk::style layout Vertical.TProgressbar {Vertical.Progressbar.trough -sticky nswe -children {
  Vertical.Progressbar.pbar -side bottom -sticky we
}}
image create photo pyimage26 -width 12 -height 12
image create photo pyimage27 -width 12 -height 12
ttk::style element create success.Striped.Horizontal.Progressbar.pbar image {pyimage26 } -width 12 -sticky ew
ttk::style layout success.Striped.Horizontal.TProgressbar {success.Striped.Horizontal.Progressbar.trough -sticky nsew -children {
  success.Striped.Horizontal.Progressbar.pbar -side left -sticky ns
}}
ttk::style configure success.Striped.Horizontal.TProgressbar -troughcolor #F8F9FA -thickness 12 -bordercolor #F8F9FA -borderwidth 1
ttk::style element create success.Striped.Vertical.Progressbar.pbar image {pyimage27 } -width 12 -sticky ns
ttk::style layout success.Striped.Vertical.TProgressbar {success.Striped.Vertical.Progressbar.trough -sticky nsew -children {
  success.Striped.Vertical.Progressbar.pbar -side bottom -sticky ew
}}
ttk::style configure success.Striped.Vertical.TProgressbar -troughcolor #F8F9FA -bordercolor #F8F9FA -thickness 12 -borderwidth 1
ttk::style configure Metersubtxt.TLabel -foreground #adb5bd -background #ffffff
ttk::style configure info.Meter.TLabel -foreground #17a2b8 -background #ffffff -space #F8F9FA
ttk::style lookup info.Meter.TLabel -foreground
ttk::style lookup info.Meter.TLabel -background
ttk::style lookup info.Meter.TLabel -space
image create photo pyimage28 -width 149 -height 149
image create photo pyimage29 -width 28 -height 9
image create photo pyimage30 -width 28 -height 9
image create photo pyimage31 -width 28 -height 9
image create photo pyimage32 -width 9 -height 28
image create photo pyimage33 -width 9 -height 28
image create photo pyimage34 -width 9 -height 28
ttk::style configure danger.Round.Horizontal.TScrollbar -troughcolor #F8F9FA -darkcolor #F8F9FA -bordercolor #F8F9FA -lightcolor #F8F9FA -arrowcolor #d9534f -arrowsize 11 -background #F8F9FA -relief flat -borderwidth 0
ttk::style element create danger.Round.Horizontal.TScrollbar.thumb image {pyimage29 pressed pyimage30 active pyimage31} -border 9 -padding 0 -sticky ew
ttk::style layout danger.Round.Horizontal.TScrollbar {Horizontal.Scrollbar.trough -sticky we -children {
  Horizontal.Scrollbar.leftarrow -side left -sticky {}
  Horizontal.Scrollbar.rightarrow -side right -sticky {}
  danger.Round.Horizontal.TScrollbar.thumb -expand 1 -sticky nswe
}}
ttk::style configure danger.Round.Horizontal.TScrollbar -arrowcolor #d9534f
ttk::style map danger.Round.Horizontal.TScrollbar -arrowcolor {pressed #ce4e4b active #e35752}
ttk::style configure danger.Round.Vertical.TScrollbar -troughcolor #F8F9FA -darkcolor #F8F9FA -bordercolor #F8F9FA -lightcolor #F8F9FA -arrowcolor #d9534f -arrowsize 11 -background #F8F9FA -relief flat
ttk::style element create danger.Round.Vertical.TScrollbar.thumb image {pyimage32 pressed pyimage33 active pyimage34} -border 9 -padding 0 -sticky ns
ttk::style layout danger.Round.Vertical.TScrollbar {Vertical.Scrollbar.trough -sticky ns -children {
  Vertical.Scrollbar.uparrow -side top -sticky {}
  Vertical.Scrollbar.downarrow -side bottom -sticky {}
  danger.Round.Vertical.TScrollbar.thumb -expand 1 -sticky nswe
}}
ttk::style configure danger.Round.Vertical.TScrollbar -arrowcolor #d9534f
ttk::style map danger.Round.Vertical.TScrollbar -arrowcolor {pressed #ce4e4b active #e35752}
ttk::style configure TButton -foreground #ffffff -background #4582ec -bordercolor #4582ec -darkcolor #4582ec -lightcolor #4582ec -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5} -anchor center
ttk::style map TButton -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed} -bordercolor {disabled #eaebeb} -darkcolor {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed} -lightcolor {disabled #eaebeb {pressed !disabled} #6a9bef {hover !disabled} #578eed}
ttk::style configure secondary.TMenubutton -foreground #ffffff -background #adb5bd -bordercolor #adb5bd -darkcolor #adb5bd -lightcolor #adb5bd -arrowsize 4 -arrowcolor #ffffff -arrowpadding {0 0 15 0} -relief raised -focusthickness 0 -focuscolor #ffffff -padding {10 5}
ttk::style map secondary.TMenubutton -arrowcolor {disabled #c2c3c5} -foreground {disabled #c2c3c5} -background {disabled #eaebeb {pressed !disabled} #bdc3ca {hover !disabled} #b5bcc3} -bordercolor {disabled #eaebeb {pressed !disabled} #bdc3ca {hover !disabled} #b5bcc3} -darkcolor {disabled #eaebeb {pressed !disabled} #bdc3ca {hover !disabled} #b5bcc3} -lightcolor {disabled #eaebeb {pressed !disabled} #bdc3ca {hover !disabled} #b5bcc3}
ttk::style configure success.Toolbutton -foreground #ffffff -background #bfbfbf -bordercolor #bfbfbf -darkcolor #bfbfbf -lightcolor #bfbfbf -relief raised -focusthickness 0 -focuscolor {} -padding {10 5} -anchor center
ttk::style map success.Toolbutton -foreground {disabled #c2c3c5 hover #ffffff selected #ffffff} -background {disabled #eaebeb {pressed !disabled} #02b875 {selected !disabled} #02b875 {hover !disabled} #02b875} -bordercolor {disabled #eaebeb {pressed !disabled} #02b875 {selected !disabled} #02b875 {hover !disabled} #02b875} -darkcolor {disabled #eaebeb {pressed !disabled} #02b875 {selected !disabled} #02b875 {hover !disabled} #02b875} -lightcolor {disabled #eaebeb {pressed !disabled} #02b875 {selected !disabled} #02b875 {hover !disabled} #02b875}
ttk::style configure info.Outline.TButton -foreground #17a2b8 -background #ffffff -bordercolor #17a2b8 -darkcolor #ffffff -lightcolor #ffffff -relief raised -focusthickness 0 -focuscolor #17a2b8 -padding {10 5} -anchor center
ttk::style map info.Outline.TButton -foreground {disabled #c2c3c5 {pressed !disabled} #ffffff {hover !disabled} #ffffff} -background {{pressed !disabled} #17a2b8 {hover !disabled} #17a2b8} -bordercolor {disabled #c2c3c5 {pressed !disabled} #17a2b8 {hover !disabled} #17a2b8} -focuscolor {{pressed !disabled} #ffffff {hover !disabled} #ffffff} -darkcolor {{pressed !disabled} #17a2b8 {hover !disabled} #17a2b8} -lightcolor {{pressed !disabled} #17a2b8 {hover !disabled} #17a2b8}
ttk::style configure warning.Outline.TMenubutton -foreground #f0ad4e -background #ffffff -bordercolor #f0ad4e -darkcolor #ffffff -lightcolor #ffffff -relief raised -focusthickness 0 -focuscolor #f0ad4e -padding {10 5} -arrowcolor #f0ad4e -arrowpadding {0 0 15 0} -arrowsize 4
ttk::style map warning.Outline.TMenubutton -foreground {disabled #c2c3c5 {pressed !disabled} #ffffff {hover !disabled} #ffffff} -background {{pressed !disabled} #f0ad4e {hover !disabled} #f0ad4e} -bordercolor {disabled #c2c3c5 pressed #f0ad4e hover #f0ad4e} -darkcolor {{pressed !disabled} #f0ad4e {hover !disabled} #f0ad4e} -lightcolor {{pressed !disabled} #f0ad4e {hover !disabled} #f0ad4e} -arrowcolor {disabled #c2c3c5 pressed #ffffff hover #ffffff}
ttk::style configure success.Outline.Toolbutton -foreground #02b875 -background #ffffff -bordercolor #02b875 -darkcolor #ffffff -lightcolor #ffffff -relief raised -focusthickness 0 -focuscolor #02b875 -padding {10 5} -anchor center -arrowcolor #02b875 -arrowpadding {0 0 15 0} -arrowsize 3
ttk::style map success.Outline.Toolbutton -foreground {disabled #c2c3c5 {pressed !disabled} #ffffff {selected !disabled} #ffffff {hover !disabled} #ffffff} -background {{pressed !disabled} #02b875 {selected !disabled} #02b875 {hover !disabled} #02b875} -bordercolor {disabled #c2c3c5 {pressed !disabled} #02b875 {selected !disabled} #02b875 {hover !disabled} #02b875} -darkcolor {disabled #ffffff {pressed !disabled} #02b875 {selected !disabled} #02b875 {hover !disabled} #02b875} -lightcolor {disabled #ffffff {pressed !disabled} #02b875 {selected !disabled} #02b875 {hover !disabled} #02b875}
image create photo pyimage35 -width 24 -height 15
image create photo pyimage36 -width 24 -height 15
image create photo pyimage37 -width 24 -height 15
image create photo pyimage38 -width 24 -height 15
ttk::style element create success.Round.Toggle.indicator image {pyimage36 {disabled selected} pyimage37 disabled pyimage38 !selected pyimage35} -width 28 -border 4 -sticky w
ttk::style configure success.Round.Toggle -relief flat -borderwidth 0 -padding 0 -foreground #343a40 -background #ffffff
ttk::style map success.Round.Toggle -foreground {disabled #c2c3c5} -background {selected #ffffff}
ttk::style layout success.Round.Toggle {Toolbutton.border -sticky nsew -children {
  Toolbutton.padding -sticky nsew -children {
    success.Round.Toggle.indicator -side left
    Toolbutton.label -side left
  }
}}
image create photo pyimage39 -width 24 -height 15
image create photo pyimage40 -width 24 -height 15
image create photo pyimage41 -width 24 -height 15
image create photo pyimage42 -width 24 -height 15
ttk::style element create Square.Toggle.indicator image {pyimage40 {disabled selected} pyimage42 disabled pyimage41 !selected pyimage39} -width 28 -border 4 -sticky w
ttk::style layout Square.Toggle {Toolbutton.border -sticky nsew -children {
  Toolbutton.padding -sticky nsew -children {
    Square.Toggle.indicator -side left
    Toolbutton.label -side left
  }
}}
ttk::style configure Square.Toggle -relief flat -borderwidth 0 -foreground #343a40
ttk::style map Square.Toggle -foreground {disabled #c2c3c5} -background {selected #ffffff !selected #ffffff}
ttk::style configure TEntry -bordercolor #bfbfbf -darkcolor #fff -lightcolor #fff -fieldbackground #fff -foreground #343a40 -insertcolor #343a40 -padding 5
ttk::style map TEntry -foreground {disabled #bfbfbf} -fieldbackground {readonly #F8F9FA} -bordercolor {invalid #d9534f {focus !disabled} #4582ec {hover !disabled} #4582ec} -lightcolor {{focus invalid} #d9534f {focus !disabled} #4582ec readonly #F8F9FA} -darkcolor {{focus invalid} #d9534f {focus !disabled} #4582ec readonly #F8F9FA}
ttk::style element create TSpinbox.uparrow from default
ttk::style element create TSpinbox.downarrow from default
ttk::style layout TSpinbox {TSpinbox.field -side top -sticky ew -children {
  null -side right -sticky {} -children {
    TSpinbox.uparrow -side top -sticky e
    TSpinbox.downarrow -side bottom -sticky e
  }
  TSpinbox.padding -sticky nsew -children {
    TSpinbox.textarea -sticky nsew
  }
}}
ttk::style configure TSpinbox -bordercolor #bfbfbf -darkcolor #fff -lightcolor #fff -fieldbackground #fff -foreground #343a40 -borderwidth 0 -background #fff -relief flat -arrowcolor #343a40 -insertcolor #343a40 -arrowsize 12 -padding {10 5}
ttk::style map TSpinbox -foreground {disabled #bfbfbf} -fieldbackground {readonly #F8F9FA} -background {readonly #F8F9FA} -lightcolor {{focus invalid} #d9534f {focus !disabled} #4582ec readonly #F8F9FA} -darkcolor {{focus invalid} #d9534f {focus !disabled} #4582ec readonly #F8F9FA} -bordercolor {invalid #d9534f {focus !disabled} #4582ec {hover !disabled} #4582ec} -arrowcolor {{disabled !disabled} #bfbfbf {pressed !disabled} #4582ec {hover !disabled} #4582ec}
image create photo pyimage43 -width 21 -height 22
ttk::style configure Date.TButton -foreground #ffffff -background #4582ec -bordercolor #4582ec -darkcolor #4582ec -lightcolor #4582ec -relief raised -focusthickness 0 -focuscolor #ffffff -padding {2 2} -anchor center -image pyimage43
ttk::style map Date.TButton -foreground {disabled #bfbfbf} -background {disabled #bfbfbf {pressed !disabled} #3e74d4 {hover !disabled} #4685f2} -bordercolor {disabled #bfbfbf} -darkcolor {disabled #bfbfbf {pressed !disabled} #3e74d4 {hover !disabled} #4685f2} -lightcolor {disabled #bfbfbf {pressed !disabled} #3e74d4 {hover !disabled} #4685f2}
ttk::style lookup info.Meter.TLabel -foreground
ttk::style lookup info.Meter.TLabel -background
ttk::style lookup info.Meter.TLabel -space
image create photo pyimage44 -width 149 -height 149
image delete pyimage28

*/
