package node

// Node is the core interface representing a renderable and updatable entity
type Node interface {
	// Update processes logic and incoming events
	Update(events []Event)
	// Draw renders the node
	Draw()
	// Active indicates whether the Node is active and should update
	Active() bool
	// Visible indicates whether the Node is visible and should be drawn
	Visible() bool
	// Tags returns the slice of string tags attached to the Node
	Tags() []string
	// SetTags replaces the Node's tags with the provided slice
	SetTags([]string)
	// Parent returns the parent of the Node
	Parent() Parent
	// SetParent sets the parent of the Node and manages membership in children lists
	SetParent(Parent) bool
	// AddedToParent is called when the Node is added to a parent (this gets automatically called)
	AddedToParent(Parent)
	// RemovedFromParent is called when the Node is removed from a parent (this gets automatically called)
	RemovedFromParent(Parent)
}

// BaseNode provides a default implementation of Node
type BaseNode struct {
	// parent holds the parent of this Node
	parent Parent
	// active determines if this Node is updated
	active bool
	// active determines if this Node is drawn
	visible bool
	// tags is a slice of string tags for this Node
	tags []string
}

// NewBaseNode constructs and returns a new BaseNode with defaults (active and visible)
func NewBaseNode() *BaseNode {
	return &BaseNode{
		active:  true,
		visible: true,
		tags:    []string{},
	}
}

// Update processes events, executing RunEvent payloads if present
func (n *BaseNode) Update(events []Event) {
	for _, e := range events {
		if runEvent, ok := e.(*RunEvent); ok {
			runEvent.Payload(n)
		}
	}
}

// Draw does nothing by default but can be overridden
func (n *BaseNode) Draw() {}

// Active indicates whether the Node is active and should update
func (n *BaseNode) Active() bool {
	return n.active
}

// Visible indicates whether the Node is visible and should be drawn
func (n *BaseNode) Visible() bool {
	return n.visible
}

// Tags returns the slice of string tags attached to the Node
func (n *BaseNode) Tags() []string {
	return n.tags
}

// SetTags replaces the Node's tags with the provided slice
func (n *BaseNode) SetTags(t []string) {
	n.tags = t
}

// Parent returns the parent of the Node
func (n *BaseNode) Parent() Parent {
	return n.parent
}

// SetParent sets the parent of the Node and manages membership in children lists
func (n *BaseNode) SetParent(p Parent) bool {
	if n.parent == p {
		return false
	}

	if n.parent != nil {
		if bp, ok := n.parent.(*BaseParent); ok {
			newChildren := []Node{}
			for _, c := range bp.children {
				if c != n {
					newChildren = append(newChildren, c)
				}
			}
			bp.children = newChildren
		}
		n.RemovedFromParent(n.parent)
	}

	n.parent = p

	if n.parent != nil {
		if bp, ok := n.parent.(*BaseParent); ok {
			found := false
			for _, c := range bp.children {
				if c == n {
					found = true
					break
				}
			}
			if !found {
				bp.children = append(bp.children, n)
			}
		}
		n.AddedToParent(n.parent)
	}

	return true
}

// AddedToParent is called when the Node is added to a parent (this gets automatically called)
func (n *BaseNode) AddedToParent(p Parent) {}

// RemovedFromParent is called when the Node is removed from a parent (this gets automatically called)
func (n *BaseNode) RemovedFromParent(p Parent) {}
