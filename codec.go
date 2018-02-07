package gosu

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"reflect"

	"github.com/bnch/uleb128"
)

type BinaryOsuCodec interface {
	BinaryOsuMarshaler
	BinaryOsuUnmarshaler
}

// Unmarshaler interface for Osu Binary blobs
type BinaryOsuUnmarshaler interface {
	UnmarshalBinary(buf io.Reader) error
}

// Marshaler interface for Osu Binary blobs
type BinaryOsuMarshaler interface {
	MarshalBinary(buf io.Writer) error
}

// Unmarshal functions
// -----------------------------------------------------------------------------

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
		err := this.Len.UnmarshalBinary(buf)
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

func (this *DateTime) UnmarshalBinary(buf io.Reader) error {
	return binary.Read(buf, binary.LittleEndian, &this.Value)
}

func (this *OsuDb) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *IntDoublePair) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *TimingPoint) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *BeatMap) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *CollectionDb) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *CollectionDbElement) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *ScoresDb) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *ScoresDbBeatMap) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *ScoresDbBeatMapScore) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *PresenceDb) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

func (this *PlayerPresence) UnmarshalBinary(buf io.Reader) error {
	return UnmarshalAny(this, buf)
}

// Marshal functions
// -----------------------------------------------------------------------------
func (this *Byte) MarshalBinary(buf io.Writer) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Short) MarshalBinary(buf io.Writer) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Int) MarshalBinary(buf io.Writer) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Long) MarshalBinary(buf io.Writer) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *ULEB128) MarshalBinary(buf io.Writer) error {
	got := uleb128.Marshal(int(uint64(*this)))
	buf.Write(got)
	return nil
}

func (this *Single) MarshalBinary(buf io.Writer) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Double) MarshalBinary(buf io.Writer) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Boolean) MarshalBinary(buf io.Writer) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *String) MarshalBinary(buf io.Writer) error {
	err := binary.Write(buf, binary.LittleEndian, &this.Cond)
	if err != nil {
		return err
	}

	if this.Cond == 0xb {
		// TODO(jordanyu): Add error handling logic
		err := this.Len.MarshalBinary(buf)
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

func (this *DateTime) MarshalBinary(buf io.Writer) error {
	return binary.Write(buf, binary.LittleEndian, &this.Value)
}

func (this *OsuDb) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *IntDoublePair) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *TimingPoint) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *BeatMap) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *CollectionDb) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *CollectionDbElement) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *ScoresDb) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *ScoresDbBeatMap) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *ScoresDbBeatMapScore) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *PresenceDb) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
}

func (this *PlayerPresence) MarshalBinary(buf io.Writer) error {
	return MarshalAny(this, buf)
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
// This will loop through every field and call the 'UnmarshalBinary' method
// on the type passing in the 'buf' which is the source of all the bytes.
// A special case is where a field is a slice. In this case:
// 1. We look up the field name like "Num<SliceFieldName>" and retrieve the
//    number of elements to exepect from the stream
// 2. Create a new slice
// 3. Iterate through each slice element and run the UnmarshalBinary method
// Args:
//   db: The object to unmarshal
///  buf: The buffer in which we retrieve bytes to unmarshal
func UnmarshalAny(db interface{}, buf io.Reader) error {
	dbVal := reflect.ValueOf(db).Elem()

	for i := 0; i < dbVal.NumField(); i++ {
		currentMutableField := reflect.ValueOf(db).Elem().Field(i)

		switch dbVal.Field(i).Type().Kind() {
		case reflect.Slice:
			numFieldName := "Num" + dbVal.Type().Field(i).Name
			numElements := dbVal.FieldByName(numFieldName).Interface()
			// TODO(jordanyu): This is hack, we can't always assume it will be an Int
			// when we have a slice of element afterwards
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
		default:
			// This is just a primitive field so just run the simple unmarhsal
			ret := Invoke(
				currentMutableField.Addr().Interface(), "UnmarshalBinary", buf)

			err := ret[0].Interface()
			if err != nil {
				return err.(error)
			}
		}
	}

	return nil
}

// Use reflection to marshal all the fields in the given interface
// This will loop through every field and call the 'MarshalBinary' method
// on the type passing in the 'buf' which is the source of all the bytes.
// A special case is where a field is a slice. In this case:
// 1. We look up the field name like "Num<SliceFieldName>" and retrieve the
//    number of elements to exepect from the stream
// 2. Create a new slice
// 3. Iterate through each slice element and run the MarshalBinary method
// Args:
//   db: The object to unmarshal
///  buf: The buffer in which to write the marshalled bytes
func MarshalAny(db interface{}, buf io.Writer) error {
	dbVal := reflect.ValueOf(db).Elem()

	for i := 0; i < dbVal.NumField(); i++ {
		currentMutableField := reflect.ValueOf(db).Elem().Field(i)

		switch dbVal.Field(i).Type().Kind() {
		case reflect.Slice:
			numFieldName := "Num" + dbVal.Type().Field(i).Name
			numElements := dbVal.FieldByName(numFieldName).Interface()
			// TODO(jordanyu): This is hack, we can't always assume it will be an Int
			// when we have a slice of element afterwards
			intNumElements := int(numElements.(Int))

			// iterate through each element of the slice and marshal the struct
			for j := 0; j < intNumElements; j++ {
				ret := Invoke(currentMutableField.Index(j).Addr().Interface(),
					"MarshalBinary", buf)

				err := ret[0].Interface()
				if err != nil {
					return err.(error)
				}
			}
		default:
			// This is just a primitive field so just run the simple marshal
			ret := Invoke(
				currentMutableField.Addr().Interface(), "MarshalBinary", buf)

			err := ret[0].Interface()
			if err != nil {
				return err.(error)
			}
		}
	}

	return nil
}