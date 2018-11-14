package treemap

import (
	"fmt"
)

func ExampleColor_Decode() {
	for _, tc := range []string{
		"0xff00ff",
		"0x",
		"0x00ff",
		"0xffff",
		"0x00ffff",
		"0xzzzzzz",
	} {
		c, err := Color(tc).Decode()
		if err != nil {
			fmt.Printf("!!! %s: %s\n", tc, err.Error())
			continue
		}
		r, g, b, a := c.RGBA()
		fmt.Printf("%s: {%d, %d, %d, %d}\n", tc, r, g, b, a)
	}
	// output:
	// 0xff00ff: {65535, 0, 65535, 65535}
	// 0x: {0, 0, 0, 65535}
	// 0x00ff: {0, 65535, 0, 65535}
	// 0xffff: {65535, 65535, 0, 65535}
	// 0x00ffff: {0, 65535, 65535, 65535}
	// !!! 0xzzzzzz: encoding/hex: invalid byte: U+007A 'z'
}
