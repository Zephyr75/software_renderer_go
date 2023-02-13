package ui

import (
	"image"
	"image/color"

	"github.com/go-gl/glfw/v3.3/glfw"

	"io/ioutil"
	"log"

	"github.com/goki/freetype"
)


type Text struct {
	Properties *Properties
}

func (text Text) Draw(img *image.RGBA, window *glfw.Window) {
	text.Properties.Draw(img, window)

	drawText(img, []string{"Hello, World!"}, "JBMono.ttf", 30, text.Properties.Center.X, 0)

}


func (text Text) SetProperties(size Size, center Point) {
	text.Properties.MaxSize = size
	text.Properties.Center = center
}

func (text Text) Debug() {
	println(text.Properties.Center.Y)
}

func drawText(img *image.RGBA, text []string, font string, fontSize float64, x, y int) {
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
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(fontColor))

	// Draw the text
	pt := freetype.Pt(x, y+int(c.PointToFixed(fontSize)>>6))
	for _, s := range text {
		_, err := c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(fontSize * 1.5)
	}
}