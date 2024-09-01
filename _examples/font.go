package main

import . "modernc.org/tk9.0"

func main() {
	t := Text(Wrap("none"))
	for i, v := range FontFamilies() {
		t.TagConfigure(t.TagAdd(t.Insert(LC{i + 1, 0}, v+"\n"), LC{i + 2, 0}), Fnt(NewFont(Family(v))))
	}
	Pack(t, TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	App.Configure(Padx("4m"), Pady("3m")).Center().Wait()
}
