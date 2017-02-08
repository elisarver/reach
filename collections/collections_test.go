package collections

import (
	"fmt"
	"testing"

	"github.com/elisarver/reach/testhelp"
)

func TestDropEmpties(t *testing.T) {
	input := []string{"not empty", "", "also not empty"}
	expected := fmt.Sprintf("%q", []string{"not empty", "also not empty"})
	actual := fmt.Sprintf("%q", DropEmpties(input))
	r := testhelp.Reporter(t, "dropEmpties")
	r.Compare(expected, actual)
}
