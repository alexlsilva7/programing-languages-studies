package main

import "fmt"

func main() {
	var arrayInteiros [3]int
	arrayInteiros[0] = 1
	arrayInteiros[1] = 2
	arrayInteiros[2] = 3
	// percorrer o array
	fmt.Printf("O array de inteiros é %v\n", arrayInteiros)
	for i := 0; i < len(arrayInteiros); i++ {
		fmt.Printf("O valor do array na posição %d é %d\n", i, arrayInteiros[i])
	}
	// percorrer o array com range
	for i, v := range arrayInteiros {
		fmt.Printf("O valor do array na posição %d é %d\n", i, v)
	}

	// slices

	sliceInteiros := []int{1, 2, 3}
	fmt.Printf("O slice de inteiros é %v, len=%d, cap=%d\n", sliceInteiros, len(sliceInteiros), cap(sliceInteiros))

	sliceInteiros = append(sliceInteiros, 4) // adiciona um elemento ao slice e dobra a capacidade por padrão
	fmt.Printf("O slice de inteiros é %v, len=%d, cap=%d\n", sliceInteiros, len(sliceInteiros), cap(sliceInteiros))
}
