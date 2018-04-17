package cycle

import (
	"reflect"
	"unsafe"
)

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}

func IsCycle(x interface{}) bool {
	seen := make(map[comparison]bool)
	return isCycle(reflect.ValueOf(x), seen)
}

func isCycle(x reflect.Value, seen map[comparison]bool) bool {
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := comparison{xptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCycle(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if !isCycle(x.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCycle(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}
	return false
}
