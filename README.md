# go-nacl: Conveniently work with NativeClient Go

`go-nacl` sets the environment, and `-exec=` argument where
appropriate, so that your builds use NativeClient (NaCl). It handles
details like `GOARCH=amd64p32` and `go test -exec=sel_ldr_x86_64` for
you. Whenever you want to use NaCl, you just type `go-nacl` instead of
`go`.

``` console
$ go-nacl build
$ go-nacl test
$ go-nacl run simple.go
```
