// Copyright 2011 Sascha Peilicke. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tribool

import (
    "testing"
)

func TestIdentity(t *testing.T) {
    a, b, c := True, Indeterminate, False

    if !a.True() || a.Indeterminate() || a.False() { t.Fail() }
    if b.True() || !b.Indeterminate() || b.False() { t.Fail() }
    if c.True() || c.Indeterminate() || !c.False() { t.Fail() }
}

func TestEquality(t *testing.T) {
    a, b, c := True, True, False

    if a != b { t.Fail() }
    if !a.Equal(b) { t.Fail() }
    if a == c { t.Fail() }
    if a.Equal(c) { t.Fail() }
}

func TestNot(t *testing.T) {
    a, b, c := True, Indeterminate, False

    if a.Not() != False { t.Fail() }
    if b.Not() != Indeterminate { t.Fail() }
    if c.Not() != True { t.Fail() }
}

func TestAnd(t *testing.T) {
    a, b, c := True, Indeterminate, False

    if a.And(a) != True { t.Fail() }
    if a.And(b) != Indeterminate { t.Fail() }
    if a.And(c) != False { t.Fail() }
    if b.And(a) != Indeterminate { t.Fail() }
    if b.And(b) != Indeterminate { t.Fail() }
    if b.And(c) != False { t.Fail() }
    if c.And(a) != False { t.Fail() }
    if c.And(b) != False { t.Fail() }
    if c.And(c) != False { t.Fail() }
}

func TestOr(t *testing.T) {
    a, b, c := True, Indeterminate, False

    if a.Or(a) != True { t.Fail() }
    if a.Or(b) != True { t.Fail() }
    if a.Or(c) != True { t.Fail() }
    if b.Or(a) != True { t.Fail() }
    if b.Or(b) != Indeterminate { t.Fail() }
    if b.Or(c) != Indeterminate { t.Fail() }
    if c.Or(a) != True { t.Fail() }
    if c.Or(b) != Indeterminate { t.Fail() }
    if c.Or(c) != False { t.Fail() }
}

func TestConvertFromBool(t *testing.T) {
    if FromBool(true) != True { t.Fail() }
    if FromBool(false) != False { t.Fail() }
}

func TestConvertFromString(t *testing.T) {
    if FromString("true") != True { t.Fail() }
    if FromString("indeterminate") != Indeterminate { t.Fail() }
    if FromString("false") != False { t.Fail() }
}

func TestConvertToString(t *testing.T) {
    if True.String() != "true" { t.Fail() }
    if Indeterminate.String() != "indeterminate" { t.Fail() }
    if False.String() != "false" { t.Fail() }
}
