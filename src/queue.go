package planning

import "sort"

type PriorityQueue []*NamedNode

func (q *PriorityQueue) Push(i *NamedNode) {

  //fmt.Println("pushing node, length: ", len(*q), "cap: ",cap(*q))
  *q = append(*q, i)
  //fmt.Println("pushed node, new length: ", len(*q), "cap: ",cap(*q))
}

func (q *PriorityQueue) Sort() {
  sort.Sort(*q)
}

func (q *PriorityQueue) Pop() (i *NamedNode) {
  sort.Sort(*q)
  //fmt.Println("queue length: %s, queue cap: %s", len(*q), cap(*q));
  i, *q = (*q)[0], (*q)[1:]
  return
}

func (q PriorityQueue) Len() int {
  return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
  return (*q[i]).PathCost() < (*q[j]).PathCost() 
}

func (q PriorityQueue) Swap(i, j int) {
  q[i], q[j] = q[j], q[i]
}
