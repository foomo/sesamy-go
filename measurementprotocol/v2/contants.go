package v2

import (
	"regexp"
)

var (
	RegexProduct = regexp.MustCompile(`pr([1-9]|[1-9][0-9]|1[0-9]{2}|200)`)
)

const (
	ParameterItem = "pr"
)
