package planning

//import "fmt"

// knowledgebase for tripplets of subject, predicate and object. each tripplet hat a given
// weight. 
type Knowledgebase interface {
	Set(subject string, predicate string, object string, value interface{}) bool
	Query(subject string, predicate string, object string) (value interface{}, exists bool)
	Subjects() []string
}

type Subject interface {
	Set(predicate string, object string, value interface{}) bool
	Query(predicate string, object string) (value interface{}, exists bool)
}

type SubjectFactory interface {
	Fetch(identifier string) (Subject, bool)
	Build(identifier string, predicates map[string]map[string]interface{}) Subject
}

type HashSubjectFactory struct {
  Objects map[string]map[string]map[string]interface{}
}

func (hsf *HashSubjectFactory) Fetch (identifier string) (subject Subject, ok bool) {
  if relations, ok := hsf.Objects[identifier]; ok {
    return hsf.Build(identifier, relations), ok
  }
  return
}

func (hsf *HashSubjectFactory) Build (identifier string, relations map[string]map[string]interface{}) Subject {
  subject := &HashSubject{identifier, relations}
  hsf.Objects[identifier] = relations
  return subject
}

func NewHashSubjectFactory() *HashSubjectFactory {
	return new(HashSubjectFactory)
}

type HashKnowledgebase struct {
	factory SubjectFactory
	Objects map[string]*HashSubject
}

func NewHashKnowlegebase() *HashKnowledgebase {
	factory := new(HashSubjectFactory)
	subjects := make(map[string]*HashSubject)

	return &HashKnowledgebase{factory, subjects}
}

func (hkb *HashKnowledgebase) SetObjects(objects map[string]*HashSubject) {
	hkb.Objects = objects
}

func (hkb *HashKnowledgebase) Subjects() []string {
	return []string{}
}

func (hkb *HashKnowledgebase) Set(subject string, predicate string, object string, value interface{}) bool {
	s, _ := hkb.fetch(subject)
	return s.Set(predicate, object, value)
}

func (hkb *HashKnowledgebase) Intersect(kb Knowledgebase) {
}

func (hkb *HashKnowledgebase) fetch(subject string) (Subject, bool) {
	var s Subject
	var ok bool
        // check if subject is available locally
	if s, ok = hkb.Objects[subject]; !ok {
		s, ok = hkb.factory.Fetch(subject)
                if ok {
		  hkb.Objects[subject] = s.(*HashSubject)
                }
	}

	return s, ok
}

func (hkb *HashKnowledgebase) Query(subject string, predicate string, object string) (answer interface{}, ok bool) {
        var s Subject
        if s, ok = hkb.fetch(subject); ok {
          answer, ok = s.Query(predicate, object)
        }

	return
}

func (hkb *HashKnowledgebase) SetSubjectFactory (f SubjectFactory) {
  hkb.factory = f
}

type HashSubject struct {
	identifier string
	relations  map[string]map[string]interface{}
}

func (hs *HashSubject) Set(predicate string, object string, value interface{}) bool {
	var predicateRelations map[string]interface{}
	var ok bool

	// check if there is already a map for the predicate
	if predicateRelations, ok = hs.relations[predicate]; !ok {
		// create one otherwise
		predicateRelations = make(map[string]interface{})
		hs.relations[predicate] = predicateRelations
	}
	predicateRelations[object] = value

	return true
}

func (hs *HashSubject) Query(predicate string, object string) (value interface{}, exists bool) {
	value = nil
	exists = false

//        fmt.Println("   looking up", predicate, object, "in",  hs.relations)
	predicateRelations, ok := hs.relations[predicate]

	if ok {
		value, exists = predicateRelations[object]
	}

	return
}
