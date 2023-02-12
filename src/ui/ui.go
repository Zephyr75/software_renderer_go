package ui

type ScaleType byte

const (
	Pixel    ScaleType = 0
	Relative ScaleType = 1
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
	Center      Alignment = 0
	Top         Alignment = 1
	Bottom      Alignment = 2
	Left        Alignment = 3
	Right       Alignment = 4
	TopLeft     Alignment = 5
	TopRight    Alignment = 6
	BottomLeft  Alignment = 7
	BottomRight Alignment = 8
)

/*
Size
*/
type Size struct {
	Scale  ScaleType
	Width  int
	Height int
}

type UIElement struct {
	Alignment Alignment
	Padding   Padding
	Size      Size
	Child     *UIElement
}

func Button(alignment Alignment, padding Padding, size Size, child *UIElement) UIElement {
	return UIElement{
		Alignment: alignment,
		Padding:   padding,
		Size:      size,
		Child:     child,
	}
}

func (element UIElement) Draw(screen []byte) {
	// TODO
	println("Draw")
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
