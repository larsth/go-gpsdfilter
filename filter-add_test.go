package gpsdfilter

import (
	"fmt"
	"testing"

	"github.com/juju/errors"
)

func TestFilterAddNil(t *testing.T) {
	var (
		f             = New()
		r       *Rule = nil
		wantErr       = errors.Annotate(ErrNilRule,
			`Error: adding a filter rule`)
		gotErr error
		s      string
		ok     bool
	)
	gotErr = f.Add(r)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestFilterAddMapNotInitialized(t *testing.T) {
	var (
		f       = new(Filter)
		r       = &Rule{}
		wantErr = errors.Annotate(ErrFilterMapNotInitialized,
			"Error: adding a filer rule")
		gotErr error
		s      string
		ok     bool
	)
	gotErr = f.Add(r)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestFilterAdd1(t *testing.T) {
	const (
		class      = "PPS"
		doLog      = true
		filterType = TypeIgnore
	)
	var (
		f       *Filter
		r       *Rule
		wantErr error = nil
		gotErr  error
		s       string
		ok      bool
	)
	f, r = filterSetup(class, doLog, filterType)
	gotErr = f.Add(r)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestFilterAdd2(t *testing.T) {
	const (
		class      = "PPS"
		doLog      = true
		filterType = TypeIgnore
	)
	var (
		f  *Filter
		rI *Rule
		r  *Rule
		ok bool
		sI string
	)
	f, rI = filterSetup(class, doLog, filterType)
	_ = f.Add(rI)
	fmt.Printf("%#v\n", f.rules)
	if r, ok = f.rules[class]; false == ok {
		sI = fmt.Sprintf("%#v", rI)
		t.Errorf("Want rule: '%s', but got *Rule: <nil>", sI)
	}
	t.Logf("\n\tRule is ok:\n\t\t Got rule '%#v', and\n\t\t want rule: '%#v'", r, rI)
}
