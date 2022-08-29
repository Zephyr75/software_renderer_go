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
	w := myApp.NewWindow("Image")

	

	// for i := 0; i < 500; i++ {
	// 	go src.Set(250, i, color.Black)
	// 	go src.Set(i, i, color.Black)
	// }
	// time.Sleep(10 * time.Millisecond)

	go drawWorld(&w)

	//w.SetContent(image)

	w.ShowAndRun()
}


func drawWorld(window *fyne.Window) {
	cube := mesh.Cube(geometry.VectorNew(100, 0, 0), geometry.VectorNew(100, 0, 0), geometry.VectorNew(400, 400, 400))
	
	src := image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	for x := 0; x < utilities.RESOLUTION_X; x++ {
		for y := 0; y < utilities.RESOLUTION_Y; y++ {
			src.Set(x, y, color.Black)
		}
	}

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

	for i := 0; i < 1000; i++ {
		cube.Draw(src, &cam, []render.Light{light})
		cube.Translate(geometry.VectorNew(0, 1, 0))
		image := canvas.NewImageFromImage(src)
		image.FillMode = canvas.ImageFillOriginal
		(*window).SetContent(image)
		(*window).Canvas().Refresh(image)
		//image.Refresh()
		
		fmt.Println("fps: ", 1000 / time.Since(start).Milliseconds())
		start = time.Now()
	}

}