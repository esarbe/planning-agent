package planning

import "testing"
import "fmt"


func TestPush(t *testing.T) {
  fmt.Println("testing queue push...")
  q := &PriorityQueue{}

  for i :=0; i < 100; i+=1 {
    n := NamedNode {nil, nil, float32(i), nil}
    q.Push(&n)
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
    fmt.Println("queue does not have expected length")
  }

  fmt.Println("done.")
}


func TestSort(t *testing.T) {
  fmt.Println("testing queue sort...")

  n1 := NamedNode{nil, nil, float32(1), nil}
  n2 := NamedNode{nil, nil, float32(2), nil}
  n3 := NamedNode{nil, nil, float32(3), nil}
  n4 := NamedNode{nil, nil, float32(4), nil}

  q := PriorityQueue {&n1, &n2, &n3, &n4}

  n := q.Pop()

  if n.PathCost() != 1 {
    fmt.Println("expected pathcost: 1, actual:", n.PathCost())
    t.FailNow()
  }

  n4.SetPathCost(float32(0.1))
  q.Sort()

  n = q.Pop()

  if n.PathCost() != 0.1 {
    fmt.Println("expected pathcost: 0.1 , actual:", n.PathCost())
    t.FailNow()
  }

  fmt.Println("done.")
}

