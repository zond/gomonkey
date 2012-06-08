
#include <js/jsapi.h>

extern JSContext*
NewContext(JSRuntime* runtime);

extern JSObject*
NewGlobalObject(JSContext *context);

extern void
DestroyContext(JSContext *context);
