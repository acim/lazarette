package k8s //nolint:testpackage

import "testing"

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
			if out := jsonPointerEscape(tt.in); out != tt.out {
				t.Errorf("got %q, want %q", out, tt.out)
			}
		})
	}
}
