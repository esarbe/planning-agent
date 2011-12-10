package "problem"

import "state"

type Action interface {
  IsAllowed(a State) bool //asks whether the given action is possible this state
  Do(a State) State //execute action and return new state
}

type Problem interface {
  Child(parent State, action Action) (child Node)
  Inital() State
  Goal() State
}

type Solver interface {
  Solve(problem Problem) (solution Node) 
}

