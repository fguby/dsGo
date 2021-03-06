/*
	A queue gives you a FIFO or first-in firs-out order.
	The element you inserted first is also the first one to come out again.
*/
package queue

import (
	"sync"

	. "github.com/zrcoder/dsGo"
	base "github.com/zrcoder/dsGo/base/queue"
)

type Queen struct {
	sync.Mutex
	inner *base.Queen
}

func New() *Queen {
	return &Queen{inner: base.New()}
}

// put a given item into the queue
func (q *Queen) Enqueue(item Any) {
	q.Lock()
	q.inner.Enqueue(item)
	q.Unlock()
}

// remove the first item from the queue and returns it
func (q *Queen) Dequeue() Any {
	q.Lock()
	item := q.inner.Dequeue()
	q.Unlock()
	return item
}
