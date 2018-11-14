package treemap

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
)

// NewTree returns a Block with the received info with all the injected children
// already positioned and sized. This function is intended to be used when programmatically
// building the root node of a tree, but it is safe to be called more than once. The returned Block
// can also be used as part of a bigger treemap.
func NewTree(ctx context.Context, info BlockInfo, children ...*Block) *Block {
	b := NewBlock(info, children...)
	prepareNode(ctx, b, 0)
	return b
}

// NewBlock returns a Block with the received info, containing the injected children
// but without setting the node details. This function is intended to be used when programmatically
// building all the nodes to be placed in a tree but the root.
func NewBlock(info BlockInfo, children ...*Block) *Block {
	return &Block{
		BlockInfo: info,
		BlockNode: BlockNode{},
		Children:  children,
	}
}

// BlockNode contains information regarding the size and position of every Block in the tree map
type BlockNode struct {
	Depth    float64  `json:"depth"`
	Height   float64  `json:"height"`
	Width    float64  `json:"width"`
	Position Position `json:"position"`
}

// BlockInfo contains all the user defined details of every Block in the tree map
type BlockInfo struct {
	Name  string `json:"name"`
	Dimm1 int    `json:"dimm1"`
	Dimm2 int    `json:"dimm2"`
	Dimm3 int    `json:"dimm3"`
	Color Color  `json:"color,omitempty"`
}

// Generate implements the quick.Generator interface (https://golang.org/pkg/testing/quick/#Generator)
func (b BlockInfo) Generate(rand *rand.Rand, size int) reflect.Value {
	name := make([]byte, size)
	color := make([]byte, 3)
	rand.Read(name)
	rand.Read(color)

	return reflect.ValueOf(BlockInfo{
		Name:  hex.EncodeToString(name),
		Dimm1: rand.Intn(2*size + 1),
		Dimm2: rand.Intn(2*size + 1),
		Dimm3: rand.Intn(2*size + 1),
		Color: Color(fmt.Sprintf("0x%x", color)),
	})
}

// Block is a tree of blocks. Every Block contains a list of children, a BlockInfo with the input values
// and a BlockNode with all the information regardind the block position and its dimensions
type Block struct {
	BlockNode
	BlockInfo
	Children []*Block `json:"children,omitempty"`
}

// String implements the fmt.Stringer interace by returning an indentated
// json serialization of the Block
func (b *Block) String() string {
	buf, _ := json.MarshalIndent(b, "", "\t")
	return string(buf)
}

// Generate implements the quick.Generator interface (https://golang.org/pkg/testing/quick/#Generator)
func (b *Block) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(generateTree(rand, size))
}

func prepareNode(ctx context.Context, b *Block, z float64) {
	b.Position.Z = z
	b.Height = float64(b.Dimm3) + 3

	if b.Children == nil || len(b.Children) == 0 {
		b.Width = float64(b.Dimm1) + 3
		b.Depth = float64(b.Dimm2) + 3
		return
	}

	tiler := NewTiler(len(b.Children))
	for _, child := range b.Children {
		select {
		case <-ctx.Done():
			return
		default:
		}
		prepareNode(ctx, child, z+b.Height)
		cp := tiler.NextPosition(child.Width, child.Depth)
		child.Position.X = cp.X
		child.Position.Y = cp.Y
	}

	bounds := tiler.GetBounds()
	b.Width, b.Depth = bounds.X, bounds.Y

	for _, child := range b.Children {
		child.Position.X -= b.Width / 2.0
		child.Position.Y -= b.Depth / 2.0
	}

	b.Width += float64(b.Dimm1)
	b.Depth += float64(b.Dimm2)
}

func generateTree(rand *rand.Rand, size int) *Block {
	infoTree := (&TreeInfo{}).Generate(rand, size).Interface().(*TreeInfo)
	return infoTree.Tree(context.TODO())
}
