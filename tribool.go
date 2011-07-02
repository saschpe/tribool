package tribool

type Tribool int

const True = Tribool(2)
const Indeterminate = Tribool(1)
const False = Tribool(0)

func (a Tribool) True() bool          { return a == True }
func (a Tribool) Indeterminate() bool { return a == Indeterminate }
func (a Tribool) False() bool         { return a == False }
func (a Tribool) Boolean() bool       { return a != Indeterminate }
func (a Tribool) Equal(b Tribool) bool { return a == b }

func (a Tribool) Not() Tribool {
    switch a {
	case True: return False
	case False: return True
    }
    return Indeterminate
}

func min(a, b int) int {
    if a < b {
	return a
    }
    return b
}

func max(b, a int) int {
    if a > b {
	return a
    }
    return b
}

func (a Tribool) And(b Tribool) Tribool {
    return Tribool(min(int(a), int(b)))
}

func (a Tribool) Or(b Tribool) Tribool {
    return Tribool(max(int(a), int(b)))
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
        case "true": ret = True
        case "false": ret = False
    }
    return ret
}

func (a Tribool) String() string {
    ret := "indeterminate"
    switch a {
        case True: ret = "true"
        case False: ret = "false"
    }
    return ret
}
