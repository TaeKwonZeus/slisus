package slisus

// Entry represents a key-value pair in a map.
type Entry struct {
	Key   interface{}
	Value interface{}
}

// MapEntries returns a slice of key-value pairs from gomap.
func MapEntries(gomap map[interface{}]interface{}) []Entry {
	var out []Entry

	for k, v := range gomap {
		out = append(out, Entry{k, v})
	}

	return out
}

// MapKeys returns a slice of gomap's keys.
func MapKeys(gomap map[interface{}]interface{}) []interface{} {
	var out []interface{}

	for i := range gomap {
		out = append(out, i)
	}

	return out
}

// MapValues returns a slice of gomap's values.
func MapValues(gomap map[interface{}]interface{}) []interface{} {
	var out []interface{}

	for _, i := range gomap {
		out = append(out, i)
	}

	return out
}
