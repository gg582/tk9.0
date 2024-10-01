// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
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

func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func TestTmp(t *testing.T) {
	ErrorMode = CollectErrors
	themes := StyleThemeNames()
	slices.Sort(themes)
	for _, theme := range themes {
		fmt.Printf("\n//\n// # %q theme style guide", theme)
		StyleThemeUse(theme)
		styles := StyleThemeStyles()
		slices.Sort(styles)
		elements := StyleElementNames()
		slices.Sort(elements)
		for _, element := range elements {
			fmt.Printf("\n//\n// %q style element options:\n//", element)
			options := StyleElementOptions(element)
			slices.Sort(options)
			for _, option := range options {
				fmt.Printf("\n//  - [%s]", capitalize(option[1:]))
			}
		}
		fmt.Printf("\n//\n// %q theme style list", theme)
		needsep := false
		for _, style := range styles {
			if needsep {
				fmt.Printf("\n//\n// -\n//\n//")
			}
			needsep = true
			fmt.Printf("\n//\n//  %s", style)
			if s := StyleLayout(style); s != "" {
				needsep = false
				fmt.Printf("\n//\n// Layout: %s", s)
			}
			if s := strings.TrimSpace(StyleMap(style)); s != "" {
				needsep = false
				s = strings.Join(strings.Fields(s), " ")
				fmt.Printf("\n//\n// Style map: %s", s)
			}
		}
	}
}

