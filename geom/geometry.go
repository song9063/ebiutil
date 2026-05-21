package geom

type Point struct {
	X, Y float64
}

func (p Point) ToPoint32() Point32 {
	return Point32{X: float32(p.X), Y: float32(p.Y)}
}

func NewPointFromInt(x, y int) Point {
	return Point{X: float64(x), Y: float64(y)}
}

type Point32 struct {
	X, Y float32
}

func (p Point32) ToPoint64() Point {
	return Point{X: float64(p.X), Y: float64(p.Y)}
}

type Rect struct {
	Point
	W, H float64
}

func NewRectFromInt(x, y, w, h int) Rect {
	return Rect{Point: NewPointFromInt(x, y), W: float64(w), H: float64(h)}
}

func (r Rect) Right() float64 {
	return r.X + r.W
}
func (r Rect) Bottom() float64 {
	return r.Y + r.H
}
func (r Rect) CenterX() float64 {
	return r.X + r.W/2
}
func (r Rect) CenterY() float64 {
	return r.Y + r.H/2
}
func (r Rect) Center() Point {
	return Point{
		X: r.CenterX(),
		Y: r.CenterY(),
	}
}
func (r Rect) HitTest(p Point) bool {
	return (p.X >= r.X && p.X <= r.Right() &&
		p.Y >= r.Y && p.Y <= r.Bottom())
}
