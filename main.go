package main

import (
	"image"
	"image/color"

	//"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	//"fyne.io/fyne/v2/theme"

	"overdrive/mesh"
	"overdrive/render"
	"overdrive/utilities"
	"overdrive/geometry"

)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	cube := mesh.Cube(geometry.VectorZero(), geometry.VectorZero(), geometry.VectorNew(400, 400, 400))

	
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

	cube.Draw(src, &cam, []render.Light{light})

	// for i := 0; i < 500; i++ {
	// 	go src.Set(250, i, color.Black)
	// 	go src.Set(i, i, color.Black)
	// }
	// time.Sleep(10 * time.Millisecond)


	image := canvas.NewImageFromImage(src)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
}
