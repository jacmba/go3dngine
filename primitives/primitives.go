package primitives

import (
	"go3dngine/buffer"
	"go3dngine/types"
	"image/color"
	"math"
)

// DrawLine draw a line from p1 to p2 with color c on given buffer b
func DrawLine(p1, p2 types.Vector2, c color.Color, b *buffer.Buffer) {
	deltaX := p2.X - p1.X
	deltaY := p2.Y - p1.Y

	steps := math.Abs(deltaY)
	if deltaX > deltaY {
		steps = math.Abs(deltaX)
	}

	xIncrement := deltaX / steps
	yIncrement := deltaY / steps

	x := p1.X
	y := p1.Y
	for i := 0; i < int(steps); i++ {
		x += xIncrement
		y += yIncrement
		b.DrawPixel(int(x), int(y), c)
	}
}

// DrawTriangle draws a triangle given as 3 2D vertices array
func DrawTriangle(v [3]types.Vector2, c color.Color, b *buffer.Buffer) {
	f := func(i, j int) {
		DrawLine(v[i], v[j], c, b)
	}

	f(0, 1)
	f(1, 2)
	f(2, 0)
}
