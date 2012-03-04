package collection

import "sort"

type PriorityQueue struct {
  HashQueue
  compare Comparison
}

func NewPriorityQueue(c Comparison) (queue *PriorityQueue) {
  queue = new(PriorityQueue)
  queue.compare = c
  queue.HashQueue = *NewHashQueue()
  return
}

func (q *PriorityQueue) Push(c Comparable) {
  q.HashQueue.Push(c)
  q.Sort()
}

func (q *PriorityQueue) Sort() {
  sort.Sort(q)
}

func (q *PriorityQueue) Less(i, j int) bool {
  return q.compare(q.Slice[i], q.Slice[j]) < 0
}

