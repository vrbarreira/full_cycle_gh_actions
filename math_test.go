package main

import "testing"

func TestSoma(t *testing.T) {
	a := 5
	b := 3
	resultado := Soma(a, b)
	esperado := 8
	if resultado != esperado {
		t.Errorf("esperado %d, resultado %d", esperado, resultado)
	}
}