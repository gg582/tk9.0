package main

import . "modernc.org/tk9.0"
import _ "embed"

//go:embed gotk.png
var icon []byte
var k = TkScaling() * 72 / 600

const tex = `$Q(\xi) = \lambda_1 y_1^2 \sum_{i=2}^n \sum_{j=2}^n y_i b_{ij} y_j$`

func main() {
	var scroll *TScrollbarWidget
	t := Text(Fnt("helvetica 10"), Height(10), Yscrollcommand(func(e *Event) { e.ScrollSet(scroll) }), Setgrid(true), Wrap("word"), Padx("4p"), Pady("4p"), Pady("12p"))
	scroll = TScrollbar(Command(func(e *Event) { e.Yview(t) }))
	Grid(t, Sticky("news"), Pady("2m"))
	Grid(scroll, Row(0), Column(1), Sticky("nes"), Pady("2m"))
	GridRowConfigure(App, 0, Weight(1))
	GridColumnConfigure(App, 0, Weight(1))
	Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	t.TagConfigure("c", Justify("center"))
	t.TagConfigure("e", Offset("-2p"))
	t.TagConfigure("t", Fnt("times"))
	t.InsertML(`<c>Hello Go + Tk`, NewPhoto(Data(icon)), Padx("4p"), `users!</c>
<br><br><c>Hello Go + Tk + <t>T<e>E</e>X</t>`, NewPhoto(Data(TeX(tex, k))), Padx("4p"), Align("top"), `users! (Text formula)</c>
<br><br><c>Hello Go + Tk + <t>T<e>E</e>X</t>`, NewPhoto(Data(TeX("$"+tex+"$", k))), Padx("4p"), Align("top"), `users! (Display formula)</c>`)
	App.Center().Wait()
}
