
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2018-02-09 01:28:03.893322923 -0800 PST m=+0.000517803

package gosu

import (
	"encoding/binary"
	"io"
)

func (this *Byte) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *Byte) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Short) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *Short) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Int) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *Int) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Long) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *Long) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Single) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *Single) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Double) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *Double) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *Boolean) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *Boolean) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *DateTime) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *DateTime) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

func (this *OsuDb) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *OsuDb) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *IntDoublePair) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *IntDoublePair) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *TimingPoint) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *TimingPoint) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *BeatMap) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *BeatMap) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *CollectionDb) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *CollectionDb) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *CollectionDbElement) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *CollectionDbElement) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *ScoresDb) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *ScoresDb) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *ScoresDbBeatMap) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *ScoresDbBeatMap) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *ScoresDbBeatMapScore) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *ScoresDbBeatMapScore) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *PresenceDb) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *PresenceDb) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

func (this *PlayerPresence) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *PlayerPresence) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}