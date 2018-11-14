package plain

import (
	"context"
	"errors"
	"image"
	"io"
	"testing"

	"github.com/kpacha/treemap"
)

func TestEncoder_erroredEncoder(t *testing.T) {
	expectedErr := errors.New("wait for me")
	_, err := newEncoder(
		treemap.NewTree(
			context.Background(),
			treemap.BlockInfo{Color: "0x00ffff"},
		),
		100,
		100,
		func(_ io.Writer, _ image.Image) error { return expectedErr },
	)
	if err != expectedErr {
		t.Errorf("unexpected err: %v", err)
	}
}

func TestEncoder_wrongColorEncoding(t *testing.T) {
	b := treemap.NewBlock(treemap.BlockInfo{
		Name:  "b",
		Dimm1: 1,
		Dimm2: 10,
		Dimm3: 5,
		Color: "zzzzzzzzzzz",
	})
	_, err := NewPNG(treemap.NewTree(context.Background(), treemap.BlockInfo{
		Name:  "root",
		Dimm1: 1,
		Dimm2: 10,
		Dimm3: 5,
		Color: "0x00ffff",
	}, b), 100, 100)
	if err == nil || err.Error() != "encoding/hex: invalid byte: U+007A 'z'" {
		t.Errorf("unexpected err: %v", err)
	}
}
