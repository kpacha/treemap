package plain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/kpacha/treemap"
)

var block *treemap.Block

func init() {
	data, err := ioutil.ReadFile("../docs/tree.json")
	if err == nil && json.Unmarshal(data, block) == nil {
		return
	}

	b := &treemap.Block{}
	block = b.Generate(rand.New(rand.NewSource(123)), 6).Interface().(*treemap.Block)
	ioutil.WriteFile("../docs/tree.json", []byte(block.String()), 0664)
}

func ExampleNewPNG() {
	encoder, err := NewPNG(block, 1024, 768)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf := new(bytes.Buffer)
	fmt.Println(encoder.WriteTo(buf))
	fmt.Println(ioutil.WriteFile("../docs/plain.png", buf.Bytes(), 0644))

	// output:
	// 6350 <nil>
	// <nil>
}

func ExampleNewJPEG() {
	encoder, err := NewJPEG(block, 1024, 768)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf := new(bytes.Buffer)
	fmt.Println(encoder.WriteTo(buf))
	fmt.Println(ioutil.WriteFile("../docs/plain.jpg", buf.Bytes(), 0644))

	// output:
	// 59797 <nil>
	// <nil>
}

func ExampleNewGIF() {
	encoder, err := NewGIF(block, 1024, 768)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf := new(bytes.Buffer)
	fmt.Println(encoder.WriteTo(buf))
	fmt.Println(ioutil.WriteFile("../docs/plain.gif", buf.Bytes(), 0644))

	// output:
	// 44119 <nil>
	// <nil>
}
