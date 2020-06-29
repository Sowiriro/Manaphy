package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

type movie struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}


func dbConn() (Db *sql.DB, err error) {
	log.Printf("hozirihoziri")
	dbDriver := "mysql"
	//dbUser := "sowiriro"
	//dbPass := "password"
	//dbName := "sowiriroapp"
	Db, err = sql.Open(dbDriver, "root:password@tcp(mysql)/sowiriroapp")
	if err != nil {
		fmt.Println(err.Error())
		return Db, err
	}
	return Db, err
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "indexを表示したしん")
	return
}

func show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie showを表示したしん")
	return
}

func create(w http.ResponseWriter, r *http.Request)  {
	log.Printf("createの最初がみられた")
	Db, err := dbConn()
	if err != nil {
		return
	}
	log.Println("dbがひらけたしん！")
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var m movie
		err := decoder.Decode(&m)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Println(m.Title, m.Description)
		//title := r.FormValue("title")
		//log.Println(r.FormValue("title"))
		//description := r.FormValue("description")
		//log.Println(r.FormValue("description"))
		//log.Printf("titleとdescriptionを取ってきたしん")

		stmt, err := Db.Prepare("INSERT INTO movies(title, description) VALUES(?, ?)")
		if err != nil {
			fmt.Println(err.Error())
		}
		log.Printf("statementがでけた！！！")
		if _, err = stmt.Exec(m.Title, m.Description); err != nil {
			fmt.Println(err.Error())
		}
		defer stmt.Close()
		log.Println("INSERT: Title: " + m.Title + " | Description: " + m.Description)
	}
	defer Db.Close()
	return
}


func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie updateを表示したしん")
	return
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie deleteを表示したしん")
	return
}
