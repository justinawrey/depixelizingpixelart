package depixelize

import "image/color"

// edge directions
const (
	n = iota
	ne
	e
	se
	s
	sw
	w
	nw
)

type node struct {
	pixel *pixel
	edges [8]bool
}

func (n *node) initEdges() {
	for i := 0; i < 8; i++ {
		n.edges[i] = true
	}
}

// pixel is a 1x1 grouping of pixels
type pixel struct {
	color color.Color
}

// pixel2 is a 2x2 grouping of pixels
type pixel2 struct{}

// pixel3 is a 3x3 grouping of pixels
type pixel3 struct{}

// pixel8 is a 8x8 grouping of pixels
type pixel8 struct{}

// curvesHeuristic implements the curves heuristic and returns its vote weight.
// See https://johanneskopf.de/publications/pixelart/paper/pixel.pdf
func curvesHeuristic() (weight int) { return 0 }

// sparsePixelsHeuristic implements the sparse pixels heuristic and returns its vote weight.
// See https://johanneskopf.de/publications/pixelart/paper/pixel.pdf
func sparsePixelsHeuristic() (weight int) { return 0 }

// islandsHeuristic implements the islands heuristic and returns its vote weight.
// See https://johanneskopf.de/publications/pixelart/paper/pixel.pdf
func islandsHeuristic() (weight int) { return 0 }
