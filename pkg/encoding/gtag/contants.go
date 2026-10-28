package gtag

import (
	"regexp"
)

var (
	// RegexProduct matches the gtag.js item parameters pr1 to pr200.
	RegexProduct = regexp.MustCompile(`pr([1-9]|[1-9][0-9]|1[0-9]{2}|200)`)
)

const (
	// ParameterItem is the gtag.js parameter prefix for ecommerce items.
	ParameterItem = "pr"
)
