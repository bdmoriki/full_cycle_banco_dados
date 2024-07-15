package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Produto struct {
	ID    int `gorm:"primaryKey"`
	Nome  string
	Preco float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var produto Produto
	db.First(&produto, 1)

	produto.Nome = "Celular"
	db.Save(produto)

	db.Delete(produto)
}
