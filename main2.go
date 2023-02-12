package main

import (
	"image/color"
	"runtime"

	// "sync"

	"overdrive/src/ui"
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

	for !window.ShouldClose() {

		var w, h = window.GetSize()

		// define an array of uint8s
		var screen = make([]uint8, w*h*4)

		//Color:    color.RGBA{0, 56, 68, 255},

		parent := ui.Button{
			Properties: ui.Properties{
				Alignment: ui.AlignmentBottomLeft,
				Padding:   ui.PaddingEqual(ui.ScalePixel, 10),
				Color:     color.RGBA{255, 235, 198, 255},
			},
			Child: ui.Column{
				Properties: ui.Properties{
					Alignment: ui.AlignmentCenter,
					Color:     color.RGBA{0, 235, 198, 255}},
				Children: []ui.UIElement{
					ui.Button{
						Properties: ui.Properties{
							Alignment: ui.AlignmentLeft,
							Size:      ui.Size{Scale: ui.ScaleRelative, Width: 40, Height: 50},
							Color:     color.RGBA{0, 56, 68, 255},
						},
					},
					ui.Button{
						Properties: ui.Properties{
							Alignment: ui.AlignmentRight,
							Size:      ui.Size{Scale: ui.ScaleRelative, Width: 40, Height: 50},
							Color:     color.RGBA{0, 108, 103, 255},
						},
					},
				},
			},
		}

		parent.Draw(screen)

		gl.BindTexture(gl.TEXTURE_2D, texture)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(screen))

		gl.BlitFramebuffer(0, 0, int32(w), int32(h), 0, 0, int32(w), int32(h), gl.COLOR_BUFFER_BIT, gl.LINEAR)

		window.SwapBuffers()
		glfw.PollEvents()

		i++
		if glfw.GetTime()-time > 1 {
			println("FPS:", i)
			i = 0
			time = glfw.GetTime()
		}

	}
}
