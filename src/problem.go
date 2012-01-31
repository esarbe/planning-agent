package planning

type NoSolutionError struct {
  message string
}

type Action interface {}

type State interface {
  Hashable
}

type Node struct {
  cost float32
  state State
  action Action
  parent *Node
}

type Error struct {
  message string
}

func NewNode(parent *Node, action Action, state State) (n *Node) {
  n = new(Node)
  n.state = state
  n.action = action
  n.parent = parent
  return
}

func (n Node) Equals (other Hashable) bool {
  o := other.(*Node)
  if n.state.Equals(o.state) {
    return true
  }
  return false
}

func (n Node) Hash() string {
  return n.state.Hash()
}

func (n Node) Compare (other Comparable) int {
  o := other.(*Node)

  if n.cost < o.cost {
    return -1
  } else if n.cost == o.cost {
    return 0
  } else if n.cost > o.cost {
    return 1
  }
  panic("there is something seriously wrong..")
}

type Problem interface {
  Initial() (initial State)
  Goal() (goal State)
  IsGoal(state State) bool
  StepCost(cost float32, origin State, action Action, destination State) (stepCost float32)
  Successors(state State) (successors map[string]State)
}

type Solver interface {
  Solve(problem Problem) (solution interface{}, error NoSolutionError)
}

