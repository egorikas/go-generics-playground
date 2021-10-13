package main

func main() {
	ss := new_set[int]()
	ss.Add(123)
	print(ss.Contains(123))
}

type set[T comparable] map[T]struct{}

func new_set[T comparable]() set[T] {
	return make(set[T])
}

func (s set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s set[T]) Delete(v T) {
	delete(s, v)
}

func (s set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s set[T]) Len() int {
	return len(s)
}

func (s set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}