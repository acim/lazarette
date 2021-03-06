package k8s_test

import (
	"testing"

	"github.com/acim/lazarette/pkg/k8s"
)

func TestJsonPointerEscape(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "blank strings",
			in:   "",
			out:  "",
		},
		{
			name: "alone ~",
			in:   "~",
			out:  "~0",
		},
		{
			name: "alone /",
			in:   "/",
			out:  "~1",
		},
		{
			name: "mixed",
			in:   "~a/b~~/./~s",
			out:  "~0a~1b~0~0~1.~1~0s",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if out := k8s.EscapeJSONPointer(tt.in); out != tt.out {
				t.Errorf("got %q, want %q", out, tt.out)
			}
		})
	}
}
