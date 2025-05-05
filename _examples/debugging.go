// Build this program using -tags=tk.dmesg

package main

import . "github.com/yoonjin67/tk9.0"

func main() {
	Pack(
		TButton(Txt("Hello"), Command(func() { Destroy(App) })),
		Ipadx(10), Ipady(5), Padx(20), Pady(10),
	)
	App.Wait()
}
