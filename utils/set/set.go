package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return Set[T](make(map[T]struct{}))
}

func (s Set[T]) Add(item ...T) {
	for _, it := range item {
		if _, exists := s[it]; exists {
			continue
		}
		s[it] = struct{}{}
	}
}

func (s Set[T]) Rem(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Items() []T {
	items := make([]T, 0)
	for item := range s {
		items = append(items, item)
	}
	return items
}
