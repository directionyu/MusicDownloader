package main

// Album 专辑类型
type Album struct {
	Code  int       `json:"code"`
	Album AlbumInfo `json:"album"`
}
