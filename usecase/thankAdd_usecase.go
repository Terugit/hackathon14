package usecase

import (
	"back/dao"
	"back/model"
	"back/ulid"
	"log"
)

func ThankAdd(Thank model.Thank) error {
	FromId := Thank.From_
	ToId := Thank.To_
	Point := Thank.Point
	Message := Thank.Message
	PostAt := Thank.PostAt
	EditAt := Thank.EditAt
	if Point < 0 {
		log.Printf("Point is not correct")
		return model.IncorrectPoint
	}
	if Message == "" {
		log.Printf("Message is empty")
		return model.EmptyMessage
	}
	ThankId, err := ulid.IdGenerate()
	if err != nil {
		log.Printf("fail: ulid.New(), %v\n", err)
		return model.FailIdGenerate
	}

	if err = dao.ThankAdd(ThankId, FromId, ToId, Point, Message, PostAt, EditAt); err != nil {
		return err
	}

	return nil
}
