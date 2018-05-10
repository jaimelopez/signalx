package signalx

import (
	"os"
	"os/signal"
)

// Handle specified signals with a callback function
// Returns a function which allows to close the opened channel
func Handle(c func(), s ...os.Signal) func() {
	ch := make(chan os.Signal)

	signal.Notify(ch, s...)

	go func() {
		for range ch {
			c()
		}
	}()

	return func() {
		close(ch)
	}
}
