package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"overdrive/test"
	"fmt"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Line")

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 1
	w.SetContent(line)

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()


	fmt.Println(test.ReverseRunes("!oG ,olleH"))
}