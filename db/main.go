package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Member struct {
	Name string `json:"name"`
	Age  int    `json:"-"`
}

func (m Member) andar() {
	fmt.Println(m.Name, "andou")
}

func main() {
	m := Member{Name: "Felipe", Age: 29}
	j, err := json.Marshal(m)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(j))
}
