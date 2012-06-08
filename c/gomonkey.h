
#include <js/jsapi.h>

extern JSContext*
NewContext(JSRuntime* runtime);

extern JSObject*
NewGlobalObject(JSContext *context);

extern void
DestroyContext(JSContext *context);

extern jsval
JsNull();

extern JSObject*
Jsval2JSObject(JSContext*, jsval);

extern JSFunction*
Jsval2JSFunction(JSContext*, jsval);

extern JSString*
Jsval2JSString(JSContext*, jsval);

extern jsdouble
Jsval2jsdouble(JSContext*, jsval);

extern JSBool
Jsval2JSBool(JSContext*, jsval);

extern jsval*
AllocateJsvalArray(int);

extern void
SetJsvalArray(jsval*, int, jsval);
