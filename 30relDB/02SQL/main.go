package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "sa:sqlpass@/godata")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/drop", dropHandler)
	http.ListenAndServe(":8080", nil)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "Yeah, db connected")
	check(err)
}

func usersHandler(w http.ResponseWriter, request *http.Request) {

	rows, err := db.Query(`SELECT usareName FROM users;`)
	check(err)

	var s, name string
	s = "All users:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}

	fmt.Fprintln(w, s)
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`CREATE TABLE users (userName VARCHAR(100) NOT NULL , password VARCHAR(100) NOT NULL);`)
	check(err)

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, n)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`INSERT INTO users VALUES ("James", "007");`)
	check(err)

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, n)

}

func readHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(`SELECT * FROM users`)
	check(err)

	var name, pass string

	for rows.Next() {

		rows.Scan(&name, &pass)
		fmt.Fprintln(w, name, " ", pass)
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`UPDATE users SET password="007pass" WHERE password="007";`)
	check(err)

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, n)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`DELETE FROM users WHERE password="007";`)
	check(err)

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(w, n)
}

func dropHandler(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`DROP TABLE users;`)
	check(err)

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE users")
}
