package diccionario

import (
	"fmt"
	"hash/fnv"
	"math"
)

type _ESTADO int

const (
	_VACIO _ESTADO = iota
	_OCUPADO
	_BORRADO
	_TAM_INICIAL       = 23
	_CRITERIO_AGRANDAR = 0.7
	_CRITERIO_ACHICAR  = 0.3
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado _ESTADO
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

func crearCeldaHash[K comparable, V any](clave K, dato V) celdaHash[K, V] {
	return celdaHash[K, V]{clave: clave, dato: dato, estado: _OCUPADO}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashing[K comparable](clave K, tam int) int {
	hash := fnv.New32()
	hash.Write(convertirABytes(clave))
	return int(hash.Sum32()) % tam

}

func esPrimo(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	//Solo verificamos hasta la raiz cuadrada del numero
	limite := int(math.Sqrt(float64(n)))
	for i := 3; i <= limite; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func obtenerPrimoSiguiente(n int) int {
	for !esPrimo(n) {
		n++
	}
	return n
}

func buscar[K comparable, V any](tabla []celdaHash[K, V], tam int, clave K) (int, bool) {
	indice := hashing(clave, tam)
	parDatoValor := tabla[indice]

	if parDatoValor.clave == clave && parDatoValor.estado == _OCUPADO {
		return indice, true
	}

	i := 1

	for parDatoValor.estado != _VACIO {
		indice = (indice + i) % tam
		parDatoValor = tabla[indice]
		if parDatoValor.clave == clave && parDatoValor.estado == _OCUPADO {
			return indice, true
		}
		i++
	}

	return indice, false
}

func panicClaveNoPertenece() {
	panic("La clave no pertenece al diccionario")
}

func crearTabla[K comparable, V any](tam int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], tam)
}

func redimensionar[K comparable, V any](hash *hashCerrado[K, V], tamnuevo int) {
	tablaVieja := hash.tabla
	tamnuevo = obtenerPrimoSiguiente(tamnuevo)
	hash.tabla = crearTabla[K, V](tamnuevo)
	hash.cantidad = 0
	hash.tam = tamnuevo
	hash.borrados = 0
	for _, celda := range tablaVieja {
		if celda.estado == _OCUPADO {
			hash.Guardar(celda.clave, celda.dato)
		}
	}
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashCerrado[K, V])
	hash.tabla = crearTabla[K, V](_TAM_INICIAL)
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
		hash.tabla[indice] = crearCeldaHash(clave, dato)
		if float64(hash.cantidad+hash.borrados)/float64(hash.tam) >= _CRITERIO_AGRANDAR {
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
	indice, encontrado := buscar(hash.tabla, hash.tam, clave)
	if !encontrado {
		panicClaveNoPertenece()
	}
	return hash.tabla[indice].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	if float64(hash.cantidad+hash.borrados)/float64(hash.tam) < _CRITERIO_ACHICAR && hash.tam/2 > _TAM_INICIAL {
		redimensionar(hash, hash.tam/2)
	}
	indice, encontrado := buscar(hash.tabla, hash.tam, clave)
	if !encontrado {
		panicClaveNoPertenece()
	}
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
	iterador := crearIteradorDiccionario[K, V](hash)
	return iterador
}

type iterDiccionario[K comparable, V any] struct {
	hash             *hashCerrado[K, V]
	contadorIterados int
	indice           int
}

func crearIteradorDiccionario[K comparable, V any](hash *hashCerrado[K, V]) *iterDiccionario[K, V] {
	iter := new(iterDiccionario[K, V])
	iter.hash = hash
	iter.contadorIterados = 0
	iter.indice = 0
	if iter.hash.tabla[iter.indice].estado != _OCUPADO {
		buscarSig(iter)
	}
	return iter
}

func panicIteradorTerminoDeIterar[K comparable, V any](iter *iterDiccionario[K, V]) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func buscarSig[K comparable, V any](iter *iterDiccionario[K, V]) {
	for iter.indice < iter.hash.tam && iter.hash.tabla[iter.indice].estado != _OCUPADO {
		iter.indice++
	}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return iter.contadorIterados < iter.hash.cantidad
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	panicIteradorTerminoDeIterar(iter)
	iter.indice++
	buscarSig(iter)
	iter.contadorIterados++
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	panicIteradorTerminoDeIterar(iter)
	return iter.hash.tabla[iter.indice].clave, iter.hash.tabla[iter.indice].dato
}
