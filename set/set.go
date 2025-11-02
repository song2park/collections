package set

type Set[T comparable] struct {
	ele map[T]struct{}
}

func NewSet[T comparable](elements ...T) Set[T] {
	s := Set[T]{}
	for _, e := range elements {
		s.Add(e)
	}

	return s
}

func (s *Set[T]) Add(element T) {
	if s.ele == nil {
		s.ele = make(map[T]struct{})
	}
	s.ele[element] = struct{}{}
}

func (s *Set[T]) AddAll(elements ...T) {
	if s.ele == nil {
		s.ele = make(map[T]struct{})
	}
	for _, e := range elements {
		s.ele[e] = struct{}{}
	}
}

func (s *Set[T]) Remove(element T) {
	if s.ele == nil {
		return
	}
	delete(s.ele, element)
}

func (s *Set[T]) RemoveAll(elements ...T) {
	if s.ele == nil {
		return
	}
	for _, e := range elements {
		delete(s.ele, e)
	}
}

func (s *Set[T]) Contains(element T) bool {
	if s.ele == nil {
		return false
	}
	_, exists := s.ele[element]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.ele)
}

func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.ele))
	for e := range s.ele {
		result = append(result, e)
	}
	return result
}

func (s *Set[T]) Clear() {
	s.ele = make(map[T]struct{})
}