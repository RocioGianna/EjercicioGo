package main

import (
	"fmt"
	"log"
	"strconv"

	"tudai2021.com/model"
)

func TransformarCadena(input string) model.Result {
	tipo := input[0:2]
	tamanio := input[2:4]
	value := input[4:]

	size, err := strconv.Atoi(tamanio)
	if err != nil {
		log.Fatal(err)
	}
	retorno := model.NewResult(tipo, value, size)
	if validacionObjeto(retorno) {
		return retorno
	}
	retorno.Type = ""
	retorno.Value = ""
	retorno.Length = 0
	return retorno
}

func validacionObjeto(resultado model.Result) bool {
	contador := 0
	if len(resultado.Value) == resultado.Length {
		if resultado.Type == "TX" {
			for i := 0; i < resultado.Length; i++ {
				if resultado.Value[i] >= 65 && resultado.Value[i] <= 90 {
					contador++
				}
			}
		} else if resultado.Type == "NN" {
			size, err := strconv.Atoi(resultado.Value[0:])
			if err != nil {
				fmt.Print(size)
				return false
			}
			return true
		}
	}
	if contador == resultado.Length {
		return true
	}
	return false
}

func main() {

	resultado := TransformarCadena("NN042312")

	if validacionObjeto(resultado) {
		fmt.Print(resultado)
	}
}
