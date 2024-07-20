package main

import (
	"errors"
	"fmt"
)

func main() {
	println(soma(1, 2))
	a, b := retornoMulti()
	println(a, b)

	soma, err := somaErroMaior50(20, 31)
	if err != nil {
		fmt.Println(err)
	} else {
		println(soma)
	}

	fmt.Printf("soma variatica: %d\n", somaVariatica(1, 2, 3, 4, 5))

	total := func() int { // função anônima
		return somaVariatica(1, 2, 3, 4, 5, 6)
	}()

	fmt.Println(total)
}

func soma(a int, b int) int {
	return a + b
}

func retornoMulti() (int, int) {
	return 1, 2
}

func somaErroMaior50(a int, b int) (int, error) {
	soma := a + b
	if soma > 50 {
		return 0, errors.New("Erro: Soma maior que 50")
	}
	return soma, nil
}

func somaVariatica(args ...int) int {
	soma := 0
	for _, v := range args {
		soma += v
	}
	return soma
}
