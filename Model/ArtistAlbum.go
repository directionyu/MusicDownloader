package main

// ArtistAlbum 歌手专辑类型
type ArtistAlbum struct {
	Code      int         `json:"code"`
	Artist    ArtistInfo  `json:"artist"`
	HotAlbums []AlbumInfo `json:"hotAlbums"`
	More      bool        `json:"more"`
}