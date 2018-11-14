package treemap

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
)

func ExampleWalk() {
	root := generateTree(rand.New(rand.NewSource(0)), 5)

	if err := Walk(root, func(b *Block) error {
		fmt.Println(b.Name)
		return nil
	}); err != nil {
		fmt.Println(err.Error())
	}

	// output:
	// root
	// root #0
	// root #0 #0
	// root #0 #0 #0
	// root #0 #0 #0 #0
	// root #0 #0 #0 #0 #0
	// root #0 #0 #1
	// root #0 #0 #1 #0
	// root #0 #0 #1 #0 #0
	// root #0 #0 #1 #1
	// root #0 #0 #1 #1 #0
	// root #0 #3
	// root #0 #3 #0
	// root #0 #3 #0 #1
	// root #0 #3 #1
	// root #0 #3 #1 #0
	// root #0 #3 #2
	// root #0 #3 #2 #0
	// root #0 #3 #2 #0 #0
	// root #0 #3 #2 #1
	// root #1
	// root #1 #0
	// root #1 #0 #0
	// root #1 #0 #0 #0
	// root #1 #0 #1
	// root #1 #0 #2
	// root #1 #0 #2 #0
	// root #1 #0 #2 #0 #0
	// root #1 #1
	// root #1 #1 #0
	// root #1 #1 #0 #0
	// root #1 #1 #0 #0 #0
	// root #1 #1 #0 #1
	// root #1 #1 #0 #1 #0
	// root #1 #1 #1
	// root #1 #1 #1 #1
	// root #1 #1 #1 #1 #0
	// root #1 #1 #2
	// root #1 #1 #2 #0
	// root #1 #2
	// root #1 #2 #0
	// root #1 #2 #0 #0
	// root #1 #2 #0 #1
	// root #1 #2 #2
	// root #1 #2 #2 #0
	// root #1 #2 #2 #0 #0
	// root #1 #3
	// root #1 #3 #0
	// root #1 #3 #0 #1
	// root #1 #3 #1
	// root #1 #3 #1 #1
	// root #1 #3 #1 #1 #0
	// root #1 #3 #2
	// root #2
	// root #2 #0
	// root #2 #0 #0
	// root #2 #0 #0 #0
	// root #2 #0 #0 #0 #0
	// root #2 #0 #2
	// root #2 #0 #2 #0
	// root #2 #0 #2 #1
	// root #2 #0 #2 #1 #0
	// root #2 #1
	// root #2 #1 #0
	// root #2 #1 #0 #0
	// root #2 #1 #0 #0 #0
	// root #2 #1 #1
	// root #2 #1 #1 #0
	// root #2 #1 #1 #0 #0
	// root #2 #1 #1 #1
	// root #2 #1 #1 #1 #0
	// root #2 #1 #2
	// root #2 #1 #2 #1
	// root #2 #1 #2 #1 #0
	// root #2 #2
	// root #2 #2 #0
	// root #2 #2 #0 #1
	// root #2 #2 #1
	// root #2 #2 #1 #0
	// root #2 #2 #1 #0 #0
	// root #2 #3
	// root #2 #3 #0
	// root #2 #3 #0 #1
	// root #2 #3 #0 #1 #0
	// root #2 #3 #2
	// root #3
	// root #3 #1
	// root #3 #1 #0
	// root #3 #1 #0 #0
	// root #3 #1 #0 #0 #0
	// root #3 #1 #0 #1
	// root #3 #1 #0 #1 #0
	// root #3 #2
	// root #3 #2 #2
	// root #3 #2 #2 #1
	// root #3 #3
	// root #3 #3 #0
	// root #3 #3 #0 #0
	// root #3 #3 #0 #0 #0
	// root #3 #3 #0 #1
	// root #3 #3 #0 #1 #0
	// root #3 #3 #1
	// root #3 #3 #1 #0
	// root #3 #3 #1 #0 #0
	// root #3 #3 #2
	// root #3 #3 #2 #0
	// root #3 #3 #2 #0 #0
	// root #3 #3 #2 #1
	// root #3 #3 #2 #1 #0
}

func ExampleWalk_errored() {
	root := generateTree(rand.New(rand.NewSource(0)), 5)
	expectedErr := errors.New("wait for me")

	var counter int
	if err := Walk(root, func(b *Block) error {
		counter++
		if counter == 3 {
			return expectedErr
		}
		if counter > 3 {
			fmt.Println("walk funct called after throwing an error")
		}
		fmt.Println(b.Name)
		return nil
	}); err != nil {
		fmt.Println("got error:", err.Error())
	}
	fmt.Println("total calls to the walk func:", counter)

	// output:
	// root
	// root #0
	// got error: wait for me
	// total calls to the walk func: 3
}

