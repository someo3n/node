package node

// Event represents a generic event in the system
type Event interface{}

// RunEvent is a special event carrying a payload function
type RunEvent struct {
	// Payload is the function that is executed with a Node and returns whether the event should propagate any further
	Payload func(Node) bool
}
