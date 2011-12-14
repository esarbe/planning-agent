package planning

import vector "container/vector"

func (error *NoSolutionError) String() string {
  return error.message
}

func (problem *Problem) BreadthFirstSearch() (solution Node, error Error) {

  node := &Node{problem, nil, nil, float32(0)}

  if problem.IsGoal(node.State()) {
    return node
  }

  frontier := []State{0,0}
  explored := []State{0,0}

  frontier = append(frontier, node)

  for {
    if len(frontier) == 0 {
      return nil, &NoSolutionError{"Frontier empty"} 
    }

    node, frontier = frontier[0], frontier[1, len(frontier)]
    for action := problem.Actions(node.state) {
       child = 
    }
  }
}

