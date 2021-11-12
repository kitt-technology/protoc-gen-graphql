package util

import (
	"strings"
)

func Last(path string) string {
	t := strings.Split(path, ".")
	return t[len(t)-1]
}
