package treemap

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
)

// TreeInfo is a tree of BlockInfo structs describing a treemap
type TreeInfo struct {
	BlockInfo
	Children []*TreeInfo `json:"children,omitempty"`
}

// Tree returns the initialized tree described by t
func (t *TreeInfo) Tree(ctx context.Context) *Block {
	children := make([]*Block, len(t.Children))
	for i, info := range t.Children {
		children[i] = info.Block()
	}

	return NewTree(ctx, t.BlockInfo, children...)
}

// Block returns the tree described by t without initializing the positions
// of the blocks
func (t *TreeInfo) Block() *Block {
	children := make([]*Block, len(t.Children))
	for i, info := range t.Children {
		children[i] = info.Block()
	}

	return NewBlock(t.BlockInfo, children...)
}

// String implements the fmt.Stringer interace by returning an indentated
// json serialization of the TreeInfo
func (t *TreeInfo) String() string {
	buf, _ := json.MarshalIndent(t, "", "\t")
	return string(buf)
}

// Generate implements the quick.Generator interface (https://golang.org/pkg/testing/quick/#Generator)
func (t *TreeInfo) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(generateTreeInfo(rand, "root", size))
}

func generateTreeInfo(rand *rand.Rand, name string, size int) *TreeInfo {
	children := []*TreeInfo{}
	for i := 0; i < size; i++ {
		if rand.Intn(10) > 6 {
			continue
		}
		children = append(children, generateTreeInfo(rand, fmt.Sprintf("%s #%d", name, i), size-1))
	}
	info := BlockInfo{}.Generate(rand, size).Interface().(BlockInfo)
	info.Name = name
	return &TreeInfo{
		BlockInfo: info,
		Children:  children,
	}
}
