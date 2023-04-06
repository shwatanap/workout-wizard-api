package driver

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/shwatanap/workout-wizard-api/src/config"
)

func NewDriver() *sql.DB {
	db, err := sql.Open("mysql", config.GetDbUri())
	if err != nil {
		log.Println("DB connect failed")
		panic(err)
	}

	log.Println("DB connect success")

	return db
}
