package main

import (
	"encoding/json"
	"fmt"

	"github.com/felipemjesus/iniciando-com-go/db/utils"
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

	utils.Soma(1, 1)

	fmt.Println(string(j))
}
