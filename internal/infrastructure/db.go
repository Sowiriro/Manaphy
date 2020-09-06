package infrastructure

import(
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
)

type (
	Conn struct {
		User string `required:"true"`
		Password string `required:"true"`
		Protocol string `required:"true"`
		Address string `required:"true"`
		Schema string `required:"true"`
	}
)

var DB *gorm.DB

func DBConnect() *gorm.DB {
	var conn Conn 

	DBMS := "mysql"
    USER := "root"
    PASS := "password"
    PROTOCOL := "tcp(Manaphy:3306)"
	DBNAME := "Manaphy"

	err := envconfig.Process("db", &conn)

	if err == nil {
		USER = conn.User
		PASS = conn.Password
		PROTOCOL = conn.Protocol
		ADDRESS = conn.Address
		DBNAME = conn.Schema
	}
	
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic("failed to connect to database")
	}
	
	db.LogMode(true)
    return db
}