package signalx

import (
	"os"
	"os/signal"
	"sync"
)

type handler struct {
	rnn  bool
	lock sync.Mutex
	ch   chan os.Signal
}

// Stop receiving signals for this specific handler
func (h *handler) Stop() {
	h.lock.Lock()
	defer h.lock.Unlock()

	defer func() {
		h.rnn = false
	}()

	if h.rnn {
		close(h.ch)
	}
}

// Handle specified signals with a callback function
// Returns a handler struct which allows to close the opened channel
func Handle(c func(), s ...os.Signal) *handler {
	var ch = make(chan os.Signal, len(s))

	signal.Notify(ch, s...)

	go func() {
		for range ch {
			c()
		}

		signal.Stop(ch)
	}()

	return &handler{rnn: true, ch: ch}
}
