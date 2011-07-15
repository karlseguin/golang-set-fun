package set

type OnDemandArraySet struct {
  data []int
}

func (this *OnDemandArraySet) Add(value int) {
  this.data = append(this.data, value)
}

func (this *OnDemandArraySet) Contains(value int) (exists bool) {
  for _, v := range this.data {
    if v == value {
      return true
    }
  }
  return false
}

func (this *OnDemandArraySet) Length() (int) {
  return len(this.data)
}
func (this *OnDemandArraySet) RemoveDuplicates() {
  length := len(this.data) - 1
  for i := 0; i < length; i++ {
    for j := i + 1; j <= length; j++ {
      if (this.data[i] == this.data[j]) {
        this.data[j] = this.data[length]
        this.data = this.data[0:length]
        length--
        j--
      }
    }
  }
}


func NewSet() (Set) {
  return &OnDemandArraySet{make([]int, 0, 100)}
}