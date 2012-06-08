
package gomonkey

/*
#cgo LDFLAGS: -L. -lgomonkey -lmozjs185
#include "c/gomonkey.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
	"runtime"
)

var jsruntime *C.JSRuntime
var scripts map[*JS]bool = make(map[*JS]bool)

func init() {
	jsruntime = C.JS_NewRuntime(1024 * 1024)
}

func Shutdown() {
	for script,_ := range(scripts) {
		script.Destroy()
	}
	C.JS_DestroyRuntime(jsruntime)
	C.JS_ShutDown()
}

type JSObject struct {
	js *JS
	object *C.JSObject
}

func NewJSObject(js *JS, object *C.JSObject) *JSObject {
	rval := &JSObject{js, object}
	C.JS_AddObjectRoot(rval.js.context, &(rval.object))
	runtime.SetFinalizer(rval, func(o *JSObject) {
		C.JS_RemoveObjectRoot(o.js.context, &(o.object))
	})
	return rval
}

func (self *JSObject) Equals(other *JSObject) bool {
	return self.js == other.js && self.object == other.object
}

type JSFunction struct {
	js *JS
	function C.jsval
}

func NewJSFunction(js *JS, function C.jsval) *JSFunction {
	rval := &JSFunction{js, function}
	C.JS_AddValueRoot(rval.js.context, &(rval.function))
	runtime.SetFinalizer(rval, func(f *JSFunction) {
		C.JS_RemoveValueRoot(f.js.context, &(f.function))
	})
	return rval
}

func (self *JSFunction) Equals(other *JSFunction) bool {
	return (self.js == other.js) && (self.function == other.function)
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
	context := C.NewContext(jsruntime)
	if context == nil {
		panic("Unable to create context!")
	}

	global := C.NewGlobalObject(context)
	if global == nil {
		panic("Unable to create global object!")
	}

	script := &JS{context, global}
	scripts[script] = true

	runtime.SetFinalizer(script, func(script *JS) {
		script.Destroy()
	})

	return script
}

func (self *JS) Destroy() {
	if _,ok := scripts[self]; ok {
		delete(scripts, self)
		C.DestroyContext(self.context)
	}
}

func (self *JS) goval2jsval(val interface{}) C.jsval {
	if val == nil {
		return C.JsNull()
	} else {
		switch val.(type) {
		case (*JSObject):
			return C.JSObject2Jsval(val.(*JSObject).object)
		case (*JSFunction):
			return val.(*JSFunction).function
		case string:
			c_string := C.CString(val.(string))
			defer C.free(unsafe.Pointer(c_string))
			js_string := C.JS_NewStringCopyZ(self.context, c_string)
			return C.JSString2Jsval(js_string)
		case float64:
			return C.Double2Jsval(C.double(val.(float64)))
		case bool:
			if val.(bool) == true {
				return C.JsTrue()
			} else {
				return C.JsFalse()
			}
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
		return NewJSFunction(self, val)
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

