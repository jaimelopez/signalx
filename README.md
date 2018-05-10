# signalx
Silly go package to handle signals

```go
import "github.com/jaimelopez/signalx"

signalx.Handle(func() {
    fmt.Println("Signal HUP received")
}, syscall.SIGHUP)
```
