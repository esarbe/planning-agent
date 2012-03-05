package command

import "testing"
import "fmt"

func TestDoCommand (t *testing.T) {
  n1 := &NopCommand{}
  index := &FunctionCommand{}

  mc := NewMetaCommand()

  indexFunction := func (arguments []string) (string, bool) {
    fmt.Println(mc.Index())
    return "", true
  }

  index.r = indexFunction

  mc.Register("set", n1)
  mc.Register("do", n1)
  mc.Register("index", index)

  fmt.Println(mc.Do([]string{"index", "d"}))

}
