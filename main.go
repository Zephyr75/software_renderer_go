package main

import (
	"image"
	"image/color"

	"overdrive/src/geometry"
	"overdrive/src/material"
	"overdrive/src/mesh"
	"overdrive/src/render"
	"overdrive/src/utilities"

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

	//Depth buffer implemented on the z-axis
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
			Position:  geometry.NewVector(200, 200, -300),
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

		// suzanne := mesh.ReadObjFile("models/cubeRetro.obj", material.GetImageFromFilePath("images/retro9.png"))
		suzanne := mesh.ReadObjFile("models/suzanne.obj", material.ColorMaterial(color.RGBA{255, 0, 0, 255}))
		// suzanne := mesh.ReadObjFile("models/suzanne.obj", material.GetImageFromFilePath("images/suzanne.png"))



		// ground := mesh.ReadObjFile("models/terrain.obj", material.ColorMaterial(color.RGBA{102, 178, 97, 255}))

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

			for i := 0; i < len(zBuffer); i++ {
				zBuffer[i] = -1
			}

			speed := 2
			cam.Position.AddAssign(geometry.NewVector(float64(speed)*float64(lsHoriz), 0, float64(speed)*float64(-lsVert)))
			if lb {
				cam.Position.AddAssign(geometry.NewVector(0, float64(speed), 0))
			}
			if rb {
				cam.Position.AddAssign(geometry.NewVector(0, float64(-speed), 0))
			}
			cam.Rotation.AddAssign(geometry.NewVector(0, 0.01*float64(rsHoriz), 0))

			//Double buffering
			viewport.Image = img
			viewport.Refresh()

			t := time.Since(start).Milliseconds()
			if t == 0 {
				t = 1
			}

			right.Text = fmt.Sprint("fps : ", 1000/t)
			right.Refresh()

			start = time.Now()
		}
	}()

	myWindow.Resize(fyne.NewSize(utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	myWindow.ShowAndRun()
}
