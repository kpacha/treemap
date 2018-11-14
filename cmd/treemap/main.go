// Treemap converts the tree definition into a treemap, exporting it as a JSON or as an image
//
// Usage:
// 	treemap -f jpeg -s volume -o tree.jpg input_file.json
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kpacha/treemap"
	"github.com/kpacha/treemap/plain"
	"github.com/kpacha/treemap/volume"
)

func main() {
	encoding := flag.String("f", "png", "encoding to use (gif,jpeg,png,none)")
	width := flag.Int("x", 720, "width")
	height := flag.Int("y", 720, "height")
	style := flag.String("s", "plain", "render package to use (plain, volume)")
	out := flag.String("o", "", "output")
	flag.Parse()

	encoders, ok := renders[strings.ToLower(*style)]
	if !ok {
		log.Fatalf("unknown package %s", *style)
	}

	encoderFn, ok := encoders[strings.ToLower(*encoding)]
	if !ok {
		log.Fatalf("unknown encoding %s", *encoding)
	}

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("input file argument required")
	}

	input := args[0]

	tree, err := process(input)
	if err != nil {
		log.Fatalf("processing (%s): %s", input, err.Error())
	}

	wt, err := encoderFn(tree, float64(*width), float64(*height))
	if err != nil {
		log.Fatal(err)
	}

	if *out == "" {
		if _, err := wt.WriteTo(os.Stdout); err != nil {
			log.Fatal(err)
		}
		return
	}

	b := new(bytes.Buffer)
	if _, err := wt.WriteTo(b); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(*out, b.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

}

func process(input string) (*treemap.Block, error) {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		return nil, err
	}

	treeInfo := treemap.TreeInfo{}
	if err := json.Unmarshal(data, &treeInfo); err != nil {
		return nil, err
	}

	return treeInfo.Tree(context.TODO()), nil
}

type encoderFunc func(*treemap.Block, float64, float64) (io.WriterTo, error)

var renders = map[string]map[string]encoderFunc{
	"plain": {
		"png":  plain.NewPNG,
		"jpeg": plain.NewJPEG,
		"gif":  plain.NewGIF,
		"none": jsonRender,
	},
	"volume": {
		"png":  volume.NewPNG,
		"jpeg": volume.NewJPEG,
		"gif":  volume.NewGIF,
		"none": jsonRender,
	},
}

func jsonRender(tree *treemap.Block, _, _ float64) (io.WriterTo, error) {
	return bytes.NewBufferString(tree.String()), nil
}
