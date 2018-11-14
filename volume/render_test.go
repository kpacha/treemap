package volume

// import (
// 	"math/rand"
// 	"testing"

// 	"github.com/kpacha/treemap"
// )

// func TestRender(t *testing.T) {
// 	b := &treemap.Block{}
// 	b = b.Generate(rand.New(rand.NewSource(0)), 5).Interface().(*treemap.Block)
// 	if err := Render(b, 750, 750); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestEncoder_wrongColorEncoding(t *testing.T) {
// 	b := treemap.NewBlock(treemap.BlockInfo{
// 		Name:  "b",
// 		Dimm1: 1,
// 		Dimm2: 10,
// 		Dimm3: 5,
// 		Color: "zzzzzzzzzzz",
// 	})
// 	err := Render(treemap.NewTree(treemap.BlockInfo{
// 		Name:  "root",
// 		Dimm1: 1,
// 		Dimm2: 10,
// 		Dimm3: 5,
// 		Color: "0x00ffff",
// 	}, b), 100, 100)
// 	if err == nil || err.Error() != "encoding/hex: invalid byte: U+007A 'z'" {
// 		t.Errorf("unexpected err: %v", err)
// 	}
// }
