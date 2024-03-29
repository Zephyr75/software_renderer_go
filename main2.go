package main

import (
	"image"
	"image/color"
	"runtime"

	"overdrive/src/ui"
	"overdrive/src/utils"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/anthonynsimon/bild/transform"
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

	for !window.ShouldClose() {
		var w, h = window.GetSize()

		img := image.NewRGBA(image.Rect(0, 0, w, h))

		//green := color.RGBA{201, 203, 163, 255}
		white := color.RGBA{0, 0, 0, 255}
		//orange := color.RGBA{226, 109, 92, 255}
		//red := color.RGBA{114, 61, 70, 255}
		brown := color.RGBA{71, 45, 48, 255}

		parent := ui.Row{
			Properties: ui.Properties{
				Alignment: ui.AlignmentCenter,
				Center: &ui.Point{
					X: w / 2,
					Y: h / 2,
				},
				MaxSize: &ui.Size{
					Scale:  ui.ScalePixel,
					Width:  w,
					Height: h,
				},
			},
			Style: ui.Style{
				Color: white,
			},
			Children: []ui.UIElement{
				ui.Button{
					Properties: ui.Properties{
						Alignment: ui.AlignmentCenter,
						Padding:   ui.PaddingEqual(ui.ScalePixel, 10),

					},
					Style: ui.Style{
						Color: brown,
					},
				},
				/*
					ui.Column{
						Properties: ui.Properties{
							Alignment: ui.AlignmentCenter,
						},
						Style: ui.Style{
							Color: white,
						},
						Children: []ui.UIElement{
							ui.Button{
								Properties: ui.Properties{
									Alignment: ui.AlignmentCenter,
									Size: ui.Size{
										Scale:  ui.ScaleRelative,
										Width:  50,
										Height: 50,
									},
									Function: func() {
										println("Button 1")
									},
								},
								Style: ui.Style{
									Color: red,
								},
							},
							ui.Button{
								Properties: ui.Properties{
									Alignment: ui.AlignmentCenter,
									Function: func() {
										println("Button 2")
									},
								},
								Style: ui.Style{
									Color: orange,
								},
							},
						},
					},
					ui.Button{
						Properties: ui.Properties{
							Alignment: ui.AlignmentCenter,
						},
						Style: ui.Style{
							Color: green,
						},
						Child: ui.Text{
							Properties: ui.Properties{
								Alignment: ui.AlignmentCenter,
								//Padding:   ui.PaddingEqual(ui.ScalePixel, 100),
								Size: ui.Size{
									Scale:  ui.ScalePixel,
									Width:  100,
									Height: 50,
								},
							},
							StyleText: ui.StyleText{
								Font: "JBMono.ttf",
								FontSize: 20,
								FontColor: color.RGBA{0, 0, 0, 255},
							},
						},
					},
				*/
			},
		}

		parent.Draw(img, window)

		/*
			exit := ui.Button{
				Properties: ui.Properties{
					Center: &ui.Point{
						X: w / 2,
						Y: h / 2,
					},
					Alignment: ui.AlignmentTopLeft,
					Size: ui.Size{
						Scale:  ui.ScalePixel,
						Width:  100,
						Height: 50,
					},
					Function: func() {
						window.SetShouldClose(true)
					},
				},
				Style: ui.Style{
					Color: color.RGBA{255, 255, 255, 255},
				},
			}

			exit.Draw(img, window)
		*/

		gl.BindTexture(gl.TEXTURE_2D, texture)

		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(transform.FlipV(img).Pix))

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
