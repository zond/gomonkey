
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
		t.Error(x, "should be 100")
	}
}

func TestTrueReturn(t *testing.T) {
	script := NewJS()
	defer script.Destroy()
	x := script.Eval("1 == 1")
	if x != true {
		t.Error(x, "should be true")
	}
}

func TestFalseReturn(t *testing.T) {
	script := NewJS()
	defer script.Destroy()
	x := script.Eval("1 != 1")
	if x != false {
		t.Error(x, "should be false")
	}
}

func TestStringReturn(t *testing.T) {
	script := NewJS()
	defer script.Destroy()
	x := script.Eval("\"hej\" + \" kompis\"")
	if x != "hej kompis" {
		t.Error(x, "should be \"hej kompis\"")
	}
}

func TestObjectReturn(t *testing.T) {
	script := NewJS()
	defer script.Destroy()
	x := script.Eval("new Object()")
	if reflect.TypeOf(x) != reflect.TypeOf(&JSObject{}) {
		t.Error(x, "should be a JSObject")
	}
}

func TestFunctionReturn(t *testing.T) {
	script := NewJS()
	defer script.Destroy()
	x := script.Eval("function() { return 1; }")
	if reflect.TypeOf(x) != reflect.TypeOf(&JSFunction{}) {
		t.Error(x, "should be a JSFunction")
	}
}