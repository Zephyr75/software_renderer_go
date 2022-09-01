package main

import (
	"image"
	"image/color"
	"math"

	"overdrive/geometry"
	"overdrive/material"
	"overdrive/mesh"
	"overdrive/render"
	"overdrive/utilities"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"fmt"
	"sync"
	"time"
)

func main() {
	// text1 := canvas.NewText("1", color.White)
	// textFps := canvas.NewText("2", color.White)

	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	img := image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))

	zBuffer := make([]float32, utilities.RESOLUTION_X*utilities.RESOLUTION_Y)

	for i := 0; i < len(zBuffer); i++ {
		zBuffer[i] = math.MaxFloat32
	}

	viewport := canvas.NewImageFromImage(img)

	// grid := container.New(layout.NewGridLayout(2), viewport, textFps, content)

	bottom := widget.NewButton("Assets browser", func() {
		fmt.Println("tapped")
	})

	right := canvas.NewText("fps", color.White)
	// middle := canvas.NewText("content", color.White)
	content := container.New(layout.NewBorderLayout(nil, bottom, nil, right),
		bottom, right, viewport)

	myCanvas.SetContent(content)

	go func() {

		cam := render.Camera{
			Position: geometry.NewVector(0, 0, -100),
			Rotation: geometry.NewVector(100, 0, 0)}
		light := render.Light{
			Position:  geometry.NewVector(100, 200, 0),
			Rotation:  geometry.ZeroVector(),
			LightType: render.Point,
			Color:     color.RGBA{0, 255, 255, 255},
			Length:    5000,
		}

		start := time.Now()

		// objects := make([]mesh.Mesh, 10)

		suzanne := mesh.ReadObjFile("obj/suzanne.obj", material.ColorMaterial(color.RGBA{0, 255, 255, 255}))
		ground := mesh.ReadObjFile("obj/ground.obj", material.ColorMaterial(color.RGBA{255, 255, 0, 255}))

		for {

			img = image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))

			// for x := 0; x < utilities.RESOLUTION_X; x++ {
			// 	for y := 0; y < utilities.RESOLUTION_Y; y++ {
			// 		img.Set(x, y, color.RGBA{107, 211, 232, 255})
			// 	}
			// }

			

			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				ground.Draw(img, zBuffer, cam, []render.Light{light})
				wg.Done()
			}()
			go func() {
				suzanne.Draw(img, zBuffer, cam, []render.Light{light})
				wg.Done()
			}()
			wg.Wait()

			for i := 0; i < len(zBuffer); i++ {
				zBuffer[i] = math.MaxFloat32
			}

			suzanne.Rotate(geometry.NewVector(0, 0.01, 0))

			viewport.Image = img
			viewport.Refresh()

			t := time.Since(start).Milliseconds()
			if t == 0 {
				t = 1
			}

			right.Text = fmt.Sprint("fps : ", 1000/t)
			right.Refresh()

			start = time.Now()
			// break
		}
	}()

	myWindow.Resize(fyne.NewSize(utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	// myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()
}

//TODO Aberty666
