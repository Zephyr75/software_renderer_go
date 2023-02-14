package ui

import (
	"fmt"
	"image"

	"github.com/go-gl/glfw/v3.3/glfw"
)


type Button struct {
	Properties Properties
	Style	   Style
	Child      UIElement
}

func (button Button) Draw(img *image.RGBA, window *glfw.Window) {

	//fmt.Println(button.Properties.MaxSize)
	//fmt.Println(button.Properties.Center)
	Draw(img, window, &button.Properties, button.Style)
	
	if button.Child != nil {
		button.Child.SetProperties(&button.Properties.Size, button.Properties.Center)
		button.Child.Draw(img, window)
	}
}


func (button Button) SetProperties(size *Size, center *Point) {
	button.Properties.MaxSize = size
	button.Properties.Center = center

	//fmt.Println(button.Properties.MaxSize)
	//fmt.Println(button.Properties.Center)

}

func (button Button) Debug() {
	fmt.Println("Debug")
	fmt.Println(button.Properties.MaxSize)
	fmt.Println(button.Properties.Center)
}