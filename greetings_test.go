package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	//El puntero de tipo testing.T, sirve para reportar el resultaod de la prueba
	name := "Juan"
	//Se crea una expresion regular, que buscar una coincidencia exacta al nombre
	//Se usa acento grave `
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Juan")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(Juan)=%q,%v, quiere coincidencia para %#q,nil`, msg, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("")=%q,%v, quiere "", error`, msg, err)
	}

}
