package controller

import (
	"back/model"
	"back/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func ThankEdit(w http.ResponseWriter, r *http.Request) {
	var Thank model.Thank
	if err := json.NewDecoder(r.Body).Decode(&Thank); err != nil {
		log.Printf("fail: json.NewDecoder(), %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//log.Print(ThankEdit)
	err := usecase.ThankEdit(Thank)
	if err != nil {
		switch err {
		case model.IncorrectRegisterName:
			{
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		case model.FailIdGenerate:
			{
				w.WriteHeader(http.StatusInternalServerError)
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
	w.WriteHeader(http.StatusOK)
}
