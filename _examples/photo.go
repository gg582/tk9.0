package main

import _ "embed"
import . "github.com/yoonjin67/tk9.0"

//go:embed gopher.png
var gopher []byte

func main() {
	Pack(Label(Image(NewPhoto(Data(gopher)))),
		TExit(),
		Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	App.Center().Wait()
}
