package planning

type Action interface {
}

type PlanningAction struct {
	precondition Predicate
	effect       Predicate
}
