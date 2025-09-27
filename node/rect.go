package node

// Rect is a rectangle struct
type Rect struct {
	X, Y, W, H float32
}

// Intersects checks if this rectangle intersects with another rectangle.
func (r Rect) Intersects(o Rect) bool {
	return r.X < o.X+o.W && r.X+r.W > o.X &&
		r.Y < o.Y+o.H && r.Y+r.H > o.Y
}
