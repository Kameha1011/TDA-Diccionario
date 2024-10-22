package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"
)

func TestDiccionario(t *testing.T) {
	dic := TDADiccionario.CrearAbb[int, int](funcionCmpInts)
	dic.Guardar(30, 1)
	dic.Guardar(20, 2)
	dic.Guardar(80, 3)
	dic.Guardar(10, 4)
	dic.Guardar(23, 5)
	dic.Guardar(70, 6)
	dic.Guardar(90, 7)
	desde := 20
	fin := 80
	iter := dic.IteradorRango(&desde, &fin)
	for iter.HaySiguiente() {
		fmt.Println(iter.VerActual())
		iter.Siguiente()
	}
}
