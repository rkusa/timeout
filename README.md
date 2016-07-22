# timeout

A timeout middleware that works well (but not exclusively) with [rkusa/web](https://github.com/rkusa/web).

[![Build Status][travis]](https://travis-ci.org/rkusa/timeout)
[![GoDoc][godoc]](https://godoc.org/github.com/rkusa/timeout)

### Example

```go
app := web.New()
app.Use(timeout.Timeout("20ms"))
```

## License

[MIT](LICENSE)

[travis]: https://img.shields.io/travis/rkusa/timeout.svg
[godoc]: http://img.shields.io/badge/godoc-reference-blue.svg
