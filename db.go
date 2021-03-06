package gosu

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
	ApproachRateByte         Byte `osu-end:"20140609"`
	CircleSizeByte           Byte `osu-end:"20140609"`
	HPDrainRateByte          Byte `osu-end:"20140609"`
	OverallDifficultyByte    Byte `osu-end:"20140609"`
	ApproachRate             Single
	CircleSize               Single
	HPDrainRate              Single
	OverallDifficulty        Single
	SliderVelocity           Double
	NumOsuStandardStarRating Int             `osu-start:"20140609"`
	OsuStandardStarRating    []IntDoublePair `osu-start:"20140609"`
	NumTaikoStarRating       Int             `osu-start:"20140609"`
	TaikoStarRating          []IntDoublePair `osu-start:"20140609"`
	NumCTBStarRating         Int             `osu-start:"20140609"`
	CTBStarRating            []IntDoublePair `osu-start:"20140609"`
	NumManiaStarRating       Int             `osu-start:"20140609"`
	ManiaStarRating          []IntDoublePair `osu-start:"20140609"`
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
	UnknownShortField        Short `osu-end:"20140609"`
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

type PresenceDb struct {
	Version    Int
	NumPlayers Int
	Players    []PlayerPresence
}

type PlayerPresence struct {
	PlayerId         Int
	PlayerName       String
	UtcOffset        Byte
	Country          Byte
	UnknownByteField Byte
	Longitude        Single
	Latitude         Single
	GlobalRank       Int
	DateModified     DateTime
}
