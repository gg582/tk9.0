// Tooltip demo

package main

import . "github.com/yoonjin67/tk9.0"

func main() {
	Pack(Tooltip(TLabel(Txt("Tooltip Label")), "This is a label"),
		TExit(),
		Ipadx(10), Ipady(5), Padx(20), Pady(10))
	App.Wait()
}
