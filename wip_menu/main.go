package main

import sasso "advent-of-code/1st_day"

func main() {
	app := NewMenu("Chose a daily challenge to run:")
	app.AddItem("Day 1 part one", "1f")
	app.AddItem("Day 1 part one", "1s")

	choice := app.Display()
	if choice == "1f" {
		sasso.Sasso()
	} else if choice == "1s" {
		sasso.Sassi()
	}
}

// app := tview.NewApplication()
// text := tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(q) to quit")
// app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
// 	if event.Rune() == 'q' {
// 		app.Stop()
// 	}
// 	return event
// })
// err := app.SetRoot(text, true).EnableMouse(false).Run()
// if err != nil {
// 	panic(err)
// }
