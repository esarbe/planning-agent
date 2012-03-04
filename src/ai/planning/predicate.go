package planning

//import "fmt"

type Proposition interface {
	State(kb Knowledgebase, parameters map[string]string) bool
	Retract(kb Knowledgebase, parameters map[string]string) bool
	Satisfied(kb Knowledgebase, parameters map[string]string) bool
}

type PropositionFactory interface {
	Build(identifier string, parameters []interface{})
}

type predicateParameters map[string]string

type Predicate struct {
	name              string
	subjectIdentifier string
	objectIdentifier  string
	state             statement
	retract           statement
	query             statement
}

// function definition for statements on a predicate
type statement func(kb Knowledgebase, subject string, object string) bool

func NewBooleanPredicate(predicate string, subjectIdentifier string, objectIdentifier string) *Predicate {
	name := predicate

	retract := func(kb Knowledgebase, subject string, object string) bool {
		return kb.Set(subject, name, object, false)
	}

	state := func(kb Knowledgebase, subject string, object string) bool {
		return kb.Set(subject, name, object, true)
	}

	query := func(kb Knowledgebase, subject string, object string) bool {
    //fmt.Println("querying:", subject, name, object)
		value, ok := kb.Query(subject, name, object)
		ok = value != nil && value.(bool) && ok
//    fmt.Println("   > ", ok)
		return ok
	}

	return &Predicate{name, subjectIdentifier, objectIdentifier, state, retract, query}
}

func extractParameters(subject string, object string, parameters map[string]string) (string, string, bool) {
	subject, sOk := parameters[subject]
	object, oOk := parameters[object]

	if !sOk || !oOk {
		return "", "", false
	}
	return subject, object, true
}

func (p *Predicate) State(kb Knowledgebase, parameters map[string]string) bool {
	return p.doWithStatement(kb, p.state, parameters)
}

func (p *Predicate) Retract(kb Knowledgebase, parameters map[string]string) bool {
	return p.doWithStatement(kb, p.retract, parameters)
}

func (p *Predicate) Satisfied(kb Knowledgebase, parameters map[string]string) bool {
	return p.doWithStatement(kb, p.query, parameters)
}

func (p *Predicate) doWithStatement(kb Knowledgebase, s statement, parameters predicateParameters) bool {
	subject, object, ok := extractParameters(p.subjectIdentifier, p.objectIdentifier, parameters)
	if ok {
    //fmt.Println("querying for: ", subject, object)
		ok = s(kb, subject, object)
	}

	return ok
}
