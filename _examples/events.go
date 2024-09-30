package main

import (
	"fmt"

	. "modernc.org/tk9.0"
)

func h(e *Event) {
	fmt.Println(e.W)
}

func main() {
	b1 := TButton(Txt("Hello"), Command(h))
	b2 := TButton(Txt("World"), Command(h))
	fmt.Printf("Button Hello = %s, button World = %s\n", b1, b2)
	opts := Opts{Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m")}
	Grid(b1, b2, opts)
	Grid(TExit(), Columnspan(2), opts)
	App.Wait()
}
