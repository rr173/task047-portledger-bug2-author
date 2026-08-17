package registry

import (
	"testing"
	"time"
)

func TestProbeStateCopyPreservesInternalPorts(t *testing.T) {
	r := New()
	when := time.Date(2026, 8, 16, 10, 0, 0, 0, time.UTC)
	if _, err := r.Submit("h", when, []int{22, 80}); err != nil {
		t.Fatalf("submit: %v", err)
	}
	state := r.State("h")
	state.CurrentPorts[0] = 9999
	got := r.State("h").CurrentPorts
	if len(got) != 2 || got[0] != 22 || got[1] != 80 {
		t.Fatalf("internal state changed through returned slice: %v", got)
	}
}
