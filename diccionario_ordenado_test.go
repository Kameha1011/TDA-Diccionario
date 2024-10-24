package diccionario_test

import (
	"fmt"
	"math/rand/v2"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_VOLUMEN_CHICO  = 1000
	_VOLUMEN_GRANDE = 10000
)

var funcionCmpInts = func(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

var funcionCmpStrings = func(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func poblarArr(arr []int, max int) {
	hash := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		for {
			n := rand.IntN(max)
			if _, ok := hash[n]; !ok {
				arr[i] = n
				hash[n] = 1
				break
			}
		}
	}
}

func ordenarArr(arr []int) {
	if len(arr) <= 1 {
		return
	}
	medio := len(arr) / 2
	izq := make([]int, medio)
	der := make([]int, len(arr)-medio)
	copy(izq, arr[:medio])
	copy(der, arr[medio:])
	ordenarArr(izq)
	ordenarArr(der)
	i, j, k := 0, 0, 0
	for i < len(izq) && j < len(der) {
		if izq[i] < der[j] {
			arr[k] = izq[i]
			i++
		} else {
			arr[k] = der[j]
			j++
		}
		k++
	}
	for i < len(izq) {
		arr[k] = izq[i]
		i++
		k++
	}
	for j < len(der) {
		arr[k] = der[j]
		j++
		k++
	}
}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Abb vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](funcionCmpInts)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElementOrdenado(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](funcionCmpStrings)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoOrdenado(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoDatoHopscotchOrdenado(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	arrClaves := make([]int, _VOLUMEN_CHICO) // Creamos las claves de forma aleatoria para que al insertarlo no nos quede lineal
	poblarArr(arrClaves, 2*_VOLUMEN_CHICO)

	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	for i := 0; i < _VOLUMEN_CHICO; i++ {
		dic.Guardar(arrClaves[i], i)
	}
	for i := 0; i < _VOLUMEN_CHICO; i++ {
		dic.Guardar(arrClaves[i], 2*i)
	}
	ok := true
	for i := 0; i < _VOLUMEN_CHICO && ok; i++ {
		ok = dic.Obtener(arrClaves[i]) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorradosOrdenado(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericasOrdenado(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](funcionCmpInts)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestClaveVaciaOrdenado(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNuloOrdenado(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](funcionCmpStrings)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestCadenaLargaParticularOrdenado(t *testing.T) {
	t.Log("Guardamos claves largas y verificamos que se guarden correctamente")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestGuardarYBorrarRepetidasVecesOrdenado(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces")

	arrClaves := make([]int, _VOLUMEN_CHICO)
	poblarArr(arrClaves, 2*_VOLUMEN_CHICO)
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	for i := 0; i < _VOLUMEN_CHICO; i++ {
		dic.Guardar(arrClaves[i], i)
		require.True(t, dic.Pertenece(arrClaves[i]))
		dic.Borrar(arrClaves[i])
		require.False(t, dic.Pertenece(arrClaves[i]))
	}
}

func TestIteradorInternoOrdenadoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](funcionCmpStrings)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoOrdenadoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](funcionCmpStrings)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoOrdenado(t *testing.T) {
	t.Log("Valida que los datos sean recorridos de forma ordenada comparando con un array de claves ordenado")
	arrClaves := make([]int, _VOLUMEN_CHICO)
	poblarArr(arrClaves, 2*_VOLUMEN_CHICO)
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)

	for i := 0; i < _VOLUMEN_CHICO; i++ {
		dic.Guardar(arrClaves[i], i)
	}

	arrClavesOrdenado := make([]int, _VOLUMEN_CHICO)
	copy(arrClavesOrdenado, arrClaves)
	ordenarArr(arrClavesOrdenado)

	require.EqualValues(t, _VOLUMEN_CHICO, dic.Cantidad())

	arrClavesRecorridas := make([]int, 0, _VOLUMEN_CHICO)

	dic.Iterar(func(clave int, _ int) bool {
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
		return true
	})

	require.EqualValues(t, arrClavesOrdenado, arrClavesRecorridas)

}

func TestIteradorInternoRango(t *testing.T) {
	t.Log("Itersa sobre un rango de claves y valida que las claves y el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{6, 7, 10, 11, 13}
	for _, c := range claves {
		dic.Guardar(c, c)
	}

	inicio := 5
	fin := 13
	arrClavesRecorridas := make([]int, 0, 10)
	dic.IterarRango(&inicio, &fin, func(clave int, _ int) bool {
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
		return true
	})
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)

}

func TestIteradorInternoRangoCorteFin(t *testing.T) {
	t.Log("Itera desde el inicio hasta un fin arbitrario y valida que las claves y el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{3, 6, 7, 10, 11, 13}
	for _, c := range claves {
		dic.Guardar(c, c)
	}
	fin := 13
	arrClavesRecorridas := make([]int, 0, 10)
	dic.IterarRango(nil, &fin, func(clave int, _ int) bool {
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
		return true
	})
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)

}

func TestIteradorInternoRangoCorteInicio(t *testing.T) {
	t.Log("Itera desde un inicio arbitrario hasta el fin del ABB, y valida que las claves y el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{10, 11, 13, 15}
	for _, c := range claves {
		dic.Guardar(c, c)
	}
	inicio := 10
	arrClavesRecorridas := make([]int, 0, 10)
	dic.IterarRango(&inicio, nil, func(clave int, _ int) bool {
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
		return true
	})
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)

}

func TestIteradorInternoRangoSinCorte(t *testing.T) {
	t.Log("Itera sobre todo el ABB y valida que las claves y el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{3, 6, 7, 10, 11, 13, 15}
	for _, c := range claves {
		dic.Guardar(c, c)
	}
	arrClavesRecorridas := make([]int, 0, 10)
	dic.IterarRango(nil, nil, func(clave int, _ int) bool {
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
		return true
	})
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)

}

func TestIteradorInternoOrdenadoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](funcionCmpStrings)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumenOrdenado(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](funcionCmpStrings)

	claves := make([]string, n)
	arrClaves := make([]int, n)
	poblarArr(arrClaves, n*2)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el Abb */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", arrClaves[i])
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenOrdenado(b, n)
			}
		})
	}
}

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](funcionCmpStrings)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterarRango(t *testing.T) {
	t.Log("Iteramos sobre un rango de claves, y validamos que las claves y el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{6, 7, 10, 11, 13}
	arrClavesRecorridas := make([]int, 0, 10)
	for _, c := range claves {
		dic.Guardar(c, c)
	}

	inicio := 5
	fin := 13
	for iter := dic.IteradorRango(&inicio, &fin); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
	}
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)
}

