# timeout

A timeout middleware for [rkgo/web](https://github.com/rkgo/web)

[![Build Status][drone]](https://ci.rkusa.st/rkgo/timeout)
[![GoDoc][godoc]](https://godoc.org/github.com/rkgo/timeout)

### Example

```go
app := web.New()
app.Use(timeout.Timeout("20ms"))
```


[drone]: http://ci.rkusa.st/api/badges/rkgo/timeout/status.svg?tyle=flat-square
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square