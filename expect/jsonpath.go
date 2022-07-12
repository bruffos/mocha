package expect

import (
	"github.com/vitorsalgado/mocha/x/jsonx"
)

// JSONPath applies the provided matcher to the JSON field value in the given path.
// Example:
//	JSONPath("address.city", EqualTo("Santiago"))
func JSONPath(p string, matcher Matcher) Matcher {
	m := Matcher{}
	m.Name = "JSONPath"
	m.Matches = func(v any, args Args) (bool, error) {
		value, err := jsonx.Reach(p, v)
		if err != nil || value == nil {
			return false, err
		}

		return matcher.Matches(value, args)
	}

	return m
}
