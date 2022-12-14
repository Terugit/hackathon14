package usecase

import (
	"back/dao"
	"back/model"
)

func UserEdit(user model.User) error {
	id := user.Id
	name := user.Name

	if id == "" {
		return model.EmptyId
	}
	if name == "" {
		return model.EmptyName
	}

	if err := dao.UserEdit(id, name); err != nil {
		return err
	}

	return nil
}
