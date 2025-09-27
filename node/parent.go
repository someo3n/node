package node

// Parent is a Node which stores children nodes
type Parent interface {
	Node
	// Children returns the children of this Parent
	Children() []Node
	// HandleChildEvents is usually called by children of Parent and the Parent calls HandleChildEvents to its Parent, and so on
	HandleChildEvents(events []Event, origins ...Node)
}

// BaseParent is the base implementation of Parent
type BaseParent struct {
	*BaseNode
	// children is the children slice of the BaseParent's children
	children []Node
}

// NewBaseParent creates a new BaseParent with an empty children slice
func NewBaseParent() *BaseParent {
	return &BaseParent{
		BaseNode: NewBaseNode(),
		children: []Node{},
	}
}

// Update propagates the update to its children and also passes on the events
func (p *BaseParent) Update(events []Event) {
	p.doRunEvents(events)
	for _, child := range p.children {
		if child != nil && child.Active() {
			child.Update(events)
		}
	}
}

// Draw propagates the draw event to its children
func (p *BaseParent) Draw() {
	for _, child := range p.children {
		if child != nil && child.Visible() {
			child.Draw()
		}
	}
}

// Children returns the children of this Parent
func (p *BaseParent) Children() []Node {
	return p.children
}

// HandleChildEvents is usually called by children of Parent and the Parent calls HandleChildEvents to its Parent, and so on
func (p *BaseParent) HandleChildEvents(events []Event, origins ...Node) {
	if p.Parent() == nil {
		return
	}

	p.doRunEvents(events)
	parent := p.Parent()
	parent.HandleChildEvents(events, append(origins, p)...)
}

// doRunEvents does run events
func (p *BaseParent) doRunEvents(events []Event) {
	doneSomething := false
	for i := 0; i < len(events); i++ {
		if events[i] == nil {
			continue
		}
		if runEvent, ok := events[i].(*RunEvent); ok {
			propagate := runEvent.Payload(p)
			if !propagate {
				events[i] = nil
				doneSomething = true
			}
		}
	}

	if doneSomething {
		newEvents := []Event{}
		for _, e := range events {
			if e != nil {
				newEvents = append(newEvents, e)
			}
		}
		events = events[:0]
		events = append(events, newEvents...)
	}
}
