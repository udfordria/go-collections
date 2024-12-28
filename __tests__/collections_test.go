package __tests__

import (
	"strings"
	"testing"

	"github.com/udfordria/go-collections/lists"
)

func TestMap(t *testing.T) {
	a := []string{"a", "b", "c"}
	res := lists.Map(a, func(a string) string { return a + "." })
	t.Log(strings.Join(res, ","))
}
