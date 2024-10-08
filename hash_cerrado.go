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
	_TAM_INICIAL          = 20
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

// busca un indice dado dentro de la tabla hash, primero verifica que la clave que encontro sea igual a la dada y que el estado sea OCUPADO retornando el indice y verdadero,
// de lo contrario, irá buscando en las siguientes posiciones de manera ciclica hasta encontrar una celda que tenga igual clave y estado ocupado.
// si no se encuentra esta celda, el bucle terminará cuando encuentre una celda vacia y te retornará el indice de ésta (este seria el caso de una colisión al guardar una
// nueva clave).
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
	//tamnuevo = obtenerPrimoSiguiente(tamnuevo)
	hash.tabla = make([]celdaHash[K, V], tamnuevo)
	hash.cantidad = 0
	hash.tam = tamnuevo
	hash.borrados = 0
	for _, celda := range tablaVieja {
		if celda.estado == _OCUPADO {
			hash.Guardar(celda.clave, celda.dato)
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
		} else {
			hash.cantidad++
		}
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
	if float64(hash.cantidad+hash.borrados)/float64(hash.tam) < 1-_CRITERIO_REDIMENSION && hash.tam/2 > _TAM_INICIAL {
		redimensionar(hash, hash.tam/2)
	}
	indice, _ := buscar(hash.tabla, hash.tam, clave)
	hash.tabla[indice].estado = _BORRADO
	hash.cantidad--
	hash.borrados++
	return hash.tabla[indice].dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	i := 0
	condicion := true
	for i < hash.tam && condicion {
		if hash.tabla[i].estado == _OCUPADO {
			condicion = visitar(hash.tabla[i].clave, hash.tabla[i].dato)
		}
		i++
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIteradorDiccionario[K, V](hash)
}

type iterDiccionario[K comparable, V any] struct {
	hash   *hashCerrado[K, V]
	actual int
	indice int
}

func panicIteradorTerminoDeIterar[K comparable, V any](iter *iterDiccionario[K, V]) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

// buscarPrimero, solamente se llegaría a ejecutar cuando invocas el primer iterador
// y te ubica el indice en el primer elemento del dicc
func buscarPrimero[K comparable, V any](iter *iterDiccionario[K, V]) {
	for iter.indice < iter.hash.tam && iter.hash.tabla[iter.indice].estado != _OCUPADO {
		iter.indice++
	}
}

func crearIteradorDiccionario[K comparable, V any](hash *hashCerrado[K, V]) *iterDiccionario[K, V] {
	iter := new(iterDiccionario[K, V])
	iter.hash = hash
	iter.actual = 0
	iter.indice = 0
	if iter.hash.tabla[iter.indice].estado != _OCUPADO {
		buscarPrimero(iter)
	}
	return iter
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	// solamente necesito iterar hasta que haya pasado por la todos los elementos ocupados del dicc
	return iter.actual < iter.hash.cantidad
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	panicIteradorTerminoDeIterar(iter)
	// acá busco el índice, y sumo el "actual" una única vez, da igual si no encontró elemento, si sucede esto es que el dicc terminó.
	iter.indice++
	for iter.indice < iter.hash.tam && iter.hash.tabla[iter.indice].estado != _OCUPADO {
		iter.indice++
	}
	iter.actual++
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	panicIteradorTerminoDeIterar(iter)
	return iter.hash.tabla[iter.indice].clave, iter.hash.tabla[iter.indice].dato
}
