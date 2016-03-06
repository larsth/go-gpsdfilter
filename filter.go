//Package gpsdfilter is for filtering gpsd JSON documents
package gpsdfilter

import (
	"encoding/json"
	"sync"

	"github.com/juju/errors"
)

const numberOfGpsdClassTypes = 16

//Rule is a type that descripe how to filter a gpsd JSON document
type Rule struct {
	Class string
	DoLog bool
	Type  FilterType
}

//Filter contains all the filter rules of type *Rule
type Filter struct {
	mutex sync.Mutex
	rules map[string]*Rule
}

//New initializes a type Filter, and returns a pointer to it
func New() *Filter {
	f := new(Filter)
	f.rules = make(map[string]*Rule, numberOfGpsdClassTypes)

	return f
}

//Add adds a *Rule
//Error ErrRuleIsNil is returned if the *Rule is nil.
//Error ErrFilterRulesMapNotInitialized is returned if Filter had not been
//initialized.
//If there are no errors, then the nil error is returned.
func (f *Filter) Add(r *Rule) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if r == nil {
		return errors.Annotate(ErrNilRule, `Error: adding a filter rule`)
	}
	if f.rules == nil {
		return errors.Annotate(ErrFilterMapNotInitialized,
			"Error: adding a filer rule")
	}
	f.rules[r.Class] = r

	return nil
}

//Filter takes a byte slice as input, and returns a *Rule and an error
// The error ErrFilterNoSuchRule is returned, if the rule is unknown.
// The error ErrFilterMapNotInitialized is returned, if a *Filter had not been
// initialized.
//
//  An error from the json parser can also be returned.
//
// If there is an error, then *Rule is nil.
// If there are no errors, then the error is nil
func (f *Filter) Filter(p []byte) (*Rule, error) {
	var (
		c   Class
		err error
	)

	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(p) == 0 {
		return nil, errors.Trace(ErrEmptyOrNilByteSlice)
	}

	if err = json.Unmarshal(p, &c); err != nil {
		annotatedErr := errors.Annotate(err,
			"JSON document Unmarshal error")
		return nil, annotatedErr
	}
	r, err := f.classUnsafe(c.Class)
	if err != nil {
		return r, errors.Trace(err)
	}
	return r, nil
}

func (f *Filter) classUnsafe(class string) (*Rule, error) {
	var (
		rule *Rule
		ok   bool
	)

	if f.rules == nil {
		annotatedErr := errors.Annotate(ErrFilterMapNotInitialized,
			"empty rules map")
		return nil, annotatedErr
	}
	if rule, ok = f.rules[class]; false == ok {
		annotatedErr := errors.Annotate(ErrFilterNoSuchRule,
			"filter rule not found")
		return nil, annotatedErr
	}
	return rule, nil
}

//Class takes a class of type string, fx. "TPV", as input, and returns a
// *Rule and an error
// The error ErrFilterNoSuchRule is returned if the rule is unknown.
// The error ErrFilterMapNotInitialized is type had not been initialized.
//  An error from the json paser can also be returned.
// If there is an error, then *Rule is nil.
// If there are no errors, then the error is nil
func (f *Filter) Class(class string) (*Rule, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.classUnsafe(class)
}
