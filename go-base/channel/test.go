package channel

import "sync"

var mu = sync.Mutex{} //互斥锁

var rwmu = sync.RWMutex{}

var ch = make(chan struct{}, 1)

func UseMutex() {
	mu.Lock()
	mu.Unlock()
}

func UseRWMutex() {
	rwmu.Lock()
	rwmu.Unlock()
}

func UseChan() {
	ch <- struct{}{}
	<-ch
}
