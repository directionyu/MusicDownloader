package main

// MusicType 类型
type MusicType struct {
	Bitrate     int     `json:"bitrate"`
	DfsID       int64   `json:"dfsId"`
	DfsIDStr    string  `json:"dfsId_str"`
	Extension   string  `json:"extension"`
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	PlayTime    int     `json:"playTime"`
	Size        int64   `json:"size"`
	Sr          int     `json:"sr"`
	VolumeDelta float64 `json:"volumeDelta"`
}