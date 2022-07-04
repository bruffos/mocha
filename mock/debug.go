package mock

import (
	"fmt"

	"github.com/vitorsalgado/mocha/internal/stylize"
	"github.com/vitorsalgado/mocha/matchers"
)

// Debug wraps a matchers.Matcher adding debugging logs.
// The return value will be the same of the provided matchers.Matcher.
// Debug is used internally by Mocha.
func Debug[V any](name string, mk Mock, m matchers.Matcher[V]) matchers.Matcher[V] {
	return func(v V, params matchers.Args) (bool, error) {
		result, err := m(v, params)

		fmt.Printf("\"%s\" received %v\n", name, stylize.Gray(fmt.Sprintf("%v", v)))

		if err != nil {
			fmt.Print(stylize.RedBright(
				fmt.Sprintf("%s - an error ocurred. reason: %s\n", stylize.Bold(fmt.Sprintf("\"%s\"", name)), err.Error())))
			fmt.Printf("mock %d %s\n", mk.ID, mk.Name)
		} else if !result {
			fmt.Print(stylize.RedBright(
				fmt.Sprintf("%s - did not match\n", stylize.Bold(fmt.Sprintf("\"%s\"", name)))))
			fmt.Printf("mock %d %s\n", mk.ID, mk.Name)
		}

		return result, err
	}
}
