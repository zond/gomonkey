
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
	object *C.JSObject
}

type JSFunction struct {
	js *JS
	function *C.JSFunction
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

func (self *JS) jsval2goval(val C.jsval) interface{} {
	t := C.JS_TypeOfValue(self.context, val)
	if t == C.JSTYPE_VOID {
		return nil
	} else if t == C.JSTYPE_OBJECT {
		return &JSObject{self, C.Jsval2JSObject(self.context, val)}
	} else if t == C.JSTYPE_FUNCTION {
		return &JSFunction{self, C.Jsval2JSFunction(self.context, val)}
	} else if t == C.JSTYPE_STRING {
		return C.GoString(C.JS_EncodeString(self.context, C.Jsval2JSString(self.context, val)))
	} else if t == C.JSTYPE_NUMBER {
		return float64(C.Jsval2jsdouble(self.context, val))
	} else if t == C.JSTYPE_BOOLEAN {
		return C.Jsval2JSBool(self.context, val) == C.JS_TRUE
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

