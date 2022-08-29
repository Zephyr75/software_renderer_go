package main

import (
	"fmt"
	"image"
	"image/color"

	//"fyne.io/fyne/v2"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	//"fyne.io/fyne/v2/theme"

	"overdrive/geometry"
	"overdrive/mesh"
	"overdrive/render"
	"overdrive/utilities"
	"time"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Overdrive")
	myCanvas := myWindow.Canvas()

	src := image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	img := canvas.NewImageFromImage(src)
	myCanvas.SetContent(img)

	go func() {
		cube := mesh.Cube(geometry.VectorNew(100, 0, 0), geometry.VectorNew(100, 0, 0), geometry.VectorNew(400, 400, 400))
		cam := render.Camera{
			Position: geometry.VectorZero(), 
			Rotation: geometry.VectorNew(0, 0, -800)}
		light := render.Light{
			Position: geometry.VectorZero(),
			Rotation: geometry.VectorZero(),
			LightType: render.Ambient,
			Color: color.White,
			Length: 0,
		}
	
		start := time.Now()
		src = image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))

		for i := 0; i < 10000; i++ {
			cube.Draw(src, &cam, []render.Light{light})
			cube.Translate(geometry.VectorNew(0, 1, 0))
			img = canvas.NewImageFromImage(src)
			img.FillMode = canvas.ImageFillOriginal
			//img.FillColor = green
			img.Refresh()
			
			fmt.Println("fps: ", 1000 / time.Since(start).Milliseconds())
			start = time.Now()
		}
	}()

	myWindow.Resize(fyne.NewSize(1000, 500))
	myWindow.ShowAndRun()
}
