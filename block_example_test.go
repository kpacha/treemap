package treemap

import (
	"context"
	"fmt"
)

func ExampleNewTree() {
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

	fmt.Println(b.String())

	// output:
	// {
	// 	"depth": 28,
	// 	"height": 13,
	// 	"width": 21,
	// 	"position": {
	// 		"x": 0,
	// 		"y": 0
	// 	},
	// 	"name": "root",
	// 	"dimm1": 5,
	// 	"dimm2": 5,
	// 	"dimm3": 10,
	// 	"color": "0x0000ff",
	// 	"children": [
	// 		{
	// 			"depth": 13,
	// 			"height": 8,
	// 			"width": 4,
	// 			"position": {
	// 				"x": -4.5,
	// 				"y": -3.5,
	// 				"z": 13
	// 			},
	// 			"name": "b1",
	// 			"dimm1": 1,
	// 			"dimm2": 10,
	// 			"dimm3": 5,
	// 			"color": "0x00ff00"
	// 		},
	// 		{
	// 			"depth": 4,
	// 			"height": 5,
	// 			"width": 13,
	// 			"position": {
	// 				"x": 0,
	// 				"y": 8,
	// 				"z": 13
	// 			},
	// 			"name": "b2",
	// 			"dimm1": 10,
	// 			"dimm2": 1,
	// 			"dimm3": 2,
	// 			"color": "0xff0000"
	// 		}
	// 	]
	// }
}

func ExampleNewBlock() {
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
	b := NewBlock(BlockInfo{
		Name:  "root",
		Dimm1: 5,
		Dimm2: 5,
		Dimm3: 10,
		Color: "0x0000ff",
	}, b1, b2)

	fmt.Println(b.String())

	// output:
	// {
	// 	"depth": 0,
	// 	"height": 0,
	// 	"width": 0,
	// 	"position": {
	// 		"x": 0,
	// 		"y": 0
	// 	},
	// 	"name": "root",
	// 	"dimm1": 5,
	// 	"dimm2": 5,
	// 	"dimm3": 10,
	// 	"color": "0x0000ff",
	// 	"children": [
	// 		{
	// 			"depth": 0,
	// 			"height": 0,
	// 			"width": 0,
	// 			"position": {
	// 				"x": 0,
	// 				"y": 0
	// 			},
	// 			"name": "b1",
	// 			"dimm1": 1,
	// 			"dimm2": 10,
	// 			"dimm3": 5,
	// 			"color": "0x00ff00"
	// 		},
	// 		{
	// 			"depth": 0,
	// 			"height": 0,
	// 			"width": 0,
	// 			"position": {
	// 				"x": 0,
	// 				"y": 0
	// 			},
	// 			"name": "b2",
	// 			"dimm1": 10,
	// 			"dimm2": 1,
	// 			"dimm3": 2,
	// 			"color": "0xff0000"
	// 		}
	// 	]
	// }
}
