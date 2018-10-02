package matematica

import "errors"

//Calculo faz calculos genéricos
func Calculo(funcao func(int, int) (int, error), x int, y int) (int, error) {
	return funcao(x, y)
}

//Multiplicador faz a multiplicação
func Multiplicador(x int, y int) (int, error) {
	return -1, errors.New("Multiplicador erro forçado")
	//return x * y, nil
}
