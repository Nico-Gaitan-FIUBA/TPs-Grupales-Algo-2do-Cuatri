package diccionario

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hash[K, V]{}
}
