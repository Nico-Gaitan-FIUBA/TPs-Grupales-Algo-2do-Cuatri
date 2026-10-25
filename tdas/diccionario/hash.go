package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	TAM_INICIAL            = 11
	FACTOR_DE_CARGA_MAXIMO = 3
	OCUPACION_MINIMA       = 0.25
	FACTOR_DE_REDIMENSION  = 2
)

type parClaveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	cantidad int
}

type iterDiccionario[K comparable, V any] struct {
	tabla           *hashAbierto[K, V]
	posTabla        int
	iterListaActual TDALista.IteradorLista[parClaveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		tabla:    make([]TDALista.Lista[parClaveValor[K, V]], TAM_INICIAL),
		cantidad: 0,
	}
}

func (h *hashAbierto[K, V]) crearIteradorExterno() IterDiccionario[K, V] {
	iterador := &iterDiccionario[K, V]{tabla: h}
	iterador.buscarListaValida()
	return iterador
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

func (h *hashAbierto[K, V]) buscarClaveYLista(clave K) (int, TDALista.Lista[parClaveValor[K, V]]) {
	claveHasheada := jenkinsHash(string(convertirABytes(clave)), len(h.tabla))
	return int(claveHasheada), h.tabla[int(claveHasheada)]
}

func (h *hashAbierto[K, V]) existeLaLista(lista TDALista.Lista[parClaveValor[K, V]]) bool {
	return lista != nil
}

func (h *hashAbierto[K, V]) actualizarValor(clave K, dato V, lista TDALista.Lista[parClaveValor[K, V]]) {
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			iter.Borrar()
			iter.Insertar(parClaveValor[K, V]{clave: clave, valor: dato})
		}
	}
}

func (h *hashAbierto[K, V]) redimensionar(tamanio int) {
	nuevaTabla := make([]TDALista.Lista[parClaveValor[K, V]], tamanio)
	for _, lista := range h.tabla {
		if h.existeLaLista(lista) {
			for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
				elem := iter.VerActual()
				claveHasheada := jenkinsHash(string(convertirABytes(elem.clave)), tamanio)
				if nuevaTabla[claveHasheada] == nil {
					nuevaTabla[claveHasheada] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
				}
				nuevaTabla[claveHasheada].InsertarUltimo(elem)
			}
		}
	}
	h.tabla = nuevaTabla
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	claveHasheada, lista := h.buscarClaveYLista(clave)

	if !h.existeLaLista(lista) {
		lista = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
		h.tabla[claveHasheada] = lista
	}

	if h.Pertenece(clave) {
		h.actualizarValor(clave, dato, lista)
		return
	} else {
		lista.InsertarUltimo(parClaveValor[K, V]{clave: clave, valor: dato})
		h.cantidad++
	}

	if float64(FACTOR_DE_CARGA_MAXIMO) <= float64(h.cantidad)/float64(len(h.tabla)) {
		h.redimensionar(len(h.tabla) * FACTOR_DE_REDIMENSION)
	}
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, lista := h.buscarClaveYLista(clave)

	if !h.existeLaLista(lista) {
		return false
	}
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			return true
		}
	}
	return false
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	_, lista := h.buscarClaveYLista(clave)

	if !h.existeLaLista(lista) {
		panic("La clave no pertenece al diccionario")
	}
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			return iter.VerActual().valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	_, lista := h.buscarClaveYLista(clave)
	if !h.existeLaLista(lista) {
		panic("La clave no pertenece al diccionario")
	}
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		if iter.VerActual().clave == clave {
			borrado := iter.VerActual().valor
			iter.Borrar()
			h.cantidad--
			if float64(h.cantidad) <= float64(len(h.tabla))*OCUPACION_MINIMA && len(h.tabla)/FACTOR_DE_REDIMENSION >= TAM_INICIAL {
				h.redimensionar(len(h.tabla) / FACTOR_DE_REDIMENSION)
			}
			return borrado
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	debeContinuar := true
	for i := 0; i < len(h.tabla) && debeContinuar; i++ {
		lista := h.tabla[i]
		if h.existeLaLista(lista) {
			lista.Iterar(func(par parClaveValor[K, V]) bool {
				if !visitar(par.clave, par.valor) {
					debeContinuar = false
				}
				return debeContinuar
			})
		}
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	return h.crearIteradorExterno()
}

func (i *iterDiccionario[K, V]) HayAlgoMas() bool {
	return i.posTabla < len(i.tabla.tabla)
}

func (i *iterDiccionario[K, V]) VerActual() (K, V) {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return i.iterListaActual.VerActual().clave, i.iterListaActual.VerActual().valor
}

func (i *iterDiccionario[K, V]) Avanzar() {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	i.iterListaActual.Avanzar()
	if !i.iterListaActual.HayAlgoMas() {
		i.posTabla++
		i.buscarListaValida()
	}

}

func (iter *iterDiccionario[K, V]) buscarListaValida() {

	for iter.posTabla < len(iter.tabla.tabla) {

		if iter.tabla.existeLaLista(iter.tabla.tabla[iter.posTabla]) {
			iterador := iter.tabla.tabla[iter.posTabla].Iterador()

			if iterador.HayAlgoMas() {
				iter.iterListaActual = iterador
				break
			}
		}

		iter.posTabla++
		if iter.posTabla >= len(iter.tabla.tabla) {
			break
		}

	}

}
