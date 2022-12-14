package dao

import (
	"back/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//func InsertUserDb(name string, id string) error {
//	tx, err := Db.Begin()
//	if err != nil {
//		log.Printf("fail: db.Begin, %v\n", err)
//		return err
//	}
//	if _, err := tx.Exec("INSERT INTO users (id, name, point) VALUES (?,?,0)", id, name); err != nil {
//		log.Printf("fail: tx.Exec(), %v\n", err)
//		if err := tx.Rollback(); err != nil {
//			log.Printf("fail: tx.Rollback(), %v\n", err)
//		}
//		return err
//	}
//	if err := tx.Commit(); err != nil {
//		log.Printf("fail: tx.Commit(), %v\n", err)
//		return err
//	}
//	return nil
//}

func UserEdit(id string, name string) error {
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}
	if _, err := tx.Exec("UPDATE users SET name =? WHERE id=? ", name, id); err != nil {
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

func UserAdd(id string, name string) error {
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}
	if _, err := tx.Exec("INSERT INTO users (id, name, point) VALUES (?,?,0) ", id, name); err != nil {
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

//func SearchUserDb(name string) (users []model.User, error error) {
//	rows, err := Db.Query("SELECT id, name, age FROM users ")
//	if err != nil {
//		log.Printf("fail: db.Query, %v\n", err)
//		return nil, err
//	}
//	users = make([]model.User, 0)
//	for rows.Next() {
//		var u model.User
//		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
//			log.Printf("fail: rows.Scan, %v\n", err)
//			if err := rows.Close(); err != nil {
//				log.Printf("fail: rows.Close(), %v\n", err)
//			}
//			return nil, err
//		}
//		users = append(users, u)
//	}
//
//	return users, nil
//}

// UserFetchAll 全てのメンバーを表示
func UserFetchAll(id string) (users []model.UserPoint, error error) {
	if id == "all" {
		rows, err := Db.Query("SELECT id, name, point FROM users")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			return nil, err
		}
		users = make([]model.UserPoint, 0)
		for rows.Next() {
			var u model.UserPoint
			err := rows.Scan(&u.Id, &u.Name, &u.Point)
			if err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)
				if err := rows.Close(); err != nil {
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				return nil, err
			}
			users = append(users, u)
		}
	}
	if id != "all" {
		rows, err := Db.Query("SELECT id, name, point FROM users WHERE id=? ", id)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			return nil, err
		}
		users = make([]model.UserPoint, 0)
		for rows.Next() {
			var u model.UserPoint
			err := rows.Scan(&u.Id, &u.Name, &u.Point)
			if err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)
				if err := rows.Close(); err != nil {
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				return nil, err
			}
			//u.Point, err = CurrentUserPoint(u.Id)
			//if err != nil {
			//	log.Printf("fail: UserPoint, %v\n", err)
			//}
			users = append(users, u)
		}
	}
	return users, nil

}

// UserRanking ポイント高い順
func UserRanking() (users []model.UserRank, error error) {
	rows, err := Db.Query("SELECT id, name, point FROM users ORDER BY point DESC")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}
	users = make([]model.UserRank, 0)
	i := 1
	for rows.Next() {
		var u model.UserRank
		u.Rank = i
		err := rows.Scan(&u.Id, &u.Name, &u.Point)
		if err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil {
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		i++
		users = append(users, u)
	}
	return users, nil
}

// UserSearch idからメンバー検索
func UserSearch(id string) (user model.UserPoint, error error) {
	rows, err := Db.Query("SELECT id, name, point FROM users WHERE id =?", id)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return model.UserPoint{}, err
	}
	rows.Next()
	if err := rows.Scan(&user.Id, &user.Name, &user.Point); err != nil {
		log.Printf("fail: rows.Scan, %v\n", err)
		if err := rows.Close(); err != nil {
			log.Printf("fail: rows.Close(), %v\n", err)
		}
		//user.Point, err = CurrentUserPoint(user.Id)
		return model.UserPoint{}, err
	}
	return user, nil
}
