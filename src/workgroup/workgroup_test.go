package workgroup

import "testing"

func assert(t *testing.T, want, got error) {
	t.Helper()
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
