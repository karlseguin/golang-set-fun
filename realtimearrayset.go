package set

type RealtimeArraySet struct {
  data []int
}

func (this *RealtimeArraySet) Add(value int) {
  if this.Contains(value) == false {
    this.data = append(this.data, value)
  }
}

func (this *RealtimeArraySet) Contains(value int) (exists bool) {
  for _, v := range this.data {
    if v == value {
      return true
    }
  }
  return false
}

func (this *RealtimeArraySet) Length() (int) {
  return len(this.data)
}
func (this *RealtimeArraySet) RemoveDuplicates() {}


func NewSet() (Set) {
  return &RealtimeArraySet{make([]int, 0, 100)}
}