package planning

type Item interface{}

type PlanningState interface{}

type PlanningAction interface{
  GetPrecondition()
}



type Evaluator interface {
  Evaluate(PlanningState) bool
}

type ConstraintBuilder struct {
  Evaluator
}

type ClassicalPlanningProblem struct {

}

type NotConstraint struct {
  Evaluator
}

func (e *NotConstraint) Evaluate(s PlanningState) bool {
  return !e.Evaluator.Evaluate(s)
}

func Not(e Evaluator) Evaluator {
  return &NotConstraint{e}
}