func ExampleWalk_contextCanceled() {
	root := generateTree(rand.New(rand.NewSource(0)), 5)

	ok := WalkFunc(func(b *Block) error {
		fmt.Println(b.Name)
		return nil
	})
	ko := WalkFunc(func(b *Block) error {
		fmt.Println("the walk function should not been called for", b.Name)
		return nil
	})

	if err := Walk(root, ok.WithContext(context.Background())); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("canceled context?")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	fmt.Println(Walk(root, ko.WithContext(ctx)))

	// output:
	// root
	// root #0
	// root #0 #0
	// root #0 #0 #0
	// root #0 #0 #0 #0
	// root #0 #0 #0 #0 #0
	// root #0 #0 #1
	// root #0 #0 #1 #0
	// root #0 #0 #1 #0 #0
	// root #0 #0 #1 #1
	// root #0 #0 #1 #1 #0
	// root #0 #3
	// root #0 #3 #0
	// root #0 #3 #0 #1
	// root #0 #3 #1
	// root #0 #3 #1 #0
	// root #0 #3 #2
	// root #0 #3 #2 #0
	// root #0 #3 #2 #0 #0
	// root #0 #3 #2 #1
	// root #1
	// root #1 #0
	// root #1 #0 #0
	// root #1 #0 #0 #0
	// root #1 #0 #1
	// root #1 #0 #2
	// root #1 #0 #2 #0
	// root #1 #0 #2 #0 #0
	// root #1 #1
	// root #1 #1 #0
	// root #1 #1 #0 #0
	// root #1 #1 #0 #0 #0
	// root #1 #1 #0 #1
	// root #1 #1 #0 #1 #0
	// root #1 #1 #1
	// root #1 #1 #1 #1
	// root #1 #1 #1 #1 #0
	// root #1 #1 #2
	// root #1 #1 #2 #0
	// root #1 #2
	// root #1 #2 #0
	// root #1 #2 #0 #0
	// root #1 #2 #0 #1
	// root #1 #2 #2
	// root #1 #2 #2 #0
	// root #1 #2 #2 #0 #0
	// root #1 #3
	// root #1 #3 #0
	// root #1 #3 #0 #1
	// root #1 #3 #1
	// root #1 #3 #1 #1
	// root #1 #3 #1 #1 #0
	// root #1 #3 #2
	// root #2
	// root #2 #0
	// root #2 #0 #0
	// root #2 #0 #0 #0
	// root #2 #0 #0 #0 #0
	// root #2 #0 #2
	// root #2 #0 #2 #0
	// root #2 #0 #2 #1
	// root #2 #0 #2 #1 #0
	// root #2 #1
	// root #2 #1 #0
	// root #2 #1 #0 #0
	// root #2 #1 #0 #0 #0
	// root #2 #1 #1
	// root #2 #1 #1 #0
	// root #2 #1 #1 #0 #0
	// root #2 #1 #1 #1
	// root #2 #1 #1 #1 #0
	// root #2 #1 #2
	// root #2 #1 #2 #1
	// root #2 #1 #2 #1 #0
	// root #2 #2
	// root #2 #2 #0
	// root #2 #2 #0 #1
	// root #2 #2 #1
	// root #2 #2 #1 #0
	// root #2 #2 #1 #0 #0
	// root #2 #3
	// root #2 #3 #0
	// root #2 #3 #0 #1
	// root #2 #3 #0 #1 #0
	// root #2 #3 #2
	// root #3
	// root #3 #1
	// root #3 #1 #0
	// root #3 #1 #0 #0
	// root #3 #1 #0 #0 #0
	// root #3 #1 #0 #1
	// root #3 #1 #0 #1 #0
	// root #3 #2
	// root #3 #2 #2
	// root #3 #2 #2 #1
	// root #3 #3
	// root #3 #3 #0
	// root #3 #3 #0 #0
	// root #3 #3 #0 #0 #0
	// root #3 #3 #0 #1
	// root #3 #3 #0 #1 #0
	// root #3 #3 #1
	// root #3 #3 #1 #0
	// root #3 #3 #1 #0 #0
	// root #3 #3 #2
	// root #3 #3 #2 #0
	// root #3 #3 #2 #0 #0
	// root #3 #3 #2 #1
	// root #3 #3 #2 #1 #0
	// canceled context?
	// context canceled
}
