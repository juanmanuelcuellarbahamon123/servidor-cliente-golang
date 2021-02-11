package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func DBConnection() {

	driver := "mysql"
	user := "root"
	password := ""
	database := "prueba"
	host := "localhost"
	port := "3306"

	URI := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database

	db, err := sqlx.Open(driver, URI)

	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	DBClient = db

}
