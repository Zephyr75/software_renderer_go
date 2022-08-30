package main

import (
	"image"
	"image/color"

	"overdrive/geometry"
	"overdrive/mesh"
	"overdrive/render"
	"overdrive/utilities"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"fmt"
	"sync"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	img := image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	viewport := canvas.NewImageFromImage(img)
	myCanvas.SetContent(viewport)

	go func() {

		cube := mesh.Cube(geometry.VectorNew(0, 0, 500), geometry.VectorZero(), geometry.VectorNew(400, 400, 400))
		cam := render.Camera{
			Position: geometry.VectorNew(0, 0, 0),
			Rotation: geometry.VectorZero()}
		light := render.Light{
			Position:  geometry.VectorZero(),
			Rotation:  geometry.VectorNew(0, 0, -800),
			LightType: render.Directional,
			Color:     color.RGBA{255, 255, 255, 255},
			Length:    0,
		}

		start := time.Now()

		//make an array of 20 filled with cube
		cubes := make([]mesh.Mesh, 50)

		for {
			// fmt.Println("cube: ", cube.Position)
			// fmt.Println("cam: ", cam.Position)

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

			} //TODO Aberty666

			wg.Wait()

			//cube.Translate(geometry.VectorNew(0, 0, 1))
			//cube.Rotate(geometry.VectorNew(0, 0.01, 0))

			viewport.Image = img
			viewport.Refresh()

			t := time.Since(start).Milliseconds()
			if t == 0 {
				t = 1
			}
			fmt.Println("fps: ", 1000/t)
			start = time.Now()
			// break
		}
	}()

	myWindow.Resize(fyne.NewSize(utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	myWindow.ShowAndRun()
}
