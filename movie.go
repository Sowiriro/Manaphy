package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

//func dbConn() (db *sql.DB, err error) {
//	log.Printf("hozirihoziri")
//	dbDriver := "mysql"
//	dbUser := "sowiriro"
//	dbPass := "password"
//	dbName := "sowiriroapp"
//	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
//	if err != nil {
//		return db, err
//	}
//	return db, err
//}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "indexを表示したしん")
	return
}

func show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie showを表示したしん")
	return
}

//func create(w http.ResponseWriter, r *http.Request)  {
//	log.Printf("createの最初がみられた")
//	db, err := dbConn()
//	log.Printf("dbがひらけた！")
//	if err != nil {
//		return
//	}
//	if r.Method == "POST" {
//		title := r.FormValue("title")
//		description := r.FormValue("description")
//		statement, err := db.Prepare("INSERT INTO movies (id, title, description) VALUES($1, $2, $3)")
//		if err != nil {
//			panic(err.Error())
//		}
//		statement.Exec(title, description)
//		log.Println("INSERT: Title: " + title + " | Description: " + description)
//	}
//	defer db.Close()
//	http.Redirect(w, r, "/", 301)
//	return
//}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie updateを表示したしん")
	return
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "movie deleteを表示したしん")
	return
}
