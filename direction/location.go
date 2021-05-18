package direction

import (
	"sort"
)

// ----------------------------------------------------------------------------------------------------------------------------
// Direction
// ----------------------------------------------------------------------------------------------------------------------------
type Direction int

const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

func (p Direction) String() string {
	return directionMap[p].string
}

func (p Direction) Opposite() Direction {
	return directionMap[p].opposite
}

// ----------------------------------------------------------------------------------------------------------------------------
// Corner
// ----------------------------------------------------------------------------------------------------------------------------

type Corner int

const (
	TopLeft     Corner = 0
	TopRight    Corner = 1
	BottomLeft  Corner = 2
	BottomRight Corner = 3
)

func (c Corner) String() string {
	return cornerMap[c].string
}

func (c Corner) Opposite() Corner {
	return cornerMap[c].opposite
}

func (c Corner) Parts() []Direction {
	return cornerMap[c].parts
}

func MakeCorner(p []Direction) Corner {
	sort.Slice(p, func(i, j int) bool {return int(p[i]) < int(p[j])	})
	if p[0] == Up {
		if p[1] == Left {
			return TopLeft
		} else if p[1] == Right {
			return TopRight
		}
	}
	if p[0] == Down {
		if p[1] == Left {
			return BottomLeft
		} else if p[1] == Right {
			return BottomRight
		}
	}
	panic("invalid corner parts")
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
type posInfo struct {
	string
	opposite Direction
}

var directionMap = map[Direction]posInfo{
	Up:    {"Up", Down},
	Down:  {"Down", Up},
	Left:  {"Left", Right},
	Right: {"Right", Left},
}

type cornerInfo struct {
	string
	opposite Corner
	parts    []Direction
}

var cornerMap = map[Corner]cornerInfo{
	TopLeft:     {"TopLeft", BottomRight, []Direction{Up, Left}},
	TopRight:    {"TopRight", BottomLeft, []Direction{Up, Right}},
	BottomLeft:  {"BottomLeft", TopRight, []Direction{Down, Left}},
	BottomRight: {"BottomRight", TopLeft, []Direction{Down, Right}},
}
