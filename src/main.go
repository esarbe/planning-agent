package main

import "strings"
import "os"
import "fmt"
import json "encoding/json"
import . "ai/planning"
import cmd "command"
import "mainloop"

type Map map[string]interface{}

func main() {
  run := true

  fileLocation := "/home/raphael/source/planning-agent/data/test-state1.json"
  var problem HashPlanningProblem

  file, error := os.Open(fileLocation)
  if error != nil {
    fmt.Println("could not open file '", fileLocation, "': ", error)
    panic("")
  }

  d := json.NewDecoder(file)

  if error = d.Decode(&problem); error != nil {
    fmt.Println("could not decode ", error)
    panic("")
  }

  init := NewHashKnowlegebase()
  init.SetSubjectFactory(problem.Init)

  mc := cmd.NewMetaCommand()

  index := func (arguments []string) (string, bool) {
    fmt.Println(mc.Index())
    return "", true
  }

  set := func (arguments []string) (string, bool) {
    if len(arguments) < 3 {
      return "not enough arguments, need: <subject> <predicate> <object>", false
    }
    predicate := NewBooleanPredicate(arguments[1], "x", "y")
    predicate.State(init, map[string]string{"x": arguments[0], "y": arguments[2]})
    return "", true
  }

  unset := func (arguments []string) (string, bool) {
    if len(arguments) < 3 {
      return "not enough arguments, need: <subject> <predicate> <object>", false
    }
    predicate := NewBooleanPredicate(arguments[1], "x", "y")
    predicate.Retract(init, map[string]string{"x": arguments[0], "y": arguments[2]})
    return "", true
  }

  query := func (arguments []string) (string, bool) {
    if len(arguments) < 3 {
      return "not enough arguments, need: <subject> <predicate> <object>", false
    }
    predicate := NewBooleanPredicate(arguments[1], "x", "y")
    ok := predicate.Satisfied(init, map[string]string{"x": arguments[0], "y": arguments[2]})
    return fmt.Sprintln(strings.Join(arguments, " "), ok), ok
  }

  show := func (arguments []string) (string, bool) {
    if len(arguments) < 1 {
      return "not enough arguments, need: <object>", false
    }

    subject, ok := init.Objects[arguments[0]]
    return fmt.Sprint(subject), ok
  }

  quit := func (arguments []string) (string, bool) {
    run = false;
    return "quitting", true
  }

  mc.RegisterFunction("index", index)
  mc.RegisterFunction("quit", quit)
  mc.RegisterFunction("set", set)
  mc.RegisterFunction("unset", unset)
  mc.RegisterFunction("query", query)
  mc.RegisterFunction("show", show)

  loop := mainloop.NewConsoleMainloop(">")

  loop.RunFunction = func() bool {
    loop.Flush()
    data, _, error := loop.ReadLine()

    if error != nil {
      fmt.Println("error:", error)
    }

    input := strings.TrimSpace(string(data))
    result, _ := mc.Do(strings.Split(input, " "))
    loop.WriteString(result)

    return run
  }

  mainloop.Loop(loop)

}


