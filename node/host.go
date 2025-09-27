package node

// Host actually gets the loop going
type Host interface {
	// TargetFps returns the target FPS
	TargetFps() uint
	// Width returns the width of the screen
	Width() uint
	// Height returns the height of the screen
	Height() uint
	// Caption returns the caption of the screen
	Caption() string
	// Node returns the Node that the Host uses
	Node() Node
	// SetNode changes the Node that the Host uses
	SetNode(n Node)
	// Run starts the loop
	Run()
}