//
// # "alt" theme style guide
//
// "Checkbutton.indicator" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Foreground]
//  - [Indicatorcolor]
//  - [Indicatormargin]
//  - [Lightcolor]
//  - [Shadecolor]
//
// "Combobox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "Menubutton.indicator" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Direction]
//
// "Radiobutton.indicator" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Foreground]
//  - [Indicatorcolor]
//  - [Indicatormargin]
//  - [Lightcolor]
//  - [Shadecolor]
//
// "Spinbox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "Spinbox.uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "Treeitem.indicator" style element options:
//
//  - [Foreground]
//  - [Indicatormargins]
//  - [Size]
//
// "arrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "border" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Borderwidth]
//  - [Default]
//  - [Relief]
//
// "downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "field" style element options:
//
//  - [Bordercolor]
//  - [Fieldbackground]
//  - [Focuscolor]
//  - [Focuswidth]
//
// "leftarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "rightarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "slider" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Borderwidth]
//  - [Orient]
//  - [Sliderrelief]
//  - [Sliderthickness]
//
// "thumb" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Orient]
//  - [Relief]
//  - [Width]
//
// "uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "alt" theme style list
//
//  .
//
// Style map: -foreground {disabled #a3a3a3} -background {disabled #d9d9d9 active #ececec} -embossed {disabled 1}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.border -sticky nswe -border 1 -children {Button.focus -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
// Style map: -highlightcolor {alternate black} -relief { {pressed !disabled} sunken {active !disabled} raised }
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
// Style map: -indicatorcolor {pressed #d9d9d9 alternate #aaaaaa disabled #d9d9d9}
//
//  TCombobox
//
// Layout: Combobox.field -sticky nswe -children {Combobox.downarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9} -arrowcolor {disabled #a3a3a3}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TMenubutton
//
// Layout: Menubutton.border -sticky nswe -children {Menubutton.focus -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -side left -sticky {}}}}
//
//  TNotebook
//
// Layout: Notebook.client -sticky nswe
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -expand {selected {1.5p 1.5p 0.75p 0}} -background {selected #d9d9d9}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
// Style map: -indicatorcolor {pressed #d9d9d9 alternate #aaaaaa disabled #d9d9d9}
//
//  TScale
//
// -
//
//
//
//  TScrollbar
//
// -
//
//
//
//  TSpinbox
//
// Layout: Spinbox.field -side top -sticky we -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9} -arrowcolor {disabled #a3a3a3}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -relief {disabled flat selected sunken pressed sunken active raised} -background {pressed #c3c3c3 active #ececec}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled #a3a3a3 selected #ffffff} -background {disabled #d9d9d9 selected #4a6984}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
// # "aqua" theme style guide
//
// "Button.button" style element options:
//
//
// "Checkbutton.button" style element options:
//
//
// "Combobox.button" style element options:
//
//
// "DisclosureButton.button" style element options:
//
//
// "Entry.field" style element options:
//
//  - [Background]
//  - [Fieldbackground]
//
// "GradientButton.button" style element options:
//
//
// "HelpButton.button" style element options:
//
//
// "Horizontal.Scrollbar.leftarrow" style element options:
//
//  - [Orient]
//
// "Horizontal.Scrollbar.rightarrow" style element options:
//
//  - [Orient]
//
// "Horizontal.Scrollbar.thumb" style element options:
//
//  - [Orient]
//
// "Horizontal.Scrollbar.trough" style element options:
//
//  - [Orient]
//
// "InlineButton.button" style element options:
//
//
// "Labelframe.border" style element options:
//
//
// "Menubutton.button" style element options:
//
//
// "Notebook.client" style element options:
//
//
// "Notebook.tab" style element options:
//
//
// "Progressbar.track" style element options:
//
//  - [Maximum]
//  - [Mode]
//  - [Orient]
//  - [Phase]
//  - [Value]
//
// "Radiobutton.button" style element options:
//
//
// "RecessedButton.button" style element options:
//
//
// "RoundedRectButton.button" style element options:
//
//
// "Scale.slider" style element options:
//
//
// "Scale.trough" style element options:
//
//  - [From]
//  - [Orient]
//  - [To]
//  - [Value]
//
// "Searchbox.field" style element options:
//
//  - [Background]
//  - [Fieldbackground]
//
// "SidebarButton.button" style element options:
//
//
// "Spinbox.downarrow" style element options:
//
//
// "Spinbox.field" style element options:
//
//  - [Background]
//  - [Fieldbackground]
//
// "Spinbox.uparrow" style element options:
//
//
// "Toolbar.background" style element options:
//
//
// "Toolbutton.border" style element options:
//
//
// "Treeheading.cell" style element options:
//
//
// "Treeitem.indicator" style element options:
//
//
// "Treeview.treearea" style element options:
//
//
// "Vertical.Scrollbar.downarrow" style element options:
//
//  - [Orient]
//
// "Vertical.Scrollbar.thumb" style element options:
//
//  - [Orient]
//
// "Vertical.Scrollbar.trough" style element options:
//
//  - [Orient]
//
// "Vertical.Scrollbar.uparrow" style element options:
//
//  - [Orient]
//
// "background" style element options:
//
//
// "field" style element options:
//
//  - [Fieldbackground]
//
// "fill" style element options:
//
//
// "hseparator" style element options:
//
//
// "separator" style element options:
//
//
// "sizegrip" style element options:
//
//
// "vseparator" style element options:
//
//
// "aqua" theme style list
//
//  .
//
// Style map: -selectforeground { background systemSelectedTextColor !focus systemSelectedTextColor} -foreground { disabled systemDisabledControlTextColor background systemLabelColor} -selectbackground { background systemSelectedTextBackgroundColor !focus systemSelectedTextBackgroundColor}
//
//  DisclosureButton
//
// Layout: DisclosureButton.button -sticky nswe
//
//  GradientButton
//
// Layout: GradientButton.button -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.image -side right -sticky {} Treeheading.text -side top -sticky {}
//
//  HelpButton
//
// Layout: HelpButton.button -sticky nswe
//
//  Horizontal.TScrollbar
//
// Layout: Horizontal.Scrollbar.trough -sticky we -children {Horizontal.Scrollbar.thumb -sticky nswe Horizontal.Scrollbar.rightarrow -side right -sticky {} Horizontal.Scrollbar.leftarrow -side right -sticky {}}
//
//  ImageButton
//
// Layout: Button.padding -sticky nswe -children {Button.label -sticky nswe}
//
// Style map: -foreground { pressed systemLabelColor !pressed systemSecondaryLabelColor }
//
//  InlineButton
//
// Layout: InlineButton.button -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}
//
// Style map: -foreground { disabled systemWindowBackgroundColor }
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -side left -sticky {}}
//
//  Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  RecessedButton
//
// Layout: RecessedButton.button -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}
//
// Style map: -font { selected RecessedFont active RecessedFont pressed RecessedFont } -foreground { {disabled selected} systemWindowBackgroundColor3 {disabled !selected} systemDisabledControlTextColor selected systemTextBackgroundColor active white pressed white }
//
//  RoundedRectButton
//
// Layout: RoundedRectButton.button -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}
//
//  Searchbox
//
// Layout: Searchbox.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
//  SidebarButton
//
// Layout: SidebarButton.button -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}
//
// Style map: -foreground { {disabled selected} systemWindowBackgroundColor3 {disabled !selected} systemDisabledControlTextColor selected systemTextColor active systemTextColor pressed systemTextColor }
//
//  TButton
//
// Layout: Button.button -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}
//
// Style map: -foreground { pressed white {alternate !pressed !background} white disabled systemDisabledControlTextColor}
//
//  TCheckbutton
//
// Layout: Checkbutton.button -sticky nswe -children {Checkbutton.padding -sticky nswe -children {Checkbutton.label -side left -sticky {}}}
//
//  TCombobox
//
// Layout: Combobox.button -sticky nswe -children {Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}
//
// Style map: -foreground { disabled systemDisabledControlTextColor } -selectbackground { !focus systemUnemphasizedSelectedTextBackgroundColor }
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -foreground { disabled systemDisabledControlTextColor } -selectbackground { !focus systemUnemphasizedSelectedTextBackgroundColor }
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TLabelframe.Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  TMenubutton
//
// Layout: Menubutton.button -sticky nswe -children {Menubutton.padding -sticky nswe -children {Menubutton.label -side left -sticky {}}}
//
//  TNotebook
//
// Layout: Notebook.client -sticky nswe
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -sticky nswe -children {Notebook.label -sticky nswe}}
//
// Style map: -foreground { {background !selected} systemControlTextColor {background selected} black {!background selected} systemSelectedTabTextColor disabled systemDisabledControlTextColor}
//
//  TProgressbar
//
// Layout: Progressbar.track -sticky nswe
//
//  TRadiobutton
//
// Layout: Radiobutton.button -sticky nswe -children {Radiobutton.padding -sticky nswe -children {Radiobutton.label -side left -sticky {}}}
//
//  TScrollbar
//
// -
//
//
//
//  TSpinbox
//
// Layout: Spinbox.buttons -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.field -sticky we -children {Spinbox.textarea -sticky we}
//
// Style map: -foreground { disabled systemDisabledControlTextColor } -selectbackground { !focus systemUnemphasizedSelectedTextBackgroundColor }
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -sticky nswe -children {Notebook.label -sticky nswe}}
//
//  Toolbar
//
// Layout: Toolbar.background -sticky nswe
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -background { selected systemSelectedTextBackgroundColor }
//
//  Vertical.TScrollbar
//
// Layout: Vertical.Scrollbar.trough -sticky ns -children {Vertical.Scrollbar.thumb -sticky nswe Vertical.Scrollbar.downarrow -side bottom -sticky {} Vertical.Scrollbar.uparrow -side bottom -sticky {}}
//
// # "clam" theme style guide
//
// "Checkbutton.indicator" style element options:
//
//  - [Indicatorbackground]
//  - [Indicatorforeground]
//  - [Indicatormargin]
//  - [Lowerbordercolor]
//  - [Upperbordercolor]
//
// "Combobox.field" style element options:
//
//  - [Bordercolor]
//  - [Fieldbackground]
//  - [Lightcolor]
//
// "Radiobutton.indicator" style element options:
//
//  - [Indicatorbackground]
//  - [Indicatorforeground]
//  - [Indicatormargin]
//  - [Lowerbordercolor]
//  - [Upperbordercolor]
//
// "Spinbox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "Spinbox.uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "arrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "bar" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "border" style element options:
//
//  - [Bordercolor]
//  - [Borderwidth]
//  - [Darkcolor]
//  - [Lightcolor]
//  - [Relief]
//
// "client" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Lightcolor]
//
// "downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "field" style element options:
//
//  - [Bordercolor]
//  - [Fieldbackground]
//  - [Lightcolor]
//
// "hgrip" style element options:
//
//  - [Bordercolor]
//  - [Gripsize]
//  - [Lightcolor]
//
// "leftarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "pbar" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "rightarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "slider" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "tab" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Lightcolor]
//
// "thumb" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "trough" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "vgrip" style element options:
//
//  - [Bordercolor]
//  - [Gripsize]
//  - [Lightcolor]
//
// "clam" theme style list
//
//  .
//
// Style map: -selectforeground {!focus white} -foreground {disabled #999999} -selectbackground {!focus #9e9a91} -background {disabled #dcdad5 active #eeebe7}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Horizontal.Sash
//
// Layout: Sash.hsash -sticky nswe -children {Sash.hgrip -sticky nswe}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Sash
//
// -
//
//
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.border -sticky nswe -border 1 -children {Button.focus -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
// Style map: -lightcolor {pressed #bab5ab} -background {disabled #dcdad5 pressed #bab5ab active #eeebe7} -bordercolor {alternate #000000} -darkcolor {pressed #bab5ab}
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
// Style map: -indicatorbackground {pressed #dcdad5 {!disabled alternate} #5895bc {disabled alternate} #a0a0a0 disabled #dcdad5}
//
//  TCombobox
//
// Layout: Combobox.downarrow -side right -sticky ns Combobox.field -sticky nswe -children {Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}
//
// Style map: -foreground {{readonly focus} #ffffff} -fieldbackground {{readonly focus} #4a6984 readonly #dcdad5} -background {active #eeebe7 pressed #eeebe7} -bordercolor {focus #4a6984} -arrowcolor {disabled #999999}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -lightcolor {focus #6f9dc6} -background {readonly #dcdad5} -bordercolor {focus #4a6984}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TMenubutton
//
// Layout: Menubutton.border -sticky nswe -children {Menubutton.focus -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -side left -sticky {}}}}
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -lightcolor {selected #eeebe7 {} #cfcdc8} -padding {selected {4.5p 3p 4.5p 1.5p}} -background {selected #dcdad5 {} #bab5ab}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
// Style map: -indicatorbackground {pressed #dcdad5 {!disabled alternate} #5895bc {disabled alternate} #a0a0a0 disabled #dcdad5}
//
//  TScale
//
// -
//
//
//
//  TScrollbar
//
// -
//
//
//
//  TSpinbox
//
// Layout: Spinbox.field -side top -sticky we -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}
//
// Style map: -background {readonly #dcdad5} -bordercolor {focus #4a6984} -arrowcolor {disabled #999999}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -lightcolor {pressed #bab5ab} -relief {disabled flat selected sunken pressed sunken active raised} -background {disabled #dcdad5 pressed #bab5ab active #eeebe7} -darkcolor {pressed #bab5ab}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled #999999 selected #ffffff} -background {disabled #dcdad5 selected #4a6984} -bordercolor {focus #4a6984}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  Vertical.Sash
//
// Layout: Sash.vsash -sticky nswe -children {Sash.vgrip -sticky nswe}
//
// # "classic" theme style guide
//
// "Button.border" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Default]
//  - [Relief]
//
// "Checkbutton.indicator" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Indicatorcolor]
//  - [Indicatormargin]
//  - [Indicatorrelief]
//  - [Indicatorsize]
//
// "Combobox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Menubutton.indicator" style element options:
//
//  - [Background]
//  - [Indicatorborderwidth]
//  - [Indicatorheight]
//  - [Indicatormargin]
//  - [Indicatorrelief]
//  - [Indicatorwidth]
//
// "Radiobutton.indicator" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Indicatorcolor]
//  - [Indicatormargin]
//  - [Indicatorrelief]
//  - [Indicatorsize]
//
// "Spinbox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Spinbox.uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "arrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "downarrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "highlight" style element options:
//
//  - [Default]
//  - [Highlightcolor]
//  - [Highlightthickness]
//
// "hsash" style element options:
//
//  - [Background]
//  - [Handlepad]
//  - [Handlesize]
//  - [Sashpad]
//  - [Sashrelief]
//  - [Sashthickness]
//
// "leftarrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "rightarrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "slider" style element options:
//
//  - [Background]
//  - [Orient]
//  - [Sliderborderwidth]
//  - [Sliderlength]
//  - [Sliderrelief]
//  - [Sliderthickness]
//
// "uparrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "vsash" style element options:
//
//  - [Background]
//  - [Handlepad]
//  - [Handlesize]
//  - [Sashpad]
//  - [Sashrelief]
//  - [Sashthickness]
//
// "classic" theme style list
//
//  .
//
// Style map: -highlightcolor {focus black} -foreground {disabled #a3a3a3} -background {disabled #d9d9d9 active #ececec}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Horizontal.TScale
//
// Layout: Horizontal.Scale.highlight -sticky nswe -children {Horizontal.Scale.trough -sticky nswe -children {Horizontal.Scale.slider -side left -sticky {}}}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Sash
//
// -
//
//
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.highlight -sticky nswe -children {Button.border -sticky nswe -border 1 -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
// Style map: -relief {{!disabled pressed} sunken}
//
//  TCheckbutton
//
// Layout: Checkbutton.highlight -sticky nswe -children {Checkbutton.border -sticky nswe -children {Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.label -side left -sticky nswe}}}
//
// Style map: -indicatorrelief {alternate raised selected sunken pressed sunken} -indicatorcolor {pressed #d9d9d9 alternate #b05e5e selected #b03060}
//
//  TCombobox
//
// Layout: Combobox.highlight -sticky nswe -children {Combobox.field -sticky nswe -children {Combobox.downarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  TEntry
//
// Layout: Entry.highlight -sticky nswe -children {Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TMenubutton
//
// Layout: Menubutton.highlight -sticky nswe -children {Menubutton.border -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -sticky {}}}}
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -background {selected #d9d9d9}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.highlight -sticky nswe -children {Radiobutton.border -sticky nswe -children {Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.label -side left -sticky nswe}}}
//
// Style map: -indicatorrelief {alternate raised selected sunken pressed sunken} -indicatorcolor {pressed #d9d9d9 alternate #b05e5e selected #b03060}
//
//  TScale
//
// Style map: -sliderrelief {{pressed !disabled} sunken}
//
//  TScrollbar
//
// Style map: -relief {{pressed !disabled} sunken}
//
//  TSpinbox
//
// Layout: Spinbox.highlight -sticky nswe -children {Spinbox.field -sticky nswe -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.focus -sticky nswe -children {Toolbutton.border -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -relief {disabled flat selected sunken pressed sunken active raised} -background {pressed #b3b3b3 active #ececec}
//
//  Treeview
//
// Layout: Treeview.highlight -sticky nswe -children {Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}}
//
// Style map: -foreground {disabled #a3a3a3 selected #000000} -background {disabled #d9d9d9 selected #c3c3c3}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  Vertical.TScale
//
// Layout: Vertical.Scale.highlight -sticky nswe -children {Vertical.Scale.trough -sticky nswe -children {Vertical.Scale.slider -side top -sticky {}}}
//
// # "default" theme style guide
//
// "" style element options:
//
//
// "Checkbutton.indicator" style element options:
//
//  - [Bordercolor]
//  - [Indicatorbackground]
//  - [Indicatorforeground]
//  - [Indicatormargin]
//
// "Combobox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Menubutton.indicator" style element options:
//
//  - [Arrowcolor]
//  - [Arrowpadding]
//  - [Arrowsize]
//
// "Radiobutton.indicator" style element options:
//
//  - [Bordercolor]
//  - [Indicatorbackground]
//  - [Indicatorforeground]
//  - [Indicatormargin]
//
// "Spinbox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Spinbox.uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Treeheading.cell" style element options:
//
//  - [Background]
//  - [Rownumber]
//
// "Treeitem.indicator" style element options:
//
//  - [Foreground]
//  - [Indicatormargins]
//  - [Indicatorsize]
//
// "Treeitem.row" style element options:
//
//  - [Background]
//  - [Rownumber]
//
// "Treeitem.separator" style element options:
//
//  - [Background]
//  - [Rownumber]
//
// "arrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "background" style element options:
//
//  - [Background]
//
// "border" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "client" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Highlight]
//  - [Highlightcolor]
//
// "ctext" style element options:
//
//  - [Anchor]
//  - [Embossed]
//  - [Font]
//  - [Foreground]
//  - [Justify]
//  - [Text]
//  - [Underline]
//  - [Width]
//  - [Wraplength]
//
// "downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "field" style element options:
//
//  - [Borderwidth]
//  - [Fieldbackground]
//  - [Focuscolor]
//  - [Focuswidth]
//
// "fill" style element options:
//
//  - [Background]
//
// "focus" style element options:
//
//  - [Focuscolor]
//  - [Focussolid]
//  - [Focusthickness]
//
// "hsash" style element options:
//
//  - [Sashthickness]
//
// "hseparator" style element options:
//
//  - [Background]
//  - [Orient]
//
// "image" style element options:
//
//  - [Background]
//  - [Image]
//  - [Stipple]
//
// "indicator" style element options:
//
//
// "label" style element options:
//
//  - [Anchor]
//  - [Background]
//  - [Compound]
//  - [Embossed]
//  - [Font]
//  - [Foreground]
//  - [Image]
//  - [Justify]
//  - [Space]
//  - [Stipple]
//  - [Text]
//  - [Underline]
//  - [Width]
//  - [Wraplength]
//
// "leftarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "padding" style element options:
//
//  - [Padding]
//  - [Relief]
//  - [Shiftrelief]
//
// "pbar" style element options:
//
//  - [Background]
//  - [Barsize]
//  - [Borderwidth]
//  - [Orient]
//  - [Pbarrelief]
//  - [Thickness]
//
// "rightarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "separator" style element options:
//
//  - [Background]
//  - [Orient]
//
// "sizegrip" style element options:
//
//  - [Background]
//  - [Gripsize]
//
// "slider" style element options:
//
//  - [Bordercolor]
//  - [Innercolor]
//  - [Orient]
//  - [Outercolor]
//
// "tab" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Highlight]
//  - [Highlightcolor]
//
// "text" style element options:
//
//  - [Anchor]
//  - [Embossed]
//  - [Font]
//  - [Foreground]
//  - [Justify]
//  - [Text]
//  - [Underline]
//  - [Width]
//  - [Wraplength]
//
// "textarea" style element options:
//
//  - [Font]
//  - [Width]
//
// "thumb" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Orient]
//  - [Relief]
//  - [Width]
//
// "treearea" style element options:
//
//
// "trough" style element options:
//
//  - [Groovewidth]
//  - [Orient]
//  - [Troughborderwidth]
//  - [Troughcolor]
//  - [Troughrelief]
//
// "uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "vsash" style element options:
//
//  - [Sashthickness]
//
// "vseparator" style element options:
//
//  - [Background]
//  - [Orient]
//
// "default" theme style list
//
//  .
//
// Style map: -foreground {disabled #a3a3a3} -background {disabled #d9d9d9 active #ececec}
//
//  Cell
//
// Layout: Treedata.padding -sticky nswe -children {Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Horizontal.Sash
//
// Layout: Sash.hsash -sticky we
//
//  Horizontal.TProgressbar
//
// Layout: Horizontal.Progressbar.trough -sticky nswe -children {Horizontal.Progressbar.pbar -side left -sticky ns Horizontal.Progressbar.ctext -side left -sticky {}}
//
//  Horizontal.TScale
//
// Layout: Horizontal.Scale.focus -sticky nswe -children {Horizontal.Scale.padding -sticky nswe -children {Horizontal.Scale.trough -sticky nswe -children {Horizontal.Scale.slider -side left -sticky {}}}}
//
//  Horizontal.TScrollbar
//
// Layout: Horizontal.Scrollbar.trough -sticky we -children {Horizontal.Scrollbar.leftarrow -side left -sticky {} Horizontal.Scrollbar.rightarrow -side right -sticky {} Horizontal.Scrollbar.thumb -sticky nswe}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  Row
//
// Layout: Treeitem.row -sticky nswe
//
//  Sash
//
// -
//
//
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.border -sticky nswe -border 1 -children {Button.focus -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
// Style map: -relief {{!disabled pressed} sunken}
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
// Style map: -indicatorbackground {{alternate disabled} #a3a3a3 {alternate pressed} #5895bc alternate #4a6984 {selected disabled} #a3a3a3 {selected pressed} #5895bc selected #4a6984 disabled #d9d9d9 pressed #c3c3c3}
//
//  TCombobox
//
// Layout: Combobox.field -sticky nswe -children {Combobox.downarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9} -arrowcolor {disabled #a3a3a3}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  TFrame
//
// Layout: Frame.border -sticky nswe
//
//  TLabel
//
// Layout: Label.border -sticky nswe -border 1 -children {Label.padding -sticky nswe -border 1 -children {Label.label -sticky nswe}}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TLabelframe.Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  TMenubutton
//
// Layout: Menubutton.border -sticky nswe -children {Menubutton.focus -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -side left -sticky {}}}}
//
// Style map: -arrowcolor {disabled #a3a3a3}
//
//  TNotebook
//
// Layout: Notebook.client -sticky nswe
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -highlightcolor {selected #4a6984} -highlight {selected 1} -background {selected #d9d9d9}
//
//  TPanedwindow
//
// Layout: Panedwindow.background -sticky {}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
// Style map: -indicatorbackground {{alternate disabled} #a3a3a3 {alternate pressed} #5895bc alternate #4a6984 {selected disabled} #a3a3a3 {selected pressed} #5895bc selected #4a6984 disabled #d9d9d9 pressed #c3c3c3}
//
//  TScale
//
// Style map: -outercolor {active #ececec}
//
//  TScrollbar
//
// Style map: -arrowcolor {disabled #a3a3a3}
//
//  TSeparator
//
// Layout: Separator.separator -sticky nswe
//
//  TSizegrip
//
// Layout: Sizegrip.sizegrip -side bottom -sticky se
//
//  TSpinbox
//
// Layout: Spinbox.field -side top -sticky we -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9} -arrowcolor {disabled #a3a3a3}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -relief {disabled flat selected sunken pressed sunken active raised} -background {pressed #c3c3c3 active #ececec}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled #a3a3a3 selected #ffffff} -background {disabled #d9d9d9 selected #4a6984}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  Vertical.Sash
//
// Layout: Sash.vsash -sticky ns
//
//  Vertical.TProgressbar
//
// Layout: Vertical.Progressbar.trough -sticky nswe -children {Vertical.Progressbar.pbar -side bottom -sticky we}
//
//  Vertical.TScale
//
// Layout: Vertical.Scale.focus -sticky nswe -children {Vertical.Scale.padding -sticky nswe -children {Vertical.Scale.trough -sticky nswe -children {Vertical.Scale.slider -side top -sticky {}}}}
//
//  Vertical.TScrollbar
//
// Layout: Vertical.Scrollbar.trough -sticky ns -children {Vertical.Scrollbar.uparrow -side top -sticky {} Vertical.Scrollbar.downarrow -side bottom -sticky {} Vertical.Scrollbar.thumb -sticky nswe}PASS

//
// # "alt" theme style guide
//
// "Checkbutton.indicator" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Foreground]
//  - [Indicatorcolor]
//  - [Indicatormargin]
//  - [Lightcolor]
//  - [Shadecolor]
//
// "Combobox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "Menubutton.indicator" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Direction]
//
// "Radiobutton.indicator" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Foreground]
//  - [Indicatorcolor]
//  - [Indicatormargin]
//  - [Lightcolor]
//  - [Shadecolor]
//
// "Spinbox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "Spinbox.uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "Treeitem.indicator" style element options:
//
//  - [Foreground]
//  - [Indicatormargins]
//  - [Size]
//
// "arrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "border" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Borderwidth]
//  - [Default]
//  - [Relief]
//
// "downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "field" style element options:
//
//  - [Bordercolor]
//  - [Fieldbackground]
//  - [Focuscolor]
//  - [Focuswidth]
//
// "leftarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "rightarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "slider" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Borderwidth]
//  - [Orient]
//  - [Sliderrelief]
//  - [Sliderthickness]
//
// "thumb" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Orient]
//  - [Relief]
//  - [Width]
//
// "uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Relief]
//
// "alt" theme style list
//
//  .
//
// Style map: -foreground {disabled #a3a3a3} -background {disabled #d9d9d9 active #ececec} -embossed {disabled 1}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.border -sticky nswe -border 1 -children {Button.focus -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
// Style map: -highlightcolor {alternate black} -relief { {pressed !disabled} sunken {active !disabled} raised }
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
// Style map: -indicatorcolor {pressed #d9d9d9 alternate #aaaaaa disabled #d9d9d9}
//
//  TCombobox
//
// Layout: Combobox.field -sticky nswe -children {Combobox.downarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9} -arrowcolor {disabled #a3a3a3}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TMenubutton
//
// Layout: Menubutton.border -sticky nswe -children {Menubutton.focus -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -side left -sticky {}}}}
//
//  TNotebook
//
// Layout: Notebook.client -sticky nswe
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -expand {selected {1.5p 1.5p 0.75p 0}} -background {selected #d9d9d9}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
// Style map: -indicatorcolor {pressed #d9d9d9 alternate #aaaaaa disabled #d9d9d9}
//
//  TScale
//
// -
//
//
//
//  TScrollbar
//
// -
//
//
//
//  TSpinbox
//
// Layout: Spinbox.field -side top -sticky we -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9} -arrowcolor {disabled #a3a3a3}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -relief {disabled flat selected sunken pressed sunken active raised} -background {pressed #c3c3c3 active #ececec}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled #a3a3a3 selected #ffffff} -background {disabled #d9d9d9 selected #4a6984}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
// # "clam" theme style guide
//
// "Checkbutton.indicator" style element options:
//
//  - [Indicatorbackground]
//  - [Indicatorforeground]
//  - [Indicatormargin]
//  - [Lowerbordercolor]
//  - [Upperbordercolor]
//
// "Combobox.field" style element options:
//
//  - [Bordercolor]
//  - [Fieldbackground]
//  - [Lightcolor]
//
// "Radiobutton.indicator" style element options:
//
//  - [Indicatorbackground]
//  - [Indicatorforeground]
//  - [Indicatormargin]
//  - [Lowerbordercolor]
//  - [Upperbordercolor]
//
// "Spinbox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "Spinbox.uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "arrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "bar" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "border" style element options:
//
//  - [Bordercolor]
//  - [Borderwidth]
//  - [Darkcolor]
//  - [Lightcolor]
//  - [Relief]
//
// "client" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Lightcolor]
//
// "downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "field" style element options:
//
//  - [Bordercolor]
//  - [Fieldbackground]
//  - [Lightcolor]
//
// "hgrip" style element options:
//
//  - [Bordercolor]
//  - [Gripsize]
//  - [Lightcolor]
//
// "leftarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "pbar" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "rightarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "slider" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "tab" style element options:
//
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Lightcolor]
//
// "thumb" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "trough" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Bordercolor]
//  - [Darkcolor]
//  - [Gripsize]
//  - [Lightcolor]
//  - [Orient]
//  - [Sliderlength]
//  - [Troughcolor]
//
// "vgrip" style element options:
//
//  - [Bordercolor]
//  - [Gripsize]
//  - [Lightcolor]
//
// "clam" theme style list
//
//  .
//
// Style map: -selectforeground {!focus white} -foreground {disabled #999999} -selectbackground {!focus #9e9a91} -background {disabled #dcdad5 active #eeebe7}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Horizontal.Sash
//
// Layout: Sash.hsash -sticky nswe -children {Sash.hgrip -sticky nswe}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Sash
//
// -
//
//
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.border -sticky nswe -border 1 -children {Button.focus -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
// Style map: -lightcolor {pressed #bab5ab} -background {disabled #dcdad5 pressed #bab5ab active #eeebe7} -bordercolor {alternate #000000} -darkcolor {pressed #bab5ab}
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
// Style map: -indicatorbackground {pressed #dcdad5 {!disabled alternate} #5895bc {disabled alternate} #a0a0a0 disabled #dcdad5}
//
//  TCombobox
//
// Layout: Combobox.downarrow -side right -sticky ns Combobox.field -sticky nswe -children {Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}
//
// Style map: -foreground {{readonly focus} #ffffff} -fieldbackground {{readonly focus} #4a6984 readonly #dcdad5} -background {active #eeebe7 pressed #eeebe7} -bordercolor {focus #4a6984} -arrowcolor {disabled #999999}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -lightcolor {focus #6f9dc6} -background {readonly #dcdad5} -bordercolor {focus #4a6984}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TMenubutton
//
// Layout: Menubutton.border -sticky nswe -children {Menubutton.focus -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -side left -sticky {}}}}
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -lightcolor {selected #eeebe7 {} #cfcdc8} -padding {selected {4.5p 3p 4.5p 1.5p}} -background {selected #dcdad5 {} #bab5ab}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
// Style map: -indicatorbackground {pressed #dcdad5 {!disabled alternate} #5895bc {disabled alternate} #a0a0a0 disabled #dcdad5}
//
//  TScale
//
// -
//
//
//
//  TScrollbar
//
// -
//
//
//
//  TSpinbox
//
// Layout: Spinbox.field -side top -sticky we -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}
//
// Style map: -background {readonly #dcdad5} -bordercolor {focus #4a6984} -arrowcolor {disabled #999999}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -lightcolor {pressed #bab5ab} -relief {disabled flat selected sunken pressed sunken active raised} -background {disabled #dcdad5 pressed #bab5ab active #eeebe7} -darkcolor {pressed #bab5ab}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled #999999 selected #ffffff} -background {disabled #dcdad5 selected #4a6984} -bordercolor {focus #4a6984}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  Vertical.Sash
//
// Layout: Sash.vsash -sticky nswe -children {Sash.vgrip -sticky nswe}
//
// # "classic" theme style guide
//
// "Button.border" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Default]
//  - [Relief]
//
// "Checkbutton.indicator" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Indicatorcolor]
//  - [Indicatormargin]
//  - [Indicatorrelief]
//  - [Indicatorsize]
//
// "Combobox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Menubutton.indicator" style element options:
//
//  - [Background]
//  - [Indicatorborderwidth]
//  - [Indicatorheight]
//  - [Indicatormargin]
//  - [Indicatorrelief]
//  - [Indicatorwidth]
//
// "Radiobutton.indicator" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Indicatorcolor]
//  - [Indicatormargin]
//  - [Indicatorrelief]
//  - [Indicatorsize]
//
// "Spinbox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Spinbox.uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "arrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "downarrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "highlight" style element options:
//
//  - [Default]
//  - [Highlightcolor]
//  - [Highlightthickness]
//
// "hsash" style element options:
//
//  - [Background]
//  - [Handlepad]
//  - [Handlesize]
//  - [Sashpad]
//  - [Sashrelief]
//  - [Sashthickness]
//
// "leftarrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "rightarrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "slider" style element options:
//
//  - [Background]
//  - [Orient]
//  - [Sliderborderwidth]
//  - [Sliderlength]
//  - [Sliderrelief]
//  - [Sliderthickness]
//
// "uparrow" style element options:
//
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "vsash" style element options:
//
//  - [Background]
//  - [Handlepad]
//  - [Handlesize]
//  - [Sashpad]
//  - [Sashrelief]
//  - [Sashthickness]
//
// "classic" theme style list
//
//  .
//
// Style map: -highlightcolor {focus black} -foreground {disabled #a3a3a3} -background {disabled #d9d9d9 active #ececec}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Horizontal.TScale
//
// Layout: Horizontal.Scale.highlight -sticky nswe -children {Horizontal.Scale.trough -sticky nswe -children {Horizontal.Scale.slider -side left -sticky {}}}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Sash
//
// -
//
//
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.highlight -sticky nswe -children {Button.border -sticky nswe -border 1 -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
// Style map: -relief {{!disabled pressed} sunken}
//
//  TCheckbutton
//
// Layout: Checkbutton.highlight -sticky nswe -children {Checkbutton.border -sticky nswe -children {Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.label -side left -sticky nswe}}}
//
// Style map: -indicatorrelief {alternate raised selected sunken pressed sunken} -indicatorcolor {pressed #d9d9d9 alternate #b05e5e selected #b03060}
//
//  TCombobox
//
// Layout: Combobox.highlight -sticky nswe -children {Combobox.field -sticky nswe -children {Combobox.downarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  TEntry
//
// Layout: Entry.highlight -sticky nswe -children {Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TMenubutton
//
// Layout: Menubutton.highlight -sticky nswe -children {Menubutton.border -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -sticky {}}}}
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -background {selected #d9d9d9}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.highlight -sticky nswe -children {Radiobutton.border -sticky nswe -children {Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.label -side left -sticky nswe}}}
//
// Style map: -indicatorrelief {alternate raised selected sunken pressed sunken} -indicatorcolor {pressed #d9d9d9 alternate #b05e5e selected #b03060}
//
//  TScale
//
// Style map: -sliderrelief {{pressed !disabled} sunken}
//
//  TScrollbar
//
// Style map: -relief {{pressed !disabled} sunken}
//
//  TSpinbox
//
// Layout: Spinbox.highlight -sticky nswe -children {Spinbox.field -sticky nswe -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.focus -sticky nswe -children {Toolbutton.border -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -relief {disabled flat selected sunken pressed sunken active raised} -background {pressed #b3b3b3 active #ececec}
//
//  Treeview
//
// Layout: Treeview.highlight -sticky nswe -children {Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}}
//
// Style map: -foreground {disabled #a3a3a3 selected #000000} -background {disabled #d9d9d9 selected #c3c3c3}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  Vertical.TScale
//
// Layout: Vertical.Scale.highlight -sticky nswe -children {Vertical.Scale.trough -sticky nswe -children {Vertical.Scale.slider -side top -sticky {}}}
//
// # "default" theme style guide
//
// "" style element options:
//
//
// "Checkbutton.indicator" style element options:
//
//  - [Bordercolor]
//  - [Indicatorbackground]
//  - [Indicatorforeground]
//  - [Indicatormargin]
//
// "Combobox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Menubutton.indicator" style element options:
//
//  - [Arrowcolor]
//  - [Arrowpadding]
//  - [Arrowsize]
//
// "Radiobutton.indicator" style element options:
//
//  - [Bordercolor]
//  - [Indicatorbackground]
//  - [Indicatorforeground]
//  - [Indicatormargin]
//
// "Spinbox.downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Spinbox.uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Treeheading.cell" style element options:
//
//  - [Background]
//  - [Rownumber]
//
// "Treeitem.indicator" style element options:
//
//  - [Foreground]
//  - [Indicatormargins]
//  - [Indicatorsize]
//
// "Treeitem.row" style element options:
//
//  - [Background]
//  - [Rownumber]
//
// "Treeitem.separator" style element options:
//
//  - [Background]
//  - [Rownumber]
//
// "arrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "background" style element options:
//
//  - [Background]
//
// "border" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "client" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Highlight]
//  - [Highlightcolor]
//
// "ctext" style element options:
//
//  - [Anchor]
//  - [Embossed]
//  - [Font]
//  - [Foreground]
//  - [Justify]
//  - [Text]
//  - [Underline]
//  - [Width]
//  - [Wraplength]
//
// "downarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "field" style element options:
//
//  - [Borderwidth]
//  - [Fieldbackground]
//  - [Focuscolor]
//  - [Focuswidth]
//
// "fill" style element options:
//
//  - [Background]
//
// "focus" style element options:
//
//  - [Focuscolor]
//  - [Focussolid]
//  - [Focusthickness]
//
// "hsash" style element options:
//
//  - [Sashthickness]
//
// "hseparator" style element options:
//
//  - [Background]
//  - [Orient]
//
// "image" style element options:
//
//  - [Background]
//  - [Image]
//  - [Stipple]
//
// "indicator" style element options:
//
//
// "label" style element options:
//
//  - [Anchor]
//  - [Background]
//  - [Compound]
//  - [Embossed]
//  - [Font]
//  - [Foreground]
//  - [Image]
//  - [Justify]
//  - [Space]
//  - [Stipple]
//  - [Text]
//  - [Underline]
//  - [Width]
//  - [Wraplength]
//
// "leftarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "padding" style element options:
//
//  - [Padding]
//  - [Relief]
//  - [Shiftrelief]
//
// "pbar" style element options:
//
//  - [Background]
//  - [Barsize]
//  - [Borderwidth]
//  - [Orient]
//  - [Pbarrelief]
//  - [Thickness]
//
// "rightarrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "separator" style element options:
//
//  - [Background]
//  - [Orient]
//
// "sizegrip" style element options:
//
//  - [Background]
//  - [Gripsize]
//
// "slider" style element options:
//
//  - [Bordercolor]
//  - [Innercolor]
//  - [Orient]
//  - [Outercolor]
//
// "tab" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Highlight]
//  - [Highlightcolor]
//
// "text" style element options:
//
//  - [Anchor]
//  - [Embossed]
//  - [Font]
//  - [Foreground]
//  - [Justify]
//  - [Text]
//  - [Underline]
//  - [Width]
//  - [Wraplength]
//
// "textarea" style element options:
//
//  - [Font]
//  - [Width]
//
// "thumb" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Orient]
//  - [Relief]
//  - [Width]
//
// "treearea" style element options:
//
//
// "trough" style element options:
//
//  - [Groovewidth]
//  - [Orient]
//  - [Troughborderwidth]
//  - [Troughcolor]
//  - [Troughrelief]
//
// "uparrow" style element options:
//
//  - [Arrowcolor]
//  - [Arrowsize]
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "vsash" style element options:
//
//  - [Sashthickness]
//
// "vseparator" style element options:
//
//  - [Background]
//  - [Orient]
//
// "default" theme style list
//
//  .
//
// Style map: -foreground {disabled #a3a3a3} -background {disabled #d9d9d9 active #ececec}
//
//  Cell
//
// Layout: Treedata.padding -sticky nswe -children {Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Horizontal.Sash
//
// Layout: Sash.hsash -sticky we
//
//  Horizontal.TProgressbar
//
// Layout: Horizontal.Progressbar.trough -sticky nswe -children {Horizontal.Progressbar.pbar -side left -sticky ns Horizontal.Progressbar.ctext -side left -sticky {}}
//
//  Horizontal.TScale
//
// Layout: Horizontal.Scale.focus -sticky nswe -children {Horizontal.Scale.padding -sticky nswe -children {Horizontal.Scale.trough -sticky nswe -children {Horizontal.Scale.slider -side left -sticky {}}}}
//
//  Horizontal.TScrollbar
//
// Layout: Horizontal.Scrollbar.trough -sticky we -children {Horizontal.Scrollbar.leftarrow -side left -sticky {} Horizontal.Scrollbar.rightarrow -side right -sticky {} Horizontal.Scrollbar.thumb -sticky nswe}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  Row
//
// Layout: Treeitem.row -sticky nswe
//
//  Sash
//
// -
//
//
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.border -sticky nswe -border 1 -children {Button.focus -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
// Style map: -relief {{!disabled pressed} sunken}
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
// Style map: -indicatorbackground {{alternate disabled} #a3a3a3 {alternate pressed} #5895bc alternate #4a6984 {selected disabled} #a3a3a3 {selected pressed} #5895bc selected #4a6984 disabled #d9d9d9 pressed #c3c3c3}
//
//  TCombobox
//
// Layout: Combobox.field -sticky nswe -children {Combobox.downarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9} -arrowcolor {disabled #a3a3a3}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9}
//
//  TFrame
//
// Layout: Frame.border -sticky nswe
//
//  TLabel
//
// Layout: Label.border -sticky nswe -border 1 -children {Label.padding -sticky nswe -border 1 -children {Label.label -sticky nswe}}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TMenubutton
//
// Layout: Menubutton.border -sticky nswe -children {Menubutton.focus -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -side left -sticky {}}}}
//
// Style map: -arrowcolor {disabled #a3a3a3}
//
//  TNotebook
//
// Layout: Notebook.client -sticky nswe
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -highlightcolor {selected #4a6984} -highlight {selected 1} -background {selected #d9d9d9}
//
//  TPanedwindow
//
// Layout: Panedwindow.background -sticky {}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
// Style map: -indicatorbackground {{alternate disabled} #a3a3a3 {alternate pressed} #5895bc alternate #4a6984 {selected disabled} #a3a3a3 {selected pressed} #5895bc selected #4a6984 disabled #d9d9d9 pressed #c3c3c3}
//
//  TScale
//
// Style map: -outercolor {active #ececec}
//
//  TScrollbar
//
// Style map: -arrowcolor {disabled #a3a3a3}
//
//  TSeparator
//
// Layout: Separator.separator -sticky nswe
//
//  TSizegrip
//
// Layout: Sizegrip.sizegrip -side bottom -sticky se
//
//  TSpinbox
//
// Layout: Spinbox.field -side top -sticky we -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}
//
// Style map: -fieldbackground {readonly #d9d9d9 disabled #d9d9d9} -arrowcolor {disabled #a3a3a3}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -relief {disabled flat selected sunken pressed sunken active raised} -background {pressed #c3c3c3 active #ececec}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled #a3a3a3 selected #ffffff} -background {disabled #d9d9d9 selected #4a6984}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  Vertical.Sash
//
// Layout: Sash.vsash -sticky ns
//
//  Vertical.TProgressbar
//
// Layout: Vertical.Progressbar.trough -sticky nswe -children {Vertical.Progressbar.pbar -side bottom -sticky we}
//
//  Vertical.TScale
//
// Layout: Vertical.Scale.focus -sticky nswe -children {Vertical.Scale.padding -sticky nswe -children {Vertical.Scale.trough -sticky nswe -children {Vertical.Scale.slider -side top -sticky {}}}}
//
//  Vertical.TScrollbar
//
// Layout: Vertical.Scrollbar.trough -sticky ns -children {Vertical.Scrollbar.uparrow -side top -sticky {} Vertical.Scrollbar.downarrow -side bottom -sticky {} Vertical.Scrollbar.thumb -sticky nswe}
//
// # "vista" theme style guide
//
// "Combobox.background" style element options:
//
//
// "Combobox.border" style element options:
//
//
// "Combobox.rightdownarrow" style element options:
//
//
// "ComboboxPopdownFrame.background" style element options:
//
//
// "Entry.background" style element options:
//
//
// "Entry.field" style element options:
//
//
// "Horizontal.Progressbar.pbar" style element options:
//
//
// "Horizontal.Scale.slider" style element options:
//
//
// "Horizontal.Scrollbar.grip" style element options:
//
//
// "Horizontal.Scrollbar.leftarrow" style element options:
//
//
// "Horizontal.Scrollbar.rightarrow" style element options:
//
//
// "Horizontal.Scrollbar.thumb" style element options:
//
//
// "Horizontal.Scrollbar.trough" style element options:
//
//
// "Menubutton.dropdown" style element options:
//
//
// "Spinbox.background" style element options:
//
//
// "Spinbox.downarrow" style element options:
//
//
// "Spinbox.field" style element options:
//
//
// "Spinbox.innerbg" style element options:
//
//
// "Spinbox.uparrow" style element options:
//
//
// "Vertical.Progressbar.pbar" style element options:
//
//
// "Vertical.Scale.slider" style element options:
//
//
// "Vertical.Scrollbar.downarrow" style element options:
//
//
// "Vertical.Scrollbar.grip" style element options:
//
//
// "Vertical.Scrollbar.thumb" style element options:
//
//
// "Vertical.Scrollbar.trough" style element options:
//
//
// "Vertical.Scrollbar.uparrow" style element options:
//
//
// "vista" theme style list
//
//  .
//
// Style map: -foreground {disabled SystemGrayText}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.background -sticky nswe -border 1 -children {ComboboxPopdownFrame.padding -sticky nswe}
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Horizontal.TProgressbar
//
// Layout: Horizontal.Progressbar.trough -sticky nswe -children {Horizontal.Progressbar.pbar -side left -sticky ns Horizontal.Progressbar.ctext -sticky nswe}
//
//  Horizontal.TScale
//
// Layout: Scale.focus -sticky nswe -children {Horizontal.Scale.trough -sticky nswe -children {Horizontal.Scale.track -sticky we Horizontal.Scale.slider -side left -sticky {}}}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.button -sticky nswe -children {Button.focus -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
//  TCombobox
//
// Layout: Combobox.border -sticky nswe -children {Combobox.rightdownarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.background -sticky nswe -children {Combobox.focus -sticky nswe -children {Combobox.textarea -sticky nswe}}}}
//
// Style map: -focusfill {{readonly focus} SystemHighlight} -foreground {disabled SystemGrayText {readonly focus} SystemHighlightText} -selectforeground {!focus SystemWindowText} -selectbackground {!focus SystemWindow}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -children {Entry.background -sticky nswe -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}}
//
// Style map: -selectforeground {!focus SystemWindowText} -selectbackground {!focus SystemWindow}
//
//  TLabelframe.Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  TMenubutton
//
// Layout: Menubutton.dropdown -side right -sticky ns Menubutton.button -sticky nswe -children {Menubutton.padding -sticky we -children {Menubutton.label -sticky {}}}
//
//  TNotebook
//
// Layout: Notebook.client -sticky nswe
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -expand {selected {2 2 2 2}}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
//  TScale
//
// -
//
//
//
//  TSpinbox
//
// Layout: Spinbox.field -sticky nswe -children {Spinbox.background -sticky nswe -children {Spinbox.padding -sticky nswe -children {Spinbox.innerbg -sticky nswe -children {Spinbox.textarea -sticky nswe}} Spinbox.uparrow -side top -sticky nse Spinbox.downarrow -side bottom -sticky nse}}
//
// Style map: -selectforeground {!focus SystemWindowText} -selectbackground {!focus SystemWindow}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled SystemGrayText selected SystemHighlightText} -background {disabled SystemButtonFace selected SystemHighlight}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  Vertical.TProgressbar
//
// Layout: Vertical.Progressbar.trough -sticky nswe -children {Vertical.Progressbar.pbar -side bottom -sticky we}
//
//  Vertical.TScale
//
// Layout: Scale.focus -sticky nswe -children {Vertical.Scale.trough -sticky nswe -children {Vertical.Scale.track -sticky ns Vertical.Scale.slider -side top -sticky {}}}
//
// # "winnative" theme style guide
//
// "Button.border" style element options:
//
//  - [Default]
//  - [Highlightcolor]
//  - [Relief]
//
// "Checkbutton.indicator" style element options:
//
//
// "Combobox.focus" style element options:
//
//  - [Focusfill]
//
// "ComboboxPopdownFrame.border" style element options:
//
//  - [Background]
//  - [Borderwidth]
//  - [Relief]
//
// "Radiobutton.indicator" style element options:
//
//
// "Scrollbar.trough" style element options:
//
//
// "Spinbox.downarrow" style element options:
//
//
// "Spinbox.uparrow" style element options:
//
//
// "border" style element options:
//
//  - [Relief]
//
// "client" style element options:
//
//
// "downarrow" style element options:
//
//
// "field" style element options:
//
//  - [Fieldbackground]
//
// "focus" style element options:
//
//
// "leftarrow" style element options:
//
//
// "rightarrow" style element options:
//
//
// "sizegrip" style element options:
//
//
// "slider" style element options:
//
//  - [Orient]
//
// "tab" style element options:
//
//  - [Background]
//  - [Borderwidth]
//
// "thumb" style element options:
//
//  - [Orient]
//
// "uparrow" style element options:
//
//
// "winnative" theme style list
//
//  .
//
// Style map: -foreground {disabled SystemGrayText} -embossed {disabled 1}
//
//  ComboboxPopdownFrame
//
// Layout: ComboboxPopdownFrame.border -sticky nswe
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.border -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}
//
// Style map: -relief {{!disabled pressed} sunken}
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
//  TCombobox
//
// Layout: Combobox.field -sticky nswe -children {Combobox.downarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.focus -sticky nswe -children {Combobox.textarea -sticky nswe}}}
//
// Style map: -focusfill {{readonly focus} SystemHighlight} -foreground {disabled SystemGrayText {readonly focus} SystemHighlightText} -selectforeground {!focus SystemWindowText} -fieldbackground {readonly SystemButtonFace disabled SystemButtonFace} -selectbackground {!focus SystemWindow}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -selectforeground {!focus SystemWindowText} -selectbackground {!focus SystemWindow} -fieldbackground {readonly SystemButtonFace disabled SystemButtonFace}
//
//  TLabelframe
//
// Layout: Labelframe.border -sticky nswe
//
//  TLabelframe.Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  TMenubutton
//
// Layout: Menubutton.border -sticky nswe -children {Menubutton.focus -sticky nswe -children {Menubutton.indicator -side right -sticky {} Menubutton.padding -sticky we -children {Menubutton.label -side left -sticky {}}}}
//
//  TNotebook
//
// Layout: Notebook.client -sticky nswe
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -expand {selected {2 2 2 0}}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
//  TScale
//
// -
//
//
//
//  TSpinbox
//
// Layout: Spinbox.field -side top -sticky we -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
// Style map: -relief {disabled flat selected sunken pressed sunken active raised}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled SystemGrayText selected SystemHighlightText} -background {disabled SystemButtonFace selected SystemHighlight}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
// # "xpnative" theme style guide
//
// "Button.button" style element options:
//
//
// "Checkbutton.indicator" style element options:
//
//
// "Combobox.downarrow" style element options:
//
//
// "Combobox.field" style element options:
//
//
// "Entry.field" style element options:
//
//
// "Horizontal.Progressbar.pbar" style element options:
//
//
// "Horizontal.Progressbar.trough" style element options:
//
//
// "Horizontal.Scale.slider" style element options:
//
//
// "Horizontal.Scale.track" style element options:
//
//
// "Horizontal.Scrollbar.grip" style element options:
//
//
// "Horizontal.Scrollbar.thumb" style element options:
//
//
// "Horizontal.Scrollbar.trough" style element options:
//
//
// "Labelframe.border" style element options:
//
//
// "Menubutton.button" style element options:
//
//
// "Menubutton.dropdown" style element options:
//
//
// "NotebookPane.background" style element options:
//
//
// "Radiobutton.indicator" style element options:
//
//
// "Scale.trough" style element options:
//
//
// "Scrollbar.downarrow" style element options:
//
//
// "Scrollbar.leftarrow" style element options:
//
//
// "Scrollbar.rightarrow" style element options:
//
//
// "Scrollbar.uparrow" style element options:
//
//
// "Spinbox.downarrow" style element options:
//
//
// "Spinbox.field" style element options:
//
//
// "Spinbox.uparrow" style element options:
//
//
// "Toolbutton.border" style element options:
//
//
// "Treeheading.border" style element options:
//
//
// "Treeitem.indicator" style element options:
//
//
// "Treeview.field" style element options:
//
//
// "Vertical.Progressbar.pbar" style element options:
//
//
// "Vertical.Progressbar.trough" style element options:
//
//
// "Vertical.Scale.slider" style element options:
//
//
// "Vertical.Scale.track" style element options:
//
//
// "Vertical.Scrollbar.grip" style element options:
//
//
// "Vertical.Scrollbar.thumb" style element options:
//
//
// "Vertical.Scrollbar.trough" style element options:
//
//
// "client" style element options:
//
//
// "sizegrip" style element options:
//
//
// "tab" style element options:
//
//
// "xpnative" theme style list
//
//  .
//
// Style map: -foreground {disabled SystemGrayText}
//
//  Heading
//
// Layout: Treeheading.cell -sticky nswe Treeheading.border -sticky nswe -children {Treeheading.padding -sticky nswe -children {Treeheading.image -side right -sticky {} Treeheading.text -sticky we}}
//
//  Horizontal.TScale
//
// Layout: Scale.focus -sticky nswe -children {Horizontal.Scale.trough -sticky nswe -children {Horizontal.Scale.track -sticky we Horizontal.Scale.slider -side left -sticky {}}}
//
//  Horizontal.TScrollbar
//
// Layout: Horizontal.Scrollbar.trough -sticky we -children {Horizontal.Scrollbar.leftarrow -side left -sticky {} Horizontal.Scrollbar.rightarrow -side right -sticky {} Horizontal.Scrollbar.thumb -sticky nswe -unit 1 -children {Horizontal.Scrollbar.grip -sticky {}}}
//
//  Item
//
// Layout: Treeitem.padding -sticky nswe -children {Treeitem.indicator -side left -sticky {} Treeitem.image -side left -sticky {} Treeitem.text -sticky nswe}
//
//  Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  TButton
//
// Layout: Button.button -sticky nswe -children {Button.focus -sticky nswe -children {Button.padding -sticky nswe -children {Button.label -sticky nswe}}}
//
//  TCheckbutton
//
// Layout: Checkbutton.padding -sticky nswe -children {Checkbutton.indicator -side left -sticky {} Checkbutton.focus -side left -sticky w -children {Checkbutton.label -sticky nswe}}
//
//  TCombobox
//
// Layout: Combobox.field -sticky nswe -children {Combobox.downarrow -side right -sticky ns Combobox.padding -sticky nswe -children {Combobox.focus -sticky nswe -children {Combobox.textarea -sticky nswe}}}
//
// Style map: -focusfill {{readonly focus} SystemHighlight} -foreground {disabled SystemGrayText {readonly focus} SystemHighlightText} -selectforeground {!focus SystemWindowText} -selectbackground {!focus SystemWindow}
//
//  TEntry
//
// Layout: Entry.field -sticky nswe -border 1 -children {Entry.padding -sticky nswe -children {Entry.textarea -sticky nswe}}
//
// Style map: -selectforeground {!focus SystemWindowText} -selectbackground {!focus SystemWindow}
//
//  TLabelframe.Label
//
// Layout: Label.fill -sticky nswe -children {Label.text -sticky nswe}
//
//  TMenubutton
//
// Layout: Menubutton.dropdown -side right -sticky ns Menubutton.button -sticky nswe -children {Menubutton.padding -sticky we -children {Menubutton.label -sticky {}}}
//
//  TNotebook
//
// Layout: Notebook.client -sticky nswe
//
//  TNotebook.Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
// Style map: -expand {selected {2 2 2 2}}
//
//  TProgressbar
//
// -
//
//
//
//  TRadiobutton
//
// Layout: Radiobutton.padding -sticky nswe -children {Radiobutton.indicator -side left -sticky {} Radiobutton.focus -side left -sticky {} -children {Radiobutton.label -sticky nswe}}
//
//  TScale
//
// -
//
//
//
//  TScrollbar
//
// -
//
//
//
//  TSpinbox
//
// Layout: Spinbox.field -side top -sticky we -children {null -side right -sticky {} -children {Spinbox.uparrow -side top -sticky e Spinbox.downarrow -side bottom -sticky e} Spinbox.padding -sticky nswe -children {Spinbox.textarea -sticky nswe}}
//
// Style map: -selectforeground {!focus SystemWindowText} -selectbackground {!focus SystemWindow}
//
//  Tab
//
// Layout: Notebook.tab -sticky nswe -children {Notebook.padding -side top -sticky nswe -children {Notebook.focus -side top -sticky nswe -children {Notebook.label -side top -sticky {}}}}
//
//  Toolbutton
//
// Layout: Toolbutton.border -sticky nswe -children {Toolbutton.focus -sticky nswe -children {Toolbutton.padding -sticky nswe -children {Toolbutton.label -sticky nswe}}}
//
//  Treeview
//
// Layout: Treeview.field -sticky nswe -border 1 -children {Treeview.padding -sticky nswe -children {Treeview.treearea -sticky nswe}}
//
// Style map: -foreground {disabled SystemGrayText selected SystemHighlightText} -background {disabled SystemButtonFace selected SystemHighlight}
//
//  Treeview.Separator
//
// Layout: Treeitem.separator -sticky nswe
//
//  Vertical.TScale
//
// Layout: Scale.focus -sticky nswe -children {Vertical.Scale.trough -sticky nswe -children {Vertical.Scale.track -sticky ns Vertical.Scale.slider -side top -sticky {}}}
//
//  Vertical.TScrollbar
//
// Layout: Vertical.Scrollbar.trough -sticky ns -children {Vertical.Scrollbar.uparrow -side top -sticky {} Vertical.Scrollbar.downarrow -side bottom -sticky {} Vertical.Scrollbar.thumb -sticky nswe -unit 1 -children {Vertical.Scrollbar.grip -sticky {}}}PASS
