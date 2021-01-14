package models

import (
	"log"
	db "main/database"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"create_time"`
	salt       string
}

func (u *User) AddUser() (id int64, err error) {
	exec, err := db.SqlDB.Exec("INSERT INTO blog_user(email,name,password) VALUES (?,?,?)", u.Email, u.Name, u.Password)
	if err != nil {
		return
	}
	id, err = exec.LastInsertId()
	return
}

func (u *User) LoginUser() (id int64, err error) {
	exec, err := db.SqlDB.Exec("SELECT email,password FROM blog_user where email=? and password=?", u.Email, u.Password)
	if err != nil {
		log.Print(err)
		return
	}
	affected, err := exec.RowsAffected()
	log.Print(affected)
	return

}
