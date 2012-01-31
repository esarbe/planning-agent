package planning

import "sort"

type PriorityQueue struct {
  HashQueue
}

func NewPriorityQueue() (queue *PriorityQueue) {
  queue = new(PriorityQueue)
  queue.HashQueue = *NewHashQueue()
  return
}

func (q *PriorityQueue) Push(c Comparable) {
  q.HashQueue.Push(c)
  q.Sort()
}

func (q *PriorityQueue) Sort() {
  sort.Sort(q.Slice)
}
