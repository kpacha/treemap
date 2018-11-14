package treemap

import (
	"math"
)

// Position is an X, Y, Z tuple.
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z,omitempty"`
}

// Add returns the 3D vector p + q
func (p Position) Add(q Position) Position {
	return Position{X: p.X + q.X, Y: p.Y + q.Y, Z: p.Z + q.Z}
}

// Tiler is a quick implementation of a solution for the tiling problem. It
// takes an amount of tiles to place in one plane starting at {0, 0} and
// identifies where a tile should be placed and what is the smaller rectangle
// containing all the passed tiles.
type Tiler struct {
	tiles        int
	margin       float64
	dimension    int
	xReference   float64
	yReference   float64
	currentIndex int
	maxWidth     float64
	maxHeight    float64
	bounds       Position
}

const (
	defaultMargin = 3
)

// NewTiler returns a tiler expecting to place totalTiles with a default margin
func NewTiler(totalTiles int) *Tiler {
	return NewTilerWithMargin(totalTiles, defaultMargin)
}

// NewTilerWithMargin returns a tiler expecting to place totalTiles with the injected margin
func NewTilerWithMargin(totalTiles int, margin float64) *Tiler {
	tiler := &Tiler{
		tiles:     totalTiles,
		margin:    margin,
		dimension: int(math.Ceil(math.Sqrt(float64(totalTiles)))),
	}

	return tiler
}

// GetBounds returns the size of the rectangle containing all the tiles placed so far
func (g *Tiler) GetBounds() Position {
	return Position{
		X: g.maxWidth + g.margin,
		Y: g.maxHeight + g.margin,
	}
}

// NextPosition calculates where a tile of the passed dimensions should be placed
func (g *Tiler) NextPosition(width, height float64) Position {
	g.currentIndex++

	if g.currentIndex > g.dimension && g.yReference+height >= g.maxWidth {
		g.currentIndex = 0
		g.yReference = 0
		g.xReference = g.maxWidth + g.margin
	}

	position := Position{X: g.xReference + (width+g.margin)/2, Y: g.yReference + (height+g.margin)/2}

	if g.xReference+width > g.maxWidth {
		g.maxWidth = g.xReference + width
	}

	if g.yReference+height > g.maxHeight {
		g.maxHeight = g.yReference + height
	}

	g.yReference += height + g.margin

	return position
}
