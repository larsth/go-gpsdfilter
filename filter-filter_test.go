package gpsdfilter

import (
	//"fmt"
	//"strings"
	"fmt"
	"strings"
	"testing"

	"github.com/juju/errors"
)

func TestFilterFilterJsonError(t *testing.T) {
	var (
		f       = New()
		p       = []byte(`{class:"TPV`)
		wantErr = errors.Errorf("%s %s",
			`JSON document Unmarshal error: invalid character 'c' looking`,
			`for beginning of object key string`)
		gotErr error
		s      string
		ok     bool
	)
	_, gotErr = f.Filter(p)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestFilterFilterMapNotInitialized(t *testing.T) {
	var (
		f       *Filter = &Filter{}
		p               = []byte(`{"class":"TPV"}`)
		wantErr         = errors.Annotate(ErrFilterMapNotInitialized,
			"empty rules map")
		gotErr error
		s      string
		ok     bool
	)
	_, gotErr = f.Filter(p)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestFilterFilterNoSuchRule(t *testing.T) {
	var (
		f       = New()
		p       = []byte(`{"class":"TPV"}`)
		wantErr = errors.Annotate(ErrFilterNoSuchRule,
			"filter rule not found")
		gotErr error
		s      string
		ok     bool
	)

	_, gotErr = f.Filter(p)
	if s, ok = errorTest(gotErr, wantErr); false == ok {
		t.Error(s)
	}
}

func TestFilterFilter(t *testing.T) {
	const (
		class      = "TPV"
		doLog      = true
		filterType = TypeIgnore
	)
	var (
		f                      *Filter
		wantRule               *Rule
		p                      = []byte(`{"class":"TPV"}`)
		gotRule                *Rule
		wantErr                error = nil
		gotErr                 error
		s, sWantRule, sGotRule string
		ok                     bool
	)

	f, wantRule = filterSetup(class, doLog, filterType)
	f.Add(wantRule)
	gotRule, gotErr = f.Filter(p)
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
