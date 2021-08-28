package slisus

import "reflect"

// Contains checks if slice contains elem.
func Contains(slice []interface{}, elem interface{}) bool {
	for _, i := range slice {
		if reflect.DeepEqual(i, elem) {
			return true
		}
	}

	return false
}

// ContainsAll checks if slice contains all elems.
func ContainsAll(slice []interface{}, elems []interface{}) bool {
	for _, i := range elems {
		if !Contains(slice, i) {
			return false
		}
	}

	return true
}

// MapContainsKey checks if gomap's keys contain key.
func MapContainsKey(gomap map[interface{}]interface{}, key interface{}) bool {
	for i := range gomap {
		if reflect.DeepEqual(i, key) {
			return true
		}
	}

	return false
}

// MapContainsValue checks if gomap's values contain value.
func MapContainsValue(gomap map[interface{}]interface{}, value interface{}) bool {
	for _, i := range gomap {
		if reflect.DeepEqual(i, value) {
			return true
		}
	}

	return false
}
