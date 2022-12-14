package usecase

import (
	"back/dao"
	"back/model"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

func genID() (string, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id, err := ulid.New(ulid.Timestamp(t), entropy)
	return id.String(), err
}

func UserAdd(user model.User) error {
	name := user.Name

	if name == "" {
		return model.EmptyName
	}
	id, err := genID()
	if err != nil {
		log.Printf("fail: ulid.New(), %v\n", err)
		return model.FailGenID
	}

	if err := dao.UserAdd(id, name); err != nil {
		return err
	}

	return nil
}
