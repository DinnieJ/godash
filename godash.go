package godash

import "fmt"

func Sprf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func If[T any](condition bool, _vtrue, _vfalse T) T {
	if condition {
		return _vtrue
	}
	return _vfalse
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Contain[T comparable](_lst_value []T, _haystack T) bool {
	for _, v := range _lst_value {
		if v == _haystack {
			return true
		}
	}
	return false
}
