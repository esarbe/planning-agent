package planning

import "testing"
import "fmt"

func TestQueue(t *testing.T) {
  q := &PriorityQueue{}


  n1 := NamedNode{nil, nil, float32(1), nil}
  n2 := NamedNode{nil, nil, float32(2), nil}
  n3 := NamedNode{nil, nil, float32(3), nil}
  n4 := NamedNode{nil, nil, float32(4), nil}

  s := PriorityQueue {&n1, &n2, &n3, &n4}


  s.Push(&n1)

  q.Push(&n4)
  q.Push(&n1)
  q.Push(&n3)
  q.Push(&n2)

  fmt.Println("addr q:", q)

  n := q.Pop()
  fmt.Println("addr q:", q)
  if q.Len() != 3 {
    fmt.Println("length expected: 3, actual:", q.Len())
    t.FailNow()
  }

  if n.PathCost() != 1 {
    fmt.Println("expected pathcost: 1, actual:", n.PathCost())
    t.FailNow()
  }

  fmt.Println("addr q:", q)
  n4.SetPathCost(float32(0.1))
  fmt.Println("n4 pathcost:", n4.PathCost())
  q.Sort()

  fmt.Println("addr q:", q)
  for _, n0 := range *q {
    fmt.Println("n0 is : ", n0.PathCost())
  }

  n = q.Pop()

  if n.PathCost() != 0.1 {
    fmt.Println("expected pathcost: 0.1 , actual:", n.PathCost())
    t.FailNow()
  }

  if q.Len() != 2 {
    fmt.Println("length expected: 2, actual:", q.Len())
    t.FailNow()
  }

}

