package k8s

import (
	"strings"

	v1 "k8s.io/api/core/v1"
)

// JsonPointerEscape escapes JSON Pointer according to https://tools.ietf.org/html/rfc6901#section-3.
func EscapeJSONPointer(s string) string {
	s = strings.ReplaceAll(s, "~", "~0")
	s = strings.ReplaceAll(s, "/", "~1")

	return s
}

// ToStringPtr converts given string to string pointer.
func ToStringPtr(s string) *string {
	return &s
}

// ToPersistentVolumeReclaimPolicyPtr converts given string to v1.PersistentVolumeReclaimPolicy pointer.
func ToPVReclaimPolicyPtr(s string) *v1.PersistentVolumeReclaimPolicy {
	t := v1.PersistentVolumeReclaimPolicy(s)

	return &t
}
