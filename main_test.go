package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {

		TransformarCadena(testData.Input)

		prueba := TransformarCadena(testData.Input)
		// acá agregar chequeos propios del test por ejemplo:
		assert.Equal(t, prueba.Type, testData.Type, testData.Success)
		assert.Equal(t, prueba.Value, testData.Value, testData.Success)
		assert.Equal(t, prueba.Length, testData.Length, testData.Success)

		prueba = TransformarCadena("TX04AQS3")
		assert.NotEqual(t, prueba.Type, "TX", "no funciona, deberia devolver vacio")
		assert.NotEqual(t, prueba.Length, 4, "no deberia funcionar ya que retorna 0")

		prueba = TransformarCadena("TX04AQWE")
		assert.Equal(t, prueba.Value, "AQWE", "funciona")
		assert.Equal(t, prueba.Length, 4, "funciona")
	}
}
