package ui


type Button struct {
	Properties *Properties
	Child      UIElement
}

func (button Button) Draw(screen []byte) {
	button.Properties.Draw(screen)
	
	if button.Child != nil {
		button.Child.SetProperties(button.Properties.Size, button.Properties.Center)
		button.Child.Draw(screen)
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