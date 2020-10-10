package primitives

import (
	"go3dngine/buffer"
	"go3dngine/types"
	"image/color"
	"math"
)

// DrawLine draw a line from p1 to p2 with color c on given buffer b
func DrawLine(p1, p2 types.Vector2, c color.Color, b *buffer.Buffer) {
	deltaX := int(math.Abs(p2.X - p1.X))
	deltaY := -int(math.Abs(p2.Y - p1.Y))

	sX := 1
	sY := 1

	if p2.X <= p1.X {
		sX = -1
	}

	if p2.Y <= p1.Y {
		sY = -1
	}

	err := deltaX + deltaY

	x := int(p1.X)
	y := int(p1.Y)

	for {
		b.DrawPixel(x, y, c)
		if x == int(p2.X) && y == int(p2.Y) {
			break
		}
		e2 := 2 * err

		if e2 >= deltaY {
			err += deltaY
			x += sX
		}

		if e2 <= deltaX {
			err += deltaX
			y += sY
		}
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
