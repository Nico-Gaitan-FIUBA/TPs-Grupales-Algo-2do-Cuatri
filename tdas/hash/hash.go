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
