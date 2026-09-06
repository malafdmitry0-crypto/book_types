// Package containers показывает обобщённые типы данных: структуру с параметрами
// типов, её методы и функции пакета, которым нужны более сильные ограничения.
package containers

import (
	"cmp"
	"fmt"
	"slices"
)

// Stack хранит элементы одного типа T. Нулевое значение готово к использованию.
type Stack[T any] struct {
	items []T
}

// Push кладёт значение на вершину стека.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop снимает верхний элемент. Второй результат отличает пустой стек
// от стека, в котором лежит нулевое значение T.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	var zero T
	s.items[len(s.items)-1] = zero
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Len возвращает число элементов.
func (s Stack[T]) Len() int { return len(s.items) }

// String печатает элементы от дна к вершине; Stack[T] реализует fmt.Stringer
// при любом T, потому что fmt.Sprint принимает любое значение.
func (s Stack[T]) String() string { return fmt.Sprint(s.items) }

// MapStack строит новый стек, применяя f к каждому элементу.
// До Go 1.27 метод не может объявить собственный параметр R,
// поэтому преобразование T → R живёт в функции пакета.
func MapStack[T, R any](s Stack[T], f func(T) R) Stack[R] {
	out := Stack[R]{items: make([]R, len(s.items))}
	for i, item := range s.items {
		out.items[i] = f(item)
	}
	return out
}

// Container описывает то, что умеет любой стек или очередь с элементами T.
type Container[T any] interface {
	Len() int
	Push(T)
}

// Fill кладёт значения в любой контейнер; *Stack[T] подходит, Stack[T] — нет,
// потому что Push объявлен на указателе.
func Fill[T any](c Container[T], values ...T) {
	for _, v := range values {
		c.Push(v)
	}
}

// Set хранит множество значений сравнимого типа K.
// Нулевое значение непригодно для записи: карта внутри равна nil.
type Set[K comparable] struct {
	m map[K]struct{}
}

// NewSet создаёт пустое множество, готовое к Add.
func NewSet[K comparable]() Set[K] {
	return Set[K]{m: make(map[K]struct{})}
}

// Add добавляет значение; повторное добавление ничего не меняет.
func (s Set[K]) Add(v K) { s.m[v] = struct{}{} }

// Has сообщает, есть ли значение в множестве.
func (s Set[K]) Has(v K) bool {
	_, ok := s.m[v]
	return ok
}

// Len возвращает число элементов.
func (s Set[K]) Len() int { return len(s.m) }

// SortedKeys возвращает элементы по возрастанию. Это функция пакета, а не метод:
// метод Set[K] знает о K только то, что K comparable, и не может потребовать порядок.
func SortedKeys[K cmp.Ordered](s Set[K]) []K {
	keys := make([]K, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
