package lista_test

import (
	"testing"

	TDALista "tdas/lista"

	"github.com/stretchr/testify/require"
)

const (
	CANTIDAD_ELEMENTOS_VOLUMEN      = 10000
	CANTIDAD_VARIOS_ELEMENTOS_INT   = 10
	CANTIDAD_VARIOS_ELEMENTOS_FLOAT = 10.0
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}

type animal struct {
	tipo  string
	patas int
}

func TestListaInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[animal]()
	lista.InsertarPrimero(animal{tipo: "perro", patas: 4})
	require.False(t, lista.EstaVacia())
	require.Equal(t, animal{tipo: "perro", patas: 4}, lista.VerPrimero())
	require.Equal(t, 1, lista.Largo())
}

func TestListaInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	lista.InsertarUltimo(25.99)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 25.99, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

func TestListaBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("Algoritmos y Estructuras de Datos")
	require.Equal(t, "Algoritmos y Estructuras de Datos", lista.BorrarPrimero())
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
	lista := TDALista.CrearListaEnlazada[animal]()
	lista.InsertarUltimo(animal{tipo: "perro", patas: 4})
	require.Equal(t, animal{tipo: "perro", patas: 4}, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo(animal{tipo: "gato", patas: 4})
	require.Equal(t, animal{tipo: "gato", patas: 4}, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	lista.InsertarUltimo(animal{tipo: "pajaro", patas: 2})
	require.Equal(t, animal{tipo: "pajaro", patas: 2}, lista.VerUltimo())
	require.Equal(t, animal{tipo: "perro", patas: 4}, lista.VerPrimero())
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
	lista := TDALista.CrearListaEnlazada[[]int]()
	lista.InsertarPrimero([]int{1, 2, 3})
	lista.InsertarPrimero([]int{4, 5, 6})
	lista.InsertarPrimero([]int{7, 8, 9})
	require.Equal(t, []int{7, 8, 9}, lista.VerPrimero())
	require.Equal(t, []int{1, 2, 3}, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, []int{7, 8, 9}, lista.BorrarPrimero())
	require.Equal(t, []int{4, 5, 6}, lista.VerPrimero())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, []int{4, 5, 6}, lista.BorrarPrimero())
	require.Equal(t, []int{1, 2, 3}, lista.VerPrimero())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, []int{1, 2, 3}, lista.BorrarPrimero())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo())
}

func TestListaInsertarPrimeroMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < CANTIDAD_ELEMENTOS_VOLUMEN; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())

		require.Equal(t, i+1, lista.Largo())
	}
	require.Equal(t, 0, lista.VerUltimo())
}

func TestListaInsertarUltimoMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < CANTIDAD_ELEMENTOS_VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
	}
	require.Equal(t, 0, lista.VerPrimero())
}

func TestListaBorrarPrimeroMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < CANTIDAD_ELEMENTOS_VOLUMEN; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, i+1, lista.Largo())
	}

	for i := CANTIDAD_ELEMENTOS_VOLUMEN - 1; !lista.EstaVacia(); i-- {
		require.Equal(t, i, lista.BorrarPrimero())
		require.Equal(t, i, lista.Largo())
	}
}

