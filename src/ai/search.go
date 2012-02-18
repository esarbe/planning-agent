package ai

import . "collection"

func (error *NoSolutionError) String() string {
  return error.message
}

type Error string

func equals(lhs Comparable, rhs Comparable) bool {
  ls := lhs.(*Node).state
  rs := rhs.(*Node).state
  return ls.Equals(rs)
}

func BreadthFirstSearch(problem Problem) (solution Node, error Error) {
  return Search(problem, NewFifoQueue())
}

func AStarSearch(problem HeuristicProblem) (solution Node, error Error) {

  compareNodes := func(lhs Comparable, rhs Comparable) float32 {
    lhsNode := lhs.(*Node)
    rhsNode := rhs.(*Node)

    lhsCost := problem.StepCost(lhsNode.Parent().Cost(),
                  lhsNode.Parent().State(),
                  lhsNode.Action(),
                  lhsNode.State()) + problem.H(lhsNode.State())

    rhsCost := problem.StepCost(rhsNode.Parent().Cost(),
                  rhsNode.Parent().State(),
                  rhsNode.Action(),
                  rhsNode.State()) + problem.H(rhsNode.State())

    return rhsCost - lhsCost
  }

  return Search(problem, NewPriorityQueue(compareNodes))
}

func Search(problem Problem, frontier Queue) (solution Node, error Error) {

  node := NewNode(nil, nil, problem.Initial())
  explored := NewHashQueue()

  frontier.Push(node)

  for {
    if frontier.Len() == 0 {
      error = "No solution found"
      return
    }

    node = frontier.Pop().(*Node)

    if problem.IsGoal(node.state) {
      solution = *node
      return
    }

    for action, nextState := range(problem.Successors(node.state)) {
      nextNode := NewNode(node, action, nextState)
      if !frontier.Contains(nextNode) && !explored.Contains(nextNode) {
        frontier.Push(nextNode)
      }
    }
    explored.Push(node)
  }

  panic("how did you get out of that infinite loop?!")
}

