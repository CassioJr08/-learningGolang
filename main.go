package main

import "fmt" //pacote que tem a função Println

const frances = "francês"
const prefixHelloFrances = "Bonjour, "
const espanhol = "espanhol"
const prefixHelloEspanhol = "Hola, "
const prefixHelloPortugues = "Hello, "

func Hello(nome string, idioma string) string {

	if nome == "" {
		nome = "Word"
	}

	prefixo := prefixHelloPortugues

	switch idioma {
	case espanhol:
		prefixo = prefixHelloEspanhol
	case frances:
		prefixo = prefixHelloFrances
	}

	return prefixo + nome

}
func main() {
	fmt.Println(Hello("junior", "francês"))
}
