package depixelize

// latticeGraph is a graph representing the pixels of an image
type latticeGraph struct{}

// pixel is a 1x1 grouping of pixels
type pixel struct{}

// pixel2 is a 2x2 grouping of pixels
type pixel2 struct{}

// pixel3 is a 3x3 grouping of pixels
type pixel3 struct{}

// pixel8 is a 8x8 grouping of pixels
type pixel8 struct{}

func (lg *latticeGraph) getWeight(p2 *pixel2) {}

// curvesHeuristic implements the curves heuristic and returns its vote weight.
// See https://johanneskopf.de/publications/pixelart/paper/pixel.pdf
func curvesHeuristic(p2 *pixel2, lg *latticeGraph) (weight int) { return 0 }

// sparsePixelsHeuristic implements the sparse pixels heuristic and returns its vote weight.
// See https://johanneskopf.de/publications/pixelart/paper/pixel.pdf
func sparsePixelsHeuristic(p2 *pixel2, window *pixel8) (weight int) { return 0 }

// islandsHeuristic implements the islands heuristic and returns its vote weight.
// See https://johanneskopf.de/publications/pixelart/paper/pixel.pdf
func islandsHeuristic(p2 *pixel2) (weight int) { return 0 }
