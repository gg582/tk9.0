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
	buttonFocusDecoratorCorner = 9.0 / 96 // The rounded corner is 9px on a 96 DPI display.
	buttonFocusDecorator       = 4.0 / 96 // 4px on a 96 DPI display.
)

var (
	corners = map[cornerKey][4]*Img{}
	tiles   = map[string]*Img{}
)

type ButtonColor int

const (
	_ ButtonColor = iota
	ButtonFace
	ButtonText
)

type ButtonColors map[ButtonColor]string

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

func getTile(color string) (r *Img) {
	if ex, ok := tiles[color]; ok {
		return ex
	}

	r = NewPhoto(Width(8), Height(8),
		Data(fmt.Sprintf(`<svg width="8" height="8" fill=%q><rect width="8" height="8" fill=%[1]q/></svg>`, color)))
	tiles[color] = r
	return r
}

// ButtonStyle defines a button style. ATM only when using the "default" theme.
func ButtonStyle(style string, scheme ButtonColors, background string) string {
	width := TkScaling() * 72 * buttonFocusDecoratorCorner
	stroke := TkScaling() * 72 * buttonFocusDecorator
	r := width - stroke/2
	clip := width - stroke
	corners := getCorners(round(width), round(clip), round(r), round(stroke), scheme[ButtonFace], background, background)
	q1 := "Q1." + style + ".corner"
	q2 := "Q2." + style + ".corner"
	q3 := "Q3." + style + ".corner"
	q4 := "Q4." + style + ".corner"
	StyleElementCreate(q1, "image", corners[0])
	StyleElementCreate(q2, "image", corners[1])
	StyleElementCreate(q3, "image", corners[2])
	StyleElementCreate(q4, "image", corners[3])
	tile := "Tile." + style + ".tile"
	t := getTile(scheme[ButtonFace])
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
					"Button.label", Sticky("nswe"),
				))))
	StyleConfigure(style, Compound(true))
	StyleConfigure(style, Foreground(scheme[ButtonText]))
	StyleConfigure(style, Background(background))
	StyleConfigure(style, Borderwidth(5))
	StyleConfigure(style, Relief("flat"))
	StyleConfigure(style, Padding("0"))
	StyleConfigure(style, FocusColor(Black))
	StyleConfigure(style, FocusThickness(0))
	StyleConfigure(style, FocusSolid("solid"))
	return style
}
