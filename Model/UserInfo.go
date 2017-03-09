package main

// UserInfo 类型
type UserInfo struct {
	AccountStatus      int      `json:"accountStatus"`
	AuthStatus         int      `json:"authStatus"`
	Authority          int      `json:"authority"`
	AvatarImgID        int      `json:"avatarImgId"`
	AvatarImgIDStr     string   `json:"avatarImgIdStr"`
	AvatarURL          string   `json:"avatarUrl"`
	BackgroundImgID    int      `json:"backgroundImgId"`
	BackgroundImgIDStr string   `json:"backgroundImgIdStr"`
	BackgroundURL      string   `json:"backgroundUrl"`
	Birthday           int64    `json:"birthday"`
	Brand              string   `json:"brand"`
	City               int      `json:"city"`
	DefaultAvatar      bool     `json:"defaultAvatar"`
	Description        string   `json:"description"`
	DetailDescription  string   `json:"detailDescription"`
	DJStatus           int      `json:"djStatus"`
	ExpertTags         []string `json:"expertTags"`
	Followed           bool     `json:"followed"`
	Gender             int      `json:"gender"`
	Mutual             bool     `json:"mutual"`
	Nickname           string   `json:"nickname"`
	Province           int      `json:"province"`
	RemarkName         string   `json:"remarkName"`
	Signature          string   `json:"signature"`
	UserID             int      `json:"userId"`
	UserType           int      `json:"userType"`
	VipType            int      `json:"vipType"`
}