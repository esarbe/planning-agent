package ai 

import "fmt"
import "collection"

type TowerOfHanoi struct {
  pegs int
  initial HanoiState
  goal HanoiState
}

type HanoiAction struct {
  origin int
  target int
}

type HanoiState []int

func (h HanoiState) Equals (other collection.Hashable) bool {
  o := other.(HanoiState)
  if h.Hash() == o.Hash() {
    return true
  }
  return false
}

func (s HanoiState) Hash() (hash string) {
  hash = fmt.Sprintln(s)
  return
}

func (a HanoiAction) Hash() (hash string) {
  hash = fmt.Sprintln(a)
  return
}

func NewTowerOfHanoi (pegs int, initial []int, goal []int) (t *TowerOfHanoi){
  t = new(TowerOfHanoi)
  t.initial = initial
  t.goal = goal
  t.pegs = pegs
  return
}

// return the first hit
func topblock(s []int, peg int) (top int) {
  top = -1
  for i := range(s) {
    if s[i] == peg {
      top = i
      break
    }
  }
  return
}

func NewHanoiState(blocks int) HanoiState {
  s := make([]int, blocks)
  return s
}

func (t TowerOfHanoi) Goal() (goal State){
  goal = t.goal
  return
}

func (t TowerOfHanoi) H(state State) float32 {
  s := state.(HanoiState)
  h := float32(0)

  for i := range(s) {
    if s[i] != t.goal[i] {
      h = h + 1.0 // 1 penalty for each block on the wrong peg
    }
  }

  return h
}

func (t TowerOfHanoi) IsGoal (goal State) bool {
  return t.goal.Equals(goal)
}

func (t TowerOfHanoi) StepCost(cost float32,
  origin State,
  action Action,
  destination State) (stepCost float32) {
  return cost + float32(1)
}

func (t TowerOfHanoi) Initial() (initial State) {
  initial = t.initial
  return
}

func (t TowerOfHanoi) Successors (state State) (successors map[string]State)  {
  s := state.(HanoiState)
  successors = make(map[string]State)

  // iterate over all possible movements from peg ... [1] 
  for i := 0; i < t.pegs; i++ {
    t_i := topblock(s, i)
    if (t_i == -1) {
      // skip if there is no block on this peg
      continue
    }
    // [1] ...to peg
    for j := 0; j < t.pegs; j++ {
      // skip if there would be no state change  
      if i == j {
        continue
      }
      // get the topmost disk of peg 'j'
      t_j := topblock(s, j)
      if t_i < t_j || t_j == -1 {
        action := &HanoiAction{i,j}
        next := NewHanoiState(len(s))
        //fmt.Println("length of new state vector is:", len(next), " cap is", cap(next))
        //fmt.Println("setting ", t_i, "th to peg ", j)
        copy(next, s)
        next[t_i] = j
        successors[action.Hash()] = next
      }
    }
  }

  return
}

