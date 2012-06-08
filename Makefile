
CFLAGS=`js-config --cflags`
LIBS=`js-config --libs`
DYNLIB=libgomonkey.dylib

all:	c/gomonkey.c c/gomonkey.h
	${CC} -fPIC -shared -o ${DYNLIB} ${CFLAGS} ${LIBS} c/gomonkey.c

clean:
	rm ${DYNLIB}