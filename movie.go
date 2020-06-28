package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func dbConn() (Db *sql.DB, err error) {
	log.Printf("hozirihoziri")
	dbDriver := "mysql"
	dbUser := "sowiriro"
	dbPass := "password"
	dbName := "sowiriroapp"
	Db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
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
		title := r.FormValue("title")
		log.Println(r.FormValue("title"))
		description := r.FormValue("description")
		log.Println(r.FormValue("description"))
		log.Printf("titleとdescriptionを取ってきたしん")
		stmt, err := Db.Prepare("INSERT INTO movies(title, description) VALUES(?, ?)")
		if err != nil {
			fmt.Println(err.Error())
		}
		log.Printf("statementがでけた！！！")
		stmt.Exec(title, description)
		log.Println("INSERT: Title: " + title + " | Description: " + description)
	}
	defer Db.Close()
	http.Redirect(w, r, "/", 301)
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
