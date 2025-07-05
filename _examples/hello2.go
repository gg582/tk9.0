// hello.go with some padding added.

package main

import . "github.com/gg582/tk9.0"

func main() {
	Pack(TExit(Txt("Hello")), Ipadx(10), Ipady(5), Padx(20), Pady(10))
	App.Wait()
}
