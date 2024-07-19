package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"João":  1200.0,
		"Maria": 1500.0,
		"Pedro": 800.0,
	}
	fmt.Println(salarios)
	salarios["Alexandre"] = 2000.0
	delete(salarios, "Pedro")
	fmt.Println(salarios)

	salarios2 := make(map[string]float64)
	fmt.Printf("O map salarios2 é %v, len=%d\n", salarios2, len(salarios2))
	keys := make([]string, 0, len(salarios2))
	for k := range salarios2 {
		keys = append(keys, k)
	}
	fmt.Printf("As chaves do map salarios2 são %v\n", keys)

	for k, v := range salarios {
		fmt.Printf("O salário de %s é %.2f\n", k, v)
	}
}
