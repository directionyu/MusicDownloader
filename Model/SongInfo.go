package main

// SongInfo 类型
type SongInfo struct {
	Album           AlbumInfo    `json:"album"`
	Alias           []string     `json:"alias"`
	Artists         []ArtistInfo `json:"artists"`
	Audition        string       `json:"audition"`
	BMusic          MusicType    `json:"bMusic"`
	CommentThreadID string       `json:"commentThreadId"`
	CopyFrom        string       `json:"copyFrom"`
	CopyrightID     int          `json:"copyrightId"`
	Crbt            string       `json:"crbt"`
	DayPlays        int          `json:"dayPlays"`
	Disc            string       `json:"disc"`
	Duration        int          `json:"duration"`
	Fee             int          `json:"fee"`
	Ftype           int          `json:"ftype"`
	HMusic          MusicType    `json:"hMusic"`
	HearTime        int          `json:"hearTime"`
	ID              int          `json:"id"`
	LMusic          MusicType    `json:"lMusic"`
	MMusic          MusicType    `json:"mMusic"`
	Mp3URL          string       `json:"mp3Url"`
	MVID            int          `json:"mvid"`
	Name            string       `json:"name"`
	No              int          `json:"no"`
	PlayedNum       int          `json:"playedNum"`
	Popularity      float64      `json:"popularity"`
	Position        int          `json:"position"`
	Ringtone        string       `json:"ringtone"`
	RtURL           string       `json:"rtUrl"`
	RtURLs          []string     `json:"rtUrls"`
	Rtype           int          `json:"rtype"`
	RURL            string       `json:"rurl"`
	Score           int          `json:"score"`
	Starred         bool         `json:"starred"`
	StarredNum      int          `json:"starredNum"`
	Status          int          `json:"status"`
}
