package cola_prioridad

const (
	PRIMER_POS         = 0
	UN_ELEMENTO        = 1
	TAM_INICIAL        = 10
	FACTOR_REDIMENSION = 2
	FACTOR_REDUCCION   = 4
)

type heap[T any] struct {
	arr         []T
	cantidad    int
	funcion_cmp func(T, T) int
}

// La función de comparación, recibe dos claves y devuelve:

//     Un entero menor que 0 si la primera clave es menor que la segunda. --> negativo si 1era < 2da
//     Un entero mayor que 0 si la primera clave es mayor que la segunda. --> positivo si 1era > 2da
//     0 si ambas claves son iguales.

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{arr: make([]T, TAM_INICIAL), cantidad: 0, funcion_cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	copia := make([]T, len(arreglo))
	copy(copia, arreglo)

	h := &heap[T]{arr: copia, cantidad: len(copia), funcion_cmp: funcion_cmp}

	heapify(h.arr[:h.cantidad], h.funcion_cmp)
	return h
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		swap(elementos, 0, i)
		downHeap(elementos[:i], funcion_cmp, elementos[0], 0)
	}
}

func heapify[T any](arr []T, funcion_cmp func(T, T) int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		downHeap(arr, funcion_cmp, arr[i], i)
	}
}

func buscoPosHijoIzquierdo(indice int) int {
	posHijoIzq := 2*indice + 1
	return posHijoIzq
}

func buscoPosHijoDerecho(indice int) int {
	posHijoDer := 2*indice + 2
	return posHijoDer
}

func buscoPosPadre[T any](arr []T, indice int) (T, int) {
	posPadre := (indice - 1) / 2
	padre := arr[posPadre]
	return padre, posPadre
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func upHeap[T any](arr []T, cmp func(T, T) int, elem T, pos int) {
	padre, posPadre := buscoPosPadre(arr, pos)
	if cmp(elem, padre) <= 0 {
		return
	}
	swap(arr, pos, posPadre)
	upHeap(arr, cmp, elem, posPadre)
}

func downHeap[T any](arr []T, cmp func(T, T) int, elem T, pos int) {
	posHijoIzquierdo := buscoPosHijoIzquierdo(pos)
	posHijoDerecho := buscoPosHijoDerecho(pos)
	posHijoMayor := posHijoDerecho
	posDelMayor := pos
	hijoMayor := elem

	switch {
	case posHijoIzquierdo >= len(arr):
		return
	case posHijoDerecho >= len(arr):
		hijoIzquierdo := arr[posHijoIzquierdo]
		hijoMayor = hijoIzquierdo
		posHijoMayor = posHijoIzquierdo
	default:
		hijoIzquierdo := arr[posHijoIzquierdo]
		hijoDerecho := arr[posHijoDerecho]
		if cmp(hijoIzquierdo, hijoDerecho) > 0 {
			hijoMayor = hijoIzquierdo
			posHijoMayor = posHijoIzquierdo
		} else {
			hijoMayor = hijoDerecho
		}
	}
	if cmp(hijoMayor, elem) > 0 {
		posDelMayor = posHijoMayor
		swap(arr, pos, posDelMayor)
		downHeap(arr, cmp, elem, posDelMayor)
	}
}

func (h *heap[T]) redimensionar(tam int) {
	nuevoSlice := make([]T, tam)
	copy(nuevoSlice, h.arr[:h.cantidad])
	h.arr = nuevoSlice
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == PRIMER_POS
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == len(h.arr) {
		nuevoTam := len(h.arr) * FACTOR_REDIMENSION
		if nuevoTam == 0 {
			nuevoTam = TAM_INICIAL
		}

		h.redimensionar(nuevoTam)
	}

	h.arr[h.cantidad] = elem
	h.cantidad++
	if h.cantidad > UN_ELEMENTO {
		upHeap(h.arr[:h.cantidad], h.funcion_cmp, elem, h.cantidad-1)
	}
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.arr[PRIMER_POS]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}

	elemDesencolado := h.arr[PRIMER_POS]
	h.cantidad--

	swap(h.arr, PRIMER_POS, h.cantidad)
	primerElem := h.arr[PRIMER_POS]

	if h.cantidad > UN_ELEMENTO {
		downHeap(h.arr[:h.cantidad], h.funcion_cmp, primerElem, PRIMER_POS)
	}

	if h.cantidad <= len(h.arr)/FACTOR_REDUCCION && len(h.arr) > TAM_INICIAL {
		h.redimensionar(len(h.arr) / FACTOR_REDIMENSION)
	}

	return elemDesencolado
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}
