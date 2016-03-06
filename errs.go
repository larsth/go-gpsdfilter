package gpsdfilter

import (
	"errors"
)

var (
	//ErrRuleIsNil is an error with the text: "*Rule is a nil Pointer"
	ErrNilRule = errors.New(`*Rule is a nil Pointer`)

	//ErrFilterNoSuchRule is an error withe the text: "Filter: No such Rule
	ErrFilterNoSuchRule = errors.New("Filter: No such Rule")

	//ErrFilterRulesMapNotInitialized
	ErrFilterMapNotInitialized = errors.New("Filter: Had not been initialized")

	//ErrEmptyOrNilByteSlice is an error with the text:
	// `empty or nil byte slice`
	ErrEmptyOrNilByteSlice = errors.New(`empty or nil byte slice`)
)
