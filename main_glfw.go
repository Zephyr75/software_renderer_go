package main

import (
	"image"
	"image/color"
	"runtime"

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

    window, err := glfw.CreateWindow(640, 480, "My Window", nil, nil)
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

    for !window.ShouldClose() {

        var w, h = window.GetSize()

        var img = image.NewRGBA(image.Rect(0, 0, w, h))

        // -------------------------
        // MODIFY OR LOAD IMAGE HERE
        // -------------------------
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				img.Set(x, y, color.RGBA{uint8(x + i), uint8(y + i), 0, 255})
			}
		}
		i++

		img.Set(0, 0, color.RGBA{255, 0, 0, 255})

        gl.BindTexture(gl.TEXTURE_2D, texture)
        gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))

        gl.BlitFramebuffer(0, 0, int32(w), int32(h), 0, 0, int32(w), int32(h), gl.COLOR_BUFFER_BIT, gl.LINEAR)

        window.SwapBuffers()
        glfw.PollEvents()
    }
}