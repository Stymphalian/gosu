package gosu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"reflect"

	"github.com/bnch/uleb128"
	"github.com/kr/pretty"
)

// Type aliases so that the number of bytes match what osu is expecting in
// all of the structs. These match up to the spec exepct for:
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

// Unmarshaler interface for Osu Binary blobs
type BinaryOsuUnmarshaler interface {
	UnmarshalBinary(buf io.Reader) error
}

// Marshaler interface for Osu Binary blobs
type BinaryOsuMarshaler interface {
	MarshalBinary() (buf io.Writer, err error)
}

// All the unmarshalling methods for the 'primitive' types above
// ----------------------------------------------------------------------------

func (this *Byte) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, this)
}
func (this *Short) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, this)
}
func (this *Int) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, this)
}
func (this *Long) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, this)
}
func (this *ULEB128) UnmarshalBinary(buf io.Reader) error {
	total := uleb128.UnmarshalReader(buf)
	*this = ULEB128(uint64(total))

	// TODO(jordanyu): Add error handling logic
	return nil
}
func (this *Single) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, this)
}
func (this *Double) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, this)
}
func (this *Boolean) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *String) UnmarshalBinary(buf io.Reader) error {
	err := binary.Read(buf, binary.LittleEndian, &this.Cond)
	if err != nil {
		return err
	}

	if this.Cond == 0xb {
		// TODO(jordanyu): Add error handling logic
		this.Len.UnmarshalBinary(buf)

		var stringBuf bytes.Buffer
		for i := ULEB128(0); i < this.Len; i++ {
			var b []byte = make([]byte, 1, 1)
			n, err := buf.Read(b)
			if n != 1 || err != nil {
				log.Fatal("Failed to read string byte.")
			}
			stringBuf.WriteByte(b[0])
		}
		this.Text = stringBuf.String()
	}
	return nil
}

func (this *DateTime) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, &this.Value)
}

// Lots of little helper methods
//-----------------------------------------------------------------------------

// A helper method for Pretty printing any object
func PrettyPrint(v interface{}) {
	fmt.Printf("%# v", pretty.Formatter(v))
}

// Invoke call a method on the given interface using reflection
// Args:
//   any: Any object for which you want to call a method on. Be sure to pass
//     in the pointer if the method belongs to the pointer interface
//   name: The method name to invoke.
//   args: Vardiadic number of args to pass to the method.
func Invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

// User reflection to unmarshal all the fields in the given interface
// This will loop through every field and call the 'UnmarshalBinary' method
// on the type passing in the 'buf' which is the source of all the bytes
// A special case is where a field is a slice. In this case there is a special
// semantic
// 1. We look up the field name like "Num<SliceFieldName" and retrieve the number
//    of elements to exepect from the stream
// 2. Create a new slice
// 3. Iterate through each slice eleement and run the UnmarshalBinary method
// Args:
//   db: The object to unmarshal
///  buf: The buffer in which we retrieve bytes to unmarshal
func UnmarshalAny(db interface{}, buf io.Reader) error {
	dbVal := reflect.ValueOf(db).Elem()

	for i := 0; i < dbVal.NumField(); i++ {
		currentMutableField := reflect.ValueOf(db).Elem().Field(i)

		if dbVal.Field(i).Type().Kind() == reflect.Slice {
			numFieldName := "Num" + dbVal.Type().Field(i).Name
			numElements := dbVal.FieldByName(numFieldName).Interface()
			// TODO(jordanyu): This is hack, we can't always assume it will be an Int
			// when we have a slice of elements afterwards
			intNumElements := int(numElements.(Int))

			// Create a new slice of appropriate size and then run the unmarshal
			// function over each element in the slice.
			sliceElems := reflect.MakeSlice(currentMutableField.Type(),
				int(intNumElements), int(intNumElements))
			currentMutableField.Set(sliceElems)
			for j := 0; j < intNumElements; j++ {
				ret := Invoke(currentMutableField.Index(j).Addr().Interface(),
					"UnmarshalBinary", buf)

				err := ret[0].Interface()
				if err != nil {
					return err.(error)
				}
			}

		} else {
			// This is just a primitive field so just run the simple unmarhsal
			ret := Invoke(currentMutableField.Addr().Interface(), "UnmarshalBinary", buf)

			err := ret[0].Interface()
			if err != nil {
				return err.(error)
			}
		}
	}

	return nil
}
