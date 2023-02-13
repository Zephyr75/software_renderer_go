package ui

import (
	"image"

	"github.com/go-gl/glfw/v3.3/glfw"
)


type Button struct {
	Properties *Properties
	Style	   Style
	Child      UIElement
}

func (button Button) Draw(img *image.RGBA, window *glfw.Window) {
	button.Properties.Draw(img, window)
	
	if button.Child != nil {
		button.Child.SetProperties(button.Properties.Size, button.Properties.Center)
		button.Child.Draw(img, window)
	}
}


func (button Button) SetProperties(size Size, center Point) {
	button.Properties.MaxSize = size
	button.Properties.Center = center
	//println("Button: ", center.X, " ", center.Y, " ", size.Width, " ", size.Height)
}

func (button Button) Debug() {
	println(button.Properties.Center.Y)
}