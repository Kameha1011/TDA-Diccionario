package diccionario_test

import (
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

var funcionCmp = func(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func TestDiccionarioVacio(t *testing.T) {
	dic := TDADiccionario.CrearAbb[int, int](funcionCmp)
	require.Equal(t, 0, dic.Cantidad())
	dic.Guardar(87, 1)
	dic.Guardar(23, 2)
	dic.Guardar(45, 3)
	dic.Guardar(12, 4)
	dic.Guardar(98, 5)
	dic.Guardar(34, 6)
	dic.Guardar(56, 7)
	dic.Guardar(78, 8)
	dic.Guardar(90, 9)
	dic.Guardar(1, 10)
	require.Equal(t, 10, dic.Cantidad())
	require.Equal(t, 3, dic.Borrar(45))
	require.False(t, dic.Pertenece(45))
	require.True(t, dic.Pertenece(78))
	require.True(t, dic.Pertenece(56))
	require.Equal(t, 7, dic.Obtener(56))
	require.Equal(t, 8, dic.Obtener(78))
	require.True(t, dic.Pertenece(34))
	require.Equal(t, 9, dic.Cantidad())
	require.Equal(t, 6, dic.Borrar(34))
}
