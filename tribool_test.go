package tribool

import (
    "fmt"
    "testing"
)

func TestTribool(tp *testing.T) {
    t1 := New("true")
    t2 := New("false")

    fmt.Println(t1)
    fmt.Println(t1.True())
    fmt.Println(t1.False())
    fmt.Println(t1.Indeterminate())
    fmt.Println(t1.Not())

    fmt.Println(t1.And(t2))
    fmt.Println(t1.Or(t2))
}

func TestPrinting(tp *testing.T) {

}
