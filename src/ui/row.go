package ui

import (
	"fmt"
	"image"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Row struct {
	Properties Properties
	Style	   Style
	Children   []UIElement
}

func (row Row) Draw(img *image.RGBA, window *glfw.Window) {
	
	Draw(img, window, &row.Properties, row.Style)

	for child := range row.Children {

		size := Size{
			Scale:  row.Properties.Size.Scale,
			Width:  row.Properties.Size.Width / len(row.Children),
			Height: row.Properties.Size.Height,
		}

		point := Point{
			X: row.Properties.Center.X - row.Properties.MaxSize.Width/2 + (2*child+1)*row.Properties.MaxSize.Width/(len(row.Children)*2),
			Y: row.Properties.Center.Y,
		}

		maxSizes = append(maxSizes, size)
		centers = append(centers, point)

		row.Children[child].SetProperties(
			&maxSizes[len(maxSizes)-1],
			&centers[len(centers)-1],
		)

		fmt.Println(size)
		fmt.Println(point)

		row.Children[child].Debug()
		row.Children[child].Draw(img, window)
	}
}


func (row Row) SetProperties(size *Size, center *Point) {
	row.Properties.MaxSize = size
	row.Properties.Center = center
}

func (row Row) Debug() {
	println(row.Properties.Center.Y)
}
