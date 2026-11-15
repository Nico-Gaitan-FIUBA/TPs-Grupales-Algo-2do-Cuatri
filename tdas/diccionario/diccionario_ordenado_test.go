package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func comparoClaves[K int | string](clave1, clave2 K) int {
	if clave1 < clave2 {
		return -1
	} else if clave1 > clave2 {
		return 1
	}
	return 0
}

func insertarBalanceado(dic TDADiccionario.DiccionarioOrdenado[string, *int], claves []string, valores []int, inicio, fin int) {
	if inicio > fin {
		return
	}

	medio := (inicio + fin) / 2

	dic.Guardar(claves[medio], &valores[medio])

	insertarBalanceado(dic, claves, valores, inicio, medio-1)
	insertarBalanceado(dic, claves, valores, medio+1, fin)
}

func buscar2(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](comparoClaves[int])
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](comparoClaves[string])
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDate(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoDateHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	for i := 0; i < 500; i++ {
		dic.Guardar(fmt.Sprintf("clave%d", i), fmt.Sprintf("valor%d", i))
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(fmt.Sprintf("clave%d", i), fmt.Sprintf("nuevo_valor%d", i))
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(fmt.Sprintf("clave%d", i)) == fmt.Sprintf("nuevo_valor%d", i)
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutilizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumerics(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](comparoClaves[int])
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestWithKeysStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionario.CrearABB[avanzado, int](func(clave1, clave2 avanzado) int {
		return comparoClaves(clave1.w, clave2.w)
	})
	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 9, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 11, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestKeyEmpty(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValueNull(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](comparoClaves[string])
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestLongChainParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestGuardarYBorrarReiteradasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar " +
		"los borrados para agrandar en un Hash Cerrado. Si no se agranda, muy probablemente se quede en un ciclo " +
		"infinito")

	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func TestIterInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas con el iterador interno")
	dic := TDADiccionario.CrearABB[string, *int](comparoClaves[string])
	dic.Guardar("Vaca", nil)
	dic.Guardar("Gato", nil)
	dic.Guardar("Perro", nil)
	dic.Guardar("Elefante", nil)
	dic.Guardar("Panda", nil)

	var recorridas []string

	dic.Iterar(func(clave string, dato *int) bool {
		recorridas = append(recorridas, clave)
		return true
	})

	require.EqualValues(t, 5, len(recorridas))
	require.EqualValues(t, "Elefante", recorridas[0])
	require.EqualValues(t, "Gato", recorridas[1])
	require.EqualValues(t, "Panda", recorridas[2])
	require.EqualValues(t, "Perro", recorridas[3])
	require.EqualValues(t, "Vaca", recorridas[4])
}

func TestIterInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente (y una única vez) con el iterador interno")

	dic := TDADiccionario.CrearABB[string, int](comparoClaves[string])
	dic.Guardar("Gato", 6)
	dic.Guardar("Perro", 2)
	dic.Guardar("Vaca", 3)
	dic.Guardar("Burrito", 4)
	dic.Guardar("Hamster", 5)

	var valores []int

	dic.Iterar(func(_ string, dato int) bool {
		valores = append(valores, dato)
		return true
	})

	require.EqualValues(t, 5, len(valores))
	require.EqualValues(t, 4, valores[0])
	require.EqualValues(t, 6, valores[1])
	require.EqualValues(t, 5, valores[2])
	require.EqualValues(t, 2, valores[3])
	require.EqualValues(t, 3, valores[4])
}

func TestIterInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridos correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")

	dic := TDADiccionario.CrearABB[string, int](comparoClaves[string])
	dic.Guardar("Elefante", 7)
	dic.Guardar("Gato", 6)
	dic.Guardar("Perro", 2)
	dic.Guardar("Vaca", 3)
	dic.Guardar("Burrito", 4)
	dic.Guardar("Hamster", 5)

	dic.Borrar("Elefante")

	var valores []int

	dic.Iterar(func(_ string, dato int) bool {
		valores = append(valores, dato)
		return true
	})

	require.EqualValues(t, 5, len(valores))
	require.EqualValues(t, 4, valores[0])
	require.EqualValues(t, 6, valores[1])
	require.EqualValues(t, 5, valores[2])
	require.EqualValues(t, 2, valores[3])
	require.EqualValues(t, 3, valores[4])
}

