package depixelize

import c "image/color"

// TODO: better init cycle
func main() {
	// p is a shortly named function used to make development pixel art less verbose
	p := func(color c.Color) *node {
		return &node{
			pixel: &pixel{color},
		}
	}

	// circle is a 10x10 circle, black on white, used for development
	circle := graph{
		contents: [][]*node{
			{p(c.White), p(c.White), p(c.White), p(c.White), p(c.Black), p(c.Black), p(c.White), p(c.White), p(c.White), p(c.White)},
			{p(c.White), p(c.White), p(c.Black), p(c.Black), p(c.White), p(c.White), p(c.Black), p(c.Black), p(c.White), p(c.White)},
			{p(c.White), p(c.Black), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.Black), p(c.White)},
			{p(c.White), p(c.Black), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.Black), p(c.White)},
			{p(c.Black), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.Black)},
			{p(c.Black), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.Black)},
			{p(c.White), p(c.Black), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.Black), p(c.White)},
			{p(c.White), p(c.Black), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.White), p(c.Black), p(c.White)},
			{p(c.White), p(c.White), p(c.Black), p(c.Black), p(c.White), p(c.White), p(c.Black), p(c.Black), p(c.White), p(c.White)},
			{p(c.White), p(c.White), p(c.White), p(c.White), p(c.Black), p(c.Black), p(c.White), p(c.White), p(c.White), p(c.White)},
		},
		h: 10,
		w: 10,
	}

	// 1. each node has a reference to its parent graph
	// 2. each node is aware of its location in the graph
	// 3. each node needs to initialized with all edges being true
	circle.traverse(func(n *node, i, j int) {
		n.setParent(circle)
		n.setLocation(i, j)
		n.initEdges()
	})

	circle.disconnectDissimilar()
	circle.resolveNode2Cases()
}
