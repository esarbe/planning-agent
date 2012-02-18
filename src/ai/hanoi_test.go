package ai

import "testing"
import "fmt"

func TestNewHanoi(t *testing.T) {
  tower := &TowerOfHanoi{3, HanoiState{0,0,0,0}, HanoiState{2,2,2,2}}
  fmt.Println("built new tower")

  successors := tower.Successors(HanoiState{0,0,0,0})

  fmt.Println("Successors:", successors)
}
