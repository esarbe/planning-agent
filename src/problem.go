package planning

type NoSolutionError struct {
  message string
}

type Action interface {
  IsAllowed(a State) bool //asks whether the given action is possible this state
  Do(a State) State //execute action and return new state
}

type Problem interface {
  Child(parent State, action Action) (child Node)
  Inital() (initial State)
  Goal() (goal State)
  IsGoal(state State) (isGoal bool)
  StepCost(state State) (stepCost float32)
  Actions(state State) (actions []Action)
}

type Solver interface {
  Solve(problem Problem) (solution Node, error NoSolutionError)
}

