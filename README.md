# gomonkey

A Go wrapper around SpiderMonkey.

## Installing

To install gomonkey you need a recent (>= 1.8.5) https://developer.mozilla.org/en/SpiderMonkey/ and have the `js-config` program in your `$PATH`.

    > go get github.com/zond/gomonkey
    # github.com/zond/gomonkey
    ld: library not found for -lgomonkey
    collect2: ld returned 1 exit status

The error at the end is expected, for some reason cgo is unable to link properly against recent SpiderMonkeys.

To fix this, do

    > cd $GOPATH/src/github.com/zond/gomonkey
    > make
    cc -fPIC -shared -o libgomonkey.dylib -I/usr/local/Cellar/spidermonkey/1.8.5/include/js -I/usr/local/Cellar/nspr/4.8.8/include/nspr -L/usr/local/Cellar/spidermonkey/1.8.5/lib -lmozjs185  c/gomonkey.c
    > cd -
    > go get github.com/zond/gomonkey

The `Makefile` is *really* trivial and probably won't work on your system (patches are welcome), but if it works you should now have a functional installation!

## Testing

`go test`

## What else?

Currently there isn't much more to do, even if you *could* look at https://github.com/zond/gomonkey/blob/master/gomonkey_test.go and figure out how to use this for other things than testing itself ;)
