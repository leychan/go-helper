package go_helper

func IsSet(m map[string]interface{}, k string) bool {
	if _, ok := m[k]; ok {
		return true
	}
	return false
}
