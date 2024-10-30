package throttler

import (
	"fmt"
	"net/http"
	"time"
)

type Throttler struct {
	maxQPS   int
	requests int
}

func New(maxQPS, requests int) *Throttler {
	return &Throttler{
		maxQPS:   maxQPS,
		requests: requests,
	}
}

func (t *Throttler) Start(url string) {
	requestQueue := make(chan int, t.requests)
	tokens := make(chan struct{}, t.maxQPS)

	// fill bucket depends on QPS
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(t.maxQPS))
		defer ticker.Stop()

		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default: // reject token if bucket is full
			}
		}
	}()

	// fill request queue
	go func() {
		for i := 1; i <= t.requests; i++ {
			requestQueue <- i
		}

		close(requestQueue)
	}()

	// send requests from the queue using throttling algorithm
	for reqID := range requestQueue {
		<-tokens

		go func(id int) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Request %d failed: %v\n", id, err)
				return
			}

			defer resp.Body.Close()
			fmt.Printf("Request %d succeeded.\n", id)
		}(reqID)
	}
}
