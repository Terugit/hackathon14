package controller

import (
	"back/model"
	"back/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func UserAdd(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("fail: json.NewDecoder(), %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := usecase.UserAdd(user)
	if err != nil {
		switch err {
		case model.EmptyId:
			{
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		case model.EmptyName:
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

	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("fail: w.Write(bytes) in userAdd_controller.go, line 50, %v\n", err)
		return
	}
}
