//Package gpsdfilter is for filtering gpsd JSON documents
package gpsdfilter

import (
	"encoding/json"
	"sync"
)

//Rule is a type that descripe how to filter a gpsd JSON document
type Rule struct {
	Class string
	DoLog bool
	Type  Type
}

//Filter contains all the filter rules of type *Rule
type Filter struct {
	mutex sync.Mutex
	rules map[string]*Rule
}

//AddRule adds a *Rule
//Error ErrRuleIsNil is returned if the *Rule is nil, otherwise
//the nil error is returned
func (f *Filter) AddRule(r *Rule) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if r == nil {
		return ErrRuleIsNil
	}
	f.rules[r.Class] = r

	return nil
}

//Filter takes a byte slice as input, and returns a *Rule and an error
// The Error ErrFilterNoSuchRule is returned if the rule is unknown.
//  An error from the json paser can also be returned.
// If there is an error, then *Rule is nil.
// If there are no errors, then the error is nil
func (f *Filter) Filter(p []byte) (*Rule, error) {
	var (
		c   Class
		err error
	)
	if err = json.Unmarshal(p, &c); err != nil {
		return nil, err
	}
	return f.FilterClass(c.Class)
}

//FilterClass takes a class of type string, fx. "TPV", as input, and returns a
// *Rule and an error
// The Error ErrFilterNoSuchRule is returned if the rule is unknown.
//  An error from the json paser can also be returned.
// If there is an error, then *Rule is nil.
// If there are no errors, then the error is nil
func (f *Filter) FilterClass(class string) (*Rule, error) {
	var (
		rule *Rule
		ok   bool
	)
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if rule, ok = f.rules[class]; ok == false {
		return nil, ErrFilterNoSuchRule
	}
	return rule, nil
}
