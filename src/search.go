package "search"
import "state"
import "problem"


type BreadthFirst struct {
  frontier map[State] bool
  known [State] bool
}

func Min(actions []Action) (best Action) {
  var min float32 = math.MaxFloat32
  for _, action := range actions {
    if action.Cost() < min {
      best = action
    }
  }
  return
  }

func Min(states []WeightedState ) (best WeightedState) {
  var min float = math.MaxFloat32
  for _, state := range states {
    if (state.Weight() < min {
      best = state
    }
  }
  return
}

func (b BreadthFirst) Solve (start WeightedState, end WeightedState) []Action {
  for _, action in range start.ListActions() {
    next := action.Do(start)

    currentWeight := start.Weight() + action.Cost() 

    if next.Weight() > currentWeight {
      next.SetWeight(currentWeight)
    }

    if !b.frontier[next] && !b.known[next] {
      b.frontier[next] = true 
    }
  }

  b.frontier[start] = false, false
  next := Min(b.frontier 
}

