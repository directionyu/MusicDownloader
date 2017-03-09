package main

// Song 歌曲类型
type Song struct {
	Code       int           `json:"code"`
	Equalizers []interface{} `json:"equalizers"`
	Songs      []SongInfo    `json:"songs"`
}
