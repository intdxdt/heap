package heap

const nilIndex = -9

func pop(a []interface{}) (interface{}, []interface{}) {
	var v interface{}
	var n int
	if len(a) == 0 {
		return nil, a
	}
	n = len(a) - 1
	v, a[n] = a[n], nil
	return v, a[:n]
}

// is index nil type
func isNil(v int) bool {
	return v == nilIndex
}
