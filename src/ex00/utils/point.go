package utils

import "strings"

func FinderPoint(ext string) string {
	if strings.Contains(ext, ".") {
		return strings.TrimLeft(ext, ".")
	}
	return ext
}
