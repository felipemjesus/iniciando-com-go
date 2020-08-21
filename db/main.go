package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Member struct {
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Member{})
}
