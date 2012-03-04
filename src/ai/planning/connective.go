package planning

type Connective struct {
	p0 Proposition
	p1 Proposition
}

type Conjunction []Proposition

type Negation struct {
	Proposition
}

type Disjunction Connective

func Not(p Proposition) Proposition {
	return &Negation{p}
}

//func And(p ...Proposition) Proposition {
//	c := &Connective{}
//	c.propositions = p
//	return c
//}

func And(p0 Proposition, p1 Proposition) Proposition {
	c := &Connective{p0, p1}
	return c
}

func (c *Connective) State(kb Knowledgebase, parameters map[string]string) bool {
	ok0 := c.p0.State(kb, parameters)
	ok1 := c.p1.State(kb, parameters)
	return ok0 && ok1
}

func (c *Connective) Retract(kb Knowledgebase, parameters map[string]string) bool {
	ok0 := c.p0.Retract(kb, parameters)
	ok1 := c.p1.Retract(kb, parameters)
	return ok0 && ok1
}

func (c *Connective) Satisfied(kb Knowledgebase, parameters map[string]string) bool {
	ok0 := c.p0.Satisfied(kb, parameters)
	ok1 := c.p1.Satisfied(kb, parameters)

	return ok0 && ok1
}

func (n *Negation) State(kb Knowledgebase, parameters map[string]string) bool {
	return n.Proposition.Retract(kb, parameters)
}

func (n *Negation) Retract(kb Knowledgebase, parameters map[string]string) bool {
	return n.Proposition.State(kb, parameters)
}

func (n *Negation) Satisfied(kb Knowledgebase, parameters map[string]string) bool {
	return !n.Proposition.Satisfied(kb, parameters)
}
