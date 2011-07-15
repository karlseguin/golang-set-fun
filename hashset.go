package set

type HashSet struct {
  data map[int]bool
}

func (this *HashSet) Add(value int) {
  this.data[value] = true
}

func (this *HashSet) Contains(value int) (exists bool) {
  _, exists = this.data[value]
  return
}

func (this *HashSet) Length() (int) {
  return len(this.data)
}
func (this *HashSet) RemoveDuplicates() {}


func NewSet() (Set) {
  return &HashSet{make(map[int] bool)}
}