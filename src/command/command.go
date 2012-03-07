package command

import "fmt"
import "strings"

type Command interface {
  Index() []string
  Do(arguments []string) (string, bool)
}

// A command containing other commands
type MetaCommand struct {
  commands map[string]Command
}

func NewMetaCommand () *MetaCommand {
  commands := make(map[string]Command)
  return &MetaCommand{commands}
}

func (mc *MetaCommand) Register (key string, c Command) bool {
  mc.commands[key] = c
  return true
}

func (mc *MetaCommand) Index () ([]string) {
  commands := []string{}
  for i, _ := range mc.commands {
    commands = append(commands, i)
  }

  return commands
}

func (mc *MetaCommand) Do (arguments []string) (string, bool) {
  c := arguments[0]

  if strings.TrimSpace(c) == "" {
    return "", true
  }

  command, ok := mc.commands[c]

  if !ok {
    return fmt.Sprintln(c, ": command not found"), false
  }

  return command.Do(arguments[1:])
}

func (mc *MetaCommand) RegisterFunction (key string, r RunFunction) bool {
  c := &FunctionCommand{}
  c.r = r
  return mc.Register(key, c)
}

type RunFunction func (arguments []string) (string, bool)

type NopCommand struct { bool }

type FunctionCommand struct {
  NopCommand
  r RunFunction
}

func (fc *FunctionCommand) Do (arguments []string) (string, bool) {
  return fc.r(arguments)
}

func (nc *NopCommand) Do (arguments []string) (string, bool) {
  return "", true
}

func (nc *NopCommand) Index() []string {
  return nil;
}




