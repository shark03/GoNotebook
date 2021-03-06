package main
import . "fmt"
import "reflect"

func main() {
  v := reflect.ValueOf([]int{0, 2, 4, 6, 8})
  for i := 0; i < v.Len(); i++ {
    x := v.Index(i)
    Printf("%v: %v\n", i, x)
    x.SetInt(int64(i))
  }
  for i := 0; i < v.Len(); i++ {
    Printf("%v: %v\n", i, v.Index(i))
  }
}
