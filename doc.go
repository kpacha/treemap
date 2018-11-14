// Package treemap implements functions and structs for building treemaps
//
// TreeMaps are an alternative way to display properties of structured layered data
// (such as filesystems). Every data point can contain up to 4 different dimmensions:
// 3 numericals and a color encoded one.
//
// Consider the folowing description of a simple, hypotetical code analyzer report. In this case
// the dimm1 is the number of lines of code contained by the node, dimm2 is the number of types,
// dimm3, the number of functions and the color is encoding the complexity of the node (CCN-like mesure)
//
// 	{
// 		"name": "mypackage",
// 		"dimm1": 1,
// 		"dimm2": 2,
// 		"dimm3": 3,
// 		"color": "0xff00ff",
// 		"children": [
// 			{
// 				"name": "mypackage/subpackage0",
// 				"dimm1": 5,
// 				"dimm2": 6,
// 				"dimm3": 7,
// 				"color": "0x00ffff"
// 			},
// 			{
// 				"name": "mypackage/subpackage1",
// 				"dimm1": 10,
// 				"dimm2": 8,
// 				"dimm3": 3,
// 				"color": "0x00ff00",
// 				"children": [
// 					{
// 						"name": "mypackage/subpackage1/subpackage0",
// 						"dimm1": 1,
// 						"dimm2": 4,
// 						"dimm3": 3,
// 						"color": "0xffff00"
// 					}
// 				]
// 			}
// 		]
// 	}
//
// treemap generates an extended version of the tree description, adding spatial coordinates and dimmensions for
// every package (block) in the tree.
//
// In this extended version, dimm1 will affect the width of the block; dimm2, its depth and dimm3 its height.
//
// The tiling algorithm is based in:
//
// 	- https://github.com/rodrigo-brito/gocity/blob/master/model/position.go
// 	- https://www.codeproject.com/Articles/210979/Fast-optimizing-rectangle-packing-algorithm-for-bu
// 	- http://www.cs.umd.edu/hcil/treemap-history/index.shtml
package treemap
