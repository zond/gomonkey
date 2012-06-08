
#include "gomonkey.h"
#include <stdio.h>

static JSClass global_class = { "global", JSCLASS_GLOBAL_FLAGS, JS_PropertyStub, JS_PropertyStub, JS_PropertyStub, JS_StrictPropertyStub, JS_EnumerateStub, JS_ResolveStub, JS_ConvertStub, JS_FinalizeStub, JSCLASS_NO_OPTIONAL_MEMBERS };

void reportError(JSContext *cx, const char *message, JSErrorReport *report) {
  fprintf(stderr, "%s:%u:%s\n", report->filename ? report->filename : "<no filename>", (unsigned int) report->lineno, message);
}

JSContext*
NewContext(JSRuntime* runtime) {
  JSContext *context = JS_NewContext(runtime, 8192);

  JS_SetOptions(context, JSOPTION_VAROBJFIX | JSOPTION_JIT | JSOPTION_METHODJIT);
  JS_SetVersion(context, JSVERSION_LATEST);
  JS_SetErrorReporter(context, reportError);

  return context;
}

JSObject*
NewGlobalObject(JSContext* context) {
  JSObject *global = JS_NewCompartmentAndGlobalObject(context, &global_class, NULL);
  JS_InitStandardClasses(context, global);

  return global;
}

void
DestroyContext(JSContext *context) {
  JS_DestroyContext(context);
}

