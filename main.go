package main

import (
	"fmt"
	c "image/color"
	"image/png"
	"os"

	dp "github.com/justinawrey/depixelizingpixelart/depixelize"
)

// TODO: better init cycle
func main() {
	// p is a shortly named function used to make development pixel art less verbose
	p := func(color c.Color) *dp.Node {
		return &dp.Node{
			Pixel: &dp.Pixel{Color: color},
		}
	}

	// circle is a 10x10 circle, black on white, used for development
	circle := dp.Graph{
		Contents: [][]*dp.Node{
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
		H:    10,
		W:    10,
		HRes: 300,
		WRes: 300,
	}

	// 1. each node has a reference to its parent graph
	// 2. each node is aware of its location in the graph
	// 3. each node needs to initialized with all edges being true
	circle.Traverse(func(n *dp.Node, i, j int) {
		n.SetParent(circle)
		n.SetLocation(i, j)
		n.InitEdges()
	})

	circle.DisconnectDissimilar()
	circle.ResolveNode2Cases()

	fi, err := os.Create("out.png")
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}

	err = png.Encode(fi, circle)
	if err != nil {
		fmt.Println("error encoding graph to png:", err)
		return
	}
}
