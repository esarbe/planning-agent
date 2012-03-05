package loading

import "os"
import "testing"
import "fmt"
import json "encoding/json"
import . "ai/planning"

type Map map[string]interface{}

func TestJsonLoading(t *testing.T) {
  fileLocation := "/home/raphael/source/planning-agent/data/test-state1.json"

  var problem HashPlanningProblem

  file, error := os.Open(fileLocation)
  if error != nil {
    fmt.Println("could not open file '", fileLocation, "': ", error)
    t.FailNow()
  }

  d := json.NewDecoder(file)

  if error = d.Decode(&problem); error != nil {
    fmt.Println("could not decode ", error)
    t.FailNow()
  }

  fmt.Println(problem)
  is := NewBooleanPredicate("is", "x", "y")
  has := NewBooleanPredicate("has", "x", "z")
  isAndHas := And(is, has)

  params0 := map[string]string{"x": "barrel", "y": "round", "z": "rim"}
  params1 := map[string]string{"x": "bar", "y": "steel" }

  ok := isAndHas.Satisfied(problem.Init, params0)
  fmt.Println("barrel is round and barrel has rim", ok)

  ok = is.Satisfied(problem.Init, params1)
  fmt.Println("bar is steel", ok)
}


