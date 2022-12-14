package controller

import (
	"back/dao"
	"back/model"
	"encoding/json"
	"log"
	"net/http"
)

func FetchAllThanks(w http.ResponseWriter, r *http.Request) {
	thanks, err := dao.FetchAllThanks()
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
	bytes, err := json.Marshal(thanks)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func FetchThanksSent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	thanks, err := dao.FetchThanksSent(id)
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
	bytes, err := json.Marshal(thanks)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func FetchThanksGot(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	thanks, err := dao.FetchThanksGot(id)
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
	bytes, err := json.Marshal(thanks)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
