package main

import "fmt"

type Endereco struct {
	logradouro string
	numero     int
}

type Pessoa interface {
	desativar()
}

type Cliente struct {
	nome  string
	idade int
	ativo bool
	Endereco
}

func (c *Cliente) desativar() { // método de Cliente
	c.ativo = false
}

func desativacao(p Pessoa) { // função que recebe uma interface Pessoa, que é implementada por Cliente
	p.desativar()
}

func main() {
	wesley := Cliente{"Wesley", 22, true, Endereco{"Rua 1", 123}}
	wesley.desativar()
	fmt.Printf("%s tem %d anos, ativo=%t\n", wesley.nome, wesley.idade, wesley.ativo)
}
