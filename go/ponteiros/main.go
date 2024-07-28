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
}
