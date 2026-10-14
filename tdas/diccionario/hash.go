package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	TAM_INICIAL            = 11
	CANTIDAD_INICIAL       = 0
	FACTOR_DE_CARGA_MINIMO = 2
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
	tam      int
	cantidad int
}

type iterDiccionario[K comparable, V any] struct {
	tabla       *hashAbierto[K, V]
	actualHash  int
	listaActual TDALista.IteradorLista[parClaveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		tabla:    make([]TDALista.Lista[parClaveValor[K, V]], TAM_INICIAL),
		tam:      TAM_INICIAL,
		cantidad: CANTIDAD_INICIAL,
	}
}

func (h *hashAbierto[K, V]) crearIteradorExterno() IterDiccionario[K, V] {
	iterador := &iterDiccionario[K, V]{tabla: h}
	for _, lista := range h.tabla {
		if !h.existeLaLista(lista) {
			iterador.actualHash++
			if iterador.actualHash >= h.tam {
				iterador.actualHash = h.tam
				break
			}
		} else {
			iterador.listaActual = h.tabla[iterador.actualHash].Iterador()
			if iterador.listaActual.HayAlgoMas() {
				break
			}
			iterador.actualHash++
		}
	}
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
	claveHasheada := jenkinsHash(string(convertirABytes(clave)), h.tam)
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
	h.tam = tamanio
}

func (h *hashAbierto[K, V]) buscarActualHash() int {
	for pos, lista := range h.tabla {
		if h.existeLaLista(lista) {
			return pos
		}
	}
	return -1
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
	condicionParaRedimension := float64(FACTOR_DE_CARGA_MINIMO) <= float64(h.cantidad)/float64(h.tam) &&
		float64(FACTOR_DE_CARGA_MAXIMO) >= float64(h.cantidad)/float64(h.tam)

	if condicionParaRedimension {
		h.redimensionar(h.tam * FACTOR_DE_REDIMENSION)
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
			condicionParaRedimension := float64(h.cantidad) == float64(h.tam)*OCUPACION_MINIMA && h.tam/FACTOR_DE_REDIMENSION >= TAM_INICIAL
			if condicionParaRedimension {
				h.redimensionar(h.tam * FACTOR_DE_REDIMENSION)
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
	for i := 0; i < h.tam && debeContinuar; i++ {
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
	return i.actualHash < i.tabla.tam
}

func (i *iterDiccionario[K, V]) VerActual() (K, V) {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return i.listaActual.VerActual().clave, i.listaActual.VerActual().valor
}

func (i *iterDiccionario[K, V]) Avanzar() {
	if !i.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	i.listaActual.Avanzar()

	for !i.listaActual.HayAlgoMas() {
		i.actualHash++
		if i.actualHash >= i.tabla.tam {
			break
		}
		lista := i.tabla.tabla[i.actualHash]
		if lista != nil {
			i.listaActual = lista.Iterador()
		}
	}
}
