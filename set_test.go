package set
import(
  "testing"
  "rand"
)

func TestCanAddUniqueValues(t *testing.T) {
  values := []int{5,4,3,9,1,2,8} 
  set := NewSet()
  for _,v := range values {
    set.Add(v)
  }
  
  Assert(t, set.Length() == len(values), "Wrong length")
  for _,v := range values {
    Assert(t, set.Contains(v), "Missing value")
  }
}


func TestUniqueRemovesDuplicates(t *testing.T) {
  values := []int{5,4,3,5,2,1,2,3,4,5,1,2,2} 
  uniques := []int{5,4,3,2,1}
  set := NewSet()
  for _,v := range values {
    set.Add(v)
  }
  set.RemoveDuplicates()
  Assert(t, set.Length() == len(uniques), "Wrong length")
  for _,v := range uniques {
    Assert(t, set.Contains(v), "Missing value")
  }
}
func Assert(t *testing.T, condition bool, message string) {
  if (!condition) {
    t.Fatal(message)
  }
}

func benchmark(b *testing.B, size int, fill int) {
  for i := 0; i < b.N; i++ {
    set := NewSet()
    for j := 0; j < size; j++ {
      set.Add(rand.Int() % fill)
    }
  }
}

func BenchmarkSmallSetWithFewCollisions(b *testing.B){
  benchmark(b, 100, 700)
}
func BenchmarkSmallSetWithMoreCollisions(b *testing.B){
 benchmark(b, 100, 100)
}
func BenchmarkSmallSetWithManyCollisions(b *testing.B){
 benchmark(b, 100, 25)
}
func BenchmarkMediumSetWithFewCollisions(b *testing.B){
  benchmark(b, 5000, 35000)
}
func BenchmarkMediumSetWithMoreCollisions(b *testing.B){
 benchmark(b, 5000, 5000)
}
func BenchmarkMediumSetWithManyCollisions(b *testing.B){
 benchmark(b, 5000, 1250)
}
func BenchmarkLargeSetWithFewCollisions(b *testing.B){
  benchmark(b, 100000, 700000)
}
func BenchmarkLargeSetWithMoreCollisions(b *testing.B){
 benchmark(b, 100000, 100000)
}
func BenchmarkLargeSetWithManyCollisions(b *testing.B){
 benchmark(b, 100000, 25000)
}