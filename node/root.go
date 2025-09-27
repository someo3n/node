package node

// Root is kind of BaseParent that just draws the last child
type Root struct {
	*BaseParent
}

// NewRoot creates a new Root
func NewRoot() *Root {
	return &Root{
		BaseParent: NewBaseParent(),
	}
}

// Draw draws only the last child
func (r *Root) Draw() {
	for _, child := range r.Children() {
		if child != nil && child.Visible() {
			child.Draw()
		}
	}
}
