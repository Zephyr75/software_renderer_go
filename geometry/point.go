package geometry

/**
 * 2D point to be drawn on screen
 */
type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}