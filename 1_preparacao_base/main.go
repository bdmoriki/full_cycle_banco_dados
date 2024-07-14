package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Produto struct {
	ID    string
	Nome  string
	Preco float64
}

func NewProduto(nome string, preco float64) *Produto {
	return &Produto{
		ID:    uuid.New().String(),
		Nome:  nome,
		Preco: preco}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	produto := NewProduto("Notebook", 12000)
	err = inserirProduto(db, produto)
	if err != nil {
		panic(err)
	}
}

func inserirProduto(db *sql.DB, produto *Produto) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(produto.ID, produto.Nome, produto.Preco)
	if err != nil {
		return err
	}
	return nil
}
