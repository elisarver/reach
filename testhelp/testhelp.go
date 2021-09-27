// Package testhelp are re-usable test helpers for all modules in reach.
package testhelp

import (
	"fmt"
	"reflect"
	"testing"
)

// R is a reporting testing.T...
type R struct {
	t        *testing.T
	instance string
}

// Reporter returns an R that reports on a test instance in its logging.
func Reporter(t *testing.T, instance string) *R {
	return &R{t: t, instance: instance}
}

// Errorf reports an error on a test to the enclosing R...
func (r *R) Errorf(format string, values ...interface{}) {
	r.t.Errorf("%s:%s", r.instance, fmt.Sprintf(format, values...))
}

// Compare reports on differing values in a result to its enclosing R...
func (r *R) Compare(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		r.Errorf("Expected:\n\t%s\nGot:\n\t%s\n", expected, actual)
	}
}
