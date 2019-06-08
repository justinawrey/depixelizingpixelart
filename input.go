package depixelize

import (
	c "image/color"
)

// p is a shortly named function used to make development pixel art less verbose
func p(color c.Color) *pixel {
	return &pixel{color}
}

// circle is a 10x10 circle, black on white, used for development
var circle = [][]*pixel{
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
}
