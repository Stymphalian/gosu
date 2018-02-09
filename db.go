package gosu

// The DB binary file formats supported by Osu!
// See https://github.com/ppy/osu-wiki/blob/master/wiki/osu!_File_Formats/Db_(file_format)/en.md
// for the spec of these fields.

// Valid osu struct tags:
// `osu-version:""` - Denotes that thie field is the version of the  binary file
// `osu-end:"YYYYMMDD"` - Tells to skip this field if the current osu-version
//    is greater than this version string. This is because this field has been
//    removed.
// `osu-start:"YYYYMMDD"` - Tells to skip this field if the current osu-version
//    is less than this version string. This is because this field has not yet
//    been populated yet.
//

type OsuDb struct {
	Version         uint32 `osu-version:""`
	FolderCount     uint32
	AccountUnlocked bool
	Datetime        uint64
	PlayerName      string
	Beatmaps        []BeatMap
	_               uint32
}

type IntDoublePair struct {
	_           byte
	IntValue    uint32
	_           byte
	DoubleValue float64
}

type TimingPoint struct {
	BPM         float64
	OffsetMsec  float64
	IsInherited bool
}

type BeatMap struct {
	SizeOfBeatmapBytes      uint32
	ArtistName              string
	ArtistNameUnicode       string
	SongTitle               string
	SongTitleUnicode        string
	CreatorName             string
	Difficulty              string
	AudioFileName           string
	Md5                     string
	OsuFileName             string
	RankedStatus            byte
	NumHitCircles           uint16
	NumOfSliders            uint16
	NumOfSpinners           uint16
	LastModTimeTicks        uint64
	ApproachRateOld         byte    `osu-end:"20140609"`
	CircleSizeOld           byte    `osu-end:"20140609"`
	HPDrainRateOld          byte    `osu-end:"20140609"`
	OverallDifficultyOld    byte    `osu-end:"20140609"`
	ApproachRate            float32 `osu-start:"20140609"`
	CircleSize              float32 `osu-start:"20140609"`
	HPDrainRate             float32 `osu-start:"20140609"`
	OverallDifficulty       float32 `osu-start:"20140609"`
	SliderVelocity          float64
	OsuStandardStarRating   []IntDoublePair `osu-start:"20140609"`
	TaikoStarRating         []IntDoublePair `osu-start:"20140609"`
	CTBStarRating           []IntDoublePair `osu-start:"20140609"`
	ManiaStarRating         []IntDoublePair `osu-start:"20140609"`
	DrainTimeSecs           uint32
	TotalTimeMsec           uint32
	AudioPreviewMsec        uint32
	TimingPoints            []TimingPoint
	BeatmapID               uint32
	BeatmapSetID            uint32
	ThreadID                uint32
	GradeOsuStandard        byte
	GradeTaiko              byte
	GradeCTB                byte
	GradeMania              byte
	LocalBeatmapOffset      uint16
	StackLeniency           float32
	OsuGameplayMode         byte
	SongSource              string
	SongTags                string
	OnlineOffset            uint16
	TitleFont               string
	IsPlayed                bool
	LastTimePlayed          uint64
	IsOsz2Format            bool
	RelativeFolderName      string
	LastTimeCheckedWithRepo uint64
	IgnoreBeatmapSound      bool
	IgnoreBeatmapSkin       bool
	DisableStoryboard       bool
	DisableVideo            bool
	VisualOverride          bool
	UnknownShort            uint16 `osu-end:"20140609"`
	LastModificationTime    uint32
	ManiaScrollSpeed        byte
}

type CollectionDb struct {
	Version     uint32 `osu-version:""`
	Collections []CollectionDbElement
}

type CollectionDbElement struct {
	Name             string
	BeatmapMd5Hashes []string
}

type ScoresDb struct {
	Version  uint32 `osu-version:""`
	Beatmaps []ScoresDbBeatMap
}

type ScoresDbBeatMap struct {
	Md5Hash string
	Scores  []ScoresDbBeatMapScore
}

type ScoresDbBeatMapScore struct {
	GameplayMode                 byte
	Version                      uint32
	Md5Hash                      string
	PlayerName                   string
	ReplayMd5Hash                string
	Num300                       uint16
	Num200                       uint16
	Num50                        uint16
	NumMax300                    uint16
	Num100                       uint16
	NumMiss                      uint16
	ReplayScore                  uint32
	MaxCombo                     uint16
	IsPerfectCombo               bool
	Mods                         uint32
	_                            string
	TimestampOfReplayWindowTicks uint64
	AlwaysNegativeOne            uint32
	OnlineScoreId                uint64
}

type PresenceDb struct {
	Version uint32 `osu-version:""`
	Players []PlayerPresence
}

type PlayerPresence struct {
	PlayerId         uint32
	PlayerName       string
	UtcOffset        byte
	Country          byte
	UnknownByteField byte
	Longitude        float32
	Latitude         float32
	GlobalRank       uint32
	DateModified     uint64
}
