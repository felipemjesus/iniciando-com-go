package main

import (
	"fmt"

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

	//db.Create(&Member{Name: "Felipe"})

	var members []Member
	db.Find(&members)

	fmt.Println(members)

	defer db.Close()
}
