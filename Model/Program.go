package main

// Program 单条主播类型
type Program struct {
	Code    int         `json:"code"`
	Program ProgramInfo `json:"program"`
}