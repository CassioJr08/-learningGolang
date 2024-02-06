package inteiros

import (
	"fmt"
	"testing"
)

func ExampleAdiciona() {
	soma := Adicao(1, 5)
	fmt.Println(soma)
	// Output: 6
}

func TestAdicionador(t *testing.T) {
	soma := Adicao(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
	}
}
