package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

type Member struct {
	Name string
}

func members(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var members []Member
		db.Find(&members)
		return c.JSON(http.StatusOK, members)
	}
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

	e := echo.New()
	e.GET("/members", members(db))
	e.Start("0.0.0.0:8080")
}
