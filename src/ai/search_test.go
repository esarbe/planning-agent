package ai

import "testing"
import "fmt"

func BenchmarkAStarSearch(b *testing.B) {
  h := NewTowerOfHanoi(4, []int{0,0,0,0}, []int{4,4,4,4})

  for i := 0; i < b.N; i++ {
    AStarSearch(h)
  }
}

func BenchmarkBreadthFirstSearch(b *testing.B) {
  h := NewTowerOfHanoi(4, []int{0,0,0,0}, []int{4,4,4,4})

  for i := 0; i < b.N; i++ {
    BreadthFirstSearch(h)
  }
}

func TestAStarSearch (t *testing.T) {
  fmt.Println("testing AStarSearch... ")
  h := NewTowerOfHanoi(3, []int{0,0,0,0}, []int{2,2,2,2})

  solution, error := AStarSearch(h)

  if error != "" {
    fmt.Println("no solution found!")
    t.FailNow()
  }

  length := 0
  for solution.parent != nil {
    solution = *(solution.parent)
    length += 1
  }

  if length != 15 {
    fmt.Println("solution has unexpected length of", length)
    t.FailNow()
  }

  fmt.Println("done")
}


func TestBreadthFirstSearch(t *testing.T) {
  fmt.Print("testing BreadthFirstSearch... ")
  h := NewTowerOfHanoi(3, []int{0,0,0,0}, []int{2,2,2,2})

  solution, error := BreadthFirstSearch(h)

  if error != "" {
    fmt.Println("no solution found!")
    t.FailNow()
  }

  length := 0
  for solution.parent != nil {
    solution = *(solution.parent)
    length += 1
  }

  if length != 15 {
    fmt.Println("solution has unexpected length of", length)
    t.FailNow()
  }
  fmt.Println("done")
}
