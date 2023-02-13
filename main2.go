package main

import (
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"runtime"

	// "sync"

	"overdrive/src/ui"
	"overdrive/src/utils"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/goki/freetype"
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

	img := image.NewRGBA(image.Rect(0, 0, utils.RESOLUTION_X, utils.RESOLUTION_Y))

	/////////////////////////

	font := "JBMono.ttf"
	fontSize := 30
	fontColor := color.RGBA{0, 200, 200, 255}
	dpi := 72

	// Load font
	fontBytes, err := ioutil.ReadFile(font)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)

	// Load freetype context
	c := freetype.NewContext()
	c.SetDPI(float64(dpi))
	c.SetFont(f)
	c.SetFontSize(float64(fontSize))
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(fontColor))

	text := []string{"Hello duuuuuuuuuuuuuuuuuuuuuude", "World", "!"}

	/////////////////////////

	for !window.ShouldClose() {
		var w, h = window.GetSize()

		green := color.RGBA{201, 203, 163, 255}
		yellow := color.RGBA{255, 225, 168, 255}
		orange := color.RGBA{226, 109, 92, 255}
		red := color.RGBA{114, 61, 70, 255}
		brown := color.RGBA{71, 45, 48, 255}

		parent := ui.Row{
			Properties: &ui.Properties{
				Alignment: ui.AlignmentCenter,
				Color:     yellow,
				Center: ui.Point{
					X: utils.RESOLUTION_X / 2,
					Y: utils.RESOLUTION_Y / 2,
				},
			},
			Children: []ui.UIElement{
				ui.Button{
					Properties: &ui.Properties{
						Alignment: ui.AlignmentCenter,
						Padding:   ui.PaddingEqual(ui.ScalePixel, 10),
						Color:     brown,
					},
				},
				ui.Column{
					Properties: &ui.Properties{
						Color:     yellow,
						Alignment: ui.AlignmentCenter,
					},
					Children: []ui.UIElement{
						ui.Button{
							Properties: &ui.Properties{
								Alignment: ui.AlignmentCenter,
								Color:     red,
								Size: ui.Size{
									Scale:  ui.ScaleRelative,
									Width:  50,
									Height: 50,
								},
								Function: func() {
									println("Button 1")
								},
							},
						},
						ui.Button{
							Properties: &ui.Properties{
								Alignment: ui.AlignmentCenter,
								Color:     orange,
								Function: func() {
									println("Button 2")
								},
							},
						},
					},
				},
				ui.Button{
					Properties: &ui.Properties{
						Alignment: ui.AlignmentCenter,
						Color:     green,
					},
				},
			},
		}

		parent.Draw(img, window)

		exit := ui.Button{
			Properties: &ui.Properties{
				Center: ui.Point{
					X: utils.RESOLUTION_X / 2,
					Y: utils.RESOLUTION_Y / 2,
				},
				Alignment: ui.AlignmentTopLeft,
				Color:     color.RGBA{255, 255, 255, 255},
				Size: ui.Size{
					Scale:  ui.ScalePixel,
					Width:  100,
					Height: 50,
				},
				Function: func() {
					window.SetShouldClose(true)
				},
			},
		}

		exit.Draw(img, window)

		drawText(c, text)

		gl.BindTexture(gl.TEXTURE_2D, texture)

		//get byte array from the image

		new_img := image.NewRGBA(image.Rect(0, 0, utils.RESOLUTION_X, utils.RESOLUTION_Y))

		// flip the image
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				new_img.Set(x, y, img.At(x, h-y-1))
			}
		}

		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(new_img.Pix))

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

func drawText(c *freetype.Context, text []string) {
	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(30)>>6))
	for _, s := range text {
		_, err := c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(30 * 1.5)
	}
}