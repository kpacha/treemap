package treemap

import (
	"reflect"
	"testing"
)

func TestTiler(t *testing.T) {
	tiler := NewTiler(4)

	for i, tc := range []struct {
		Input  Position
		Output Position
	}{
		{
			Input:  Position{X: 10, Y: 1},
			Output: Position{X: 6.5, Y: 2},
		},
		{
			Input:  Position{X: 1, Y: 10},
			Output: Position{X: 2, Y: 10.5},
		},
		{
			Input:  Position{X: 5, Y: 5},
			Output: Position{X: 17, Y: 4},
		},
		{
			Input:  Position{X: 5, Y: 5},
			Output: Position{X: 17, Y: 12},
		},
	} {
		position := tiler.NextPosition(tc.Input.X, tc.Input.Y)
		if !reflect.DeepEqual(position, tc.Output) {
			t.Errorf("unexpected output for input %d: %v", i, position)
		}
	}

	if !reflect.DeepEqual(Position{X: 21, Y: 17}, tiler.GetBounds()) {
		t.Errorf("unexpected bounds: %v", tiler.GetBounds())
	}
}
