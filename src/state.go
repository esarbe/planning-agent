package state

import "math"

type State interface {
  ListActions() []Action
  Name() string
}

type Node interface {
  State() State
  Parent() State
  PathCost() float32
}

type Solver interface {
  Solve(problem Problem) *[]Action
}

// implementations
type NamedState struct {
  name string
  weight float32
  actions []Action
}

type NamedAction struct {
  name string
  execute func(State) State
  validate func(State) bool
}

func (a NamedAction) Cost() float32 {
  return 0.0
}

func (a NamedAction) Allowed(s State) bool {
  if s.Name() == "StartState" {
    return true
  }
  return false
}

func (s NamedState) Name() string {
  return s.name
}

func (s NamedState) ListActions() []Action {
  return s.actions
}

func (a NamedAction) Do(s State) State {
  return nil
}

func (a NamedAction) IsAllowed(s State) bool {
  return a.validate(s)
}

func (s NamedState) Weight() float32 {
  return s.weight
}

func (s NamedState) SetWeight(weight float32) {
  s.weight = weight
  return
}

