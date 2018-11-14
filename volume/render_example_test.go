package volume

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/kpacha/treemap"
)

func ExampleNewPNG() {
	encoder, err := NewPNG(generateTree(), 1024, 768)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf := new(bytes.Buffer)
	fmt.Println(encoder.WriteTo(buf))
	fmt.Println(ioutil.WriteFile("../docs/volume.png", buf.Bytes(), 0644))

	// output:
	// 85536 <nil>
	// <nil>
}

func ExampleNewJPEG() {
	encoder, err := NewJPEG(generateTree(), 1024, 768)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf := new(bytes.Buffer)
	fmt.Println(encoder.WriteTo(buf))
	fmt.Println(ioutil.WriteFile("../docs/volume.jpg", buf.Bytes(), 0644))

	// output:
	// 50430 <nil>
	// <nil>
}

func ExampleNewGIF() {
	encoder, err := NewGIF(generateTree(), 1024, 768)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf := new(bytes.Buffer)
	fmt.Println(encoder.WriteTo(buf))
	fmt.Println(ioutil.WriteFile("../docs/volume.gif", buf.Bytes(), 0644))

	// output:
	// 52546 <nil>
	// <nil>
}

func generateTree() *treemap.Block {
	b := &treemap.Block{}
	return b.Generate(rand.New(rand.NewSource(0)), 5).Interface().(*treemap.Block)
}
