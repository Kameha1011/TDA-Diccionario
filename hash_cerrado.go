package diccionario

type Estado int

const (
	_OCUPADO Estado = iota
	_VACIO
	_BORRADO
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

func CrearHash[K comparable, V any]() *hashCerrado[K, V] {
	return nil
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	return false
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	return hash.tabla[1].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	return hash.tabla[1].dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return 3
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
