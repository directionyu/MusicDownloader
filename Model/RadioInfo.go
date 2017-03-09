package main

// RadioInfo 类型
type RadioInfo struct {
	Buyed                 bool    `json:"buyed"`
	Category              string  `json:"category"`
	CategoryID            int     `json:"categoryId"`
	CreateTime            int64   `json:"createTime"`
	Desc                  string  `json:"desc"`
	DJ                    string  `json:"dj"`
	Finished              bool    `json:"finished"`
	ID                    int     `json:"id"`
	LastProgramCreateTime int64   `json:"lastProgramCreateTime"`
	LastProgramID         int     `json:"lastProgramId"`
	LastProgramName       string  `json:"lastProgramName"`
	Name                  string  `json:"name"`
	PicURL                string  `json:"picUrl"`
	Price                 float64 `json:"price"`
	ProgramCount          int     `json:"programCount"`
	PurchaseCount         int     `json:"purchaseCount"`
	RadioFeeType          int     `json:"radioFeeType"`
	Rcmdtext              string  `json:"rcmdtext"`
	SubCount              int     `json:"subCount"`
	UnderShelf            bool    `json:"underShelf"`
	Videos                string  `json:"videos"`
}
