package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Member struct {
	Name string
}

func (m Member) andar() {
	fmt.Println(m.Name, "andou")
}

func main() {
	m := Member{Name: "Felipe"}
	m.andar()
}
