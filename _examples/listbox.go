package main

import . "github.com/yoonjin67/tk9.0"
import "fmt"

func main() {
	l := Listbox()
	l2 := Listbox()
    l.SelectMode("multiple")
    l2.SelectMode("multiple")

	l.Configure(l.Options("background", "lightgrey") + " " +
		l.Options("foreground", "darkblue") + " " +
		l.Options("selectbackground", "lightblue") + " " +
		l.Options("selectforeground", "black") + " " +
		l.Options("relief", "sunken") + " " +
		l.Options("highlightthickness", "2") + " " +
		l.Options("highlightcolor", "green") + " " +
		l.Options("highlightbackground", "white") + " " +
		l.Options("height", "8") + " " +
		l.Options("width", "15"))

	l.AddItems(0, "Apple Banana Cherry Date Elderberry Fig Grape Honeydew")

	l.ItemForeground(0, "red")
	l.ItemBackground(1, "yellow")
	l.ItemSelectForeground(3, "white")
	l.ItemSelectBackground(3, "purple")

	selectedVarStore := make([]string, 0, 100)
	b := Button(Txt("Get Selected"), Command(func() {
		l2.DeleteItems(0, len(selectedVarStore)-1)
		selectedVarStore = selectedVarStore[:0]
		sel := l.Selected()
		fmt.Println("Selected Indices:", sel)
		for i, index := range sel {
			itm := l.Get(index, index)
			if len(itm) > 0 {
				selectedVarStore = append(selectedVarStore, itm[0])
				l2.AddItems(i, itm[0])
			}
		}
		fmt.Println("Selected Items:", selectedVarStore)
	}))

	clearButton := Button(Txt("Clear List 2"), Command(func() {
		l2.DeleteItems(0, len(selectedVarStore))
		selectedVarStore = selectedVarStore[:0]
	}))

	l2.Configure(l2.Options("background", "whitesmoke") + " " +
		l2.Options("foreground", "darkgreen") + " " +
		l2.Options("height", "5") + " " +
		l2.Options("width", "15") + " " +
		l2.Options("relief", "groove"))

	Pack(TExit(), b, clearButton, l, Label(Txt("Selected Items:")), l2)
	App.Wait()
}
