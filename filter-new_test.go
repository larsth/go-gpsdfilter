package gpsdfilter

import "testing"

func TestNew(t *testing.T) {
	var (
		f *Filter
	)
	if f = New(); f == nil {
		t.Error("*gpsd-filter.Filter is nil")
	}
}
