package node

import "slices"

// CollidingNodes returns nodes that collide with the target, optionally filtered by tags and checked recursively
func CollidingNodes(nodes []Node, recursive bool, target Collider, tags ...string) []Node {
	var result []Node

	for _, child := range nodes {
		if c, ok := child.(Collider); ok {
			if target.CollidesWith(c) {
				if len(tags) == 0 || hasTag(child, tags...) {
					result = append(result, child)
				}
			}
		}

		if recursive {
			if par, ok := child.(Parent); ok {
				result = append(result, CollidingNodes(par.Children(), true, target, tags...)...)
			}
		}
	}

	return result
}

// hasTag returns true if the node has at least one of the specified tags
func hasTag(n Node, tags ...string) bool {
	nodeTags := n.Tags()
	for _, t := range tags {
		if slices.Contains(nodeTags, t) {
			return true
		}
	}
	return false
}
