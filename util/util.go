package util

import "reflect"

func SetInt64(i *int64) {
	val := reflect.ValueOf(i)
	val.Elem().SetInt(*i + 1)
}

func StrInList(l []string, d string) bool {
	var (
		ok bool
	)
	for i := 0; i < len(l); i++ {
		if l[i] == d {
			ok = true
			break
		}
	}
	return ok
}