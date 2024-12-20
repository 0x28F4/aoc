package utils

import (
	"cmp"
	"fmt"
	"strconv"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func MustLen[T any](l []T, want int) {
	if len(l) != want {
		panic(fmt.Sprintf("unexpected length of list, want %d, got %d\nitems:\n%v", want, len(l), l))
	}
}

func MustInt(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func MustEq(a, b any) {
	if a != b {
		panic(fmt.Sprintf("not equal: %v =/= %v", a, b))
	}
}

func MustNotEq(a, b any) {
	if a == b {
		panic(fmt.Sprintf("got equal values, but they shouldnt be %v", a))
	}
}

func IsSliceEq[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range b {
		if b[i] != a[i] {
			return false
		}
	}

	return true
}

func MustSliceEq[T comparable](a []T, b []T) {
	if len(a) != len(b) {
		panic(fmt.Sprintf("lengths of slices don't match: len(a)=%d, len(b)=%d\na=%v\nb=%v", len(a), len(b), a, b))
	}
	for i := range b {
		if b[i] != a[i] {
			panic(fmt.Sprintf("values of slices don't match, %v =/= %v at %d", b[i], a[i], i))
		}
	}
}

func MustNil(v any) {
	if v != nil {
		panic(fmt.Sprintf("not nil: %v", v))
	}
}

func MustNotNil(v any) {
	if v == nil {
		panic("value is nil")
	}
}

func MustFalse(v bool) {
	if v {
		panic("got true")
	}
}

func MustSmaller[T cmp.Ordered](actual T, other T) {
	if actual >= other {
		panic(fmt.Sprintf("actual %d is greater or equal than other %d", actual, other))
	}
}

func MustGreater[T cmp.Ordered](actual T, other T) {
	if actual <= other {
		panic(fmt.Sprintf("actual %d is smaller or equal than other %d", actual, other))
	}
}

func MustSmallerEq[T cmp.Ordered](actual T, other T) {
	if actual > other {
		panic(fmt.Sprintf("actual %d is greater than other %d", actual, other))
	}
}

func MustGreaterEq[T cmp.Ordered](actual T, other T) {
	if actual < other {
		panic(fmt.Sprintf("actual %d is smaller than other %d", actual, other))
	}
}

func MustTrue(v bool) {
	if !v {
		panic("got false")
	}
}

func Take[T any](s []T, index int) (T, []T) {
	if index > len(s)-1 {
		panic(fmt.Sprintf("error in Take: index out of bounds, access %d which is greater than length of %d", index, len(s)))
	}
	if index < 0 {
		panic(fmt.Sprintf("error in Take: index out of bounds, access %d", index))
	}

	var newS []T
	for i, v := range s {
		if i == index {
			continue
		}
		newS = append(newS, v)
	}
	ret := s[index]
	return ret, newS
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func StringPopLeft(s string) (item string, newS string) {
	if len(s) == 0 {
		return
	}
	item = s[0:1]
	newS = s[1:]
	return
}

func Pow(x, p int) int {
	MustGreaterEq(p, 0)
	ret := 1
	for range p {
		ret *= x
	}
	return ret
}

var Inf = int(^uint(0) >> 1)

func Distance(a []int, b []int) int {
	if len(a) != len(b) {
		return Inf
	}

	dist := 0
	for i := range a {
		dist += Abs(a[i] - b[i])
	}

	return dist
}

func DistanceBinary(a []int, b []int) int {
	dist := 0
	for i := range max(len(a), len(b)) {
		if i >= len(b) || i >= len(a) {
			dist++
			continue
		}

		if a[i] != b[i] {
			dist++
		}
	}

	return dist
}

func Min[T cmp.Ordered](s []T) (best T) {
	MustGreater(len(s), 0)
	best = s[0]
	for _, v := range s[1:] {
		if v < best {
			best = v
		}
	}
	return best
}

func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
