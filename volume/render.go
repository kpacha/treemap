// Package volume exposes functions for rendering 3D views of treemaps
package volume

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	"github.com/kpacha/treemap"
	"github.com/tidwall/pinhole"
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

// Image returns an image of the tree using the pinhole lib for drawing 3D cubes
func Image(tree *treemap.Block, width, height float64) (image.Image, error) {
	p := pinhole.New()

	max := math.Max(tree.Width, tree.Depth)
	scale := treemap.Position{X: 1 / max, Y: 1 / max, Z: 0.01}
	if err := render(tree, p, treemap.Position{}, scale); err != nil {
		return nil, err
	}

	p.Rotate(-2*math.Pi/3, 0, 0)
	p.Rotate(0, math.Pi/7, 0)
	return p.Image(int(width), int(height), nil), nil
}

func render(b *treemap.Block, pin *pinhole.Pinhole, offset, p treemap.Position) error {
	off := offset.Add(b.Position)

	c, err := treemap.Color(b.Color[2:]).Decode()
	if err != nil {
		return err
	}

	pin.Begin()
	pin.DrawCube(cubeCoord(b, off, p))
	pin.Colorize(c)
	pin.End()

	for _, child := range b.Children {
		if err = render(child, pin, off, p); err != nil {
			return err
		}
	}
	return nil
}

func cubeCoord(b *treemap.Block, offset, p treemap.Position) (minx, miny, minz, maxx, maxy, maxz float64) {
	return p.X * (offset.X - b.Width/2),
		p.Y * (offset.Y - b.Depth/2),
		p.Z * (b.Position.Z - b.Height/2),
		p.X * (offset.X + b.Width/2),
		p.Y * (offset.Y + b.Depth/2),
		p.Z * (b.Position.Z + b.Height/2)
}
