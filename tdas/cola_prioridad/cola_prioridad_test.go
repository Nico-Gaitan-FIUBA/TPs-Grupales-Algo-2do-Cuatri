package cola_prioridad_test

import (
	"testing"

	TDAColaPrioridad "tdas/cola_prioridad"

	"github.com/stretchr/testify/require"
)

func TestColaPrioridadVacia(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestColaPrioridadEncolarUnElemento(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})
	cola.Encolar(2026)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.Cantidad())
	require.Equal(t, 2026, cola.VerMax())
}

func TestColaPrioridadEncolarDosElementos(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})
	cola.Encolar(10)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 1, cola.Cantidad())
	require.Equal(t, 10, cola.VerMax())

	cola.Encolar(20)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 2, cola.Cantidad())
	require.Equal(t, 20, cola.VerMax())

}

func TestColaPrioridadEncolarVariosElementosOrdenados(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(5)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 5, cola.Cantidad())
	require.Equal(t, 5, cola.VerMax())

}

func TestColaPrioridadEncolarVariosElementosDesordenados(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})
	cola.Encolar(66)
	cola.Encolar(10)
	cola.Encolar(90)
	cola.Encolar(14)
	cola.Encolar(35)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 5, cola.Cantidad())
	require.Equal(t, 90, cola.VerMax())

}

func TestColaPrioridadDesencolarUnElemento(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})
	cola.Encolar(10)
	require.Equal(t, 10, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestColaPrioridadDesencolarDosElementos(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})
	cola.Encolar(11)
	cola.Encolar(40)
	require.Equal(t, 40, cola.Desencolar())
	require.Equal(t, 11, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestColaPrioridadDesencolarVarioslementos(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})
	cola.Encolar(11)
	cola.Encolar(40)
	cola.Encolar(20)
	cola.Encolar(35)
	cola.Encolar(66)

	require.Equal(t, 66, cola.Desencolar())
	require.Equal(t, 40, cola.Desencolar())
	require.Equal(t, 35, cola.Desencolar())
	require.Equal(t, 20, cola.Desencolar())
	require.Equal(t, 11, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })

}

func TestColaPrioridadEncolarDespuesDeVaciar(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return b - a
	})

	cola.Encolar(10)
	cola.Encolar(20)
	cola.Encolar(0)
	require.Equal(t, 0, cola.Desencolar())
	require.Equal(t, 10, cola.Desencolar())
	require.Equal(t, 20, cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })

	cola.Encolar(5)
	cola.Encolar(14)
	cola.Encolar(8)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 3, cola.Cantidad())
	require.Equal(t, 5, cola.VerMax())

}

func TestColaPrioridadConElementosRepetidos(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return b - a
	})
	cola.Encolar(10)
	cola.Encolar(10)
	cola.Encolar(20)
	cola.Encolar(20)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 4, cola.Cantidad())
	require.Equal(t, 10, cola.VerMax())
}

func TestColaPrioridadConElementosRepetidosDesencolar(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return b - a
	})
	cola.Encolar(10)
	cola.Encolar(10)
	cola.Encolar(20)
	cola.Encolar(20)

	require.Equal(t, 10, cola.Desencolar())
	require.Equal(t, 10, cola.Desencolar())
	require.Equal(t, 20, cola.Desencolar())
	require.Equal(t, 20, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestColaPrioridadHeapifyArrVacio(t *testing.T) {
	arr := []int{}
	cola := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int {
		return a - b
	})

	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}
func TestColaPrioridadHeapify(t *testing.T) {
	arr := []int{66, 10, 90, 14, 35}
	cola := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int {
		return b - a
	})

	require.False(t, cola.EstaVacia())
	require.Equal(t, 5, cola.Cantidad())
	require.Equal(t, 10, cola.VerMax())
	require.Equal(t, 10, cola.Desencolar())
	cola.Encolar(5)
	require.Equal(t, 5, cola.VerMax())
}

func TestColaPrioridadHeapifyVaciar(t *testing.T) {
	arr := []int{66, 10, 90, 14, 35}
	cola := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int {
		return a - b
	})

	require.Equal(t, 90, cola.Desencolar())
	require.Equal(t, 66, cola.Desencolar())
	require.Equal(t, 35, cola.Desencolar())
	require.Equal(t, 14, cola.Desencolar())
	require.Equal(t, 10, cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })

}

func TestColaPrioridadHeapifyEncolarDespuesDeVaciar(t *testing.T) {
	arr := []int{66, 10, 90, 14, 35}
	cola := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int {
		return a - b
	})

	require.Equal(t, 90, cola.Desencolar())
	require.Equal(t, 66, cola.Desencolar())
	require.Equal(t, 35, cola.Desencolar())
	require.Equal(t, 14, cola.Desencolar())
	require.Equal(t, 10, cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })

	cola.Encolar(5)
	cola.Encolar(14)
	cola.Encolar(8)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 3, cola.Cantidad())
	require.Equal(t, 14, cola.VerMax())

}

func TestColaPrioridadHeapifyConElementosRepetidos(t *testing.T) {
	arr := []int{10, 10, 20, 20}
	cola := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int {
		return a - b
	})

	require.False(t, cola.EstaVacia())
	require.Equal(t, 4, cola.Cantidad())
	require.Equal(t, 20, cola.VerMax())
}

func TestColaPrioridadHeapSort(t *testing.T) {
	arr := []int{66, 10, 90, 14, 35}

	cmp := func(a, b int) int { return a - b }
	TDAColaPrioridad.HeapSort(arr, cmp)
	require.Equal(t, []int{10, 14, 35, 66, 90}, arr)
}

func TestColaPrioridadVolumen(t *testing.T) {
	cola := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return a - b
	})

	for i := 0; i < 100000; i++ {
		cola.Encolar(i)
	}

	require.False(t, cola.EstaVacia())
	require.Equal(t, 100000, cola.Cantidad())
	require.Equal(t, 99999, cola.VerMax())

	for i := 99999; i >= 0; i-- {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}

func TestColaPrioridadHeapifyVolumen(t *testing.T) {
	arr := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		arr[i] = i
	}
	cola := TDAColaPrioridad.CrearHeapArr(arr, func(a, b int) int {
		return a - b
	})

	require.False(t, cola.EstaVacia())
	require.Equal(t, 100000, cola.Cantidad())
	require.Equal(t, 99999, cola.VerMax())

	for i := 99999; i >= 0; i-- {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.Equal(t, 0, cola.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerMax() })
}
