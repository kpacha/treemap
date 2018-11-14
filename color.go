package treemap

import (
	"encoding/hex"
	"image/color"
)

// Color is a wrapper over a string, useful for string to color conversions and quick
// serialization. The expected format is "0xff00ff"
type Color string

// Decode returns the color described as hex at c
func (c Color) Decode() (color.Color, error) {
	s := string(c)
	if len(s) < 3 {
		return color.Black, nil
	}
	decoded, err := hex.DecodeString(s[2:])
	if err != nil {
		return color.Black, err
	}
	res := color.NRGBA{A: 255}
	if len(decoded) > 0 {
		res.R = decoded[0]
	}
	if len(decoded) > 1 {
		res.G = decoded[1]
	}
	if len(decoded) > 2 {
		res.B = decoded[2]
	}
	return res, nil
}
