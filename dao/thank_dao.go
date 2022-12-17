package dao

import (
	"back/model"
	"log"
)

func ThankAdd(id string, from_ string, to_ string, point int, message string, postAt string, editAt string) error {
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}
	//time := time.Now()
	if _, err := tx.Exec("INSERT INTO thanks (id, from_, to_, point, message, postAt, editAt) VALUES (?,?,?,?,?,?,?)", id, from_, to_, point, message, postAt, editAt); err != nil {
		log.Printf("fail: tx.Exec(), %v\n", err)
		if err := tx.Rollback(); err != nil {
			log.Printf("fail: tx.Rollback(), %v\n", err)
		}
		return err
	}
	if _, err := tx.Exec("UPDATE  users  SET point =point+? WHERE id=? ", point, to_); err != nil {
		log.Printf("fail: tx.Exec(), %v\n", err)
		if err := tx.Rollback(); err != nil {
			log.Printf("fail: tx.Rollback(), %v\n", err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit(), %v\n", err)
		return err
	}
	return nil
}
func ThankEdit(id string, to_ string, add_point int, delete_point int, message string, editAt string) error {
	//log.Print(delete_point)
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}
	if _, err := tx.Exec("UPDATE thanks SET to_=?, point=?, message=?,editAt=? WHERE id=? ", to_, add_point, message, editAt, id); err != nil {
		log.Printf("fail: tx.Exec(), %v\n", err)
		if err := tx.Rollback(); err != nil {
			log.Printf("fail: tx.Rollback(), %v\n", err)
		}
		return err
	}
	if _, err := tx.Exec("UPDATE  users  SET point =point+? WHERE id=? ", add_point-delete_point, to_); err != nil {
		log.Printf("fail: tx.Exec(), %v\n", err)
		if err := tx.Rollback(); err != nil {
			log.Printf("fail: tx.Rollback(), %v\n", err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit(), %v\n", err)
		return err
	}
	return nil
}

func ThankDelete(id string, to_ string, point int) error {
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}
	if _, err := tx.Exec("DELETE FROM thanks WHERE id=?  ", id); err != nil {
		log.Printf("fail: tx.Exec(), %v\n", err)
		if err := tx.Rollback(); err != nil {
			log.Printf("fail: tx.Rollback(), %v\n", err)
		}
		return err
	}
	if _, err := tx.Exec("UPDATE  users  SET point = point-? WHERE id=?", point, to_); err != nil {
		log.Printf("fail: tx.Exec(), %v\n", err)
		if err := tx.Rollback(); err != nil {
			log.Printf("fail: tx.Rollback(), %v\n", err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Printf("fail: tx.Commit(), %v\n", err)
		return err
	}
	return nil
}

func FetchAllThanks() (thanks []model.Thank, error error) {
	rows, err := Db.Query("SELECT id, from_, to_, point, message, postAt, editAt FROM thanks")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}
	thanks = make([]model.Thank, 0)
	for rows.Next() {
		var u model.Thank
		if err := rows.Scan(&u.Id, &u.From_, &u.To_, &u.Point, &u.Message, &u.PostAt, &u.EditAt); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil {
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		From, err := SearchUser(u.From_)
		if err != nil {
			log.Printf("fail: SearchUser, %v\n", err)
		}
		To, err := SearchUser(u.To_)
		if err != nil {
			log.Printf("fail: SearchUser, %v\n", err)
		}
		var thank model.Thank
		thank.Id = u.Id
		thank.Point = u.Point
		thank.Message = u.Message
		thank.FromWho = From.Name
		thank.ToWho = To.Name
		thank.PostAt = u.PostAt
		thank.EditAt = u.EditAt

		thanks = append(thanks, thank)
	}

	return thanks, nil
}

func FetchThanksGot(id string) (thanks []model.Thank, error error) {
	rows, err := Db.Query("SELECT id, from_, to_, point, message, postAt, editAt FROM thanks WHERE from_ =?", id)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}
	if rows == nil {
		return nil, nil
	}
	thanks = make([]model.Thank, 0)
	for rows.Next() {
		var u model.Thank
		if err := rows.Scan(&u.Id, &u.From_, &u.To_, &u.Point, &u.Message, &u.PostAt, &u.EditAt); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil {
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		From, err := SearchUser(u.From_)
		if err != nil {
			log.Printf("fail: SearchUser, %v\n", err)
		}
		To, err := SearchUser(u.To_)
		if err != nil {
			log.Printf("fail: SearchUser, %v\n", err)
		}
		var thank model.Thank
		thank.Id = u.Id
		thank.Point = u.Point
		thank.Message = u.Message
		thank.FromWho = From.Name
		thank.ToWho = To.Name
		thank.PostAt = u.PostAt
		thank.EditAt = u.EditAt
		thank.To_ = u.To_

		thanks = append(thanks, thank)
	}

	return thanks, nil
}

func FetchThanksSent(id string) (thanks []model.Thank, error error) {
	rows, err := Db.Query("SELECT id, from_, to_, point, message, postAt, editAt FROM thanks WHERE to_ =?", id)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}
	if rows == nil {
		return nil, nil
	}
	thanks = make([]model.Thank, 0)
	for rows.Next() {
		var u model.Thank
		if err := rows.Scan(&u.Id, &u.From_, &u.To_, &u.Point, &u.Message, &u.PostAt, &u.EditAt); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil {
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		From, err := SearchUser(u.From_)
		if err != nil {
			log.Printf("fail: SearchUser, %v\n", err)
		}
		To, err := SearchUser(u.To_)
		if err != nil {
			log.Printf("fail: SearchUser, %v\n", err)
		}
		var thank model.Thank
		thank.Id = u.Id
		thank.Point = u.Point
		thank.Message = u.Message
		thank.FromWho = From.Name
		thank.ToWho = To.Name
		thank.PostAt = u.PostAt
		thank.EditAt = u.EditAt

		thanks = append(thanks, thank)
	}

	return thanks, nil
}
