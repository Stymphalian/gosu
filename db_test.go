package gosu

import (
	"os"
	"testing"
)

// TODO(jordanyu): Actually add assertions for these tests

func TestUnserialiazeOsuDb(t *testing.T) {
	path := "data/osu!.db"
	file, err := os.Open(path)
	if err != nil {
		t.Error("Failed to open file")
	}
	defer file.Close()

	var db OsuDb
	err = db.UnmarshalBinary(file)
	if err != nil {
		t.Error("Failed to unmarshal binary")
	}
}

func TestUnserialiazeCollectionDb(t *testing.T) {
	path := "data/collection.db"
	file, err := os.Open(path)
	if err != nil {
		t.Error("Failed to open file")
	}
	defer file.Close()

	var db CollectionDb
	err = db.UnmarshalBinary(file)
	if err != nil {
		t.Error("Failed to unmarshal binary")
	}
}

func TestUnserialiazeScoreDb(t *testing.T) {
	path := "data/scores.db"
	file, err := os.Open(path)
	if err != nil {
		t.Error("Failed to open file")
	}
	defer file.Close()

	var db ScoresDb
	err = db.UnmarshalBinary(file)
	if err != nil {
		t.Error("Failed to unmarshal binary")
	}
}
