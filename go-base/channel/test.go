package channel

import "sync"

var mu = sync.Mutex{}

var ch = make(chan struct{}, 1)

func UseMutex() {
	mu.Lock()
	mu.Unlock()
}

func UseChan() {
	ch <- struct{}{}
	<-ch
}
