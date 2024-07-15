package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Categoria struct {
	ID   int `gorm:"primaryKey"`
	Nome string
}

type Produto struct {
	ID           int `gorm:"primaryKey"`
	Nome         string
	Preco        float64
	CategoriaID  int
	Categoria    Categoria
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProdutoID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8,b4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Produto{}, &Categoria{}, &SerialNumber{})

	// criar categoria
	categoria := Categoria{Nome: "Eletronicos"}
	db.Save(&categoria)

	// criar produto
	produto := Produto{Nome: "Monitor", Preco: 2000, CategoriaID: categoria.ID}
	db.Save(&produto)

	// criar serial number
	serialNumber := SerialNumber{Number: "S1", ProdutoID: produto.ID}
	db.Save(&serialNumber)

	var produtos []Produto
	db.Preload("Categoria").Preload("SerialNumber").Find(&produtos)
	for _, p := range produtos {
		fmt.Println(p.Nome, p.Categoria.Nome, p.SerialNumber.Number)
	}
}
