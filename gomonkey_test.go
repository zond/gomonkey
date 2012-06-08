
package gomonkey

import (
	"testing"
	"reflect"
)

func TestBasicEvaluation(t *testing.T) {
	script := NewJS()
	x := script.Eval("x = 10; x * x;")
	if x != 100.0 {
		t.Error(x, "should be 100")
	}
}

func TestTrueReturn(t *testing.T) {
	script := NewJS()
	x := script.Eval("1 == 1")
	if x != true {
		t.Error(x, "should be true")
	}
}

func TestFalseReturn(t *testing.T) {
	script := NewJS()
	x := script.Eval("1 != 1")
	if x != false {
		t.Error(x, "should be false")
	}
}

func TestStringReturn(t *testing.T) {
	script := NewJS()
	x := script.Eval("\"hej\" + \" kompis\"")
	if x != "hej kompis" {
		t.Error(x, "should be \"hej kompis\"")
	}
}

func TestObjectReturn(t *testing.T) {
	script := NewJS()
	x := script.Eval("new Object()")
	if reflect.TypeOf(x) != reflect.TypeOf(&JSObject{}) {
		t.Error(x, "should be a JSObject")
	}
}

func TestFunctionReturn(t *testing.T) {
	script := NewJS()
	x := script.Eval("function x() { return 1; }; x")
	if reflect.TypeOf(x) != reflect.TypeOf(&JSFunction{}) {
		t.Error(x, "should be a JSFunction")
	}
	if x.(*JSFunction).Call(nil) != 1.0 {
		t.Error(x, "should return 1.0")
	}
}

func TestFunctionArguments(t *testing.T) {
	script := NewJS()
	x := script.Eval("function x(y) { return y; }; x")
	if x.(*JSFunction).Call(nil, 1.2) != 1.2 {
		t.Error(x, "should return 1.2")
	}
	if x.(*JSFunction).Call(nil, "blep") != "blep" {
		t.Error(x, "should return \"blep\"")
	}
	if !x.(*JSFunction).Call(nil, x).(*JSFunction).Equals(x.(*JSFunction)) {
		t.Error(x, "should return", x)
	}
	
}