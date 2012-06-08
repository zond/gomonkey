# gomonkey

A Go wrapper around SpiderMonkey.

## Using

To use gomonkey you must install a recent (>= 1.8.5) https://developer.mozilla.org/en/SpiderMonkey/ and have the `js-config` program in your `$PATH`.

Then you build the small helper class (necessary because for some reason cgo is unable to link properly against recent SpiderMonkeys) by `make`.

The `Makefile` is *really* trivial and probably won't work on your system. Patches are welcome.

Then you just `go test`. Currently there isn't much more to do, even if you *could* look at `gomonkey_test.go` and figure out how to use this for other things than testing itself ;)
