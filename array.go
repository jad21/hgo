package x

import "sort"

// Find: encontrar el primero que cumpla el criterio
func Find[T any](array []T, cb func(item T) bool) *T {
	for _, v := range array {
		if cb(v) {
			return &v
		}
	}
	return nil
}

// Filter devuelve un nuevo array con los elementos que cumplen con un criterio específico
func Filter[T any](array []T, criteria func(T) bool) []T {
	var result []T
	for _, v := range array {
		if criteria(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reverse invierte el orden de los elementos en un array
func Reverse[T any](array []T) []T {
	length := len(array)
	reversed := make([]T, length)
	for i, v := range array {
		reversed[length-1-i] = v
	}
	return reversed
}

// Contains verifica si un array contiene un valor específico
func Contains[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// Map aplica una función a cada elemento de un array y devuelve un nuevo array con los resultados
func Map[T any, U any](array []T, f func(T) U) []U {
	mapped := make([]U, len(array))
	for i, v := range array {
		mapped[i] = f(v)
	}
	return mapped
}

// Map aplica una función a cada elemento de un index del array y devuelve un nuevo array con los resultados
func MapIndex[T any](count int, f func(i int) T) []T {
	mapped := make([]T, count)
	for i := 0; i < count; i++ {
		mapped[i] = f(i)
	}
	return mapped
}

// Repeat repite el valor el numero de veces que se desea
func Repeat[T any](v T, count int) (results []T) {
	for i := 0; i < count; i++ {
		results = append(results, v)
	}
	return
}

// Unique devuelve un nuevo array con solo elementos únicos del array dado
func Unique[T comparable](array []T) []T {
	seen := make(map[T]struct{})
	var uniqueArray []T
	for _, v := range array {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			uniqueArray = append(uniqueArray, v)
		}
	}
	return uniqueArray
}

// UniqueBy devuelve un nuevo array con solo elementos únicos del array dado, basados en un atributo específico
func UniqueBy[T any, K comparable](array []T, keyFunc func(T) K) []T {
	seen := make(map[K]struct{})
	var uniqueArray []T
	for _, v := range array {
		key := keyFunc(v)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			uniqueArray = append(uniqueArray, v)
		}
	}
	return uniqueArray
}

// Sort ordena un array utilizando una función de comparación
func Sort[T any](array []T, less func(i, j T) bool) {
	sort.Slice(array, func(i, j int) bool {
		return less(array[i], array[j])
	})
}
