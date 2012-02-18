package collection

type FifoQueue struct {
  HashQueue
}

func NewFifoQueue() *FifoQueue {
  q := new(FifoQueue)
  q.HashQueue = *NewHashQueue()
  return q
}
