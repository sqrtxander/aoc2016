package utils

type (
	HashGrid        map[Point]bool
	BoundedHashGrid struct {
		Grid HashGrid
		W    int
		H    int
	}
)

func (g BoundedHashGrid) GetBoundedHash() string {
	result := ""
	for y := range g.H {
		for x := range g.W {
			if g.Grid[Point{X: x, Y: y}] {
				result += "#"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
    return result[:len(result)-1]
}
