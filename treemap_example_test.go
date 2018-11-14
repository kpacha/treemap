package treemap

import (
	"context"
	"encoding/json"
	"fmt"
)

func Example_fromJSONDefinition() {
	jsonDefinition := `{
	"name": "root",
	"dimm1": 1,
	"dimm2": 2,
	"dimm3": 3,
	"color": "0xff00ff",
	"children": [
		{
			"name": "root #0",
			"dimm1": 5,
			"dimm2": 6,
			"dimm3": 7,
			"color": "0x00ffff"
		},
		{
			"name": "root #1",
			"dimm1": 10,
			"dimm2": 8,
			"dimm3": 3,
			"color": "0x00ff00",
			"children": [
				{
					"name": "root #1 #1",
					"dimm1": 1,
					"dimm2": 4,
					"dimm3": 3,
					"color": "0xffff00"
				}
			]
		}
	]
}`
	tree := TreeInfo{}
	json.Unmarshal([]byte(jsonDefinition), &tree)

	fmt.Println("Get the Block")
	fmt.Println(tree.Block().String())

	fmt.Println("Get the Tree")
	fmt.Println(tree.Tree(context.Background()).String())

	// output:
	// Get the Block
	// {
	// 	"depth": 0,
	// 	"height": 0,
	// 	"width": 0,
	// 	"position": {
	// 		"x": 0,
	// 		"y": 0
	// 	},
	// 	"name": "root",
	// 	"dimm1": 1,
	// 	"dimm2": 2,
	// 	"dimm3": 3,
	// 	"color": "0xff00ff",
	// 	"children": [
	// 		{
	// 			"depth": 0,
	// 			"height": 0,
	// 			"width": 0,
	// 			"position": {
	// 				"x": 0,
	// 				"y": 0
	// 			},
	// 			"name": "root #0",
	// 			"dimm1": 5,
	// 			"dimm2": 6,
	// 			"dimm3": 7,
	// 			"color": "0x00ffff"
	// 		},
	// 		{
	// 			"depth": 0,
	// 			"height": 0,
	// 			"width": 0,
	// 			"position": {
	// 				"x": 0,
	// 				"y": 0
	// 			},
	// 			"name": "root #1",
	// 			"dimm1": 10,
	// 			"dimm2": 8,
	// 			"dimm3": 3,
	// 			"color": "0x00ff00",
	// 			"children": [
	// 				{
	// 					"depth": 0,
	// 					"height": 0,
	// 					"width": 0,
	// 					"position": {
	// 						"x": 0,
	// 						"y": 0
	// 					},
	// 					"name": "root #1 #1",
	// 					"dimm1": 1,
	// 					"dimm2": 4,
	// 					"dimm3": 3,
	// 					"color": "0xffff00"
	// 				}
	// 			]
	// 		}
	// 	]
	// }
	// Get the Tree
	// {
	// 	"depth": 35,
	// 	"height": 6,
	// 	"width": 21,
	// 	"position": {
	// 		"x": 0,
	// 		"y": 0
	// 	},
	// 	"name": "root",
	// 	"dimm1": 1,
	// 	"dimm2": 2,
	// 	"dimm3": 3,
	// 	"color": "0xff00ff",
	// 	"children": [
	// 		{
	// 			"depth": 9,
	// 			"height": 10,
	// 			"width": 8,
	// 			"position": {
	// 				"x": -4.5,
	// 				"y": -10.5,
	// 				"z": 6
	// 			},
	// 			"name": "root #0",
	// 			"dimm1": 5,
	// 			"dimm2": 6,
	// 			"dimm3": 7,
	// 			"color": "0x00ffff"
	// 		},
	// 		{
	// 			"depth": 18,
	// 			"height": 6,
	// 			"width": 17,
	// 			"position": {
	// 				"x": 0,
	// 				"y": 6,
	// 				"z": 6
	// 			},
	// 			"name": "root #1",
	// 			"dimm1": 10,
	// 			"dimm2": 8,
	// 			"dimm3": 3,
	// 			"color": "0x00ff00",
	// 			"children": [
	// 				{
	// 					"depth": 7,
	// 					"height": 6,
	// 					"width": 4,
	// 					"position": {
	// 						"x": 0,
	// 						"y": 0,
	// 						"z": 12
	// 					},
	// 					"name": "root #1 #1",
	// 					"dimm1": 1,
	// 					"dimm2": 4,
	// 					"dimm3": 3,
	// 					"color": "0xffff00"
	// 				}
	// 			]
	// 		}
	// 	]
	// }
}

func Example() {
	tree := TreeInfo{
		BlockInfo: BlockInfo{
			Name:  "root",
			Dimm1: 1,
			Dimm2: 2,
			Dimm3: 3,
			Color: "0xff00ff",
		},
		Children: []*TreeInfo{
			{
				BlockInfo: BlockInfo{
					Name:  "root #0",
					Dimm1: 5,
					Dimm2: 6,
					Dimm3: 7,
					Color: "0x00ffff",
				},
				Children: []*TreeInfo{},
			},
			{
				BlockInfo: BlockInfo{
					Name:  "root #1",
					Dimm1: 10,
					Dimm2: 8,
					Dimm3: 3,
					Color: "0x00ff00",
				},
				Children: []*TreeInfo{
					{
						BlockInfo: BlockInfo{
							Name:  "root #1 #1",
							Dimm1: 1,
							Dimm2: 4,
							Dimm3: 3,
							Color: "0xffff00",
						},
						Children: []*TreeInfo{},
					}},
			},
		},
	}

	fmt.Println(tree.Tree(context.Background()).String())

	// output:
	// {
	// 	"depth": 35,
	// 	"height": 6,
	// 	"width": 21,
	// 	"position": {
	// 		"x": 0,
	// 		"y": 0
	// 	},
	// 	"name": "root",
	// 	"dimm1": 1,
	// 	"dimm2": 2,
	// 	"dimm3": 3,
	// 	"color": "0xff00ff",
	// 	"children": [
	// 		{
	// 			"depth": 9,
	// 			"height": 10,
	// 			"width": 8,
	// 			"position": {
	// 				"x": -4.5,
	// 				"y": -10.5,
	// 				"z": 6
	// 			},
	// 			"name": "root #0",
	// 			"dimm1": 5,
	// 			"dimm2": 6,
	// 			"dimm3": 7,
	// 			"color": "0x00ffff"
	// 		},
	// 		{
	// 			"depth": 18,
	// 			"height": 6,
	// 			"width": 17,
	// 			"position": {
	// 				"x": 0,
	// 				"y": 6,
	// 				"z": 6
	// 			},
	// 			"name": "root #1",
	// 			"dimm1": 10,
	// 			"dimm2": 8,
	// 			"dimm3": 3,
	// 			"color": "0x00ff00",
	// 			"children": [
	// 				{
	// 					"depth": 7,
	// 					"height": 6,
	// 					"width": 4,
	// 					"position": {
	// 						"x": 0,
	// 						"y": 0,
	// 						"z": 12
	// 					},
	// 					"name": "root #1 #1",
	// 					"dimm1": 1,
	// 					"dimm2": 4,
	// 					"dimm3": 3,
	// 					"color": "0xffff00"
	// 				}
	// 			]
	// 		}
	// 	]
	// }
}
