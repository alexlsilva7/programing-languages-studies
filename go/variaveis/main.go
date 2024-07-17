package main

import "fmt"

var ( // declaração de variaveis globais
	a     string  = "Teste"
	b     int     = 10
	c     float64 = 0.2
	d     bool    = true
	array         = []int{1, 2, 3}
)

type ID int          // declaração de um tipo
type pessoa struct { // declaração de um tipo struct
	nome  string
	idade int
}

func main() {
	println(a, b, c, d)
	pessoa1 := pessoa{"João", 20} // declaração de uma variavel do tipo struct
	fmt.Printf("%v tem %v anos\n", pessoa1.nome, pessoa1.idade)
	fmt.Printf("O tipo da variavel pessoa1 é %T\n", pessoa1)
}

func x() {
	e := "x"         // inicialização de uma nova variavel pela primeira vez
	e = "atribuição" // atribuição de valor a variavel
	println(e)
}
