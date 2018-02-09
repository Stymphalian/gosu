package gosu

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/d4l3k/messagediff"
)

// TODO(jordanyu): Actually add assertions for these tests
func TestEmpty(t *testing.T) {
}

func _runDbTest(DataFilepath string, Db interface{}, FinalDb interface{}, t *testing.T) {
	reader := OsuBinaryReader{}
	writer := OsuBinaryWriter{}

	fmt.Println("Reading test case:", DataFilepath)

	// Unmarshal a binary osu file
	file, err := os.Open(DataFilepath)
	if err != nil {
		t.Errorf("Failed to open file %s", DataFilepath)
	}
	defer file.Close()
	// err = Db.UnmarshalOsuBinary(file)
	err = reader.Read(Db, file)
	if err != nil {
		t.Error("Failed to unmarshal binary")
	}

	// Marshal the file to a temp file
	tmpfile, err := ioutil.TempFile("", "test-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	// err = Db.MarshalOsuBinary(tmpfile)
	err = writer.Write(Db, tmpfile)
	if err != nil {
		t.Error("Failed to marshal bianry")
	}

	// Unmarshal the file we just wrote. There should be no diff.
	tmpfile.Seek(0, 0)
	// err = FinalDb.(BinaryOsuUnmarshaler).UnmarshalOsuBinary(tmpfile)
	err = reader.Read(FinalDb, tmpfile)
	if err != nil {
		t.Error("Failed to unmarshal final binary", err)
	}

	diff, equal := messagediff.PrettyDiff(Db, FinalDb)
	if !equal {
		t.Errorf("Marshal/Unmarshal failed.\n%s", diff)
	}
}
func TestMarshalUnmarshal(t *testing.T) {
	_runDbTest("data/scores.db", new(ScoresDb), new(ScoresDb), t)
	_runDbTest("data/collection.db", new(CollectionDb), new(CollectionDb), t)
	_runDbTest("data/presence.db", new(PresenceDb), new(PresenceDb), t)
	_runDbTest("data/osu!.db", new(OsuDb), new(OsuDb), t)
}
