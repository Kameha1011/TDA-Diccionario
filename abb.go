package diccionario

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
	abb *abb[K, V]
}

func crearIteradorAbb[K comparable, V any](abb *abb[K, V]) *iterAbb[K, V] {
	iter := new(iterAbb[K, V])
	iter.abb = abb
	return iter

}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIteradorAbb(abb)
}

func (iter *iterAbb[K, V]) Siguiente() {
	return
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return iter.HaySiguiente()
}
func (iter *iterAbb[K, V]) VerActual() (K, V) {
	return iter.VerActual()
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return crearIteradorAbb(abb)
}

func iterar[K comparable, V any](nodo *nodoAbb[K, V], visitar func(clave K, dato V) bool, condicion bool) bool {
	if nodo == nil {
		return true
	}
	if condicion {
		condicion = iterar(nodo.izq, visitar, condicion)
		condicion = visitar(nodo.clave, nodo.dato)
		condicion = iterar(nodo.der, visitar, condicion)
	}
	return condicion
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	return
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterar(abb.raiz, visitar, true)
}
