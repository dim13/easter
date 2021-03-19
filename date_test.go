package easter

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	d := date(2009, 314)
	if d.Day() != 10 || d.Month() != time.November || d.Year() != 2009 {
		t.Errorf("got %v, want 2009-11-10", d)
	}
}
