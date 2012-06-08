
all:	c/gomonkey.c c/gomonkey.h
	gcc -fPIC -shared -o libgomonkey.dylib c/gomonkey.c -lmozjs185