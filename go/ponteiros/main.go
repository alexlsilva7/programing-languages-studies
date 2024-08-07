package main

func main() {
	// Memória -> Endereço -> Valor
	a := 10
	// variavel -> ponteiro que tem um endereço na memória -> valor
	var ponteiro *int = &a
	println(a, ponteiro, *ponteiro) // 10 0xc000065f38 10
	*ponteiro = 20                  // altera o valor de onde o ponteiro aponta
	println(a, ponteiro, *ponteiro) // 20 0xc000065f38 20
	b := &a                         // b é um ponteiro para a
	println(b, *b)                  // 0xc000065f38 20

	println(soma(10, 20)) // 30

	var result int = 0
	somaResult(1, 2, &result)
	println(result)
}

func soma(a, b int) int { // passagem de uma cópia do valor
	return a + b
}

func somaResult(a int, b int, result *int) { // passagem de um ponteiro para o valor
	*result = a + b
}
