package node

// Collider is an interface for colliders
type Collider interface {
	// Bounds returns the bounds of this Collider
	Bounds() Rect
	// SetBounds sets the bounds of this Collider
	SetBounds(r Rect)
	// CollidesWith checks if another Collider collides with this Collider
	CollidesWith(c Collider) bool
}

// BaseCollider is the base implementation of Collider
type BaseCollider struct {
	rect Rect
}

func NewBaseCollider(rect Rect) *BaseCollider {
	return &BaseCollider{
		rect: rect,
	}
}

// Bounds returns the bounds of this Collider
func (b *BaseCollider) Bounds() Rect {
	return b.rect
}

// SetBounds sets the bounds of this Collider
func (b *BaseCollider) SetBounds(r Rect) {
	b.rect = r
}

// CollidesWith checks if another Collider collides with this Collider
func (b *BaseCollider) CollidesWith(c Collider) bool {
	return b.Bounds().Intersects(c.Bounds())
}
