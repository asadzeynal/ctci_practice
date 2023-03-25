package ch4

type GraphNode struct {
	Data     int
	Adjacent []*GraphNode
	Marked   bool
}

type Graph struct {
	Nodes []*GraphNode
}

func routeBetweenNodes(node1 *GraphNode, node2 *GraphNode) bool {
	queue := make([]*GraphNode, 1)
	queue[0] = node1
	node1.Marked = true
	for len(queue) != 0 {
		if queue[0] == node2 {
			return true
		}
		for _, a := range queue[0].Adjacent {
			if a.Marked {
				continue
			}
			a.Marked = true
			queue = append(queue, a)
		}
		queue = queue[1:]
	}

	return false
}
