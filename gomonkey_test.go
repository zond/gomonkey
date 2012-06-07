
package gomonkey

import (
	"testing"
)

func TestBasicEvaluation(t *testing.T) {
	script := NewJS()
	x := script.Eval("x = 10; x * x;")
	if x != 100.0 {
		t.Error(x,"Should be 100")
	}
	script.Destroy()
}
