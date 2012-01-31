package planning

type Queue interface {
  Pop() Comparable
  Push(c Comparable)
  Contains(c Comparable) bool
  Len() int
}

type Container interface {
  Contains(h Hashable) bool
}

type Hashable interface {
  Hash() string
  Equals(other Hashable) bool
}

type Comparable interface {
  Hashable
  Compare(other Comparable) int
}

type Slice []Comparable
type HashMap map[string]Comparable

type HashQueue struct {
  HashMap
  Slice
}

func NewHashQueue() *HashQueue {
  h := new(HashQueue)
  h.HashMap = make(map[string]Comparable)
  h.Slice = make([]Comparable, 0) 
  return h
}

func (m HashMap) Contains(c Comparable) bool {
  _, ok := m[c.Hash()]
  return ok
}

func (q *HashQueue) Pop() (ret Comparable) {
  ret = q.Slice.Pop().(Comparable)
  q.HashMap[ret.Hash()] = nil, false
  return
}

func (q *HashQueue) Push(h Comparable) {
  q.HashMap[h.Hash()] = h
  q.Slice.Push(h)
}

type Comparison func(lhs Comparable, rhs Comparable) bool
func (s Slice) Len() int {
  return len(s)
}

func (s Slice) Less(i, j int) bool {
  qi := s[i]
  qj := s[j]
  return (qi).Compare(qj) < 0
}

func (s Slice) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

func (s *Slice) Push(c Comparable) {
  *s = append(*s, c)
}

func (s *Slice) Pop() (c Comparable) {
  c, *s = (*s)[0], (*s)[1:]
  return
}

