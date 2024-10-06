package diccionario

import (
	"fmt"
	fnv "hash/fnv"
)

type Estado int

const (
	_VACIO Estado = iota
	_OCUPADO
	_BORRADO
	_TAM_INICIAL          = 23
	_CRITERIO_REDIMENSION = 0.7
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado Estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashing[K comparable](clave K, tam int) int {
	hash := fnv.New32()
	hash.Write(convertirABytes(clave))
	return int(hash.Sum32()) % tam

}

func buscar[K comparable, V any](tabla []celdaHash[K, V], tam int, clave K) (int, bool) {
	indice := hashing(clave, tam)
	parDatoValor := tabla[indice]
	if parDatoValor.clave == clave && parDatoValor.estado == _OCUPADO {
		return indice, true
	}
	i := 1
	for parDatoValor.estado != _VACIO { // && indice < 2*tam
		indice = (indice + i) % tam
		parDatoValor = tabla[indice]
		if parDatoValor.clave == clave && parDatoValor.estado == _OCUPADO {
			return indice, true
		}
		i++
	}
	return indice, false
}

func panicClaveNoPertenece[K comparable, V any](hash *hashCerrado[K, V], clave K) {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
}

func redimensionar[K comparable, V any](hash *hashCerrado[K, V], tamnuevo int) {
	tablaVieja := hash.tabla
	hash.tabla = make([]celdaHash[K, V], tamnuevo)
	hash.tam = tamnuevo
	hash.borrados = 0
	for _, celda := range tablaVieja {
		if celda.estado == _OCUPADO {
			hash.Guardar(celda.clave, celda.dato)
		}
	}
}

func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func obtenerPrimoSiguiente(n int) int {
	for {
		n++
		if esPrimo(n) {
			return n
		}
	}
}

func (hash *hashCerrado[K, V]) ObtenerPrimoSiguiente(n int) int {
	for {
		n++
		if esPrimo(n) {
			return n
		}
	}
}

func CrearHash[K comparable, V any]() *hashCerrado[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tabla = make([]celdaHash[K, V], _TAM_INICIAL)
	hash.tam = _TAM_INICIAL
	hash.cantidad = 0
	hash.borrados = 0
	return hash
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	indice, encontrado := buscar(hash.tabla, hash.tam, clave)
	if encontrado {
		hash.tabla[indice].dato = dato
		hash.tabla[indice].clave = clave
		hash.tabla[indice].estado = _OCUPADO
	} else {
		hash.tabla[indice] = celdaHash[K, V]{clave: clave, dato: dato, estado: _OCUPADO}
		if float64(hash.cantidad+hash.borrados)/float64(hash.tam) > _CRITERIO_REDIMENSION {
			redimensionar(hash, 2*hash.tam)
		}
		hash.cantidad++
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, encontrado := buscar(hash.tabla, hash.tam, clave)
	return encontrado
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	panicClaveNoPertenece(hash, clave)
	indice, _ := buscar(hash.tabla, hash.tam, clave)
	return hash.tabla[indice].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	panicClaveNoPertenece(hash, clave)
	indice, _ := buscar(hash.tabla, hash.tam, clave)
	hash.tabla[indice].estado = _BORRADO
	hash.cantidad--
	hash.borrados++
	return hash.tabla[indice].dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(func(clave K, dato V) bool) {

}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIteradorDiccionario[K, V]()
}

type iterDiccionario[K comparable, V any] struct {
	hash    *hashCerrado[K, V]
	primero *celdaHash[K, V]
	ultimo  *celdaHash[K, V]
}

func crearIteradorDiccionario[K comparable, V any]() *iterDiccionario[K, V] {
	return nil
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return false
}

func (iter *iterDiccionario[K, V]) Siguiente() {

}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	return iter.primero.clave, iter.primero.dato
}
