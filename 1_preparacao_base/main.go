package main

import "github.com/google/uuid"

type Produto struct {
	ID    string
	Nome  string
	Preco float64
}

func NewProduto(nome string, preco float64) *Produto {
	return &Produto{ID: uuid.New().String(), Nome: nome, Preco: preco}
}

func main() {

}
