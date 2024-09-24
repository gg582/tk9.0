// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The b5 package is a work in progress with no stable API yet.
package b5 // import "modernc.org/tk9.0/b5"

import (
	"fmt"
	"math"

	. "modernc.org/tk9.0"
)

//TODO const semantics = None?, Primary, Secondary, Success, Info, Warning, Danger, Dark, Light, Link

const (
	buttonFocusDecoratorCorner = 9 / 96.  // The rounded corner is 9px on a 96 DPI display.
	buttonFocusDecorator       = 4 / 96.  // 4px on a 96 DPI display.
	buttonTileHeight           = 27 / 96. // 27px on a 96 DPI display
)

var (
	corners = map[cornerKey][4]*Img{}
	tiles   = map[tileKey]*Img{}
)

type ButtonColor int

const (
	_ ButtonColor = iota
	ButtonFace
	ButtonFocus
	ButtonText
)

type ButtonColors map[ButtonColor]string

type tileKey struct {
	width  int
	height int
	color  string
}

type cornerKey struct {
	width       int
	clip        int
	r           int
	strokeWidth int
	fill        string
	stroke      string
	background  string
}

func round(n float64) int {
	return int(math.Round(n))
}

// All sizes in px
func getCorners(width, clip, r, strokeWidth int, fill, stroke, background string) (re [4]*Img) {
	k := cornerKey{width, clip, r, strokeWidth, fill, stroke, background}
	if ex, ok := corners[k]; ok {
		return ex
	}

	svg := fmt.Sprintf(`<svg>
	<rect width="%[7]d" height="%[7]d" fill=%[6]q />
	<circle r="%[2]d" cx="%[1]d" cy="%[1]d" stroke-width="%[3]d" fill=%q stroke=%q />
</svg>`,
		width, r, strokeWidth, fill, stroke, background, 2*width)
	img := NewPhoto(Data(svg))
	re[0] = NewPhoto(Width(clip), Height(clip))
	re[0].Copy(img, From(width, width-clip, width+clip, width))
	re[1] = NewPhoto(Width(clip), Height(clip))
	re[1].Copy(img, From(width-clip, width-clip, width, width))
	re[2] = NewPhoto(Width(clip), Height(clip))
	re[2].Copy(img, From(width-clip, width, width, width+clip))
	re[3] = NewPhoto(Width(clip), Height(clip))
	re[3].Copy(img, From(width, width, width+clip, width+clip))
	corners[k] = re
	return re
}

// All sizes in px
func getTile(width, height int, color string) (r *Img) {
	k := tileKey{width, height, color}
	if ex, ok := tiles[k]; ok {
		return ex
	}

	r = NewPhoto(Width(width), Height(height),
		Data(fmt.Sprintf(`<svg width="%d" height="%d" fill=%q><rect width="%[1]d" height="%d" fill=%q/></svg>`, width, height, color)))
	tiles[k] = r
	return r
}

// ButtonStyle defines a button style. ATM only when using the "default" theme.
func ButtonStyle(style string, scheme ButtonColors, background string) string {
	width := TkScaling() * 72 * buttonFocusDecoratorCorner
	stroke := TkScaling() * 72 * buttonFocusDecorator
	th := TkScaling() * 72 * buttonTileHeight
	r := width - stroke/2
	clip := width - stroke
	corners := getCorners(round(width), round(clip), round(r), round(stroke), scheme[ButtonFace], background, background)
	q1 := style + ".q1"
	q2 := style + ".q2"
	q3 := style + ".q3"
	q4 := style + ".q4"
	StyleElementCreate(q1, "image", corners[0])
	StyleElementCreate(q2, "image", corners[1])
	StyleElementCreate(q3, "image", corners[2])
	StyleElementCreate(q4, "image", corners[3])
	tile := "Tile." + style + ".tile"
	t := getTile(8, round(th), scheme[ButtonFace])
	StyleElementCreate(tile, "image", t)
	StyleLayout(style,
		"Button.border", Sticky("nswe"), Border(1), Children(
			"Button.focus", Sticky("nswe"), Children(
				"Button.padding", Sticky("nswe"), Children(
					tile,
					q1, Sticky("ne"),
					q2, Sticky("nw"),
					q3, Sticky("sw"),
					q4, Sticky("se"),
					"Button.label", Sticky("nswe")))))
	StyleConfigure(style, Background(background), Borderwidth(10), Compound(true), FocusColor(Black), FocusSolid(false),
		FocusThickness(0), Foreground(scheme[ButtonText]), Padding(0), Relief("flat"), Shiftrelief(0))
	return style
}

// package main
//
// import (
// 	"fmt"
// 	. "modernc.org/tk9.0"
// 	"modernc.org/tk9.0/b5"
// )
//
// func main() {
// 	StyleThemeUse("default")
// 	opts := Opts{Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
// 	pb := TButton(Txt("Primary"), Style(b5.ButtonStyle("primary.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#0d6efd"}, "#fff")))
// 	Grid(pb,
// 		TButton(Txt("Secondary"), Style(b5.ButtonStyle("secondary.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#6c757d"}, "#fff"))),
// 		TButton(Txt("Success"), Style(b5.ButtonStyle("sucess.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#198754"}, "#fff"))),
// 		opts)
// 	Grid(TButton(Txt("Danger"), Style(b5.ButtonStyle("sanger.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#dc3545"}, "#fff"))),
// 		TButton(Txt("Warning"), Style(b5.ButtonStyle("warning.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#ffc107"}, "#fff"))),
// 		TButton(Txt("Info"), Style(b5.ButtonStyle("info.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#0dcaf0"}, "#fff"))),
// 		opts)
// 	Grid(TButton(Txt("Light"), Style(b5.ButtonStyle("light.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#f8f9fa"}, "#fff"))),
// 		TButton(Txt("Dark"), Style(b5.ButtonStyle("dark.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#212529"}, "#fff"))),
// 		TButton(Txt("Link"), Style(b5.ButtonStyle("link.TButton", b5.ButtonColors{b5.ButtonText: "#1774fd", b5.ButtonFace: "#fff"}, "#fff"))),
// 		opts)
// 	Grid(TExit(), Columnspan(3), opts)
// 	//style2 := b5.ButtonStyle2("primary.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#0d6efd", b5.ButtonFocus: "#ff0000"}, "#fff")
// 	Bind("TButton", "<FocusIn>", Command(func(e *Event) {
// 		fmt.Println(e)
// 	}))
// 	App.Configure(Background("#fff")).Wait()
// }
