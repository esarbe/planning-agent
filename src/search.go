package planning

import "os"

func (error *NoSolutionError) String() string {
  return error.message
}

func equals(lhs Comparable, rhs Comparable) bool {
  ls := lhs.(*Node).state
  rs := rhs.(*Node).state
  return ls.Equals(rs)
}

func BreadthFirstSearch(problem Problem) (solution Node, error os.Error) {
  return Search(problem, NewPriorityQueue())
}

func Search(problem Problem, frontier Queue) (solution Node, error os.Error) {

  node := NewNode(nil, nil, problem.Initial())
  explored := NewHashQueue()

  frontier.Push(node)

  //fmt.Println("frontier length: ", frontier.Len())

  for {
    if frontier.Len() == 0 {
      error = os.NewError("No solution found")
      return
    }

    node = frontier.Pop().(*Node)

    if problem.IsGoal(node.state) {
      //fmt.Println("found solution: ", node.state)
      solution = *node
      return
    }
    //fmt.Println("iterating over states..")
    for action, nextState := range(problem.Successors(node.state)) {
      //fmt.Println("state:", nextState)
      nextNode := NewNode(node, action, nextState)
      if !frontier.Contains(nextNode) && !explored.Contains(nextNode) {
        frontier.Push(nextNode)
      }
    }
    explored.Push(node)
  }

  panic("how did you get out of that infinite loop?!")
}

