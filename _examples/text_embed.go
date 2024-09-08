package main

import . "modernc.org/tk9.0"
import _ "embed"

//go:embed gotk.png
var icon []byte
var scale = TkScaling()*72/600
const tex = `$Q(\xi) = \lambda_1 y_1^2 \sum_{i=2}^n \sum_{j=2}^n y_i b_{ij} y_j$`

func main() {
	println(scale)
	var scroll *TScrollbarWidget
	t := Text(Fnt("helvetica 10"), Yscrollcommand(func(e *Event) { e.ScrollSet(scroll) }), Setgrid(true), Wrap("word"), Padx("2m"), Pady("2m"))
	scroll = TScrollbar(Command(func(e *Event) { e.Yview(t) }))
	Grid(t, Sticky("news"), Pady("2m"))
	Grid(scroll, Row(0), Column(1), Sticky("nes"), Pady("2m"))
	GridRowConfigure(App, 0, Weight(1))
	GridColumnConfigure(App, 0, Weight(1))
	Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	t.TagConfigure("center", Justify("center"))
	t.InsertML(`Hello Go`, NewPhoto(Data(icon)), Padx("1m"),`programmers!`)
	t.InsertML(`<br><br>Hello TeX`, NewPhoto(Data(TeX(tex, scale))), Padx("1m"), `users!`)
	t.InsertML(`<br><br>Hello TeX`, NewPhoto(Data(TeX("$"+tex+"$", scale))), Padx("1m"), `users!`)
	App.Center().Wait()
}
