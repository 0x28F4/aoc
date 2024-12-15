package utils

import (
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

func MustFalse(v bool) {
	if v {
		panic("got true")
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
