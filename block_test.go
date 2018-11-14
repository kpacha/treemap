package treemap

import (
	"context"
	"math/rand"
	"testing"
	"testing/quick"
)

func TestTrees_deterministic(t *testing.T) {
	f := func(b *Block) string {
		prepareNode(context.Background(), b, 0)
		return b.String()
	}
	if err := quick.CheckEqual(f, f, nil); err != nil {
		t.Error(err)
	}
}

func TestBlock_String(t *testing.T) {
	b1 := NewBlock(BlockInfo{
		Name:  "b1",
		Dimm1: 1,
		Dimm2: 10,
		Dimm3: 5,
		Color: "0x00ff00",
	})
	b2 := NewBlock(BlockInfo{
		Name:  "b2",
		Dimm1: 10,
		Dimm2: 1,
		Dimm3: 2,
		Color: "0xff0000",
	})
	b := NewTree(context.Background(), BlockInfo{
		Name:  "root",
		Dimm1: 5,
		Dimm2: 5,
		Dimm3: 10,
		Color: "0x0000ff",
	}, b1, b2)

	if text := b.String(); customBlockDump != text {
		t.Error("unexpected result:", text)
	}

	root := generateTree(rand.New(rand.NewSource(0)), 2)
	if text := root.String(); generatedBlockDump != text {
		t.Error("unexpected result:", text)
	}
}

var (
	customBlockDump = `{
	"depth": 28,
	"height": 13,
	"width": 21,
	"position": {
		"x": 0,
		"y": 0
	},
	"name": "root",
	"dimm1": 5,
	"dimm2": 5,
	"dimm3": 10,
	"color": "0x0000ff",
	"children": [
		{
			"depth": 13,
			"height": 8,
			"width": 4,
			"position": {
				"x": -4.5,
				"y": -3.5,
				"z": 13
			},
			"name": "b1",
			"dimm1": 1,
			"dimm2": 10,
			"dimm3": 5,
			"color": "0x00ff00"
		},
		{
			"depth": 4,
			"height": 5,
			"width": 13,
			"position": {
				"x": 0,
				"y": 8,
				"z": 13
			},
			"name": "b2",
			"dimm1": 10,
			"dimm2": 1,
			"dimm3": 2,
			"color": "0xff0000"
		}
	]
}`
	generatedBlockDump = `{
	"depth": 15,
	"height": 6,
	"width": 12,
	"position": {
		"x": 0,
		"y": 0
	},
	"name": "root",
	"dimm1": 2,
	"dimm2": 4,
	"dimm3": 3,
	"color": "0xb44ce8",
	"children": [
		{
			"depth": 8,
			"height": 3,
			"width": 7,
			"position": {
				"x": 0,
				"y": 0,
				"z": 6
			},
			"name": "root #0",
			"dimm1": 1,
			"dimm2": 2,
			"dimm3": 0,
			"color": "0xf95ff6",
			"children": [
				{
					"depth": 3,
					"height": 3,
					"width": 3,
					"position": {
						"x": 0,
						"y": 0,
						"z": 9
					},
					"name": "root #0 #0",
					"dimm1": 0,
					"dimm2": 0,
					"dimm3": 0,
					"color": "0x73c86e"
				}
			]
		}
	]
}`
)
