package depixelize

import (
	"image/color"
	"math"
)

const (
	yThresh = 0.18823529411
	uThresh = 0.02745098039
	vThresh = 0.02352941176
)

func dissimilar(n1, n2 *node) bool {
	y1, u1, v1 := n1.pixel.yuv()
	y2, u2, v2 := n2.pixel.yuv()
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

type graph struct {
	contents [][]*node
	h        int
	w        int
}

func (g graph) traverse(onEach func(n *node, i, j int)) {
	for j, row := range g.contents {
		for i, node := range row {
			onEach(node, i, j)
		}
	}
}

func (g graph) traverse2(onEach func(n2 *node2)) {
	//TODO:
}

func (g graph) hasNodeAt(i, j int) bool {
	return j < g.h && i < g.w
}

func (g graph) disconnectDissimilar() {
	g.traverse(func(n *node, i, j int) {
		for i := 0; i < 8; i++ {
			if neighbour := n.getAdjacentNode(i); neighbour != nil {
				if dissimilar(n, neighbour) {
					n.setEdge(i, false)
				}
			}
		}
	})
}

type node struct {
	parent graph
	pixel  *pixel
	edges  [8]bool
	i      int
	j      int
}

type node2 struct {
	parent graph
	tl     *node
	tr     *node
	bl     *node
	br     *node
}

func (n *node) getAdjacentNode(dir int) *node {
	connection := opposites[dir]
	i, j := connection.ix+n.i, connection.jx+n.j

	if n.parent.hasNodeAt(i, j) {
		return n.parent.contents[j][i]
	}
	return nil
}

func (n *node) setEdge(dir int, to bool) {
	if neighbour := n.getAdjacentNode(dir); neighbour != nil {
		n.edges[dir] = to
		oppDir := opposites[dir].dir
		neighbour.edges[oppDir] = to
	}
}

func (n *node) setLocation(i, j int) {
	n.i = i
	n.j = j
}

func (n *node) initEdges() {
	for i := 0; i < 8; i++ {
		n.edges[i] = true
	}
}

func (n *node) setParent(g graph) {
	n.parent = g
}

// pixel is a 1x1 grouping of pixels
type pixel struct {
	color color.Color
}

// yuv returns the YUV colors of pixel p
func (p *pixel) yuv() (y, u, v uint8) {
	r, g, b, _ := p.color.RGBA()
	r8, g8, b8 := uint8(r), uint8(g), uint8(b)

	return color.RGBToYCbCr(r8, g8, b8)
}
