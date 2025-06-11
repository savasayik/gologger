# gologger

A powerful, generic, context-aware logger built on top of [zerolog](https://github.com/rs/zerolog), designed for production-grade applications in Go.

## ✨ Features

- Generic logger: `Logger[T]`
- Global logger instance
- Context injection (`WithContext`)
- Error stack logging (`WithErrorStack`)
- Structured domain logging (`StructuredLog`, `StructuredDebug`, `StructuredError`)
- HTTP middleware support (`X-Request-ID`, `X-Trace-ID`, `X-User-ID` headers)

## 📦 Installation

```bash
go get github.com/savasayik/gologger
```

## 🚀 Usage

### Initialization

```go
gologger.InitLogger(gologger.DebugLevel, "my-service")
```

### Structured Logging

```go
type Order struct {
    ID     int
    Amount float64
}

gologger.GetLogger().StructuredDebug("order_created", Order{ID: 1, Amount: 100.0})
```

### Error Logging with Stack

```go
err := errors.New("DB connection failed")
gologger.GetLogger().WithErrorStack(err, "database error")
```

### HTTP Middleware

```go
http.Handle("/", gologger.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    gologger.GetLogger().WithContext(r.Context()).Msg("request received")
    w.Write([]byte("ok"))
})))
```

## 📄 License

MIT
