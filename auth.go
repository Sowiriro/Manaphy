package main

import (
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type JwtCustomClaims struct {
	UserID      int64
	jwt.StandardClaims
}

type User struct {
	Id 			int64 		`json:"id"`
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)



type JWT struct {
	Token string `json:"token"`
}

var jwtMiddleWare = jwtmiddleware.New(jwtmiddleware.Options{
ValidationKeyGetter: func(token *jwt.Token)(interface{}, error){
return []byte("SIGNINGKEY"), nil  //何をreturnしているのか？引数になぜか(jwtの)tokenそのものを持ってきたが使ってない
// returnをしているのは、バイト型の配列に環境変数を読み込んで、入れているだけ。
},

SigningMethod: jwt.SigningMethodHS256,
})



func authenticate (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ちゃんと承認をしたしん")
	return
}


func signUp (w http.ResponseWriter, r *http.Request){
	log.Print("user createの中身が見られた")
	Db, err := dbConn()
	if err != nil {
		return
	}
	log.Println("dbがひらけた真")
	if r.Method == "POST" {

		decoder := json.NewDecoder(r.Body)
		var u User
		err := decoder.Decode(&u)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Println(u.Name, u.Email, u.Password)

		u.Password = generateHash(u.Password)

		fmt.Println("ハッシュ化されたパスワード", u.Password)

		stmt, err := Db.Prepare("INSERT INTO users(name, email, password) VALUES(?, ?, ?)")
		if err != nil {
			return
		}
		log.Printf("statementがでけた")

		_, err = stmt.Exec(u.Name, u.Email, u.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("userの完了")
		log.Println("INSERT: Name: " + u.Name + "| Email: " + u.Email + "| Password: " + u.Password)

		w.Header().Set("Content-Type", "application/json")
		responseByJSON(w, u)
		log.Println("ちゃんと送れたしん")
	}

	defer Db.Close()
	return
}


func logout(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "logoutをしただしん")
	return
}


func login(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "ログインをしただしん")

	Db, err := dbConn()
	if err != nil {
		return
	}
	log.Println("dbがひらけたしん")

	var u  User
	//var jwt JWT

	json.NewDecoder(r.Body).Decode(&u)

	password := &u.Password
	//decodeをしてきたpasswordでuserの中身を取ってくる
	user := Db.QueryRow("SELECT * FROM USERS WHERE email=?", u.Email)

	err = user.Scan(&u.Id,&u.Name, &u.Email, &u.Password)

	//ここで該当をするEmailがなかったらエラー
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}

	log.Println(u.Id, u.Email, u.Password)

	hashedPassword := u.Password
	fmt.Println("hashedPotato: ", hashedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(*password))

	if err != nil {
		fmt.Println(err.Error())
	}

	claims := &JwtCustomClaims{
		u.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	jwtFileName := "./private-key.pem"

	keyData, err := ioutil.ReadFile(jwtFileName)
	if err != nil {
		return
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return
	}

	t, err := token.SignedString(signKey)
	if err != nil {
		return
	}

	responseByJSON(w, t)
}



//func generateToken(claims jwt.Claims)(t, string, err error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	jwtFile := "./public-key.pem"
//
//	keyData, err := ioutil.ReadFile(jwtFile)
//	if err != nil {
//		return
//	}
//	//ioutil.readFileが何をしているのか？
//
//	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
//	if err != nil {
//		return
//	}
//
//	t, err = token.SignedString(signKey)
//	if err != nil {
//		return
//	}
//
//	return
//}

func generateHash(password string) string {
	result := sha256.Sum256([]byte(password))
	//sha256が何をしているのかをわかったら消す
	return hex.EncodeToString(result[:])
	//hexが何をしているのかわかったらこれを消す
}