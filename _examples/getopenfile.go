package main

import "fmt"
import . "github.com/gg582/tk9.0"

func main() {
	Pack(
		TButton(Txt("Open File..."), Command(func() {
			fmt.Printf("%q\n", GetOpenFile(
				Title("Open File"),
				Multiple(true),
				Filetypes([]FileType{{"Go files", []string{".go"}, ""}})),
			)
		})),
		TExit(),
		Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"),
	)
	App.Wait()
}
