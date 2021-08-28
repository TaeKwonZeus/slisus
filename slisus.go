// Package slisus provides utils for working with slices and maps.
package slisus

import "reflect"

// Entry represents a key-value pair in a map.
type Entry struct {
	Key   interface{}
	Value interface{}
}

// All checks if all elements in slice match predicate.
func All(slice []interface{}, predicate func(element interface{}) bool) bool {
	for _, i := range slice {
		if !predicate(i) {
			return false
		}
	}

	return true
}

// Associate returns a map of entries provided by slice mapped with transform.
func Associate(slice []interface{}, transform func(interface{}) Entry) map[interface{}]interface{} {
	out := map[interface{}]interface{}{}

	for _, i := range slice {
		entry := transform(i)
		out[entry.Key] = entry.Value
	}

	return out
}

// Contains checks if slice contains element.
func Contains(slice []interface{}, element interface{}) bool {
	for _, i := range slice {
		if reflect.DeepEqual(i, element) {
			return true
		}
	}

	return false
}

// ContainsAll checks if slice contains all elements.
func ContainsAll(slice []interface{}, elements []interface{}) bool {
	for _, i := range elements {
		if !Contains(slice, i) {
			return false
		}
	}

	return true
}

// Chunked splits slice into smaller slices each not exceeding size.
func Chunked(slice []interface{}, size int) [][]interface{} {
	out := [][]interface{}{}

	if size <= 0 {
		return out
	}

	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}

		out = append(out, slice[i:end])
	}

	return out
}

// Get returns an element from slice with a specified index.
// If index is negative, it will count from the end of slice.
// If the item is not found, Get will return nil.
func Get(slice []interface{}, index int) interface{} {
	if index < 0 {
		return Get(slice, len(slice)+index)
	}

	if index > len(slice)-1 {
		return nil
	}

	return slice[index]
}

// IndexOf returns the index of the first occurrence of element in slice, or -1 if none were found.
func IndexOf(slice []interface{}, element interface{}) int {
	for i, e := range slice {
		if reflect.DeepEqual(e, element) {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index of the last occurrence of element in slice, or -1 if none were found.
func LastIndexOf(slice []interface{}, element interface{}) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if reflect.DeepEqual(slice[i], element) {
			return i
		}
	}

	return -1
}

// Reversed returns reversed slice.
func Reversed(slice []interface{}) []interface{} {
	out := []interface{}{}

	for i := len(slice) - 1; i >= 0; i-- {
		out = append(out, slice[i])
	}

	return out
}
