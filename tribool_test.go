package tribool

import (
    "testing"
)

func TestPrinting(t *testing.T) {
    if Tribool(True).String() != "true" {
        t.Fail()
    }
    if Tribool(Indeterminate).String() != "indeterminate" {
        t.Fail()
    }
    if Tribool(False).String() != "false" {
       t.Fail()
    }
}

func TestIdentity(t *testing.T) {
    a, b, c := Tribool(True), Tribool(Indeterminate), Tribool(False)

    if !a.True() || a.Indeterminate() || a.False() {
        t.Fail()
    }
    if b.True() || !b.Indeterminate() || b.False() {
        t.Fail()
    }
    if c.True() || c.Indeterminate() || !c.False() {
        t.Fail()
    }
}

func TestAnd(t *testing.T) {

}

func TestOr(t *testing.T) {

}