func TestDiccionarioIterarRangoCorteFin(t *testing.T) {
	t.Log("Iteramos sobre un rango de claves, y validamos que las claves y el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{3, 6, 7, 10, 11, 13}
	arrClavesRecorridas := make([]int, 0, 10)
	for _, c := range claves {
		dic.Guardar(c, c)
	}

	fin := 13
	for iter := dic.IteradorRango(nil, &fin); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
	}
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)
}

func TestDiccionarioIterarRangoCorteInicio(t *testing.T) {
	t.Log("Iteramos sobre un rango de claves, y validamos que las claves y el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{10, 11, 13, 15}
	arrClavesRecorridas := make([]int, 0, 10)
	for _, c := range claves {
		dic.Guardar(c, c)
	}

	inicio := 10
	for iter := dic.IteradorRango(&inicio, nil); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
	}
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)
}

func TestDiccionarioIterarRangoSinCorte(t *testing.T) {
	t.Log("Iteramos sobre todas las claves y validamos que las claves y el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{3, 6, 7, 10, 11, 13, 15}
	arrClavesRecorridas := make([]int, 0, 10)
	for _, c := range claves {
		dic.Guardar(c, c)
	}
	for iter := dic.IteradorRango(nil, nil); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
	}
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)
}

func TestDiccionarioIterarOrden(t *testing.T) {
	t.Log("Iteramos sobre el diccionario y verificamos que el orden sea el correcto")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	arrClavesDeberianSerRecorridas := []int{3, 6, 7, 10, 11, 13, 15}
	arrClavesRecorridas := make([]int, 0, 10)
	for _, c := range claves {
		dic.Guardar(c, c)
	}
	for iter := dic.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		arrClavesRecorridas = append(arrClavesRecorridas, clave)
	}
	require.EqualValues(t, arrClavesDeberianSerRecorridas, arrClavesRecorridas)
}

func TestDiccionarioIterarRangoInvertido(t *testing.T) {
	t.Log("Iteramos sobre un rango de claves invertido (inicio mayor que fin) y validamos que no itere")
	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)
	claves := []int{15, 7, 3, 11, 6, 10, 13}
	for _, c := range claves {
		dic.Guardar(c, c)
	}
	inicio := 13
	fin := 5

	iter := dic.IteradorRango(&inicio, &fin)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorOrdenadoNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarOrdenadoTrasBorrados(t *testing.T) {
	t.Log("Prueba la iteracion de un diccionario tras borrar los elementos e insertar nuevos")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](funcionCmpStrings)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func ejecutarPruebasVolumenIteradorOrdenado(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](funcionCmpStrings)

	claves := make([]string, n)
	arrClaves := make([]int, n)
	poblarArr(arrClaves, n*2)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el Abb */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", arrClaves[i])
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorOrdenado(b, n)
			}
		})
	}
}

func TestVolumenIteradorOrdenadoCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](funcionCmpInts)

	arrClaves := make([]int, _VOLUMEN_GRANDE)
	poblarArr(arrClaves, 2*_VOLUMEN_GRANDE)

	/* Inserta 'n' parejas en el Abb */
	for i := 0; i < _VOLUMEN_GRANDE; i++ {
		dic.Guardar(arrClaves[i], i)
	}
	dic.Guardar(100, 18) // Insertamos este elemento para asegurarnos que el corte en algun momento se cumpla

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}
