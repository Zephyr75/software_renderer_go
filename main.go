package main

import (
	"image"
	"image/color"

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
	"github.com/0xcafed00d/joystick"

	"fmt"
	// "sync"
	"time"
)

func main() {
	js, err := joystick.Open(0)
	if err != nil {
		panic(err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	img := image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))

	zBuffer := make([]float32, utilities.RESOLUTION_X*utilities.RESOLUTION_Y)

	for i := 0; i < len(zBuffer); i++ {
		zBuffer[i] = -1
	}

	viewport := canvas.NewImageFromImage(img)


	bottom := widget.NewButton("Assets browser", func() {
		fmt.Println("tapped")
	})

	right := canvas.NewText("fps", color.White)
	content := container.New(layout.NewBorderLayout(nil, bottom, nil, right), bottom, right, viewport)

	myCanvas.SetContent(content)

	go func() {

		cam := render.Camera{
			Position: geometry.NewVector(0, 0, -100),
			Rotation: geometry.NewVector(100, 0, 0)}
		pointLight := render.Light{
			Position:  geometry.NewVector(100, 200, 0),
			Rotation:  geometry.ZeroVector(),
			LightType: render.Point,
			Color:     color.RGBA{255, 255, 255, 255},
			Length:    50000,
		}
		ambientLight := render.Light{
			Position:  geometry.ZeroVector(),
			Rotation:  geometry.ZeroVector(),
			LightType: render.Ambient,
			Color:     color.RGBA{100, 100, 100, 255},
			Length:    50000,
		}

		start := time.Now()

		// objects := make([]mesh.Mesh, 10)

		suzanne := mesh.ReadObjFile("obj/suzanne.obj", material.ColorMaterial(color.RGBA{55, 122, 223, 255}))
		// ground := mesh.ReadObjFile("obj/terrain.obj", material.ColorMaterial(color.RGBA{102, 178, 97, 255}))

		for {

			img = image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))

			// for x := 0; x < utilities.RESOLUTION_X; x++ {
			// 	for y := 0; y < utilities.RESOLUTION_Y; y++ {
			// 		img.Set(x, y, color.RGBA{107, 211, 232, 255})
			// 	}
			// }

			state, err := js.Read()
			if err != nil {
				panic(err)
			}

			// fmt.Println("Axis Data: %v", state.AxisData)
			// fmt.Println("Button Data: %v", state.Buttons)

			// a := (state.Buttons & 1) > 0
			// b := (state.Buttons & 2) > 0
			// x := (state.Buttons & 4) > 0
			// y := (state.Buttons & 8) > 0
			lb := (state.Buttons & 16) > 0
			rb := (state.Buttons & 32) > 0
			// fmt.Println("a:", a, "b:", b, "x:", x, "y:", y, "lb:", lb, "rb:", rb)

			lsHoriz := float64(state.AxisData[0] / 32767)
			lsVert := float64(state.AxisData[1] / 32767)
			// rsVert := float64(state.AxisData[3] / 32767)
			rsHoriz := float64(state.AxisData[4] / 32767)
			// crossHoriz := float64(state.AxisData[5] / 32767)
			// crossVert := float64(state.AxisData[6] / 32767)
			// trigger := float64(state.AxisData[2] / 32641)
			// fmt.Println("lsHoriz:", lsHoriz, "lsVert:", lsVert, "rsHoriz:", rsHoriz, "rsVert:", rsVert, "crossHoriz:", crossHoriz, "crossVert:", crossVert, "trigger:", trigger)

			js.Close()

			
			// ground.Draw(img, zBuffer, cam, []render.Light{ambientLight})
			suzanne.Draw(img, zBuffer, cam, []render.Light{ambientLight, pointLight})

			// wg := sync.WaitGroup{}
			// wg.Add(2)
			// for i := 0; i < 1; i++ {
			// 	go func() {
			// 		ground.Draw(img, zBuffer, cam, []render.Light{pointLight, ambientLight})
			// 		wg.Done()
			// 	}()
			// 	go func() {
			// 		suzanne.Draw(img, zBuffer, cam, []render.Light{pointLight, ambientLight})
			// 		wg.Done()
			// 	}()
			// }
			// wg.Wait()

			for i := 0; i < len(zBuffer); i++ {
				zBuffer[i] = -1
			}

			speed := 2
			cam.Position.AddAssign(geometry.NewVector(float64(speed)*float64(lsHoriz), 0, float64(speed)*float64(-lsVert)))
			if lb {
				// suzanne.Rotate(geometry.NewVector(0, -0.1, 0))
				cam.Position.AddAssign(geometry.NewVector(0, float64(speed), 0))
			}
			if rb {
				cam.Position.AddAssign(geometry.NewVector(0, float64(-speed), 0))
				// suzanne.Rotate(geometry.NewVector(0, 0.1, 0))
			}
			cam.Rotation.AddAssign(geometry.NewVector(0, 0.01*float64(rsHoriz), 0))

			viewport.Image = img
			viewport.Refresh()

			t := time.Since(start).Milliseconds()
			if t == 0 {
				t = 1
			}
			fmt.Println("fps:", 1000/t)
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
