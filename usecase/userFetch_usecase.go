package usecase

import (
	"back/dao"
	"back/model"
	"log"
)

func FetchUser(id string) (Users []model.User, error error) {
	user, err := dao.FetchUser(id)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}

	return user, nil
}
