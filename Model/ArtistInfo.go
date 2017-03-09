package main

// ArtistInfo 类型
type ArtistInfo struct {
	AlbumSize   int      `json:"albumSize"`
	Alias       []string `json:"alias"`
	BriefDesc   string   `json:"briefDesc"`
	ID          int      `json:"id"`
	Img1v1ID    int      `json:"img1v1Id"`
	Img1v1IDStr string   `json:"img1v1Id_str"`
	Img1v1URL   string   `json:"img1v1Url"`
	MusicSize   int      `json:"musicSize"`
	Name        string   `json:"name"`
	PicID       int      `json:"picId"`
	PicURL      string   `json:"picUrl"`
	TopicPerson int      `json:"topicPerson"`
	Trans       string   `json:"trans"`
}
