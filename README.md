# signalx

Silly go package to handle signals

```go
import "github.com/jaimelopez/signalx"

sign : = signalx.Handle(func() {
    fmt.Println("Signal HUP received")
}, syscall.SIGHUP)


// In case you want to stop receiving the specified signals:
sign.Stop()
```
