package ai

import . "collection"

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

func NewNode(parent *Node, action Action, state State) (n *Node) {
  n = new(Node)
  n.state = state
  n.action = action
  n.parent = parent
  return
}

func (n Node) Cost() float32 {
  return n.cost
}

func (n Node) State() State {
  return n.state
}

func (n Node) Action() Action {
  return n.action
}

func (n Node) Parent() *Node {
  return n.parent
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

  return int(n.cost - o.cost)
}

type Problem interface {
  Initial() (initial State)
  Goal() (goal State)
  IsGoal(state State) bool
  StepCost(cost float32, origin State, action Action, destination State) (stepCost float32)
  Successors(state State) (successors map[string]State)
}

type HeuristicProblem interface {
  Problem
  H(state State) float32
}

type Solver interface {
  Solve(problem Problem) (solution interface{}, error NoSolutionError)
}

