package main

import "testing"

func TestHellow(t *testing.T) {

	verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}
	t.Run("diz hello para as pessoas", func(t *testing.T) {
		resultado := Hello("Word")
		esperado := "Hello, Word"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Hello, Word' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Hello("")
		esperado := "Hello, Word"

		verificarMensagemCorreta(t, resultado, esperado)
	})
}
