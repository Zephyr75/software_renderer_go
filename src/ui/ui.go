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
	SetProperties(size Size, center Point)
	Debug()
}

type Properties struct {
	MaxSize   Size
	Center    Point
	Size      Size
	Alignment Alignment
	Padding   Padding
	Color     color.Color
}

type Point struct {
	X int
	Y int
}

func (props *Properties) Draw(screen []byte) {
	if props.MaxSize.Width == 0 || props.MaxSize.Height == 0 {
		props.MaxSize.Width = utils.RESOLUTION_X
		props.MaxSize.Height = utils.RESOLUTION_Y
		props.MaxSize.Scale = ScalePixel
	}

	maxWidth := props.MaxSize.Width
	maxHeight := props.MaxSize.Height
	if props.MaxSize.Scale == ScaleRelative {
		maxWidth = utils.RESOLUTION_X * props.MaxSize.Width / 100
		maxHeight = utils.RESOLUTION_Y * props.MaxSize.Height / 100
	}

	if props.Size.Width == 0 || props.Size.Height == 0 {
		props.Size.Width = props.MaxSize.Width
		props.Size.Height = props.MaxSize.Height
		props.Size.Scale = ScalePixel
	}

	centerX := props.Center.X
	centerY := props.Center.Y


	println("-----------------")
	println(centerX, " ", centerY)

	width := props.Size.Width
	height := props.Size.Height
	if props.Size.Scale == ScaleRelative {
		width = maxWidth * props.Size.Width / 100
		height = maxHeight * props.Size.Height / 100
	}

	switch props.Alignment {
	case AlignmentBottom:
		centerY += height/2 - maxHeight/2
	case AlignmentTop:
		centerY -= height/2 - maxHeight/2
	case AlignmentLeft:
		centerX -= width/2 - maxWidth/2
	case AlignmentRight:
		centerX += width/2 - maxWidth/2
	}
	println(maxWidth, " ", maxHeight)
	println(centerX, " ", centerY)
	println(width, " ", height)

	centerX -= width / 2
	centerY -= height / 2


	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			r, g, b, _ := props.Color.RGBA()
			screen[(centerX+i)*4+(centerY+j)*utils.RESOLUTION_X*4] = byte(r)
			screen[(centerX+i)*4+(centerY+j)*utils.RESOLUTION_X*4+1] = byte(g)
			screen[(centerX+i)*4+(centerY+j)*utils.RESOLUTION_X*4+2] = byte(b)
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
