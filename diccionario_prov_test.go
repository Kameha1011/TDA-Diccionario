package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"
)

func TestDiccionario(t *testing.T) {
	var arr []int
	dic := TDADiccionario.CrearAbb[int, int](funcionCmpInts)
	dic.Guardar(30, 1)
	dic.Guardar(20, 2)
	dic.Guardar(80, 3)
	dic.Guardar(10, 4)
	dic.Guardar(23, 5)
	dic.Guardar(70, 6)
	dic.Guardar(90, 7)
	dic.Iterar(func(clave int, valor int) bool {
		arr = append(arr, clave)
		fmt.Println(clave, valor)
		return clave < 30
	})
	fmt.Println(arr)
	// desde := 10

	iter := dic.IteradorRango(nil, nil)
	fmt.Println(iter)
}
