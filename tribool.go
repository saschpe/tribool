package tribool

type Tribool int

const True = 2
const Indeterminate = 1
const False = 0

func New(value string) Tribool {
    ret := Tribool(False)
    switch value {
    case "true": ret = True
    case "indeterminate": ret = Indeterminate
    }
    return ret
}

func (a Tribool) True() bool          { return a == True }
func (a Tribool) Indeterminate() bool { return a == Indeterminate }
func (a Tribool) False() bool         { return a == False }
func (a Tribool) Boolean() bool       { return a != Indeterminate }

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

func (a Tribool) String() string {
    ret := "indeterminate"
    switch a {
        case True: ret = "true"
        case False: ret = "false"
    }
    return ret
}

//func (a Tribool) toBoolean() bool, os.Error {
//}