func TestListaIterarTrue(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < CANTIDAD_VARIOS_ELEMENTOS_INT; i++ {
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
	//termina cuando encuentra al primer animal sin patas
	lista := TDALista.CrearListaEnlazada[animal]()
	animales := []animal{
		{tipo: "vaca", patas: 4},
		{tipo: "paloma", patas: 2},
		{tipo: "halcon", patas: 2},
		{tipo: "delfin", patas: 0},
		{tipo: "elefante", patas: 4},
	}
	for _, animal := range animales {
		lista.InsertarUltimo(animal)
	}

	var animalesConPatas []string
	lista.Iterar(func(a animal) bool {
		if a.patas != 0 {
			animalesConPatas = append(animalesConPatas, a.tipo)
		}
		return a.patas != 0
	})

	require.Equal(t, []string{"vaca", "paloma", "halcon"}, animalesConPatas)
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

func TestListaIterardorInsertarDespuesDeCrear(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	iterador := lista.Iterador()
	iterador.Insertar(3.14)

	require.Equal(t, 3.14, iterador.VerActual())
	require.Equal(t, true, iterador.HayAlgoMas())
	require.Equal(t, 3.14, lista.VerPrimero())
	require.Equal(t, 3.14, lista.VerUltimo())
	require.Equal(t, false, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
}

func TestListaIteradorInsertarAlFinal(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[string]()
	ingenierias := []string{"Informatica", "Civil", "Naval", "Mecanica"}

	for _, ingenieria := range ingenierias {
		lista.InsertarUltimo(ingenieria)
	}

	iterador := lista.Iterador()

	for iterador.HayAlgoMas() {
		iterador.Avanzar()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Avanzar() })

	iterador.Insertar("Quimica")

	require.Equal(t, "Quimica", lista.VerUltimo())
	require.Equal(t, "Quimica", iterador.VerActual())
}

func TestListaIteradorInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < CANTIDAD_VARIOS_ELEMENTOS_INT; i++ {
		lista.InsertarUltimo(i)
	}

	for iterador := lista.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		if iterador.VerActual() == lista.Largo()/2 {
			iterador.Insertar(20)
			break
		}
	}
	var resultado []int

	for iterador := lista.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		resultado = append(resultado, iterador.VerActual())
	}

	require.Equal(t, []int{0, 1, 2, 3, 4, 20, 5, 6, 7, 8, 9}, resultado)
}

func TestListaIteradorBorrarDespuesDeCrear(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[animal]()

	animales := []animal{
		{tipo: "vaca", patas: 4},
		{tipo: "paloma", patas: 2},
		{tipo: "halcon", patas: 2},
		{tipo: "delfin", patas: 0},
		{tipo: "elefante", patas: 4},
	}

	for _, animal := range animales {
		lista.InsertarPrimero(animal)
	}

	iterador := lista.Iterador()

	iterador.Borrar()

	require.Equal(t, animal{tipo: "delfin", patas: 0}, lista.VerPrimero())

}
func TestListaIteradorBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()

	for i := 0.0; i < CANTIDAD_VARIOS_ELEMENTOS_FLOAT; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()

	for i := 0; i < lista.Largo()-1; i++ {
		iterador.Avanzar()
	}
	require.Equal(t, 9.0, lista.VerUltimo())

	iterador.Borrar()

	require.Equal(t, 8.0, lista.VerUltimo())

}

func TestListaIteradorBorrarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < CANTIDAD_VARIOS_ELEMENTOS_INT; i++ {
		lista.InsertarUltimo(i)
	}

	for iterador := lista.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		if iterador.VerActual() == lista.Largo()/2 {
			iterador.Borrar()
			break
		}
	}
	var resultado []int

	for iterador := lista.Iterador(); iterador.HayAlgoMas(); iterador.Avanzar() {
		resultado = append(resultado, iterador.VerActual())
	}

	require.Equal(t, []int{0, 1, 2, 3, 4, 6, 7, 8, 9}, resultado)
}

func TestListaIteradorInsertarDespuesDeVaciar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()

	iterador := lista.Iterador()

	iterador.Insertar(1.1)
	iterador.Insertar(1.4)
	iterador.Insertar(1.5)

	require.Equal(t, 1.5, lista.VerPrimero())

	require.Equal(t, 1.5, iterador.Borrar())

	require.Equal(t, 1.4, lista.VerPrimero())

	iterador.Insertar(1.7)
	require.Equal(t, 1.7, iterador.VerActual())

	iterador.Borrar()
	iterador.Borrar()
	iterador.Borrar()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.True(t, lista.EstaVacia())

	iterador.Insertar(8.5)
	iterador.Insertar(8.6)
	iterador.Insertar(5.0)

	require.Equal(t, 5.0, lista.VerPrimero())
	require.Equal(t, 8.5, lista.VerUltimo())
	require.Equal(t, 5.0, iterador.VerActual())

}
