package main

import . "modernc.org/tk9.0"
import _ "embed"

//go:embed gotk.png
var icon []byte

func main() {
	var scroll *TScrollbarWidget
	t := Text(Font("helvetica", 10), Height(15), Yscrollcommand(func(e *Event) { e.ScrollSet(scroll) }), Setgrid(true), Wrap("word"), Padx("4p"), Pady("4p"), Pady("12p"))
	scroll = TScrollbar(Command(func(e *Event) { e.Yview(t) }))
	Grid(t, Sticky("news"), Pady("2m"))
	Grid(scroll, Row(0), Column(1), Sticky("nes"), Pady("2m"))
	GridRowConfigure(App, 0, Weight(1))
	GridColumnConfigure(App, 0, Weight(1))
	Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	t.TagConfigure("c", Justify("center"))
	t.TagConfigure("e", Offset("-2p"))
	t.TagConfigure("t", Font("times"))
	t.InsertML(`<c>Hello Go + Tk`, NewPhoto(Data(icon)), Padx("4p"), `users!</c>
<br><br><c>Hello Go + Tk + <t>T<e>E</e>X</t> $Q(\xi) = \lambda_1 y_1^2 \sum_{i=2}^n \sum_{j=2}^n y_i b_{ij} y_j$ users! (\$inline math\$)</c>
<br><br><c>Hello Go + Tk + <t>T<e>E</e>X</t> $$Q(\xi) = \lambda_1 y_1^2 \sum_{i=2}^n \sum_{j=2}^n y_i b_{ij} y_j$$ users! (\$\$display math\$\$)</c>
<br><br>The above exemplifies embeding pictures and <t>T<e>E</e>X</t> scripts. A text widget can also embed other widgets. For example,
when a`, TButton(Txt("<Tbutton>")), Padx("4p"), Pady("2p"), Align("center"), "and a", TEntry(Textvariable("<TEntry>"), Width(8)), Padx("4p"),
		Pady("2p"), Align("center"), `are part of the markup, they will reflow when their containing text widget is resized.`)
	App.Center().Wait()
}
