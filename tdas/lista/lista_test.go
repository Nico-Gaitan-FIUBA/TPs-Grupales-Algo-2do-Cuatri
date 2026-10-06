package lista_test

import (
	"testing"

	TDALista "tdas/lista"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}

func TestListaInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 1, lista.Largo())
}

func TestListaInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

func TestListaBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	require.Equal(t, 10, lista.BorrarPrimero())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}

func TestListaInsertarPrimeroPocosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarPrimero(20)
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 2, lista.Largo())
	lista.InsertarPrimero(30)
	require.Equal(t, 30, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
}

func TestListaInsertarUltimoPocosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo(20)
	require.Equal(t, 20, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	lista.InsertarUltimo(30)
	require.Equal(t, 30, lista.VerUltimo())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 3, lista.Largo())
}

func TestListaInsertarAlternado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 10, lista.VerPrimero())
	lista.InsertarPrimero(20)
	require.Equal(t, 20, lista.VerPrimero())
	lista.InsertarUltimo(30)
	require.Equal(t, 30, lista.VerUltimo())
	lista.InsertarPrimero(40)
	require.Equal(t, 40, lista.VerPrimero())
}

func TestListaBorrarPrimeroPocosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(20)
	lista.InsertarPrimero(30)
	require.Equal(t, 30, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 30, lista.BorrarPrimero())
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 20, lista.BorrarPrimero())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 10, lista.BorrarPrimero())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}

func TestListaInsertarPrimeroMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())

		require.Equal(t, i+1, lista.Largo())
	}
	require.Equal(t, 0, lista.VerUltimo())
}

func TestListaInsertarUltimoMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
	}
	require.Equal(t, 0, lista.VerPrimero())
}

func TestListaBorrarPrimeroMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, i+1, lista.Largo())
	}

	for i := 9999; !lista.EstaVacia(); i-- {
		require.Equal(t, i, lista.BorrarPrimero())
		require.Equal(t, i, lista.Largo())
	}
}

func TestListaIterarTrue(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}

	var resultado []int
	lista.Iterar(func(valor int) bool {
		resultado = append(resultado, valor)
		return true
	})

	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, resultado)
}

func TestListaIterarFalse(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}

	var resultado []int
	lista.Iterar(func(valor int) bool {
		resultado = append(resultado, valor)
		return valor != 5
	})

	require.Equal(t, []int{0, 1, 2, 3, 4, 5}, resultado)
}

func TestListaIteradorVerActual(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	iterador := lista.Iterador()
	require.Equal(t, 10, iterador.VerActual())
	iterador.Avanzar()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
}

func TestListaIteradorHayAlgoMas(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	iterador := lista.Iterador()
	require.Equal(t, true, iterador.HayAlgoMas())
	iterador.Avanzar()
	require.Equal(t, false, iterador.HayAlgoMas())
}

func TestListaIteradorAvanzar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	iterador := lista.Iterador()
	iterador.Avanzar()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Avanzar() })
}

func TestListaIteradorInsertar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(10)
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 10, iterador.VerActual())
	require.Equal(t, 1, lista.Largo())
	iterador.Insertar(20)
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 20, iterador.VerActual())
}

func TestListaIteradorBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(10)
	require.Equal(t, 10, iterador.Borrar())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
}

func TestListaIteradorInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}

	for iterador := lista.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		if iterador.VerActual() == 4 {
			iterador.Insertar(20)
			break
		}
		iterador.Avanzar()
	}
	var resultado []int

	for iterador := lista.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		resultado = append(resultado, iterador.VerActual())
	}

	require.Equal(t, []int{0, 1, 2, 3, 20, 4, 5, 6, 7, 8, 9}, resultado)
}

func TestListaIteradorBorrarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}

	for iterador := lista.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		if iterador.VerActual() == 4 {
			iterador.Borrar()
			break
		}
		iterador.Avanzar()
	}
	var resultado []int

	for iterador := lista.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		resultado = append(resultado, iterador.VerActual())
	}

	require.Equal(t, []int{0, 1, 2, 3, 5, 6, 7, 8, 9}, resultado)
}
