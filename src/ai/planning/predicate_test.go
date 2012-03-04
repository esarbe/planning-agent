package planning

import "testing"
import "fmt"

func TestPredicateApplyRetract(t *testing.T) {
	is := NewBooleanPredicate("is", "x", "y")
	//has := NewPredicate("has", "x", "y")

	kb := NewHashKnowledgebase()

	is.State(kb, map[string]string{"x": "steel", "y": "metal"})

	if mapPredicate, ok := kb.objects["steel"]; !ok {
		fmt.Println("map is:", kb)
		t.FailNow()
	} else if !mapPredicate.(*HashSubject).relations["is"]["metal"].(bool) {
		t.FailNow()
	}

	nIs := Not(is)
	nIs.State(kb, map[string]string{"x": "steel", "y": "metal"})
	if kb.objects["steel"].(*HashSubject).relations["is"]["metal"].(bool) {
		fmt.Println("subj", kb.objects["steel"].(*HashSubject).relations["is"]["metal"])
		t.FailNow()
	}
}

func TestPredicateAnd(t *testing.T) {
	kb := NewHashKnowledgebase()
	is := NewBooleanPredicate("is", "x", "y")
	has := NewBooleanPredicate("has", "x", "z")
	isAndHas := And(is, has)

	isAndHas.State(kb, map[string]string{"x": "barrel", "y": "round", "z": "rim"})

	if len(kb.objects["barrel"].(*HashSubject).relations) != 2 {
		fmt.Println("unexpected relations:", kb.objects["barrel"].(*HashSubject).relations)
		t.FailNow()
	}

	isAndHas.Retract(kb, map[string]string{"x": "barrel", "y": "round", "z": "rim"})

	if len(kb.objects["barrel"].(*HashSubject).relations) != 2 {
		fmt.Println("unexpected relations:", kb.objects["barrel"].(*HashSubject).relations)
		t.FailNow()
	}
}

func TestPredicateSatisfied(t *testing.T) {
	kb := NewHashKnowledgebase()
	is := NewBooleanPredicate("is", "x", "y")
	has := NewBooleanPredicate("has", "x", "z")
	isAndHas := And(is, has)

	params := map[string]string{"x": "barrel", "y": "round", "z": "rim"}

	isAndHas.State(kb, params)

	if !isAndHas.Satisfied(kb, params) {
		fmt.Println("expected proposition to be true")
		t.FailNow()
	}
}
