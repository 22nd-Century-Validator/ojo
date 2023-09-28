package util

// AppendString will append a string to a slice
// if and only if the string is not already in
// the slice.
func AppendString(keys []string, key string) []string {
	duplicate := false
	for _, v := range keys {
		if key == v {
			duplicate = true
		}
	}
	if !duplicate {
		keys = append(keys, key)
	}
	return keys
}
