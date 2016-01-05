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

//Filter takes a byte slice as input, and returns 3 types of information:
//  A boolean: If true, the gpsd JSON document should be logged, otherwise not logged
//  A Type: Tells what to do with the gpsd JSON document.
//  An error: Error ErrFilterNoSuchRule is returned if the rule is unknown.
//  An error from the json paser can also be returned, and
//  otherwise the nil error is returned.
func (f *Filter) Filter(p []byte) (bool, Type, error) {
	var (
		c   Class
		err error
	)
	if err = json.Unmarshal(p, &c); err != nil {
		return false, TypeUnknown, err
	}
	return f.FilterClass(c.Class)
}

//FilterClass takes a class of type string, fx. "TPV", as input, and returns 3
// types of information:
//  A boolean: If true, the gpsd JSON document should be logged, otherwise not logged
//  A Type: Tells what to do with the gpsd JSON document.
//  An error: Error ErrFilterNoSuchRule is returned if the rule is unknown.
//  An error from the json paser can also be ret
func (f *Filter) FilterClass(class string) (bool, Type, error) {
	var (
		rule *Rule
		ok   bool
	)
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if rule, ok = f.rules[class]; ok == false {
		return false, TypeUnknown, ErrFilterNoSuchRule
	}
	return rule.DoLog, rule.Type, nil
}
