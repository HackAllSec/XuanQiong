package controllers

import "testing"

func TestNormalizeDeleteUserID(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want uint64
		ok   bool
	}{
		{name: "number", in: float64(42), want: 42, ok: true},
		{name: "string", in: "42", want: 42, ok: true},
		{name: "invalid string", in: "abc", ok: false},
		{name: "negative number", in: float64(-1), ok: false},
		{name: "fractional number", in: float64(1.5), ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := normalizeDeleteUserID(tt.in)
			if ok != tt.ok {
				t.Fatalf("ok = %v, want %v", ok, tt.ok)
			}
			if got != tt.want {
				t.Fatalf("id = %d, want %d", got, tt.want)
			}
		})
	}
}
