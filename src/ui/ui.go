package ui

import (
	"image/color"
	"overdrive/src/utils"
)

type ScaleType byte

const (
	ScalePixel    ScaleType = 0
	ScaleRelative ScaleType = 1
)

/*
Padding
*/
type Padding struct {
	Scale  ScaleType
	Top    int
	Right  int
	Bottom int
	Left   int
}

func PaddingEqual(scale ScaleType, padding int) Padding {
	return Padding{
		Scale:  scale,
		Top:    padding,
		Right:  padding,
		Bottom: padding,
		Left:   padding,
	}
}
func PaddingSymmetric(scale ScaleType, vertical, horizontal int) Padding {
	return Padding{
		Scale:  scale,
		Top:    vertical,
		Right:  horizontal,
		Bottom: vertical,
		Left:   horizontal,
	}
}
func PaddingSideBySide(scale ScaleType, top, right, bottom, left int) Padding {
	return Padding{
		Scale:  scale,
		Top:    top,
		Right:  right,
		Bottom: bottom,
		Left:   left,
	}
}

/*
Alignment
*/
type Alignment byte

const (
	AlignmentCenter      Alignment = 0
	AlignmentTop         Alignment = 1
	AlignmentBottom      Alignment = 2
	AlignmentLeft        Alignment = 3
	AlignmentRight       Alignment = 4
	AlignmentTopLeft     Alignment = 5
	AlignmentTopRight    Alignment = 6
	AlignmentBottomLeft  Alignment = 7
	AlignmentBottomRight Alignment = 8
)

/*
Size
*/
type Size struct {
	Scale  ScaleType
	Width  int
	Height int
}

type UIElement interface {
	Draw(screen []byte)
	GetParentSize(size Size)
}


type Properties struct {
	Alignment  Alignment
	Padding    Padding
	Size       Size
	ParentSize Size
	Color      color.Color
}

func (props Properties) Draw(screen []byte) {
	maxWidth := props.ParentSize.Width
	maxHeight := props.ParentSize.Height
	if props.ParentSize.Width == 0 {
		maxWidth = utils.RESOLUTION_X
	}
	if props.ParentSize.Width == 0 {
		maxHeight = utils.RESOLUTION_Y
	}

	width := props.Size.Width
	height := props.Size.Height
	if props.Size.Width == 0 {
		width = maxWidth
	}
	if props.Size.Height == 0 {
		height = maxHeight
	}
	if props.Size.Scale == ScaleRelative {
		width = maxWidth * props.Size.Width / 100
		height = maxHeight * props.Size.Height / 100
	}

	centerX := 0
	centerY := 0

	switch props.Alignment {
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
			r, g, b, _ := props.Color.RGBA()
			screen[(centerX+i)*4+(centerY+j)*maxWidth*4] = byte(r)
			screen[(centerX+i)*4+(centerY+j)*maxWidth*4+1] = byte(g)
			screen[(centerX+i)*4+(centerY+j)*maxWidth*4+2] = byte(b)
		}
	}
}

/*
Button
Text
Row
Column

Align
--------
Center
Left
Right
Top
Bottom
Top left
Top right
Bottom left
Bottom right



Padding
--------
Pixel : All around, Symmetric, Side by side
Relative : All around, Symmetric, Side by side



Style
--------
Background color
Border color
Border width
Border radius
Shadow
Text color
Text size
Text font



Parent





Color

Border radius
*/
