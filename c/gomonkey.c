
#include "gomonkey.h"
#include <stdio.h>

static JSClass global_class = { "global", JSCLASS_GLOBAL_FLAGS, JS_PropertyStub, JS_PropertyStub, JS_PropertyStub, JS_StrictPropertyStub, JS_EnumerateStub, JS_ResolveStub, JS_ConvertStub, JS_FinalizeStub, JSCLASS_NO_OPTIONAL_MEMBERS };

void reportError(JSContext *cx, const char *message, JSErrorReport *report) {
  fprintf(stderr, "%s:%u:%s\n", report->filename ? report->filename : "<no filename>", (unsigned int) report->lineno, message);
}

jsval
JsNull() {
  return JSVAL_NULL;
}

jsval
JsTrue() {
  return JSVAL_TRUE;
}

jsval
JsFalse() {
  return JSVAL_FALSE;
}

JSObject*
Jsval2JSObject(JSContext *context, jsval val) {
  jsval out;
  if (JS_ConvertValue(context, val, JSTYPE_OBJECT, &out)) {
    return (JSObject*) out;
  } else {
    return NULL;
  }
}

JSFunction*
Jsval2JSFunction(JSContext *context, jsval val) {
  jsval out;
  if (JS_ConvertValue(context, val, JSTYPE_FUNCTION, &out)) {
    return (JSFunction*) out;
  } else {
    return NULL;
  }
}

JSString*
Jsval2JSString(JSContext *context, jsval val) {
  return JS_ValueToString(context, val);
}

jsdouble
Jsval2jsdouble(JSContext *context, jsval val) {
  jsdouble out;
  if (JS_ValueToNumber(context, val, &out)) {
    return out;
  } else {
    return 0.0;
  }
}

JSBool
Jsval2JSBool(JSContext *context, jsval val) {
  jsval out;
  if (JS_ConvertValue(context, val, JSTYPE_BOOLEAN, &out)) {
    return (JSBool) out;
  } else {
    return JS_FALSE;
  }
}

jsval*
AllocateJsvalArray(int size) {
  return (jsval*) malloc(sizeof(jsval) * size);
}

jsval
JSObject2Jsval(JSObject *obj) {
  return OBJECT_TO_JSVAL(obj);
}

jsval
JSString2Jsval(JSString *s) {
  return STRING_TO_JSVAL(s);
}

jsval
Double2Jsval(double d) {
  return DOUBLE_TO_JSVAL(d);
}

void
SetJsvalArray(jsval *ary, int index, jsval val) {
  ary[index] = val;
}

JSContext*
NewContext(JSRuntime* runtime) {
  JSContext *context = JS_NewContext(runtime, 8192);

  if (context == NULL) {
    return NULL;
  } else {
    JS_SetOptions(context, JSOPTION_VAROBJFIX | JSOPTION_JIT | JSOPTION_METHODJIT);
    JS_SetVersion(context, JSVERSION_LATEST);
    JS_SetErrorReporter(context, reportError);
    
    return context;
  }
}
  
JSObject*
NewGlobalObject(JSContext* context) {
  JSObject *global = JS_NewCompartmentAndGlobalObject(context, &global_class, NULL);
  if (global == NULL) {
    return NULL;
  } else {
    JS_InitStandardClasses(context, global);
    
    return global;
  }
}

void
DestroyContext(JSContext *context) {
  JS_DestroyContext(context);
}

