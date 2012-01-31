
type NamedNode struct {
  cost float32
}

func (n NamedNode) Compare (other Comparable) int {
  o := other.(*NamedNode)

  if n.cost < o.cost {
    return -1
  } else if n.cost == o.cost {
    return 0
  } else if n.cost > o.cost {
    return 1
  }
  panic("there is something seriously wrong..")
}


