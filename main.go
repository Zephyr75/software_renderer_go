package main

import (
	"image"
	"image/color"

	"overdrive/src/geometry"
	"overdrive/src/material"
	"overdrive/src/mesh"
	"overdrive/src/render"
	"overdrive/src/utils"

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

func main2() {
	//if err != nil {
	//	panic(err)
	//}

	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	img := image.NewRGBA(image.Rect(0, 0, utils.RESOLUTION_X, utils.RESOLUTION_Y))

	//Depth buffer implemented on the z-axis
	zBuffer := make([]float32, utils.RESOLUTION_X*utils.RESOLUTION_Y)

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

		camera := render.NewCamera(geometry.NewVector(0, 0, 0), geometry.NewVector(0, 0, 0))

		pointLight := render.PointLight(geometry.NewVector(-50, 0, 0), geometry.ZeroVector(), color.RGBA{255, 255, 255, 255}, 5000)
		ambientLight := render.AmbientLight(color.RGBA{50, 50, 50, 255})
		lights := []*render.Light{&pointLight, &ambientLight}

		start := time.Now()

		suzanne := mesh.ReadObjFile("models/suzanne2.obj", material.ReadImageFile("images/suzanne2.png"))
		//suzanne := mesh.ReadObjFile("models/suzanne2.obj", material.ColorMaterial(color.RGBA{255, 255, 255, 255}))
		suzanne.Translate(geometry.NewVector(0, 0, 100))

		// ground := mesh.ReadObjFile("models/terrain.obj", material.ColorMaterial(color.RGBA{255, 255, 255, 255}))
		// ground.Translate(geometry.NewVector(0, 100, 0))

		for {

			img = image.NewRGBA(image.Rect(0, 0, utils.RESOLUTION_X, utils.RESOLUTION_Y))

			for x := 0; x < utils.RESOLUTION_X; x++ {
				for y := 0; y < utils.RESOLUTION_Y; y++ {
					img.Set(x, y, color.RGBA{107, 211, 232, 255})
				}
			}

			// Fill all light buffers
			var wgLight sync.WaitGroup
			wgLight.Add(len(lights))
			for _, light := range lights {
				go func(light *render.Light) {
					defer wgLight.Done()
					for i := 0; i < len(light.ZBuffer); i++ {
						light.ZBuffer[i] = -1
					}
					suzanne.LightPass(light)
					// ground.LightPass(light)
				}(light)
			}
			wgLight.Wait()

			suzanne.Draw(img, zBuffer, camera, lights)
			// ground.Draw(img, zBuffer, camera, lights)

			//Reset camera zBuffer
			for i := 0; i < len(zBuffer); i++ {
				zBuffer[i] = -1
			}

			//Double buffering
			viewport.Image = img
			viewport.Refresh()

			//Compute fps count and display it on screen
			t := time.Since(start).Milliseconds()
			if t == 0 {
				t = 1
			}
			right.Text = fmt.Sprint("fps : ", 1000/t)
			right.Refresh()
			start = time.Now()

			////////////////////////////////////////////////////////////////
			//state, _ := js.Read()
			// if err != nil {
			// 	panic(err)
			// }

			// a := (state.Buttons & 1) > 0
			// b := (state.Buttons & 2) > 0
			// x := (state.Buttons & 4) > 0
			// y := (state.Buttons & 8) > 0
			//lb := (state.Buttons & 16) > 0
			//rb := (state.Buttons & 32) > 0
			// fmt.Println("a:", a, "b:", b, "x:", x, "y:", y, "lb:", lb, "rb:", rb)

			lsHoriz := 0//float64(state.AxisData[0] / 32767)
			lsVert := 0//float64(state.AxisData[1] / 32767)
			// rsVert := float64(state.AxisData[3] / 32767)
			rsHoriz := 0//float64(state.AxisData[4] / 32767)
			// crossHoriz := float64(state.AxisData[5] / 32767)
			// crossVert := float64(state.AxisData[6] / 32767)
			// trigger := float64(state.AxisData[2] / 32641)
			// fmt.Println("lsHoriz:", lsHoriz, "lsVert:", lsVert, "rsHoriz:", rsHoriz, "rsVert:", rsVert, "crossHoriz:", crossHoriz, "crossVert:", crossVert, "trigger:", trigger)


			speed := 2 //TODO: define actual methods for camera
			camera.Position.AddAssign(geometry.NewVector(float64(speed)*float64(lsHoriz), 0, float64(speed)*float64(-lsVert)))
			
			camera.Direction.AddAssign(geometry.NewVector(0, 0.01*float64(rsHoriz), 0))
		}
	}()

	myWindow.Resize(fyne.NewSize(utils.RESOLUTION_X, utils.RESOLUTION_Y))
	myWindow.ShowAndRun()
}
