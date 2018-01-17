// Copyright 2011 Sascha Peilicke. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tribool

// Tribool represents a ternary (three-valued) logic type.
type Tribool int8

// Constants for all possible values in a ternary logic.
const (
	True          = Tribool(2)
	Indeterminate = Tribool(1)
	False         = Tribool(0)
)

func (a Tribool) True() bool          { return a == True }
func (a Tribool) Indeterminate() bool { return a == Indeterminate }
func (a Tribool) False() bool         { return a == False }
func (a Tribool) Boolean() bool       { return a != Indeterminate }

func (a Tribool) Not() Tribool {
	switch a {
	case True:
		return False
	case False:
		return True
	}
	return Indeterminate
}

func (a Tribool) Equal(b Tribool) bool { return a == b }

func (a Tribool) And(b Tribool) Tribool {
	if int(a) < int(b) {
		return a
	}
	return b
}

func (a Tribool) Or(b Tribool) Tribool {
	if int(a) > int(b) {
		return a
	}
	return b
}

func FromBool(value bool) Tribool {
	ret := Tribool(False)
	if value == true {
		ret = True
	}
	return ret
}

func FromString(value string) Tribool {
	ret := Tribool(Indeterminate)
	switch value {
	case "true":
		ret = True
	case "false":
		ret = False
	}
	return ret
}

func (a Tribool) String() string {
	ret := "indeterminate"
	switch a {
	case True:
		ret = "true"
	case False:
		ret = "false"
	}
	return ret
}
