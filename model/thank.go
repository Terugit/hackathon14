package model

type Thank struct {
	Id          string `json:"id"`
	From_       string `json:"from_"`
	FromWho     string `json:"fromWho"`
	To_         string `json:"to_"`
	ToWho       string `json:"toWho"`
	Point       int    `json:"point"`
	Message     string `json:"message"`
	PostAt      string `json:"postAt"`
	EditAt      string `json:"editAt"`
	DeletePoint int    `json:"delete_point"`
}

//type ContDetail struct {
//	Id         string `json:"id"`
//	FromName   string `json:"from_name"`
//	ToName     string `json:"to_name"`
//	Point      int    `json:"point"`
//	Message    string `json:"message"`
//	PostTime   string `json:"post_time"`
//	UpdateTime string `json:"update_time"`
//}
//
//type ContFromDetail struct {
//	Id         string `json:"id"`
//	FromName   string `json:"from_name"`
//	ToId       string `json:"to_id"`
//	ToName     string `json:"to_name"`
//	Point      int    `json:"point"`
//	Message    string `json:"message"`
//	PostTime   string `json:"post_time"`
//	UpdateTime string `json:"update_time"`
//}

//// フロントから送られる情報
//type Contribute struct {
//	FromId     string `json:"from_id"`
//	ToId       string `json:"to_id"`
//	Point      int    `json:"point"`
//	Message    string `json:"message"`
//	PostTime   string `json:"post_time"`
//	UpdateTime string `json:"update_time"`
//}
//
//type ContEdit struct {
//	Id          string `json:"id"`
//	ToId        string `json:"to_id"`
//	Point       int    `json:"point"`
//	DeletePoint int    `json:"delete_point"`
//	Message     string `json:"message"`
//	UpdateTime  string `json:"update_time"`
//}
//
//type ContDelete struct {
//	Id    string `json:"id"`
//	ToId  string `json:"to_id"`
//	Point int    `json:"point"`
//}
