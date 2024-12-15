package perm

func Equal[T any](n int, picks []T) [][]T {
	var equal func(cur []T, level int) [][]T
	equal = func(cur []T, level int) [][]T {
		if level == n {
			return [][]T{cur}
		}

		var ret [][]T
		for _, p := range picks {
			var s []T
			s = append(s, cur...)
			s = append(s, p)
			ret = append(ret, equal(s, level+1)...)
		}
		return ret
	}

	return equal([]T{}, 0)
}

func EqualFunc[T any](n int, picks []T, predicate func([]T)) {
	var equal func(cur []T, level int) [][]T
	equal = func(cur []T, level int) [][]T {
		if level == n {
			predicate(cur)
		}

		var ret [][]T
		for _, p := range picks {
			var s []T
			s = append(s, cur...)
			s = append(s, p)
			ret = append(ret, equal(s, level+1)...)
		}
		return ret
	}

	equal([]T{}, 0)
}
