package diccionario

import (
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
