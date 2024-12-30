package __tests__

import (
	"testing"

	"github.com/udfordria/go-collections/iterables"
)

func TestSliceIsNotEmpty(t *testing.T) {
	a := []string{"a", "b", "c"}
	t.Log(iterables.IsNotEmpty(a))
}

func TestMapIsEmpty(t *testing.T) {
	a := map[string]int{}
	t.Log(iterables.IsEmpty(a))
}
