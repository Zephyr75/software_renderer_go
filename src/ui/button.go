package ui

import (
)

type Button struct {
	Properties Properties
	Child      UIElement
}

func (button Button) Draw(screen []byte) {
	button.Properties.Draw(screen)
	if button.Child != nil {
		button.Child.GetParentSize(button.Properties.Size)
		button.Child.Draw(screen)
	}
}

func (button Button) GetParentSize(size Size) {
	button.Properties.ParentSize = size
}
