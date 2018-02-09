// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
	"time"

	"github.com/Stymphalian/gosu"
)

type Thing struct {
	Timestamp time.Time
	Codecs    []string
	Commons   []string
}

func main() {
	f, err := os.Create("codec.auto.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	packageName := "gosu"
	thing := Thing{
		time.Now(),
		[]string{
			strings.TrimLeft(reflect.TypeOf(gosu.OsuDb{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.IntDoublePair{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.TimingPoint{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.BeatMap{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.CollectionDb{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.CollectionDbElement{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.ScoresDb{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.ScoresDbBeatMap{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.ScoresDbBeatMapScore{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.PresenceDb{}).Name(), packageName),
			strings.TrimLeft(reflect.TypeOf(gosu.PlayerPresence{}).Name(), packageName),
		},
		[]string{
			"Byte",
			"Short",
			"Int",
			"Long",
			"Single",
			"Double",
			"Boolean",
			"DateTime",
			// ULEB128
			// String
		},
	}
	packageTemplate.Execute(f, thing)
}

var packageTemplate = template.Must(template.New("").Parse(`
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at {{ .Timestamp }}

package gosu

import (
	"encoding/binary"
	"io"
)

{{- range .Commons }}

func (this *{{.}}) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return binary.Read(buf, binary.LittleEndian, this)
}

func (this *{{.}}) MarshalOsuBinary(buf io.Writer, version Int) error {
	return binary.Write(buf, binary.LittleEndian, this)
}

{{- end }}

{{- range .Codecs }}

func (this *{{.}}) UnmarshalOsuBinary(buf io.Reader, version Int) error {
	return UnmarshalAny(this, buf, version)
}

func (this *{{.}}) MarshalOsuBinary(buf io.Writer, version Int) error {
	return MarshalAny(this, buf, version)
}

{{- end }}
`))
