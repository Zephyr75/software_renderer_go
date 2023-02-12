package ui

import (
	"image/color"
	"overdrive/src/utils"
)

type Button struct {
	Alignment  Alignment
	Padding    Padding
	Size       Size
	ParentSize Size
	Color      color.Color
	Child      *UIElement
}


func (button Button) Draw(screen []byte) {
	maxWidth := button.ParentSize.Width
	maxHeight := button.ParentSize.Height
	if button.ParentSize.Width == 0 { maxWidth = utils.RESOLUTION_X }
	if button.ParentSize.Width == 0 { maxHeight = utils.RESOLUTION_Y }

	width := button.Size.Width
	height := button.Size.Height
	if button.Size.Width == 0 { width = maxWidth }
	if button.Size.Height == 0 { height = maxHeight }

	centerX := 0
	centerY := 0

	switch button.Alignment {
	case AlignmentCenter:
		centerX = maxWidth / 2
		centerY = maxHeight / 2
	case AlignmentBottom:
		centerX = maxWidth / 2
		centerY = height / 2
	case AlignmentTop:
		centerX = maxWidth / 2
		centerY = maxHeight - height/2
	case AlignmentLeft:
		centerX = width / 2
		centerY = maxHeight / 2
	case AlignmentRight:
		centerX = maxWidth - width/2
		centerY = maxHeight / 2
	case AlignmentTopLeft:
		centerX = width / 2
		centerY = maxHeight - height/2
	case AlignmentTopRight:
		centerX = maxWidth - width/2
		centerY = maxHeight - height/2
	case AlignmentBottomLeft:
		centerX = width / 2
		centerY = height / 2
	case AlignmentBottomRight:
		centerX = maxWidth - width/2
		centerY = height / 2
	default:
		panic("Invalid alignment")
	}

	centerX -= width / 2
	centerY -= height / 2

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			r, g, b, _ := button.Color.RGBA()
			screen[(centerX+i)*4+(centerY+j)*maxWidth*4] = byte(r)
			screen[(centerX+i)*4+(centerY+j)*maxWidth*4+1] = byte(g)
			screen[(centerX+i)*4+(centerY+j)*maxWidth*4+2] = byte(b)
		}
	}

}

func (button Button) GetParentSize(size Size) {
	button.ParentSize = size
}
