# timeout

A timeout middleware for [rkgo/web](https://github.com/rkgo/web)

[![Build Status][drone]](https://ci.rkusa.st/github.com/rkgo/timeout)
[![GoDoc][godoc]](https://godoc.org/github.com/rkgo/timeout)

### Example

```go
app := web.New()
app.Use(timeout.Timeout("20ms"))
```


[drone]: http://ci.rkusa.st/api/badge/github.com/rkgo/timeout/status.svg?branch=master&style=flat-square
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square