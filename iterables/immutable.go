package iterables

// Should be used
func IsEmpty(element interface{}) bool {
	switch value := element.(type) {
	case string:
		return len(value) == 0
	case map[interface{}]interface{}:
		return len(value) == 0
	case []interface{}:
		return len(value) == 0
	default:
		return false
	}
}

func IsNotEmpty(element interface{}) bool {
	return !IsEmpty(element)
}
