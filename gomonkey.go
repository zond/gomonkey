
package gomonkey

/*
#cgo LDFLAGS: -L. -lgomonkey -lmozjs185
#include "c/gomonkey.h"
*/
import "C"

var runtime *C.JSRuntime
var scripts map[*JS]bool = make(map[*JS]bool)

func init() {
	runtime = C.JS_NewRuntime(1024 * 1024)
}

func Shutdown() {
	for script,_ := range(scripts) {
		script.Destroy()
	}
	C.JS_DestroyRuntime(runtime)
	C.JS_ShutDown()
}

type JSObject struct {
	js *JS
	value *C.JSObject
}

type JSFunction struct {
	js *JS
	value *C.JSFunction
}

type JS struct {
	context *C.JSContext
	global *C.JSObject
}

func NewJS() *JS {
	context := C.NewContext(runtime)
	global := C.NewGlobalObject(context)
	script := &JS{context, global}
	scripts[script] = true
	return script
}

func (self *JS) Destroy() {
	delete(scripts, self)
	C.DestroyContext(self.context)
}
/*
func (self *JS) goval2jsval(val interface{}) C.jsval {
	if val == nil {
		return C.JSVAL_NULL
	} else {
		switch t := val.(type) {
		case *JSObject:
			var rval C.jsval
			C.JS_ValueToObject(self.context, rval, val.(*JSObject).value)
			return rval
		case *JSFunction:
			return C.JS_ValueToFunction(self.contect, val.(*JSFunction).value)
		}
	}
}
*/
func (self *JS) jsval2goval(val C.jsval) interface{} {
	t := C.JS_TypeOfValue(self.context, val)
	if t == C.JSTYPE_VOID {
		return nil
	} else if t == C.JSTYPE_OBJECT {
		var obj C.JSObject
		var obj_p = &obj
		C.JS_ValueToObject(self.context, val, &obj_p)
		return &JSObject{self, &obj}
	} else if t == C.JSTYPE_FUNCTION {
		return &JSFunction{self, C.JS_ValueToFunction(self.context, val)}
	} else if t == C.JSTYPE_STRING {
		return C.GoString(C.JS_EncodeString(self.context, C.JS_ValueToString(self.context, val)))
	} else if t == C.JSTYPE_NUMBER {
 		var rval C.jsdouble
		C.JS_ValueToNumber(self.context, val, &rval)
		return float64(rval)
	} else if t == C.JSTYPE_BOOLEAN {
		var rval C.JSBool
		C.JS_ValueToBoolean(self.context, val, &rval)
		return rval == C.JS_TRUE
	}
	return nil
}

func (self *JS) Eval(script string) interface{} {
	var rval C.jsval
	C.JS_EvaluateScript(self.context, 
                self.global,
                C.CString(script), 
                C.uintN(len(script)),
                C.CString("script"),
                1, 
                &rval)
	return self.jsval2goval(rval)
}

