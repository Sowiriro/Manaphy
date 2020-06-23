package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return db, err
	}
	return db, err
}
//
//func Index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "indexを表示したしん")
//	return
//}
//
//func Show(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "movie showを表示したしん")
//	return
//}
//
//func (movie *movie)Create(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	statement := "INSERT INTO movies (id, title, description) VALUES($1, $2, $3)"
//	stmt, err := db.Prepare(statement)
//	if err != nil {
//		return
//	}
//	defer stmt.Close()
//	err = stmt.QueryRow(movie.title, movie.description).Scan(&movie.Id)
//	return
//}
//
//func Update(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "movie updateを表示したしん")
//	return
//}
//
//func Delete(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "movie deleteを表示したしん")
//	return
//}
