package planning

import "testing"
import "fmt"

func TestAStarSearch (t *testing.T) {
  
}

func TestBreadthFirstSearch(t *testing.T) {
  fmt.Print("testing BreadthFirstSearch... ")
  h := NewTowerOfHanoi(3, []int{0,0,0,0}, []int{2,2,2,2})

  solution, error := BreadthFirstSearch(h)

  if error != nil {
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
