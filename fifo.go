package main

import (
	"container/list"
)

var (
	fifoQueue *list.List
)

func init() {
	fifoQueue = list.New()
}

func enqueueItem(c Call) {
	fifoQueue.PushBack(c)
}

func consumeQueue(cb func(c Call)) {
	for fifoQueue.Len() > 0 {
		e := fifoQueue.Front()
		cb(e.Value.(Call))
		fifoQueue.Remove(e)
	}
}
