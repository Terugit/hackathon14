package usecase

import (
	"back/dao"
	"back/model"
	"back/ulid"
	"log"
)

func UserAdd(user model.User) error {
	name := user.Name

	if name == "" {
		return model.EmptyName
	}
	id, err := ulid.IdGenerate()
	if err != nil {
		log.Printf("fail: ulid.New(), %v\n", err)
		return model.FailIdGenerate
	}

	if err := dao.UserAdd(id, name); err != nil {
		return err
	}

	return nil
}
