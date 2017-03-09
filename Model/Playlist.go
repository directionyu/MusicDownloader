package main

// Playlist 歌单类型
type Playlist struct {
	Code   int        `json:"code"`
	Result Collection `json:"result"`
}
