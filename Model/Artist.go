package main

// Artist 歌手类型
type Artist struct {
	Code     int        `json:"code"`
	Artist   ArtistInfo `json:"artist"`
	HotSongs []SongInfo `json:"hotSongs"`
	More     bool       `json:"more"`
}