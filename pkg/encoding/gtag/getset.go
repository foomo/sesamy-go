package gtag

func Set[T any](v T) *T {
	return &v
}

func Get[T any](v *T) T {
	return *v
}

func GetDefault[T any](v *T, fallback T) T {
	if v == nil {
		return fallback
	}
	return *v
}
