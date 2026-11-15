package cola_prioridad

const (
	TAM_INICIAL        = 10
	FACTOR_REDIMENSION = 2
	FACTOR_REDUCCION   = 4
)

type heap[T any] struct {
	cola        []T
	cantidad    int
	funcion_cmp cmp[T]
}
type cmp[T any] func(T, T) int

// La función de comparación, recibe dos claves y devuelve:

//     Un entero menor que 0 si la primera clave es menor que la segunda. --> negativo si 1era < 2da
//     Un entero mayor que 0 si la primera clave es mayor que la segunda. --> positivo si 1era > 2da
//     0 si ambas claves son iguales.

func CrearHeap[T any](funcion_cmp cmp[T]) ColaPrioridad[T] {
	return &heap[T]{cola: make([]T, TAM_INICIAL), cantidad: 0, funcion_cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp cmp[T]) ColaPrioridad[T] {

	if len(arreglo) == 0 {
		return CrearHeap(funcion_cmp)
	}
	copia := make([]T, len(arreglo))
	copy(copia, arreglo)

	h := &heap[T]{cola: copia, cantidad: len(copia), funcion_cmp: funcion_cmp}

	heapify(h.cola[:h.cantidad], h.funcion_cmp)
	return h
}

func HeapSort[T any](elementos []T, funcion_cmp cmp[T]) {
	heapify(elementos, funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		swap(elementos, 0, i)
		downHeap(elementos[:i], funcion_cmp, 0)
	}
}

func heapify[T any](cola []T, funcion_cmp cmp[T]) {
	for i := len(cola)/2 - 1; i >= 0; i-- {
		downHeap(cola, funcion_cmp, i)
	}
}

func calcularPosHijoIzquierdo(indice int) int {
	return 2*indice + 1
}

func calcularPosHijoDerecho(indice int) int {
	return 2*indice + 2
}

func calcularPosPadre(indice int) int {
	return (indice - 1) / 2
}

func swap[T any](cola []T, i, j int) {
	cola[i], cola[j] = cola[j], cola[i]
}

func upHeap[T any](cola []T, cmp cmp[T], elem T, pos int) {
	posPadre := calcularPosPadre(pos)
	if cmp(elem, cola[posPadre]) <= 0 {
		return
	}
	swap(cola, pos, posPadre)
	upHeap(cola, cmp, elem, posPadre)
}
func downHeap[T any](cola []T, cmp cmp[T], pos int) {
	posHijoIzquierdo := calcularPosHijoIzquierdo(pos)
	posHijoDerecho := calcularPosHijoDerecho(pos)
	posMayor := pos

	if posHijoIzquierdo < len(cola) && cmp(cola[posMayor], cola[posHijoIzquierdo]) < 0 {
		posMayor = posHijoIzquierdo
	}

	if posHijoDerecho < len(cola) && cmp(cola[posMayor], cola[posHijoDerecho]) < 0 {
		posMayor = posHijoDerecho
	}

	if posMayor != pos {
		swap(cola, posMayor, pos)
		downHeap(cola, cmp, posMayor)
	}
}

func (h *heap[T]) redimensionar(tam int) {
	nuevoSlice := make([]T, tam)
	copy(nuevoSlice, h.cola[:h.cantidad])
	h.cola = nuevoSlice
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == len(h.cola) {
		nuevoTam := len(h.cola) * FACTOR_REDIMENSION
		h.redimensionar(nuevoTam)
	}

	h.cola[h.cantidad] = elem
	h.cantidad++

	upHeap(h.cola[:h.cantidad], h.funcion_cmp, elem, h.cantidad-1)

}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.cola[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}

	elemDesencolado := h.cola[0]
	h.cantidad--

	swap(h.cola, 0, h.cantidad)

	downHeap(h.cola[:h.cantidad], h.funcion_cmp, 0)

	if h.cantidad <= len(h.cola)/FACTOR_REDUCCION && len(h.cola) > TAM_INICIAL {
		h.redimensionar(len(h.cola) / FACTOR_REDIMENSION)
	}

	return elemDesencolado
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}
