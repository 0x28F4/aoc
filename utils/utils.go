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
