package slisus

import "reflect"

// Contains checks if the specified element is contained in slice.
func Contains(slice []interface{}, elem interface{}) bool {
	for _, i := range slice {
		if reflect.DeepEqual(i, elem) {
			return true
		}
	}

	return false
}

func ContainsAll(slice []interface{}, elems []interface{}) bool {
	for _, i := range elems {
		if !Contains(slice, i) {
			return false
		}
	}

	return true
}
