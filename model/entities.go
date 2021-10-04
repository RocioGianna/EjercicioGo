package model

type Result struct {
	Type   string
	Value  string
	Length int
}

func NewResult(tipo string, value string, length int) Result {
	return Result{tipo, value, length}
}
