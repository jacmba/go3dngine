package primitives

import (
	"go3dngine/buffer"
	"go3dngine/types"
	"image/color"
)

// DrawLine draw a line from p1 to p2 with color c on given buffer b
func DrawLine(p1, p2 types.Vector2, c color.Color, b *buffer.Buffer) {
	deltaX := p2.X - p1.X
	deltaY := p2.Y - p1.Y
	slope := deltaY / deltaX

	step := 0.1
	pA := &p1
	pB := &p2

	if deltaX < 0.0 {
		pA = &p2
		pB = &p1
	}

	//if math.Abs(deltaX) > math.Abs(deltaY) {
	for x := pA.X; x < pB.X; x += step {
		y := slope*(x-pA.X) + pA.Y // Rect line ecuation
		ix := int(x)
		iy := int(y)

		b.DrawPixel(ix, iy, c)
	}
	/*} else {
		for y := pA.Y; y < pB.Y; y += step {
			x := (y-pA.Y)/slope + pA.X
			ix := int(x)
			iy := int(y)

			b.DrawPixel(ix, iy, c)
		}
	}*/
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
