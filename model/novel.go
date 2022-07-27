package model

type Novel struct {
	Name   string `json:"Name"`
	Author string `json:"Author"`
	Genre  string `json:"Genre"`
	Pages  int    `json:"Pages"`
}
