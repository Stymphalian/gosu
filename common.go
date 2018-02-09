package gosu

import (
	"fmt"

	"github.com/kr/pretty"
)

// Type aliases so that the number of bytes match what osu is expecting.
// Exceptions:
// 1. ULEB128 is aliased to uint64 (hopefully this is big enough)
// 2. String is aliased to a struct with all the required information.
type ULEB128 uint64
type String struct {
	Cond byte
	Len  ULEB128
	Text string
}

// A helper method for Pretty printing any object
func PrettyPrint(v interface{}) {
	fmt.Printf("%# v\n", pretty.Formatter(v))
}
