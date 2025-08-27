package main

import (
	"container/list"
	"sync"
)

var (
	fifoQueue     *list.List
	fifoQueueLock *sync.Mutex
)

func init() {
	fifoQueue = list.New()
	fifoQueueLock = &sync.Mutex{}
}

// enqueueItem adds a call item to the play list, in a locking and
// thread-safe manner.
func enqueueItem(c Call) {
	fifoQueueLock.Lock()
	fifoQueue.PushBack(c)
	fifoQueueLock.Unlock()
}

// consumeQueue loops through an active queue list and executes a
// callback in a locking and thread-safe manner.
func consumeQueue(cb func(c Call)) {
	for fifoQueue.Len() > 0 {
		fifoQueueLock.Lock()
		e := fifoQueue.Front()
		fifoQueueLock.Unlock()

		cb(e.Value.(Call))

		fifoQueueLock.Lock()
		fifoQueue.Remove(e)
		fifoQueueLock.Unlock()
	}
}
