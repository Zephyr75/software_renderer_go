package main

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	//"fyne.io/fyne/v2/theme"

	"fmt"
	"overdrive/test"
	"time"
	"github.com/StephaneBunel/bresenham"
	"overdrive/mesh"
	//"overdrive/render"

)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")


	vect := mesh.Vector3{
		1,
		2,
		3,
		color.Black,
	}

	//set src to a white image of 500 x 500 with a black pixel in the middle
	src := image.NewRGBA(image.Rect(0, 0, 500, 500))
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			src.Set(x, y, color.White)
		}
	}

	for i := 0; i < 500; i++ {
		go src.Set(250, i, color.Black)
		go src.Set(i, i, color.Black)
	}
	time.Sleep(10 * time.Millisecond)

	bresenham.DrawLine(src, 14, 71, 441, 317, color.Black)

	// image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	// image := canvas.NewImageFromFile(fileName)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()

	fmt.Println(test.ReverseRunes("!oG ,olleH"))
}

func drawLine(i int, w fyne.Window) {
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 1
	line.Position1 = fyne.NewPos(float32(10*i), 0)
	line.Position2 = fyne.NewPos(float32(10*i), 50)
	w.SetContent(line)
}
