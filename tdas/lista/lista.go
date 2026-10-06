package lista

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista y devuelve su valor. Si la lista tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último elemento de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Iterar recorre la lista desde el primero al último elemento, aplicando la función dada a cada uno de ellos.
	// Si la función devuelve false, se detiene la iteración. Si devuelve true, se continúa con el siguiente elemento.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo posicionado en el primer elemento de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	//VerActual devuelve el valor del elemento actual del iterador. Si el iterador no tiene un elemento actual, entra en pánico con un mensaje "El iterador termino de iterar".
	VerActual() T

	//HayAlgoMas devuelve true si el iterador tiene más elementos, false en caso contrario.
	HayAlgoMas() bool

	//Avanzar mueve el iterador al siguiente elemento. Si el iterador no tiene más elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Avanzar()

	//Insertar inserta un nuevo elemento en la posición actual del iterador.
	Insertar(T)

	//Borrar borra el elemento actual del iterador y devuelve su valor. Si el iterador no tiene más elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
