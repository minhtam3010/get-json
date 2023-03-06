package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	FullName string
}

func Conn() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/newdb")
	if err != nil {
		panic(err)
	}

	return db
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := Conn()
	defer db.Close()

	var users []User
	// GET DATA
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FullName)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	var res []byte
	res, err = json.Marshal(users)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func main() {

}
