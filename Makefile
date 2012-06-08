
CFLAGS := $(shell js-config --cflags)
LIBS := $(shell js-config --libs)
UNAME := $(shell uname)

ifeq (${UNAME}, Darwin)
 DYNLIB := libgomonkey.dylib
endif
ifeq (${UNAME}, Linux)
 DYNLIB := libgomonkey.so
endif

all:	c/gomonkey.c c/gomonkey.h
	${CC} -fPIC -shared -o ${DYNLIB} ${CFLAGS} ${LIBS} c/gomonkey.c

clean:
	rm ${DYNLIB}