package main

type DJRadio struct {
	Code     int           `json:"code"`
	Count    int           `json:"count"`
	More     bool          `json:"more"`
	Programs []ProgramInfo `json:"programs"`
}
