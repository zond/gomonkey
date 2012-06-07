
package gomonkey

import (
	"testing"
	"reflect"
)

func TestBasicEvaluation(t *testing.T) {
	script := NewJS()
	defer script.Destroy()
	x := script.Eval("x = 10; x * x;")
	if x != 100.0 {
		t.Error(x, "Should be 100")
	}
}

func TestObjectReturn(t *testing.T) {
	script := NewJS()
	defer script.Destroy()
	x := script.Eval("new Object()")
	if reflect.TypeOf(x) != reflect.TypeOf(&JSObject{}) {
		t.Error(x, "Should be a JSObject")
	}
}

func TestFunctionReturn(t *testing.T) {
	script := NewJS()
	defer script.Destroy()
	x := script.Eval("function() { return 1; }")
	if reflect.TypeOf(x) != reflect.TypeOf(&JSFunction{}) {
		t.Error(x, "Should be a JSFunction")
	}
}