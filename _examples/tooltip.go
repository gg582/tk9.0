// Tooltip demo

package main

import . "modernc.org/tk9.0"

func main() {
	l := TLabel(Txt("Tooltip Label"))
	Pack(l,
		TExit(Txt("Exit")),
		Ipadx(10), Ipady(5), Padx(20), Pady(10))
	Tooltip(l, "This is a label")
	App.Wait()
}
