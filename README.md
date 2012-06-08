# gomonkey

A Go wrapper around SpiderMonkey

## Using

To use gomonkey you must install a recent (>= 1.8.5) https://developer.mozilla.org/en/SpiderMonkey/ in a location found by gcc.

Then you build the small helper class (necessary because cgo for some reason is unable to link properly against SpiderMonkey) by `make`.

The `Makefile` is *really* trivial and probably won't work on your system. Patches are welcome.

Then you just `go test`. Currently there isn't much more to do, even if you *could* look at `gomonkey_test.go` and figure out how to use this for other things than testing itself ;)
