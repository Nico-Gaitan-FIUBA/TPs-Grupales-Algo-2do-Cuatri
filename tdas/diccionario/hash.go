package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	TAM_INICIAL      = 11
	CANTIDAD_INICIAL = 0
)

type parClaveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterDiccionario[K comparable, V any] struct {
	tabla       *hashAbierto[K, V]
	actualHash  int
	actualLista TDALista.IteradorLista[parClaveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		tabla:    make([]TDALista.Lista[parClaveValor[K, V]], TAM_INICIAL),
		tam:      TAM_INICIAL,
		cantidad: CANTIDAD_INICIAL,
	}
}

func jenkinsHash(clave string, largo int) uint32 {
	var hash uint32 = 0

	for i := 0; i < len(clave); i++ {
		hash += uint32(clave[i])
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return hash % uint32(largo)
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hashAbierto[K, V]) buscarParClaveValor(clave K) *parClaveValor[K, V] {
	claveHasheada := jenkinsHash(string(convertirABytes(clave)), h.tam)
	iterador := h.tabla[claveHasheada].Iterador()
	for iterador.HayAlgoMas() {
		actual := iterador.VerActual()
		if actual.clave == clave {
			return &actual
		}
	}
	return nil
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	if h.Pertenece(clave) {
		puntero := h.buscarParClaveValor(clave)
		puntero.valor = dato
	} else {
		claveHasheada := jenkinsHash(string(convertirABytes(clave)), h.tam)
		h.tabla[claveHasheada].InsertarUltimo(parClaveValor[K, V]{clave: clave, valor: dato})
		h.cantidad++
		// if //algo {
		// 	h.redimensionar()
		// }
	}
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	return h.buscarParClaveValor(clave) != nil
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	return
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	return
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Iterar(func(clave K, dato V) bool) {
	return
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	return
}
