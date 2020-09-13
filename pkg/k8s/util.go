package k8s

import "strings"

func JSONPointerEscape(s string) string {
	s = strings.Replace(s, "~", "~0", -1)
	s = strings.Replace(s, "/", "~1", -1)

	return s
}
