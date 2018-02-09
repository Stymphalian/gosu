package gosu

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"reflect"
	"strconv"

	"github.com/bnch/uleb128"
)

//go:generate go run tools/gen_codec.go

type BinaryOsuCodec interface {
	BinaryOsuMarshaler
	BinaryOsuUnmarshaler
}

// Unmarshaler interface for Osu Binary blobs
type BinaryOsuUnmarshaler interface {
	UnmarshalOsuBinary(buf io.Reader, version Int) error
}

// Marshaler interface for Osu Binary blobs
type BinaryOsuMarshaler interface {
	MarshalOsuBinary(buf io.Writer, version Int) error
}

// Marshal/Unmarshal function for ULEB128 and String
// The other types can be found in auto.codec.go
// -----------------------------------------------------------------------------

func (this *ULEB128) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	total := uleb128.UnmarshalReader(buf)
	*this = ULEB128(uint64(total))

	// TODO(jordanyu): Add error handling logic
	return nil
}

func (this *ULEB128) MarshalOsuBinary(buf io.Writer, version Int) error {
	got := uleb128.Marshal(int(uint64(*this)))
	buf.Write(got)
	return nil
}

func (this *String) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	err := binary.Read(buf, binary.LittleEndian, &this.Cond)
	if err != nil {
		return err
	}

	if this.Cond == 0xb {
		// TODO(jordanyu): Add error handling logic
		err := this.Len.UnmarshalOsuBinary(buf, version)
		if err != nil {
			return err
		}

		var stringBuf bytes.Buffer
		var into []byte = make([]byte, int(this.Len))
		n, err := buf.Read(into)
		if err != nil {
			return err
		}
		if n != int(this.Len) {
			return errors.New("Did not read enough bytes for string")
		}
		stringBuf.Write(into)
		this.Text = stringBuf.String()
	}
	return nil
}

func (this *String) MarshalOsuBinary(buf io.Writer, version Int) error {
	err := binary.Write(buf, binary.LittleEndian, &this.Cond)
	if err != nil {
		return err
	}

	if this.Cond == 0xb {
		// TODO(jordanyu): Add error handling logic
		err := this.Len.MarshalOsuBinary(buf, version)
		if err != nil {
			return err
		}

		stringBuf := bytes.NewBufferString(this.Text)
		n, err := buf.Write(stringBuf.Bytes())
		if err != nil {
			return err
		}
		if n != int(this.Len) {
			return errors.New("Did not read enough bytes for string")
		}
	}
	return nil
}

// Lots of little helper methods
//-----------------------------------------------------------------------------

// Invoke - call a method on the given interface using reflection
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

// Use reflection to unmarshal all the fields in the given interface
// This will loop through every field and call the 'UnmarshalOsuBinary' method
// on the type passing in the 'buf' which is the source of all the bytes.
// A special case is where a field is a slice. In this case:
// 1. We look up the field name like "Num<SliceFieldName>" and retrieve the
//    number of elements to exepect from the stream
// 2. Create a new slice
// 3. Iterate through each slice element and run the UnmarshalOsuBinary method
// Args:
//   db: The object to unmarshal
///  buf: The buffer in which we retrieve bytes to unmarshal
func UnmarshalAny(db interface{}, buf io.Reader, version Int) error {
	dbVal := reflect.ValueOf(db).Elem()
	dbType := reflect.TypeOf(db).Elem()

	for i := 0; i < dbVal.NumField(); i++ {
		currentField := dbVal.Field(i)
		currentFieldType := dbType.Field(i)

		// handles tags
		if gotVersion, ok := currentFieldType.Tag.Lookup("osu-end"); ok {
			intVersion, err := strconv.ParseUint(gotVersion, 10, 32)
			if err != nil {
				return err
			}
			if version > Int(intVersion) {
				continue
			}
		}
		if gotVersion, ok := currentFieldType.Tag.Lookup("osu-start"); ok {
			intVersion, err := strconv.ParseUint(gotVersion, 10, 32)
			if err != nil {
				return err
			}
			if version < Int(intVersion) {
				continue
			}
		}

		switch dbVal.Field(i).Type().Kind() {
		case reflect.Slice:
			numFieldName := "Num" + dbVal.Type().Field(i).Name
			numElements := dbVal.FieldByName(numFieldName).Interface()
			// TODO(jordanyu): This is hack, we can't always assume it will be an Int
			// when we have a slice of element afterwards
			intNumElements := int(numElements.(Int))

			// Create a new slice of appropriate size and then run the unmarshal
			// function over each element in the slice.
			sliceElems := reflect.MakeSlice(currentField.Type(),
				int(intNumElements), int(intNumElements))
			currentField.Set(sliceElems)
			for j := 0; j < intNumElements; j++ {
				ret := Invoke(currentField.Index(j).Addr().Interface(),
					"UnmarshalOsuBinary", buf, version)

				err := ret[0].Interface()
				if err != nil {
					return err.(error)
				}
			}
		default:
			ret := Invoke(
				currentField.Addr().Interface(), "UnmarshalOsuBinary", buf, version)

			err := ret[0].Interface()
			if err != nil {
				return err.(error)
			}
		}
	}

	return nil
}

// Use reflection to marshal all the fields in the given interface
// This will loop through every field and call the 'MarshalOsuBinary' method
// on the type passing in the 'buf' which is the source of all the bytes.
// A special case is where a field is a slice. In this case:
// 1. We look up the field name like "Num<SliceFieldName>" and retrieve the
//    number of elements to exepect from the stream
// 2. Create a new slice
// 3. Iterate through each slice element and run the MarshalOsuBinary method
// Args:
//   db: The object to unmarshal
///  buf: The buffer in which to write the marshalled bytes
func MarshalAny(db interface{}, buf io.Writer, version Int) error {
	dbVal := reflect.ValueOf(db).Elem()
	dbType := reflect.TypeOf(db).Elem()

	for i := 0; i < dbVal.NumField(); i++ {
		currentField := reflect.ValueOf(db).Elem().Field(i)
		currentFieldType := dbType.Field(i)

		// handles tags
		if gotVersion, ok := currentFieldType.Tag.Lookup("osu-end"); ok {
			intVersion, err := strconv.ParseUint(gotVersion, 10, 32)
			if err != nil {
				return err
			}
			if version >= Int(intVersion) {
				continue
			}
		}
		if gotVersion, ok := currentFieldType.Tag.Lookup("osu-start"); ok {
			intVersion, err := strconv.ParseUint(gotVersion, 10, 32)
			if err != nil {
				return err
			}
			if version < Int(intVersion) {
				continue
			}
		}

		switch dbVal.Field(i).Type().Kind() {
		case reflect.Slice:
			numFieldName := "Num" + dbVal.Type().Field(i).Name
			numElements := dbVal.FieldByName(numFieldName).Interface()
			// TODO(jordanyu): This is hack, we can't always assume it will be an Int
			// when we have a slice of element afterwards
			intNumElements := int(numElements.(Int))

			// iterate through each element of the slice and marshal the struct
			for j := 0; j < intNumElements; j++ {
				ret := Invoke(currentField.Index(j).Addr().Interface(),
					"MarshalOsuBinary", buf, version)

				err := ret[0].Interface()
				if err != nil {
					return err.(error)
				}
			}
		default:
			// This is just a primitive field so just run the simple marshal
			ret := Invoke(
				currentField.Addr().Interface(), "MarshalOsuBinary", buf, version)

			err := ret[0].Interface()
			if err != nil {
				return err.(error)
			}
		}
	}

	return nil
}
