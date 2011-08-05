package set

type ListElement struct {
  next *ListElement
  Value int
}

type LinkedListSet struct {
  head, tail *ListElement 
  length int
}

func (this *LinkedListSet) Add(value int) {
  e := &ListElement{nil, value}
  if (this.head == nil){
    this.head = e
    this.tail = e
    this.length = 1
  } else if (this.head.Value >  value) {
    e.next = this.head
    this.head = e
    this.length += 1
  } else {
    var prev *ListElement
    for e := this.head; e != nil; e = e.next {
      if (e.Value == value) {
        return
      }
      if(e.Value > value) {
        break
      }
      prev = e
    }
    this.length += 1
    e.next = prev.next
    prev.next = e
  }
}

func (this *LinkedListSet) Contains(value int) (exists bool) {
  for e := this.head; e != nil; e = e.next {
    if e.Value == value {
      return true
    }
    if e.Value > value {
      return false
    }
  }
  return false
}

func (this *LinkedListSet) Length() (int) {
  return this.length
}
func (this *LinkedListSet) RemoveDuplicates() {}

func NewSet() (Set) {
  return new (LinkedListSet)
}