package main

// AlbumInfo 类型
type AlbumInfo struct {
	Alias           []string     `json:"alias"`
	Artist          ArtistInfo   `json:"artist"`
	Artists         []ArtistInfo `json:"artists"`
	BlurPicURL      string       `json:"blurPicUrl"`
	BriefDesc       string       `json:"briefDesc"`
	CommentThreadID string       `json:"commentThreadId"`
	Company         string       `json:"company"`
	CompanyID       int          `json:"companyId"`
	CopyrightID     int          `json:"copyrightId"`
	Description     string       `json:"description"`
	ID              int          `json:"id"`
	Name            string       `json:"name"`
	OnSale          bool         `json:"onSale"`
	Paid            bool         `json:"paid"`
	Pic             int          `json:"pic"`
	PicID           int          `json:"picId"`
	PicIDStr        string       `json:"picId_str"`
	PicURL          string       `json:"picUrl"`
	PublishTime     int64        `json:"publishTime"`
	Size            int          `json:"size"`
	Songs           []SongInfo   `json:"songs"`
	Status          int          `json:"status"`
	SubType         string       `json:"subType"`
	Tags            string       `json:"tags"`
	Atype           string       `json:"type"`
}
