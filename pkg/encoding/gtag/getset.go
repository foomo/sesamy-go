package gtag

// Deprecated: use `new()` instead
func Set[T any](v T) *T { //nolint:modernize // keep with deprecation notice
	return &v
}

// Get returns the value of a pointer
func Get[T any](v *T) T {
	return *v
}

// GetDefault returns the value of a pointer, or a fallback value if the pointer is nil
func GetDefault[T any](v *T, fallback T) T {
	if v == nil {
		return fallback
	}

	return *v
}
