package main

import "fmt" //pacote que tem a função Println

const prefixHelloPortuges = "Hello, "

func Hello(nome string) string {
	if nome == "" {
		nome = "Word"
	}
	return prefixHelloPortuges + nome
}
func main() {
	fmt.Println(Hello("Word"))
}
