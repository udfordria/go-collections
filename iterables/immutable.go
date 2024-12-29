package iterables

func IsEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

func IsNotEmpty[T comparable](v T) bool {
	var zero T
	return zero != v
}
