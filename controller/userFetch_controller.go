package controller

import (
	"back/dao"
	"back/model"
	"back/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func UserFetch(w http.ResponseWriter, r *http.Request) {

	users, err := usecase.FetchUser("all")
	if err != nil {
		switch err {
		case model.EmptySearchName:
			{
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		default:
			{
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func Ranking(w http.ResponseWriter, r *http.Request) {
	users, err := dao.UserRank()
	if err != nil {
		switch err {
		case model.EmptySearchName:
			{
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		default:
			{
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
