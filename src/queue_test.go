package planning

import "testing"
import "fmt"

type intState int

func (s intState) Equals (o Hashable) bool {
  return s == o.(intState)
}

func (s intState) Hash() string {
  return fmt.Sprintln(s)
}

func (s intState) Compare (o Comparable) int {
  return int(s - o.(intState))
}

func TestSlice(t *testing.T) {
  fmt.Println("testing slice push")
}

func TestPriorityQueue(t *testing.T) {
  fmt.Println("testing priority queue...")
  q := NewPriorityQueue()

  for i :=0; i < 100; i+=1 {
    n := NewNode(nil, nil, intState(i))
    n.cost = float32(i)
    q.Push(n)
  }

  if q.Len() != 100 {
    fmt.Println("queue does not have expected length")
  }

  fmt.Println("done.")

  fmt.Println("testing queue pop...")
  for i := 0; i < 50; i+=1 {

    q.Pop()
  }

  if q.Len() != 50 {
    fmt.Println("queue does not have expected length of 50, is ", q.Len()," instead")
  }

  fmt.Println("done.")
}

