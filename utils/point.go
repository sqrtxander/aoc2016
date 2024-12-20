package utils

type Point struct {
	X int
	Y int
}

func ORIGIN() Point {
	return Point{X: 0, Y: 0}
}

type Direction complex128

const (
	NORTH Direction = complex(0, 1)
	EAST  Direction = complex(1, 0)
	SOUTH Direction = complex(0, -1)
	WEST  Direction = complex(-1, 0)
)

// rotate direction clockwise
func (d *Direction) RotateCW() {
	*d *= complex(0, -1)
}

// rotate direction counter clockwise
func (d *Direction) RotateCCW() {
	*d *= complex(0, 1)
}

// rotate direction 180 degrees
func (d *Direction) Rotate180() {
	*d *= -1
}

func (p *Point) MoveInDir(dir Direction, amount int) {
	p.X += int(real(dir)) * amount
	p.Y += int(imag(dir)) * amount
}

// the manhattan distance between a point and the origin
func (p Point) Manhattan() int {
	return Abs(p.X) + Abs(p.Y)
}

// the absolute value of an integer
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Adjacent4(p Point) [4]Point {
	return [4]Point{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
	}
}

func Adjacent8(p Point) [8]Point {
	return [8]Point{
		{X: p.X + 1, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X, Y: p.Y + 1},
	}
}

func Add(p, q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}
