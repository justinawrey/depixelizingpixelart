package depixelize

import (
	"image"
	"image/color"
	"math"
)

const (
	yThresh = 0.18823529411
	uThresh = 0.02745098039
	vThresh = 0.02352941176
)

func dissimilar(n1, n2 *Node) bool {
	y1, u1, v1 := n1.Pixel.yuv()
	y2, u2, v2 := n2.Pixel.yuv()
	y3, u3, v3 := float64(y1), float64(u1), float64(v1)
	y4, u4, v4 := float64(y2), float64(u2), float64(v2)

	return math.Abs(y3-y4) > yThresh ||
		math.Abs(u3-u4) > uThresh ||
		math.Abs(v3-v4) > vThresh
}

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

type connectionInfo struct {
	dir int
	ix  int
	jx  int
}

// opposites to edge directions and displacements
var opposites = map[int]connectionInfo{
	n:  connectionInfo{s, 0, -1},
	ne: connectionInfo{sw, 1, 1},
	e:  connectionInfo{w, 1, 0},
	se: connectionInfo{nw, 1, 1},
	s:  connectionInfo{n, 0, 1},
	sw: connectionInfo{ne, -1, 1},
	w:  connectionInfo{e, -1, 0},
	nw: connectionInfo{se, -1, -1},
}

type Graph struct {
	Contents [][]*Node
	H        int
	W        int
	HRes     int
	WRes     int
}

func (g Graph) Traverse(onEach func(n *Node, i, j int)) {
	for j, row := range g.Contents {
		for i, node := range row {
			onEach(node, i, j)
		}
	}
}

func (g Graph) traverse2(onEach func(n2 *node2)) {
	for j := 0; j < g.H-1; j++ {
		for i := 0; i < g.W-1; i++ {
			tl := g.Contents[j][i]
			tr := tl.getAdjacentNode(e)
			bl := tl.getAdjacentNode(s)
			br := tl.getAdjacentNode(se)
			n2 := &node2{g, tl, tr, bl, br}

			onEach(n2)
		}
	}
}

func (g Graph) hasNodeAt(i, j int) bool {
	return j < g.H && i < g.W
}

func (g Graph) DisconnectDissimilar() {
	g.Traverse(func(n *Node, i, j int) {
		for i := 0; i < 8; i++ {
			if neighbour := n.getAdjacentNode(i); neighbour != nil {
				if dissimilar(n, neighbour) {
					n.setEdge(i, false)
				}
			}
		}
	})
}

func (g Graph) ResolveNode2Cases() {
	g.traverse2(func(n2 *node2) {
		n2.resolve()
	})
}

func (g Graph) ColorModel() color.Model {
	return color.RGBAModel
}

func (g Graph) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: g.WRes,
			Y: g.HRes,
		},
	}
}

func (g Graph) At(x, y int) color.Color {
	pxXSize := g.WRes / g.W
	pxYSize := g.HRes / g.H
	atX := x / pxXSize
	atY := y / pxYSize

	return g.Contents[atY][atX].Pixel.Color
}

type Node struct {
	parent Graph
	Pixel  *Pixel
	edges  [8]bool
	i      int
	j      int
}

func (n *Node) getAdjacentNode(dir int) *Node {
	connection := opposites[dir]
	i, j := connection.ix+n.i, connection.jx+n.j

	if n.parent.hasNodeAt(i, j) {
		return n.parent.Contents[j][i]
	}
	return nil
}

func (n *Node) hasEdge(dir int) bool {
	neighbour := n.getAdjacentNode(dir)
	return neighbour != nil && n.edges[dir]
}

func (n *Node) setEdge(dir int, to bool) {
	if neighbour := n.getAdjacentNode(dir); neighbour != nil {
		n.edges[dir] = to
		oppDir := opposites[dir].dir
		neighbour.edges[oppDir] = to
	}
}

func (n *Node) SetLocation(i, j int) {
	n.i = i
	n.j = j
}

func (n *Node) InitEdges() {
	for i := 0; i < 8; i++ {
		n.edges[i] = true
	}
}

func (n *Node) SetParent(g Graph) {
	n.parent = g
}

func (n *Node) valence() int {
	var count int
	for i := 0; i < 8; i++ {
		if n.hasEdge(i) {
			count++
		}
	}
	return count
}

type node2 struct {
	parent Graph
	tl     *Node
	tr     *Node
	bl     *Node
	br     *Node
}

func (n2 *node2) unfold() (tl, tr, bl, br *Node) {
	return n2.tl, n2.tr, n2.bl, n2.br
}

func (n2 *node2) isFullyConnected() bool {
	tl, tr, bl, br := n2.unfold()
	return tl.hasEdge(e) && tr.hasEdge(s) && br.hasEdge(w) && bl.hasEdge(n) && tl.hasEdge(se) && bl.hasEdge(ne)
}

func (n2 *node2) isProblematic() bool {
	tl, tr, bl, br := n2.unfold()
	return !tl.hasEdge(e) && !tr.hasEdge(s) && !br.hasEdge(w) && !bl.hasEdge(n) && tl.hasEdge(se) && bl.hasEdge(ne)
}

func (n2 *node2) curvesHeuristic(first, second *Node, dir int) (weight int) {
	// TODO:
	return 1
}

func (n2 *node2) sparsePixelsHeuristic(first, second *Node, dir int) (weight int) {
	// TODO:
	return 1
}

func (n2 *node2) islandsHeuristic(first, second *Node, dir int) int {
	if first.valence() == 1 || second.valence() == 1 {
		return 5
	}
	return 0
}

func (n2 *node2) getWeight(dir int) int {
	tl, tr, bl, br := n2.unfold()
	var first *Node
	var second *Node

	if dir == se {
		first = tl
		second = br
	} else if dir == ne {
		first = bl
		second = tr
	}

	return n2.curvesHeuristic(first, second, dir) + n2.sparsePixelsHeuristic(first, second, dir) + n2.islandsHeuristic(first, second, dir)
}

func (n2 *node2) removeSEDiagonal() {
	n2.tl.setEdge(se, false)
}

func (n2 *node2) removeNEDiagonal() {
	n2.bl.setEdge(ne, false)
}

func (n2 *node2) removeDiagonals() {
	n2.removeSEDiagonal()
	n2.removeNEDiagonal()
}

func (n2 *node2) resolve() {
	if n2.isFullyConnected() {
		n2.removeDiagonals()
	} else if n2.isProblematic() {
		seWeight := n2.getWeight(se)
		neWeight := n2.getWeight(ne)

		if seWeight == neWeight {
			n2.removeDiagonals()
		} else if seWeight > neWeight {
			n2.removeNEDiagonal()
		} else {
			n2.removeSEDiagonal()
		}
	}
}

// pixel is a 1x1 grouping of pixels
type Pixel struct {
	Color color.Color
}

// yuv returns the YUV colors of pixel p
func (p *Pixel) yuv() (y, u, v uint8) {
	r, g, b, _ := p.Color.RGBA()
	r8, g8, b8 := uint8(r), uint8(g), uint8(b)

	return color.RGBToYCbCr(r8, g8, b8)
}
