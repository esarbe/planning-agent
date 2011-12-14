package planning 

type State interface {
  ListActions() []Action
  Name() string
}

type Node interface {
  State() State
  Parent() Node 
  Action() Action
  PathCost() float32
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

type NamedNode struct {
  parent Node
  action Action
  pathCost float32
  state State
}

func (n NamedNode) Parent() Node {
  return n.parent
}

func (n NamedNode) State() State {
  return n.state
}

func (n NamedNode) Action() Action {
  return n.action
}

func (n NamedNode) PathCost() float32 {
  return n.pathCost
}

func (n *NamedNode) SetPathCost(pathCost float32) {
  n.pathCost = pathCost
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

