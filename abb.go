package diccionario

import (
	TDAPila "tdas/pila"
)

type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz *nodoAbb[K, V]
	cant int
	cmp  funcCmp[K]
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = dato
	return nodo
}

func CrearAbb[K comparable, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cmp = funcionCmp
	return abb
}

func panicAbb[K comparable, V any](nodo *nodoAbb[K, V]) {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
}

func buscarNodoAbb[K comparable, V any](nodo **nodoAbb[K, V], funcionCmp func(K, K) int, clave K) **nodoAbb[K, V] {
	if *nodo == nil {
		return nodo
	}
	if funcionCmp(clave, (*nodo).clave) > 0 {
		return buscarNodoAbb(&(*nodo).der, funcionCmp, clave)
	} else if funcionCmp(clave, (*nodo).clave) < 0 {
		return buscarNodoAbb(&(*nodo).izq, funcionCmp, clave)
	} else {
		return nodo
	}
}

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	nodoGuardar := buscarNodoAbb(&(abb.raiz), abb.cmp, clave)
	if *nodoGuardar == nil {
		nuevoNodo := crearNodoAbb(clave, valor)
		*nodoGuardar = nuevoNodo
		abb.cant++
	} else {
		(*nodoGuardar).dato = valor
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return *buscarNodoAbb(&abb.raiz, abb.cmp, clave) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := *buscarNodoAbb(&abb.raiz, abb.cmp, clave)
	panicAbb(nodo)
	return nodo.dato
}
func (abb *abb[K, V]) Borrar(clave K) V {
	nodo := buscarNodoAbb(&abb.raiz, abb.cmp, clave)
	panicAbb(*nodo)
	dato := (*nodo).dato
	borrarAbb(nodo)
	abb.cant--
	return dato
}

func buscarMax[K comparable, V any](nodoPadre **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodoPadre).der == nil {
		return nodoPadre
	}
	return buscarMax(&(*nodoPadre).der)
}

func borrarAbb[K comparable, V any](nodoBorrar **nodoAbb[K, V]) {
	if (*nodoBorrar).izq == nil {
		*nodoBorrar = (*nodoBorrar).der
	} else if (*nodoBorrar).der == nil {
		*nodoBorrar = (*nodoBorrar).izq
	} else {
		nodoMax := buscarMax(&(*nodoBorrar).izq)
		(*nodoBorrar).clave = (*nodoMax).clave
		(*nodoBorrar).dato = (*nodoMax).dato
		borrarAbb(nodoMax)
	}

}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cant
}

type iterAbb[K comparable, V any] struct {
	abb  *abb[K, V]
	pila TDAPila.Pila[*nodoAbb[K, V]]
}

func crearIteradorAbb[K comparable, V any](abb *abb[K, V], desde *K, hasta *K) *iterAbb[K, V] {
	iter := new(iterAbb[K, V])
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.abb = abb
	iter.pila = pila
	iter.apilarHastaPrimero(iter.abb.raiz, desde, hasta)
	return iter
}

func panicIteradorTerminoDeIterar2[K comparable, V any](iter *iterAbb[K, V]) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iter *iterAbb[K, V]) apilarHastaPrimero(nodo *nodoAbb[K, V], desde, hasta *K) {
	if nodo == nil {
		return
	}
	if desde == nil {
		iter.pila.Apilar(nodo)
		iter.apilarHastaPrimero(nodo.izq, desde, hasta)
		return
	}
	if iter.abb.cmp(nodo.clave, *desde) >= 0 {
		iter.pila.Apilar(nodo)
		iter.apilarHastaPrimero(nodo.izq, desde, hasta)
	}
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIteradorAbb(abb, nil, nil)
}

func (iter *iterAbb[K, V]) Siguiente() {
	panicIteradorTerminoDeIterar2(iter)
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return iter.pila.EstaVacia()
}
func (iter *iterAbb[K, V]) VerActual() (K, V) {
	panicIteradorTerminoDeIterar2(iter)
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return crearIteradorAbb(abb, desde, hasta)
}

func (abb *abb[K, V]) iterar(nodo *nodoAbb[K, V], visitar func(clave K, dato V) bool, condicion bool, hasta *K) bool {
	if nodo == nil {
		return condicion
	}

	if hasta != nil && abb.cmp(nodo.clave, *hasta) > 0 {
		return condicion
	}

	condicion = abb.iterar(nodo.izq, visitar, condicion, hasta)

	if condicion {
		condicion = visitar(nodo.clave, nodo.dato)
	}

	if condicion {
		return abb.iterar(nodo.der, visitar, condicion, hasta)
	}

	return condicion
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if desde == nil {
		abb.iterar(abb.raiz, visitar, true, hasta)
		return
	}

	inicio := buscarNodoAbb(&abb.raiz, abb.cmp, *desde)
	abb.iterar(*inicio, visitar, true, hasta)
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.iterar(abb.raiz, visitar, true, nil)
}
