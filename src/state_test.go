package state

import "testing"
import "fmt"

func TestStateMachine(t *testing.T) {

  endState := &NamedState{"EndState", nil}

  endAction := &NamedAction { "EndAction", 
                              func (s State) State { return endState },
                              func (s State) bool {
                                return s.Name() == "StartState" 
                              },
                            }

  startState := &NamedState{"StartState", []Action {endAction}}
  startState.ListActions()
  
  fmt.Println("test")
  t.Error("failed");
}
