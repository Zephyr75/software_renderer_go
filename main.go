package main

import (
	"image/color"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"overdrive/utilities"
	"overdrive/mesh"
	"overdrive/geometry"
	"overdrive/render"
	
	"fmt"
	"time"
	"sync"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	img := image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	viewport := canvas.NewImageFromImage(img)
	myCanvas.SetContent(viewport)

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

		//make an array of 20 filled with cube
		cubes := make([]mesh.Mesh, 20)

		for {
			// for x := 0; x < utilities.RESOLUTION_X; x++ {
			// 	for y := 0; y < utilities.RESOLUTION_Y; y++ {
			// 		src.Set(x, y, color.Black)
			// 	}
			// }

			img = image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))

			for i := range cubes {
				cubes[i] = cube
			}

			wg := sync.WaitGroup{}

			for i := range cubes {

				wg.Add(1)
				go func(i int) {
					cubes[i].Draw(img, &cam, []render.Light{light})
					wg.Done()
				}(i)
				
			}

			wg.Wait()

			cube.Translate(geometry.VectorNew(0, 1, 0))
			cube.Rotate(geometry.VectorNew(0, 0.01, 0))

			viewport.Image = img
			viewport.Refresh()

			t := time.Since(start).Milliseconds()
			if t == 0 {
				t = 1
			}
			fmt.Println("fps: ", 1000 / t)
			start = time.Now()
		}
	}()

	myWindow.Resize(fyne.NewSize(utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	myWindow.ShowAndRun()
}
