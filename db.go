package gosu

import (
	"io"
)

// The DB binary file formats supported by Osu!
// See https://github.com/ppy/osu-wiki/blob/master/wiki/osu!_File_Formats/Db_(file_format)/en.md
// for the spec of these fields.

type OsuDb struct {
	Version         Int
	FolderCount     Int
	AccountUnlocked Boolean
	Datetime        DateTime
	PlayerName      String
	NumBeatmaps     Int
	Beatmaps        []BeatMap
	Extra           Int
}

type IntDoublePair struct {
	ExtraBeforeInt    Byte
	IntValue          Int
	ExtraBeforeDouble Byte
	DoubleValue       Double
}

type TimingPoint struct {
	BPM         Double
	OffsetMsec  Double
	IsInherited Boolean
}

type BeatMap struct {
	SizeOfBeatmapBytes       Int
	ArtistName               String
	ArtistNameUnicode        String
	SongTitle                String
	SongTitleUnicode         String
	CreatorName              String
	Difficulty               String
	AudioFileName            String
	Md5                      String
	OsuFileName              String
	RankedStatus             Byte
	NumHitCircles            Short
	NumOfSliders             Short
	NumOfSpinners            Short
	LastModTimeTicks         Long
	ApproachRate             Single
	CircleSize               Single
	HPDrainRate              Single
	OverallDifficulty        Single
	SliderVelocity           Double
	NumOsuStandardStarRating Int
	OsuStandardStarRating    []IntDoublePair
	NumTaikoStarRating       Int
	TaikoStarRating          []IntDoublePair
	NumCTBStarRating         Int
	CTBStarRating            []IntDoublePair
	NumManiaStarRating       Int
	ManiaStarRating          []IntDoublePair
	DrainTimeSecs            Int
	TotalTimeMsec            Int
	AudioPreviewMsec         Int
	NumTimingPoints          Int
	TimingPoints             []TimingPoint
	BeatmapID                Int
	BeatmapSetID             Int
	ThreadID                 Int
	GradeOsuStandard         Byte
	GradeTaiko               Byte
	GradeCTB                 Byte
	GradeMania               Byte
	LocalBeatmapOffset       Short
	StackLeniency            Single
	OsuGameplayMode          Byte
	SongSource               String
	SongTags                 String
	OnlineOffset             Short
	TitleFont                String
	IsPlayed                 Boolean
	LastTimePlayed           Long
	IsOsz2Format             Boolean
	RelativeFolderName       String
	LastTimeCheckedWithRepo  Long
	IgnoreBeatmapSound       Boolean
	IgnoreBeatmapSkin        Boolean
	DisableStoryboard        Boolean
	DisableVideo             Boolean
	VisualOverride           Boolean
	LastModificationTime     Int
	ManiaScrollSpeed         Byte
}

type CollectionDb struct {
	Version        Int
	NumCollections Int
	Collections    []CollectionDbElement
}

type CollectionDbElement struct {
	Name                String
	NumBeatmapMd5Hashes Int
	BeatmapMd5Hashes    []String
}

type ScoresDb struct {
	Version     Int
	NumBeatmaps Int
	Beatmaps    []ScoresDbBeatMap
}

type ScoresDbBeatMap struct {
	Md5Hash   String
	NumScores Int
	Scores    []ScoresDbBeatMapScore
}

type ScoresDbBeatMapScore struct {
	GameplayMode                 Byte
	Version                      Int
	Md5Hash                      String
	PlayerName                   String
	ReplayMd5Hash                String
	Num300                       Short
	Num200                       Short
	Num50                        Short
	NumMax300                    Short
	Num100                       Short
	NumMiss                      Short
	ReplayScore                  Int
	MaxCombo                     Short
	IsPerfectCombo               Boolean
	Mods                         Int
	EmptyString                  String
	TimestampOfReplayWindowTicks Long
	AlwaysNegativeOne            Int
	OnlineScoreId                Long
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
