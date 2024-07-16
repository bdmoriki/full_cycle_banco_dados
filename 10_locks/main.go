package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Categoria struct {
	ID       int `gorm:"primaryKey"`
	Nome     string
	Produtos []Produto `gorm:"many2many:produtos_categorias;"`
}

type Produto struct {
	ID         int `gorm:"primaryKey"`
	Nome       string
	Preco      float64
	Categorias []Categoria `gorm:"many2many:produtos_categorias;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8,b4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Produto{}, &Categoria{})

	tx := db.Begin()

	var c Categoria
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Nome = "Banheiro"

	tx.Debug().Save(&c)
	tx.Commit()
}
