package main

import (
	. "modernc.org/tk9.0"
	"modernc.org/tk9.0/b5"
)

func main() {
	StyleThemeUse("default")
	opts := Opts{Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	Grid(TButton(Txt("Primary"), Style(b5.ButtonStyle("Primary.B5.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#0d6efd"}, "#fff"))),
		TButton(Txt("Secondary"), Style(b5.ButtonStyle("Secondary.B5.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#6c757d"}, "#fff"))),
		TButton(Txt("Success"), Style(b5.ButtonStyle("Sucess.B5.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#198754"}, "#fff"))),
		opts)
	Grid(TButton(Txt("Danger"), Style(b5.ButtonStyle("Danger.B5.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#dc3545"}, "#fff"))),
		TButton(Txt("Warning"), Style(b5.ButtonStyle("Warning.B5.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#ffc107"}, "#fff"))),
		TButton(Txt("Info"), Style(b5.ButtonStyle("Info.B5.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#0dcaf0"}, "#fff"))),
		opts)
	Grid(TButton(Txt("Light"), Style(b5.ButtonStyle("Light.B5.TButton", b5.ButtonColors{b5.ButtonText: "#000", b5.ButtonFace: "#f8f9fa"}, "#fff"))),
		TButton(Txt("Dark"), Style(b5.ButtonStyle("Dark.B5.TButton", b5.ButtonColors{b5.ButtonText: "#fff", b5.ButtonFace: "#212529"}, "#fff"))),
		TButton(Txt("Link"), Style(b5.ButtonStyle("Link.B5.TButton", b5.ButtonColors{b5.ButtonText: "#1774fd", b5.ButtonFace: "#fff"}, "#fff"))),
		opts)
	Grid(TExit(), Columnspan(3), opts)
	App.Configure(Background("#fff")).Wait()
}
