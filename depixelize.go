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

func (g graph) hasNodeAt(i, j int) bool {
	return j < g.h && i < g.w
}

type node struct {
	parent graph
	pixel  *pixel
	edges  [8]bool
	i      int
	j      int
}

func (n *node) setEdge(dir int, to bool) {
	n.edges[dir] = to

	connection := opposites[dir]
	oppDir, i, j := connection.dir, connection.ix+n.i, connection.jx+n.j

	if n.parent.hasNodeAt(i, j) {
		n.parent.contents[j][i].edges[oppDir] = to
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
