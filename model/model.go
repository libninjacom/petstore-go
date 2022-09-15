package model

type Pet struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Pets []Pet

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
