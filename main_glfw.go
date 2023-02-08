package main

import (
	"image"
	"image/color"
	"runtime"
	"sync"

	// "sync"

	"overdrive/src/geometry"
	"overdrive/src/material"
	"overdrive/src/mesh"
	"overdrive/src/render"
	"overdrive/src/utils"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// GLFW: This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(utils.RESOLUTION_X, utils.RESOLUTION_Y, "My Window", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	var texture uint32
	{
		gl.GenTextures(1, &texture)

		gl.BindTexture(gl.TEXTURE_2D, texture)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

		gl.BindImageTexture(0, texture, 0, false, 0, gl.WRITE_ONLY, gl.RGBA8)
	}

	var framebuffer uint32
	{
		gl.GenFramebuffers(1, &framebuffer)
		gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texture, 0)

		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, framebuffer)
		gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
	}

	i := 0
	time := glfw.GetTime()

	/////////////////////////

	img := image.NewRGBA(image.Rect(0, 0, utils.RESOLUTION_X, utils.RESOLUTION_Y))

	//Depth buffer implemented on the z-axis
	zBuffer := make([]float32, utils.RESOLUTION_X*utils.RESOLUTION_Y)

	for i := 0; i < len(zBuffer); i++ {
		zBuffer[i] = -1
	}

	camera := render.NewCamera(geometry.NewVector(0, 0, 0), geometry.NewVector(0, 0, 0))

	pointLight := render.PointLight(geometry.NewVector(-50, 0, 0), geometry.ZeroVector(), color.RGBA{255, 255, 255, 255}, 5000)
	ambientLight := render.AmbientLight(color.RGBA{50, 50, 50, 255})
	lights := []*render.Light{&pointLight, &ambientLight}


	//suzanne := mesh.ReadObjFile("models/suzanne2.obj", material.ReadImageFile("images/suzanne2.png"))
	suzanne := mesh.ReadObjFile("models/suzanne2.obj", material.ColorMaterial(color.RGBA{255, 255, 255, 255}))
	suzanne.Translate(geometry.NewVector(0, 0, 100))

	/////////////////////////

	for !window.ShouldClose() {

		var w, h = window.GetSize()

		// -------------------------
		// MODIFY OR LOAD IMAGE HERE
		// -------------------------

		// define an array of uint8s
		var pixels = make([]uint8, w*h*4)

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				pixels[(x+y*w)*4+0] = uint8(x + i)
				// img.Set(x, y, color.RGBA{uint8(x + i), uint8(y + i), 0, 255})
			}
		}

		/////////////////////////

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

		/////////////////////////

		i++

		img.Set(0, 0, color.RGBA{255, 0, 0, 255})

		gl.BindTexture(gl.TEXTURE_2D, texture)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))

		gl.BlitFramebuffer(0, 0, int32(w), int32(h), 0, 0, int32(w), int32(h), gl.COLOR_BUFFER_BIT, gl.LINEAR)

		window.SwapBuffers()
		glfw.PollEvents()

		if glfw.GetTime()-time > 1 {
			println("FPS:", i)
			i = 0
			time = glfw.GetTime()
		}

	}
}
