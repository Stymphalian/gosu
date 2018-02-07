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

func TestMarshalUnmarshal(t *testing.T) {
	testcases := []struct {
		Db           BinaryOsuCodec
		FinalDb      BinaryOsuCodec
		DataFilepath string
	}{
		{new(ScoresDb), new(ScoresDb), "data/scores.db"},
		{new(CollectionDb), new(CollectionDb), "data/collection.db"},
		{new(PresenceDb), new(PresenceDb), "data/presence.db"},
		{new(OsuDb), new(OsuDb), "data/osu!.db"},
	}

	for _, testcase := range testcases {
		fmt.Println("Reading test case:", testcase.DataFilepath)

		// Unmarshal a binary osu file
		file, err := os.Open(testcase.DataFilepath)
		if err != nil {
			t.Errorf("Failed to open file %s", testcase.DataFilepath)
		}
		defer file.Close()
		err = testcase.Db.UnmarshalBinary(file)
		if err != nil {
			t.Error("Failed to unmarshal binary")
		}

		// Marshal the file to a temp file
		tmpfile, err := ioutil.TempFile("", "test-")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())
		err = testcase.Db.MarshalBinary(tmpfile)
		if err != nil {
			t.Error("Failed to marshal bianry")
		}

		// Unmarshal the file we just wrote. There should be no diff.
		tmpfile.Seek(0, 0)
		err = testcase.FinalDb.(BinaryOsuUnmarshaler).UnmarshalBinary(tmpfile)
		if err != nil {
			t.Error("Failed to unmarshal final binary", err)
		}

		diff, equal := messagediff.PrettyDiff(testcase.Db, testcase.FinalDb)
		if !equal {
			t.Errorf("Marshal/Unmarshal failed.\n%s", diff)
		}
	}
}
