package main

// Collection 类型
type Collection struct {
	AdType                int          `json:"adType"`
	Artists               []ArtistInfo `json:"artists"`
	CloudTrackCount       int          `json:"cloudTrackCount"`
	CommentCount          int          `json:"commentCount"`
	CommentThreadID       string       `json:"commentThreadId"`
	CoverImgID            int          `json:"coverImgId"`
	CoverImgIDStr         string       `json:"coverImgId_str"`
	CoverImgURL           string       `json:"coverImgUrl"`
	CreateTime            int64        `json:"createTime"`
	Creator               UserInfo     `json:"creator"`
	Description           string       `json:"description"`
	HighQuality           bool         `json:"highQuality"`
	ID                    int          `json:"id"`
	Name                  string       `json:"name"`
	NewImported           bool         `json:"newImported"`
	Ordered               bool         `json:"ordered"`
	PlayCount             int          `json:"playCount"`
	Privacy               int          `json:"privacy"`
	ShareCount            int          `json:"shareCount"`
	SpecialType           int          `json:"specialType"`
	Status                int          `json:"status"`
	Subscribed            bool         `json:"subscribed"`
	SubscribedCount       int          `json:"subscribedCount"`
	Subscribers           []string     `json:"subscribers"`
	Tags                  []string     `json:"tags"`
	TotalDuration         int          `json:"totalDuration"`
	TrackCount            int          `json:"trackCount"`
	TrackNumberUpdateTime int64        `json:"trackNumberUpdateTime"`
	TrackUpdateTime       int64        `json:"trackUpdateTime"`
	Tracks                []SongInfo   `json:"tracks"`
	UpdateTime            int64        `json:"updateTime"`
	UserID                int          `json:"userId"`
}
