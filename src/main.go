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

  mc := cmd.NewMetaCommand()
  index := func (arguments []string) (string, bool) {
    fmt.Println(mc.Index())
    return "", true
  }

  quit := func (arguments []string) (string, bool) {
    run = false;
    return "quitting", true
  }

  mc.RegisterFunction("index", index)
  mc.RegisterFunction("quit", quit)

  loop := mainloop.NewConsoleMainloop()

  loop.RunFunction = func() bool {
    data, _, error := loop.ReadLine()

    if error != nil {
      fmt.Println("error:", error)
    }

    input := strings.TrimSpace(string(data))
    result, _ := mc.Do([]string{input})
    loop.WriteString(result)
    loop.Flush()

    return run
  }

  mainloop.Loop(loop)

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

  fmt.Println(problem.Init)
  is := NewBooleanPredicate("is", "x", "y")
  has := NewBooleanPredicate("has", "x", "z")
  isAndHas := And(is, has)

  params0 := map[string]string{"x": "barrel", "y": "round", "z": "rim"}
  params1 := map[string]string{"x": "bar", "y": "steel" }
  init := NewHashKnowlegebase()
  init.SetSubjectFactory(problem.Init)

  ok := isAndHas.Satisfied(init, params0)
  fmt.Println("barrel is round and has rim", ok)

  ok = is.Satisfied(init, params1)
  fmt.Println("bar is steel", ok)
}


