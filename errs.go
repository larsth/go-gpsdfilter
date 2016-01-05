package gpsdfilter

import (
	"errors"
)

var (
	//ErrRuleIsNil is an error with the text: "*Rule is a nil Pointer"
	ErrRuleIsNil = errors.New("*Rule is a nil Pointer")
	//ErrFilterNoSuchRule is an error withe the text: "Filter: No such Rule
	ErrFilterNoSuchRule = errors.New("Filter: No such Rule")
)
