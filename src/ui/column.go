package ui

import (
)

type Column struct {
	Properties Properties
	Children   []UIElement
}

func (column Column) Draw(screen []byte) {
	column.Properties.Draw(screen)

	for child := range column.Children {
		column.Children[child].GetParentSize(column.Properties.Size)
		column.Children[child].Draw(screen)
	}

}

func (column Column) GetParentSize(size Size) {
	column.Properties.ParentSize = size
}
