package gpsdfilter

import (
	"fmt"
	"strings"
)

func errorTest(got error, want error) (s string, ok bool) {
	var (
		wantStr string
		gotStr  string
	)
	if want == nil {
		wantStr = `<nil>`
	} else {
		wantStr = want.Error()
	}
	if got == nil {
		gotStr = `<nil>`
	} else {
		gotStr = got.Error()
	}
	if strings.Compare(wantStr, gotStr) != 0 {
		format := "%s%s'\n, %s%s'"
		s1 := `Got the error '`
		s2 := `... but want the error '`
		s = fmt.Sprintf(format, s1, gotStr, s2, wantStr)
		ok = false
	} else {
		s = ""
		ok = true
	}
	return
}

func errorTestN(got error, want error, i int, tdVarName string) (s string, ok bool) {
	var (
		sTest string
	)
	sTest, ok = errorTest(got, want)
	if ok {
		format := "%s\n%s %s[%d]"
		s1 := `The test error ocurred in test: `
		s = fmt.Sprintf(format, sTest, s1, tdVarName, i)
	} else {
		s = ""
	}
	return
}

func byteSliceTest(got []byte, want []byte, i int) (ok bool, s string) {
	var (
		wantStr string
		gotStr  string
	)
	if want == nil {
		wantStr = "<empty>"
	} else {
		wantStr = string(want)
		if len(wantStr) == 0 {
			wantStr = "<empty>"
		}
	}
	if got == nil {
		gotStr = "<empty>"
	} else {
		gotStr = string(got)
		if len(gotStr) == 0 {
			gotStr = "<empty>"
		}
	}
	if strings.Compare(wantStr, gotStr) != 0 {
		format := "%s%s\n, %s%s\"\n%s %d]"
		s1 := `Got the byte slice "`
		s2 := `"... but want the byte slice "`
		s3 := `The test error ocurred in test: tdNew[`
		s = fmt.Sprintf(format, s1, gotStr, s2, wantStr, s3, i)
		ok = false
	} else {
		s = ""
		ok = true
	}
	return
}

func filterSetup(class string, doLog bool,
	filterType FilterType) (f *Filter, rI *Rule) {

	f = New()
	rI = &Rule{Class: class, DoLog: doLog, Type: filterType}
	return
}
