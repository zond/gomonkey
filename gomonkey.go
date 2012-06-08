
package gomonkey

/*
#cgo LDFLAGS: -L. -lgomonkey -lmozjs185
#include "c/gomonkey.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

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
	function C.jsval
}

func (self *JSFunction) Call(receiver *JSObject, params... interface{}) interface{} {
	var c_receiver *C.JSObject
	if receiver == nil {
		c_receiver = nil
	} else {
		c_receiver = receiver.object
	}
	var rval C.jsval
	c_params := C.AllocateJsvalArray(C.int(len(params)))
	for index, param := range(params) {
		C.SetJsvalArray(c_params, C.int(index), self.js.goval2jsval(param))
	}
	C.JS_CallFunctionValue(self.js.context, 
		c_receiver,
		self.function, 
		C.uintN(len(params)), 
		c_params,
		&rval)
	return self.js.jsval2goval(rval)
}

type JS struct {
	context *C.JSContext
	global *C.JSObject
}

func NewJS() *JS {
	context := C.NewContext(runtime)
	if context == nil {
		panic("Unable to create context!")
	}

	global := C.NewGlobalObject(context)
	if global == nil {
		panic("Unable to create global object!")
	}

	script := &JS{context, global}
	scripts[script] = true
	return script
}

func (self *JS) Destroy() {
	delete(scripts, self)
	C.DestroyContext(self.context)
}

func (self *JS) goval2jsval(val interface{}) C.jsval {
	if val == nil {
		return C.JsNull()
	} else {
		switch val.(type) {
		case (*JSObject):
		case (*JSFunction):
		case string:
		case float64:
		case bool:
		default:
			panic(fmt.Sprint("Unable to convert", val, "to jsval!"))
		}
	}
	return C.JsNull()
}

func (self *JS) jsval2goval(val C.jsval) interface{} {
	t := C.JS_TypeOfValue(self.context, val)
	if t == C.JSTYPE_VOID {
		return nil
	} else if t == C.JSTYPE_OBJECT {
		return &JSObject{self, C.Jsval2JSObject(self.context, val)}
	} else if t == C.JSTYPE_FUNCTION {
		return &JSFunction{self, val}
	} else if t == C.JSTYPE_STRING {
		c_string := C.JS_EncodeString(self.context, C.Jsval2JSString(self.context, val))
		defer C.free(unsafe.Pointer(c_string))
		return C.GoString(c_string)
	} else if t == C.JSTYPE_NUMBER {
		return float64(C.Jsval2jsdouble(self.context, val))
	} else if t == C.JSTYPE_BOOLEAN {
		return C.Jsval2JSBool(self.context, val) == C.JS_TRUE
	}
	return nil
}

func (self *JS) Eval(script string) interface{} {
	var rval C.jsval
	c_script := C.CString(script)
	defer C.free(unsafe.Pointer(c_script))
	c_scriptname := C.CString("script")
	defer C.free(unsafe.Pointer(c_scriptname))
	C.JS_EvaluateScript(self.context, 
                self.global,
                c_script, 
                C.uintN(len(script)),
                c_scriptname,
                1, 
                &rval)
	return self.jsval2goval(rval)
}

