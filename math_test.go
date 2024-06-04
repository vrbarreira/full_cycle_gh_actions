package main

import "testing"

func TestSoma(t *testing.T) {
	a := 2
	b := 3
	resultado := Soma(a, b)
	esperado := 5
	if resultado != esperado {
		t.Errorf("esperado %d, resultado %d", esperado, resultado)
	}
}