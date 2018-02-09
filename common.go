package gosu

import (
	"fmt"
	"io"

	"github.com/kr/pretty"
)

// Valid osu struct tags:
// `osu-end:"YYYYMMDD"` - Tells to skip this field if the current osu-version
//    is greater than this version string. This is because this field has been
//    removed.
// `osu-start:"YYYYMMDD"` - Tells to skip this field if the current osu-version
//    is less than this version string. This is because this field has not yet
//    been populated yet.

// Type aliases so that the number of bytes match what osu is expecting.
// Exceptions:
// 1. ULEB128 is aliased to uint64 (hopefully this is big enough)
// 2. String is aliased to a struct with all the required information.
type Byte uint8
type Short uint16
type Int uint32
type Long uint64
type ULEB128 uint64
type Single float32
type Double float64
type Boolean uint8
type String struct {
	Cond Byte
	Len  ULEB128
	Text string
}
type DateTime struct {
	Value uint64
}

// A helper method for Pretty printing any object
func PrettyPrint(v interface{}) {
	fmt.Printf("%# v\n", pretty.Formatter(v))
}

func GetVersionOfBinary(buf io.Reader) (Int, error) {
	var version Int
	if err := version.UnmarshalOsuBinary(buf, Int(0)); err != nil {
		return Int(0), err
	}
	return version, nil
}
