package main

import (
	t "art_of_development/1_primitive_types/types"
	"fmt"
)

func main() {
	// t.PrimitivesEx()
	// t.CharsRunesEx()
	t.StringsEx()

	fmt.Println("Test prefix func:", t.HasPrefixInStr("hello", "he"))

	t.RunesInStringEx()
}
