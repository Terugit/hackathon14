package usecase

import (
	"back/dao"
	"back/model"
	"log"
)

func ThankEdit(Thank model.Thank) error {
	id := Thank.Id
	to_ := Thank.To_
	point := Thank.Point
	message := Thank.Message
	editAt := Thank.EditAt
	deletePoint := Thank.DeletePoint
	if point < 0 {
		log.Printf("Point is not correct")
		return model.IncorrectRegisterName
	}
	if message == "" {
		log.Printf("Type message")
		return model.EmptyMessage
	}

	if err := dao.ThankEdit(id, to_, point, deletePoint, message, editAt); err != nil {
		return err
	}

	return nil
}