func ejecutarPruebaVol(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](comparoClaves[string])

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	insertarBalanceado(dic, claves, valores, 0, len(claves)-1)

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == &valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == &valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVol(b, n)
			}
		})
	}
}

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](comparoClaves[string])
	iter := dic.Iterador()
	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestDiccionarioOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	dic.Guardar("Vaca", "moo")
	dic.Guardar("Gato", "miau")
	dic.Guardar("Perro", "guau")
	iter := dic.Iterador()

	require.True(t, iter.HayAlgoMas())
	primero, primerValor := iter.VerActual()
	require.EqualValues(t, "Gato", primero)
	require.EqualValues(t, "miau", primerValor)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	segundo, segundoValor := iter.VerActual()
	require.EqualValues(t, "Perro", segundo)
	require.EqualValues(t, "guau", segundoValor)

	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())
	tercero, tercerValor := iter.VerActual()
	require.EqualValues(t, "Vaca", tercero)
	require.EqualValues(t, "moo", tercerValor)

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestIterNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	dic.Guardar("B", "")
	dic.Guardar("C", "")
	dic.Guardar("A", "")

	dic.Iterador()
	iter2 := dic.Iterador()
	primero, _ := iter2.VerActual()
	iter2.Avanzar()
	segundo, _ := iter2.VerActual()
	iter2.Avanzar()
	tercero, _ := iter2.VerActual()
	iter2.Avanzar()
	require.False(t, iter2.HayAlgoMas())
	require.EqualValues(t, "A", primero)
	require.EqualValues(t, "B", segundo)
	require.EqualValues(t, "C", tercero)
}

func TestPruebaDeIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del abb cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](comparoClaves[string])
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HayAlgoMas())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func ejecutarPruebasdeVolumenIterador(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](comparoClaves[string])

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	insertarBalanceado(dic, claves, valores, 0, n-1)

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HayAlgoMas())

	ok := true
	var i int
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HayAlgoMas() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		if c1 != claves[i] {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Avanzar()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HayAlgoMas(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIter(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasdeVolumenIterador(b, n)
			}
		})
	}
}

func TestVolumenIterCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	var insertar func(izq, der int)
	insertar = func(izq, der int) {
		if izq > der {
			return
		}
		medio := (izq + der) / 2
		dic.Guardar(medio, medio)
		insertar(izq, medio-1)
		insertar(medio+1, der)
	}
	insertar(0, 9999)

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false
	visitados := 0

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		require.EqualValues(t, visitados, c)
		visitados++
		if c == 100 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
	require.EqualValues(t, 101, visitados)
}

func TestIterarRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	rango := []int{2, 5}
	recorridas := []int{}

	dic.IterarRango(&rango[0], &rango[1], func(clave int, valor int) bool {
		if len(recorridas) < len(claves) {
			recorridas = append(recorridas, clave)
		}
		return true
	})

	require.EqualValues(t, 4, len(recorridas))
	require.EqualValues(t, 2, recorridas[0], "Debería haberse recorrido la clave 2")
	require.EqualValues(t, 3, recorridas[1], "Debería haberse recorrido la clave 3")
	require.EqualValues(t, 4, recorridas[2], "Debería haberse recorrido la clave 4")
	require.EqualValues(t, 5, recorridas[3], "Debería haberse recorrido la clave 5")
}

func TestIterarRangoSinLimiteInferior(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	rango := []int{0, 5}
	recorridas := []int{}

	dic.IterarRango(nil, &rango[1], func(clave int, valor int) bool {
		recorridas = append(recorridas, clave)
		return true
	})

	require.EqualValues(t, 5, len(recorridas))
	require.EqualValues(t, 1, recorridas[0], "Debería haberse recorrido la clave 1")
	require.EqualValues(t, 2, recorridas[1], "Debería haberse recorrido la clave 2")
	require.EqualValues(t, 3, recorridas[2], "Debería haberse recorrido la clave 3")
	require.EqualValues(t, 4, recorridas[3], "Debería haberse recorrido la clave 4")
	require.EqualValues(t, 5, recorridas[4], "Debería haberse recorrido la clave 5")
}

func TestIterarRangoSinLimiteSuperior(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	rango := []int{5, 0}
	recorridas := []int{}

	dic.IterarRango(&rango[0], nil, func(clave int, valor int) bool {
		recorridas = append(recorridas, clave)
		return true
	})

	require.EqualValues(t, 3, len(recorridas))
	require.EqualValues(t, 5, recorridas[0], "Debería haberse recorrido la clave 5")
	require.EqualValues(t, 6, recorridas[1], "Debería haberse recorrido la clave 6")
	require.EqualValues(t, 7, recorridas[2], "Debería haberse recorrido la clave 7")
}

