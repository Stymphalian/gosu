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

type BinaryOsuCodec interface {
	BinaryOsuMarshaler
	BinaryOsuUnmarshaler
}

// Unmarshaler interface for Osu Binary blobs
type BinaryOsuUnmarshaler interface {
	UnmarshalOsuBinary(buf io.Reader) error
}

// Marshaler interface for Osu Binary blobs
type BinaryOsuMarshaler interface {
	MarshalOsuBinary(buf io.Writer) error
}

// Marshal/Unmarshal function for ULEB128 and String
// The other types can be found in auto.codec.go
// -----------------------------------------------------------------------------

func (this *ULEB128) UnmarshalOsuBinary(buf io.Reader) error {
	total := uleb128.UnmarshalReader(buf)
	*this = ULEB128(uint64(total))

	// TODO(jordanyu): Add error handling logic
	return nil
}

func (this *ULEB128) MarshalOsuBinary(buf io.Writer) error {
	got := uleb128.Marshal(int(uint64(*this)))
	buf.Write(got)
	return nil
}

func (this *String) UnmarshalOsuBinary(buf io.Reader) error {
	err := binary.Read(buf, binary.LittleEndian, &this.Cond)
	if err != nil {
		return err
	}

	if this.Cond == 0xb {
		// TODO(jordanyu): Add error handling logic
		err := this.Len.UnmarshalOsuBinary(buf)
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

func (this *String) MarshalOsuBinary(buf io.Writer) error {
	err := binary.Write(buf, binary.LittleEndian, &this.Cond)
	if err != nil {
		return err
	}

	if this.Cond == 0xb {
		// TODO(jordanyu): Add error handling logic
		err := this.Len.MarshalOsuBinary(buf)
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

type OsuBinaryReader struct {
	version uint32
}

func (this *OsuBinaryReader) readSingleField(fieldType reflect.Type,
	mutableField reflect.Value, tags reflect.StructTag, buf io.Reader) error {

	if version, ok := tags.Lookup("osu-end"); ok {
		intVersion, err := strconv.ParseUint(version, 10, 32)
		if err != nil {
			return err
		}
		if this.version > uint32(intVersion) {
			return nil
		}
	}
	if version, ok := tags.Lookup("osu-start"); ok {
		intVersion, err := strconv.ParseUint(version, 10, 32)
		if err != nil {
			return err
		}
		if this.version < uint32(intVersion) {
			return nil
		}
	}

	switch fieldType.Kind() {
	case reflect.Uint8:
		var v uint8
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
		if mutableField.CanSet() {
			mutableField.Set(reflect.ValueOf(v))
		}
	case reflect.Uint16:
		var v uint16
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
		if mutableField.CanSet() {
			mutableField.Set(reflect.ValueOf(v))
		}
	case reflect.Uint32:
		var v uint32
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
		if mutableField.CanSet() {
			mutableField.Set(reflect.ValueOf(v))
		}
	case reflect.Uint64:
		var v uint64
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
		if mutableField.CanSet() {
			mutableField.Set(reflect.ValueOf(v))
		}
	case reflect.Float32:
		var v float32
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
		if mutableField.CanSet() {
			mutableField.Set(reflect.ValueOf(v))
		}
	case reflect.Float64:
		var v float64
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
		if mutableField.CanSet() {
			mutableField.Set(reflect.ValueOf(v))
		}
	case reflect.Bool:
		var v uint8
		if err := binary.Read(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
		if mutableField.CanSet() {
			mutableField.SetBool(v != 0)
		}
	case reflect.String:
		var v String
		if err := v.UnmarshalOsuBinary(buf); err != nil {
			return err
		}
		if v.Cond == 0x0b && v.Len > 0 {
			if mutableField.CanSet() {
				mutableField.SetString(v.Text)
			}
		}
	case reflect.Struct:
		if err := this.Read(mutableField.Interface(), buf); err != nil {
			return err
		}
	case reflect.Slice:
		var numElems uint32
		if err := binary.Read(buf, binary.LittleEndian, &numElems); err != nil {
			return err
		}
		sliceElems := reflect.MakeSlice(mutableField.Type(), int(numElems), int(numElems))
		mutableField.Set(sliceElems)
		for j := 0; j < int(numElems); j++ {
			if err := this.Read(mutableField.Index(j).Addr().Interface(), buf); err != nil {
				return err
			}
		}
	default:
		return errors.New("Unsupported field type")
	}

	if _, ok := tags.Lookup("osu-version"); ok {
		this.version = mutableField.Interface().(uint32)
	}

	return nil
}

func (this *OsuBinaryReader) Read(db interface{}, buf io.Reader) error {
	valVal := reflect.ValueOf(db).Elem()
	typeVal := reflect.TypeOf(db).Elem()

	switch typeVal.Kind() {
	case reflect.Struct:
		for i := 0; i < typeVal.NumField(); i++ {
			fieldType := typeVal.Field(i)
			mutableField := valVal.Field(i)

			if err := this.readSingleField(fieldType.Type, mutableField, fieldType.Tag, buf); err != nil {
				return err
			}
		}
	default:
		// assuming primitive type
		var tag reflect.StructTag
		if err := this.readSingleField(typeVal, valVal, tag, buf); err != nil {
			return err
		}
	}

	return nil
}

type OsuBinaryWriter struct {
	version uint32
}

func (this *OsuBinaryWriter) writeSingleField(fieldType reflect.Type,
	mutableField reflect.Value, tags reflect.StructTag, buf io.Writer) error {

	if _, ok := tags.Lookup("osu-version"); ok {
		this.version = mutableField.Interface().(uint32)
	}

	if version, ok := tags.Lookup("osu-end"); ok {
		intVersion, err := strconv.ParseUint(version, 10, 32)
		if err != nil {
			return err
		}
		if this.version > uint32(intVersion) {
			return nil
		}
	}
	if version, ok := tags.Lookup("osu-start"); ok {
		intVersion, err := strconv.ParseUint(version, 10, 32)
		if err != nil {
			return err
		}
		if this.version < uint32(intVersion) {
			return nil
		}
	}

	switch fieldType.Kind() {
	case reflect.Uint8:
		var v uint8
		if mutableField.CanInterface() {
			v = mutableField.Interface().(uint8)
		}
		if err := binary.Write(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
	case reflect.Uint16:
		var v uint16
		if mutableField.CanInterface() {
			v = mutableField.Interface().(uint16)
		}
		if err := binary.Write(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
	case reflect.Uint32:
		var v uint32
		if mutableField.CanInterface() {
			v = mutableField.Interface().(uint32)
		}
		if err := binary.Write(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
	case reflect.Uint64:
		var v uint64
		if mutableField.CanInterface() {
			v = mutableField.Interface().(uint64)
		}
		if err := binary.Write(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
	case reflect.Float32:
		var v float32
		if mutableField.CanInterface() {
			v = mutableField.Interface().(float32)
		}
		if err := binary.Write(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
	case reflect.Float64:
		var v float64
		if mutableField.CanInterface() {
			v = mutableField.Interface().(float64)
		}
		if err := binary.Write(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
	case reflect.Bool:
		var boolVal bool
		if mutableField.CanInterface() {
			boolVal = mutableField.Interface().(bool)
		}
		var v uint8 = 0
		if boolVal {
			v = 1
		}
		if err := binary.Write(buf, binary.LittleEndian, &v); err != nil {
			return err
		}
	case reflect.String:
		var str string = ""
		if mutableField.CanInterface() {
			str = mutableField.Interface().(string)
		}

		v := String{}
		if len(str) == 0 {
			v.Cond = 0
		} else {
			v.Cond = 0xb
			v.Len = ULEB128(len(str))
			v.Text = str
		}
		if err := v.MarshalOsuBinary(buf); err != nil {
			return err
		}
	case reflect.Struct:
		if err := this.Write(mutableField.Interface(), buf); err != nil {
			return err
		}
	case reflect.Slice:
		var numElems uint32 = uint32(mutableField.Len())
		if err := binary.Write(buf, binary.LittleEndian, &numElems); err != nil {
			return err
		}
		for j := 0; j < int(numElems); j++ {
			if err := this.Write(mutableField.Index(j).Addr().Interface(), buf); err != nil {
				return err
			}
		}
	default:
		return errors.New("Unsupported field type")
	}

	return nil
}

func (this *OsuBinaryWriter) Write(db interface{}, buf io.Writer) error {
	valVal := reflect.ValueOf(db).Elem()
	typeVal := reflect.TypeOf(db).Elem()

	switch typeVal.Kind() {
	case reflect.Struct:
		for i := 0; i < typeVal.NumField(); i++ {
			fieldType := typeVal.Field(i)
			mutableField := valVal.Field(i)

			if err := this.writeSingleField(fieldType.Type, mutableField, fieldType.Tag, buf); err != nil {
				return err
			}
		}
	default:
		// assuming primitive type
		var tag reflect.StructTag
		if err := this.writeSingleField(typeVal, valVal, tag, buf); err != nil {
			return err
		}
	}

	return nil
}
