// Package plain exposes functions for rendering vertical projections of treemaps
package plain

import (
	"bytes"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/kpacha/treemap"
)

// NewPNG returns the image of the received tree encoded as a PNG
func NewPNG(tree *treemap.Block, width, height float64) (io.WriterTo, error) {
	return newEncoder(tree, width, height, pngEncode)
}

// NewJPEG returns the image of the received tree encoded as a JPEG
func NewJPEG(tree *treemap.Block, width, height float64) (io.WriterTo, error) {
	return newEncoder(tree, width, height, jpegEncode)
}

// NewGIF returns the image of the received tree encoded as a GIF
func NewGIF(tree *treemap.Block, width, height float64) (io.WriterTo, error) {
	return newEncoder(tree, width, height, gifEncode)
}

type encodeFunc func(io.Writer, image.Image) error

func pngEncode(w io.Writer, i image.Image) error  { return png.Encode(w, i) }
func jpegEncode(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) }
func gifEncode(w io.Writer, i image.Image) error  { return gif.Encode(w, i, nil) }

func newEncoder(block *treemap.Block, width, height float64, enc encodeFunc) (io.WriterTo, error) {
	rectImage, err := Image(block, width, height)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := enc(buf, rectImage); err != nil {
		return nil, err
	}

	return buf, nil
}

// Image returns an image of the tree using a vertincal projection of the treemap
func Image(tree *treemap.Block, width, height float64) (image.Image, error) {
	dst := image.NewRGBA(bounds(width, height))
	if err := drawSubBlock(tree, dst, image.ZP, treemap.Position{X: width / tree.Width, Y: height / tree.Depth}); err != nil {
		return nil, err
	}
	return dst, nil
}

func drawSubBlock(b *treemap.Block, dst draw.Image, offset image.Point, p treemap.Position) error {
	off := offset.Add(image.Pt(int(b.Position.X*p.X), int(b.Position.Y*p.Y)))

	color, err := treemap.Color(b.Color[2:]).Decode()
	if err != nil {
		return err
	}
	draw.Draw(dst, bounds(b.Width*p.X, b.Depth*p.Y).Add(off), &image.Uniform{color}, image.ZP, draw.Src)

	for _, c := range b.Children {
		if err = drawSubBlock(c, dst, off, p); err != nil {
			return err
		}
	}
	return nil
}

func bounds(x, y float64) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: int(-x / 2), Y: int(-y / 2)},
		Max: image.Point{X: int(x / 2), Y: int(y / 2)},
	}
}
