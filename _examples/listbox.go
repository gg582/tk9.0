package main

import . "modernc.org/tk9.0"

func main() {
    l := Listbox()
    l2 := Listbox()
    l.AddItems(0, "white red blue")
    l.Background("grey")
    l.Foreground("blue")
    l.SelectBackground("blue")
    l.SelectForeground("skyblue")
    l.Relief("flat")
    l.ItemForeground(0,"white")
    l.ItemForeground(1,"red")
    l.ItemForeground(2,"blue")
    l.ItemSelectForeground(2,"green")
    l.HighlightThickness(1)
    l.HighlightColor("yellow")
    l.HighlightBackground("violet")
    l.SelectMode("multiple")
    l.Height(10)
    l.Width(4)
    selected_varStore := make([]string, 0, 100)
    b := Button(Txt("Get Selected Items"), Command( func() {
        l2.DeleteItems(0,len(selected_varStore)-1)
        selected_varStore = selected_varStore[:0]
        sel := l.Selected()
        for _, i := range sel {
            println(i)
        }
        for i, c := range sel {
            itm := l.Get(c,c)
            selected_varStore = append(selected_varStore, itm[0])
            l2.AddItems(i,itm[0])
        }
    }))


    
	Pack(TExit(),b,l,l2)
	App.Wait()
}
