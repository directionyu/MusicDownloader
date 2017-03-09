package main

// ProgramInfo 类型
type ProgramInfo struct {
	BdAuditStatus   int        `json:"bdAuditStatus"`
	BlurCoverURL    string     `json:"blurCoverUrl"`
	Buyed           bool       `json:"buyed"`
	Channels        []string   `json:"channels"`
	CommentCount    int        `json:"commentCount"`
	CommentThreadID string     `json:"commentThreadId"`
	CoverURL        string     `json:"coverUrl"`
	CreateTime      int64      `json:"createTime"`
	Description     string     `json:"description"`
	DJ              UserInfo   `json:"dj"`
	Duration        int        `json:"duration"`
	H5Links         string     `json:"h5Links"`
	ID              int        `json:"id"`
	IsPublish       bool       `json:"isPublish"`
	LikedCount      int        `json:"likedCount"`
	ListenerCount   int        `json:"listenerCount"`
	MainSong        SongInfo   `json:"mainSong"`
	MainTrackID     int        `json:"mainTrackId"`
	Name            string     `json:"name"`
	ProgramDesc     string     `json:"programDesc"`
	ProgramFeeType  int        `json:"programFeeType"`
	PubStatus       int        `json:"pubStatus"`
	Radio           RadioInfo  `json:"radio"`
	Reward          bool       `json:"reward"`
	SerialNum       int        `json:"serialNum"`
	ShareCount      int        `json:"shareCount"`
	Songs           []SongInfo `json:"songs"`
	Subscribed      bool       `json:"subscribed"`
	SubscribedCount int        `json:"subscribedCount"`
	TitbitImages    string     `json:"titbitImages"`
	Titbits         string     `json:"titbits"`
	TrackCount      int        `json:"trackCount"`
}