package treemap

import (
	"context"
)

// WalkFunc is the type of the function called for every Block
// visited by Walk. After the first error returned by the WalkFunc,
// the Walk is finished and the error propagated
type WalkFunc func(*Block) error

// WithContext wraps the WalfFunc w with another WalfFunc containing a check for context
// cancelations before executing the actual function w
func (w WalkFunc) WithContext(ctx context.Context) WalkFunc {
	return func(b *Block) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return w(b)
	}
}

// Walk walks the block tree rooted at root depth-first, calling walkFn for each Block in the
// tree
func Walk(root *Block, walkFn WalkFunc) error {
	if err := walkFn(root); err != nil {
		return err
	}
	for _, c := range root.Children {
		if err := Walk(c, walkFn); err != nil {
			return err
		}
	}
	return nil
}
