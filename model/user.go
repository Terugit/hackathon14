package model

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Point int    `json:"point"`
}

type UserAdd struct {
	Name string `json:"name"`
}

type userId struct {
	Id string `json:"id"`
}

type UserPoint struct {
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
