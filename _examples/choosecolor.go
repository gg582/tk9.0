package main

import . "github.com/yoonjin67/tk9.0"

func main() {
	b := Button(Txt("Choose Color"))
	b.Configure(Command(func() { b.Configure(Background(ChooseColor(Initialcolor("gray"), Title("Choose color")))) }))
	Pack(b, Exit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	App.Wait()
}
