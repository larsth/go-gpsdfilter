package gpsdfilter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/juju/errors"
)

func TestFilterClassMapNotInitialized(t *testing.T) {
	const (
		class      = "PPS"
		doLog      = true
		filterType = TypeIgnore
	)
	var (
		f       *Filter = new(Filter)
		wantErr         = errors.Annotate(ErrFilterMapNotInitialized,
			"empty rules map")
		gotErr error
		s      string
		ok     bool
	)

	_, gotErr = f.Class(class)

	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestFilterClassNoSuchRule(t *testing.T) {
	const (
		class = "PPS"
	)
	var (
		f             = New()
		wantErr error = errors.Annotate(ErrFilterNoSuchRule,
			"filter rule not found")
		gotErr error
		s      string
		ok     bool
	)

	_, gotErr = f.Class(class)

	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestFilterClass(t *testing.T) {
	const (
		class      = "PPS"
		doLog      = true
		filterType = TypeIgnore
	)
	var (
		f                      *Filter
		wantRule, gotRule      *Rule
		wantErr, gotErr        error
		s, sWantRule, sGotRule string
		ok                     bool
	)

	f, wantRule = filterSetup(class, doLog, filterType)
	f.Add(wantRule)

	gotRule, gotErr = f.Class(class)

	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}

	sWantRule = fmt.Sprintf("%#v", wantRule)
	if gotRule != nil {
		sGotRule = fmt.Sprintf("%#v", gotRule)
	} else {
		sGotRule = "<nil>"
	}
	if strings.Compare(sWantRule, sGotRule) != 0 {
		t.Errorf("Got *Rule: '%s', but want *Rule: '%s'", sGotRule, sWantRule)
	}
}
