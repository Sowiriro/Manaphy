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

type Movie struct {
	Id          int		`json:"id"`
	Title       string		`json:"title"`
	Description string			`json:"description"`
	CreatedAt   time.Time		`json:"created_at"`
	UpdatedAt   time.Time		`json:"updated_at"`
}

func responseByJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
	return
}


func dbConn() (Db *sql.DB, err error) {
	log.Printf("DB開き中！！！！！！！")
	dbDriver := "mysql"
	//dbUser := "sowiriro"
	//dbPass := "password"
	//dbName := "sowiriroapp"
	Db, err = sql.Open(dbDriver, "root:password@tcp(mysql)/sowiriroapp?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		return Db, err
	}
	return Db, err
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("indexをみたしん")
	Db, err := dbConn()
	if err != nil {
		return
	}
	log.Printf("dbが拓けたしん")
	stmt, err := Db.Query("SELECT * FROM movies ORDER BY id DESC")
	movie := Movie{}
	res := []Movie{}
	for stmt.Next(){
		var id int
		var title, description string
		var created_at, updated_at time.Time
		err = stmt.Scan(&id, &title, &description, &created_at, &updated_at)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("scan成功")
		movie.Id = id
		movie.Title = title
		movie.Description = description
		movie.CreatedAt = created_at
		movie.UpdatedAt = updated_at
		log.Printf("代入成功")


		res = append(res, movie)
		log.Printf("indexを持ってこれたと思う")
	}

	log.Print(res)
	defer stmt.Close()
	response , err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(response))
	log.Printf("errがなかったら取れてるよ！！！")
	responseByJSON(w, res)
}

func show (w http.ResponseWriter, r *http.Request) {
	log.Printf("movie showをみたしん")
	Db, err := dbConn()
	if err != nil {
		return
	}
		urlId := r.FormValue("id")
		log.Printf("urlIdをとれた")
		stmt, err := Db.Query("SELECT * FROM movies WHERE id=?", urlId)
		if err != nil {
			fmt.Println(err.Error())
		}
		log.Print("statementができた！")

	movie := Movie{}
	res := []Movie{}

	for stmt.Next(){
		var id int
		var title, description string
		var created_at, updated_at time.Time
		err = stmt.Scan(&id, &title, &description, &created_at, &updated_at)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("scan成功")
		movie.Id = id
		movie.Title = title
		movie.Description = description
		movie.CreatedAt = created_at
		movie.UpdatedAt = updated_at
		log.Printf("代入成功")
		res = append(res, movie)
		log.Printf("showを持ってこれたと思う")
	}


		defer stmt.Close()
		response, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(response))
		log.Printf("errがなかったら取れてるよ！！！")
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
		var m Movie
		err := decoder.Decode(&m)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Println(m.Title, m.Description)

		stmt, err := Db.Prepare("INSERT INTO movies(title, description) VALUES(?, ?)")
		if err != nil {
			fmt.Println(err.Error())
		}
		log.Printf("statementがでけた！！！")
		movie, err := stmt.Exec(m.Title, m.Description); if err != nil {
			fmt.Println(err.Error())
		}

		response, err := json.Marshal(movie)

		if err != nil  {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Printf(string(response))
		defer stmt.Close()
		log.Println("INSERT: Title: " + m.Title + " | Description: " + m.Description)
	}

	defer Db.Close()
	return
}


func update(w http.ResponseWriter, r *http.Request) {
	log.Printf( "movie updateをみたしん")
	Db, err := dbConn()
	if err != nil {
		return
	}
	log.Printf("DBが開けたしん")
	if r.Method == "POST"{
		//decoder := json.NewDecoder(r.Body)
		//var m Movie
		//err :=  decoder.Decode(&m)
		//if err != nil {
		//	return
		//}
		//defer r.Body.Close()
		//log.Println(m.Title, m.Description, m.Id)
		urlId := r.FormValue("id")
		title := r.FormValue("title")
		log.Println(r.FormValue("title"))
		description := r.FormValue("description")
		log.Println(r.FormValue("description"))
		log.Printf("titleとdescriptionを取ってきたしん")

		stmt, err := Db.Prepare("UPDATE movies SET title=?, description=? WHERE id=?")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer stmt.Close()
		log.Printf("statementがでけた！！！")

		movie, err := stmt.Exec(title, description, urlId); if err != nil {
			fmt.Println(err.Error())
		}

		log.Println(movie)

		res, err := json.Marshal(movie)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Printf(string(res))
		log.Println("UPDATE: Title: " + title + " | Description: " + description + " | ID: " + urlId )
	}

	log.Print("updateをすることができた")
	defer Db.Close()
	return
}

func delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("movie deleteをみたしん")
	Db, err := dbConn()
	if err != nil {
		return
	}
	if r.Method == "POST"{
	urlId := r.FormValue("id")
	log.Printf("IDを取ってきたしん")
	stmt, err := Db.Prepare("DELETE FROM movies WHERE id=?")
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("statementがでけた！！！")
	if _, err = stmt.Exec(urlId); err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	log.Println("DELETE")
}
 	log.Print("deleteをすることができた")

	return
}
