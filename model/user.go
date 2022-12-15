package model

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Point int    `json:"point"`
}

type UserRank struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Point int    `json:"point"`
	Rank  int    `json:"rank"`
}