func TestIterarRangoSinLimites(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	recorridas := []int{}

	dic.IterarRango(nil, nil, func(clave int, valor int) bool {
		if len(recorridas) < len(claves) {
			recorridas = append(recorridas, clave)
		}
		return true
	})

	require.EqualValues(t, 7, len(recorridas), "Se deberían haber recorrido todas las claves")
	require.EqualValues(t, 1, recorridas[0], "Debería haberse recorrido la clave 1")
	require.EqualValues(t, 2, recorridas[1], "Debería haberse recorrido la clave 2")
	require.EqualValues(t, 3, recorridas[2], "Debería haberse recorrido la clave 3")
	require.EqualValues(t, 4, recorridas[3], "Debería haberse recorrido la clave 4")
	require.EqualValues(t, 5, recorridas[4], "Debería haberse recorrido la clave 5")
	require.EqualValues(t, 6, recorridas[5], "Debería haberse recorrido la clave 6")
	require.EqualValues(t, 7, recorridas[6], "Debería haberse recorrido la clave 7")
}

func TestIteradorRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	rango := []int{2, 5}
	recorridas := []int{}

	iter := dic.IteradorRango(&rango[0], &rango[1])
	for iter.HayAlgoMas() {
		c, _ := iter.VerActual()
		if len(recorridas) < len(claves) {
			recorridas = append(recorridas, c)
		}
		iter.Avanzar()
	}

	require.EqualValues(t, 4, len(recorridas))
	require.EqualValues(t, 2, recorridas[0], "Debería haberse recorrido la clave 2")
	require.EqualValues(t, 3, recorridas[1], "Debería haberse recorrido la clave 3")
	require.EqualValues(t, 4, recorridas[2], "Debería haberse recorrido la clave 4")
	require.EqualValues(t, 5, recorridas[3], "Debería haberse recorrido la clave 5")
}

func TestIteradorRangoSinLimiteInferior(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	rango := []int{2, 5}
	recorridas := []int{}

	iter := dic.IteradorRango(nil, &rango[1])
	for iter.HayAlgoMas() {
		c, _ := iter.VerActual()
		if len(recorridas) < len(claves) {
			recorridas = append(recorridas, c)
		}
		iter.Avanzar()
	}

	require.EqualValues(t, 5, len(recorridas), "Se deberían haber recorrido 5 claves")
	require.EqualValues(t, 1, recorridas[0], "Debería haberse recorrido la clave 1")
	require.EqualValues(t, 2, recorridas[1], "Debería haberse recorrido la clave 2")
	require.EqualValues(t, 3, recorridas[2], "Debería haberse recorrido la clave 3")
	require.EqualValues(t, 4, recorridas[3], "Debería haberse recorrido la clave 4")
	require.EqualValues(t, 5, recorridas[4], "Debería haberse recorrido la clave 5")
}

func TestIteradorRangoSinLimiteSuperior(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	rango := []int{2, 5}
	recorridas := []int{}

	iter := dic.IteradorRango(&rango[0], nil)
	for iter.HayAlgoMas() {
		c, _ := iter.VerActual()
		if len(recorridas) < len(claves) {
			recorridas = append(recorridas, c)
		}
		iter.Avanzar()
	}

	require.EqualValues(t, 6, len(recorridas), "Se deberían haber recorrido 6 claves")
	require.EqualValues(t, 2, recorridas[0], "Debería haberse recorrido la clave 2")
	require.EqualValues(t, 3, recorridas[1], "Debería haberse recorrido la clave 3")
	require.EqualValues(t, 4, recorridas[2], "Debería haberse recorrido la clave 4")
	require.EqualValues(t, 5, recorridas[3], "Debería haberse recorrido la clave 5")
	require.EqualValues(t, 6, recorridas[4], "Debería haberse recorrido la clave 6")
	require.EqualValues(t, 7, recorridas[5], "Debería haberse recorrido la clave 7")
}

func TestIteradorRangoSinLimites(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](comparoClaves[int])

	claves := []int{4, 2, 6, 1, 3, 5, 7}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	recorridas := []int{}

	iter := dic.IteradorRango(nil, nil)
	for iter.HayAlgoMas() {
		c, _ := iter.VerActual()
		if len(recorridas) < len(claves) {
			recorridas = append(recorridas, c)
		}
		iter.Avanzar()
	}

	require.EqualValues(t, 7, len(recorridas), "Se deberían haber recorrido 7 claves")
	require.EqualValues(t, 1, recorridas[0], "Debería haberse recorrido la clave 1")
	require.EqualValues(t, 2, recorridas[1], "Debería haberse recorrido la clave 2")
	require.EqualValues(t, 3, recorridas[2], "Debería haberse recorrido la clave 3")
	require.EqualValues(t, 4, recorridas[3], "Debería haberse recorrido la clave 4")
	require.EqualValues(t, 5, recorridas[4], "Debería haberse recorrido la clave 5")
	require.EqualValues(t, 6, recorridas[5], "Debería haberse recorrido la clave 6")
	require.EqualValues(t, 7, recorridas[6], "Debería haberse recorrido la clave 7")
}
