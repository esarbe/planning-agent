package planning

import "testing"
import "fmt"

func TestKnowledgebase(t *testing.T) {
	kb := NewHashKnowledgebase()

	kb.Set("steel", "is", "metal", true)

	weight, ok := kb.Query("steel", "is", "metal")
	if !ok || weight.(bool) != true {
		fmt.Println("kb:", kb)
		t.FailNow()
	}

	kb.Set("steel", "is", "metal", false)

	isSteel, _ := kb.Query("steel", "is", "metal")

	if isSteel.(bool) {
		t.FailNow()
	}
}
