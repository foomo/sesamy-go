package measurementprotocol

import (
	"fmt"
)

func Set[T any](v T) *T {
	return &v
}

func SetInt(v int) *string {
	if v == 0 {
		return nil
	}
	return Set(fmt.Sprintf("%d", v))
}

func SetUInt(v uint) *string {
	if v == 0 {
		return nil
	}
	return Set(fmt.Sprintf("%d", v))
}

func SetInt32(v int32) *string {
	if v == 0 {
		return nil
	}
	return Set(fmt.Sprintf("%d", v))
}

func SetUInt32(v uint32) *string {
	if v == 0 {
		return nil
	}
	return Set(fmt.Sprintf("%d", v))
}

func SetInt64(v int64) *string {
	if v == 0 {
		return nil
	}
	return Set(fmt.Sprintf("%d", v))
}

func SetUInt64(v uint64) *string {
	if v == 0 {
		return nil
	}
	return Set(fmt.Sprintf("%d", v))
}

func SetFloat32(v float32) *string {
	if v == 0 {
		return nil
	}
	return Set(fmt.Sprintf("%f", v))
}

func SetFloat64(v float64) *string {
	if v == 0 {
		return nil
	}
	return Set(fmt.Sprintf("%f", v))
}

func SetString(v string) *string {
	if v == "" {
		return nil
	}
	return Set(v)
}

func SetBool(v bool) *string {
	if !v {
		return nil
	}
	return Set("1")
}

func SetStringMap(v map[string]string) map[string]string {
	if len(v) == 0 {
		return nil
	}
	return v
}

func AddStringMap(t map[string]string, k string, v *string) {
	if v == nil {
		return
	}
	t[k] = *v
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
