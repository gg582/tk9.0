package main

import (
    "fmt"
    . "modernc.org/tk9.0"
)

func main() {
    tr := TreeView()

    tr.SelectMode("extended")

    // configure the set of data‑columns by name
    tr.Columns([]string{"name", "birthyear", "guitar", "region", "style"})
    // show only the headings (no tree column text)
    tr.Configure("-show", "headings")

    // set up the heading for the implicit tree column (#0) as “Artist”
    tr.Heading("#0", "-text", Quote("Artist"), "-anchor", "w")
    // set up headings for each data‑column by its name
    tr.Heading("name",      "-text", Quote("Name"),        "-anchor", "w")
    tr.Heading("birthyear", "-text", Quote("Birth Year"), "-anchor", "center")
    tr.Heading("guitar",    "-text", Quote("Guitar"),     "-anchor", "w")
    tr.Heading("region",    "-text", Quote("Region"),     "-anchor", "w")
    tr.Heading("style",     "-text", Quote("Style"),      "-anchor", "w")

    // configure column widths & anchors
    tr.Column("#0",     "-width 140 -anchor w")
    tr.Column("name",   "-width 200 -anchor w")
    tr.Column("birthyear", "-width 80 -anchor center")
    tr.Column("guitar", "-width 200 -anchor w")
    tr.Column("region", "-width 200 -anchor w")
    tr.Column("style",  "-width 200 -anchor w")

    // insert sample blues artists: text→tree‑col, values→other cols
    artists := []struct {
        id     string
        text   string
        values []string
    }{
        {"sonhouse", Quote("Son House"),           []string{Quote("Son House"), Quote("1902"), Quote("Resonator Guitar"), Quote("Mississippi"), Quote("Delta Blues")}},
        {"williemctell", Quote("Blind Willie McTell"), []string{Quote("Blind Willie McTell"), Quote("1898"), Quote("12-String Guitar"), Quote("Georgia"), Quote("Ragtime Blues")}},
        {"mrburnside", Quote("R. L. Burnside"),      []string{Quote("R. L. Burnside"), Quote("1926"), Quote("Electric Guitar"), Quote("North Mississippi"), Quote("Hill Country Blues")}},
        {"leadbelly", Quote("Lead Belly"),          []string{Quote("Lead Belly"), Quote("1888"), Quote("Stella Jumbo 12-String"), Quote("Louisiana"), Quote("Folk/Blues")}},
    }
    for i, a := range artists {
        tr.Insert(a.id, a.text, 0, i, a.values)
    }

    // button to print selected IDs
    printBtn := Button(Txt("Show Selected"), Command(func() {
        for _, itm := range tr.Selected() {
            fmt.Println("Selected item IDs:", itm)
			fmt.Println(tr.GetColumn(itm, "guitar"))
        }
    }))

    Pack(
        TExit(),
        printBtn,
        Label(Txt("Legendary Blues Artists")),
        tr,
    )
    App.Wait()
}

