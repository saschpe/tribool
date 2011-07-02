package tribool

type Tribool interface {
    True() bool
    Indeterminate() bool
    False() bool

    Not() Tribool
    And(b Tribool) Tribool
    Or(b Tribool) Tribool
}

type tribool struct {
    value int
}

const tribool_true = 2
const tribool_indeterminate = 1
const tribool_false = 0

func New(value string) Tribool {
    var ret tribool
    switch value {
        case "true": ret.value = tribool_true
        case "indeterminate": ret.value = tribool_indeterminate
        case "false": ret.value = tribool_false
    }
    return &ret
}

func (a *tribool) True() bool { return a.value == tribool_true }
func (a *tribool) Indeterminate() bool { return a.value == tribool_indeterminate }
func (a *tribool) False() bool { return a.value == tribool_false }

func (a *tribool) Not() Tribool {
    ret := a
    switch ret.value {
        case tribool_true: ret.value = tribool_false
        case tribool_false: ret.value = tribool_true
    }
    return ret
}

func (a *tribool) And(b Tribool) Tribool {
    var ret tribool
    if a.True() {
        if b.True() {
            ret.value = tribool_true
        } else if b.Indeterminate() {
            ret.value = tribool_indeterminate
        } else {
            ret.value = tribool_false
        }
    } else if a.Indeterminate() {
        if b.False() {
            ret.value = tribool_false
        } else {
            ret.value = tribool_indeterminate
        }
    } else { // a.False()
        ret.value = tribool_false
    }
    return &ret
}

func (a *tribool) Or(b Tribool) Tribool {
    var ret tribool
    if a.True() {
        ret.value = tribool_true
    } else if a.Indeterminate() {
        if b.True() {
            ret.value = tribool_true
        } else {
            ret.value = tribool_indeterminate
        }
    } else { // a.False()
        if b.True() {
            ret.value = tribool_true
        } else if b.Indeterminate() {
            ret.value = tribool_indeterminate
        } else {
            ret.value = tribool_false
        }
    }
    return &ret
}

func (a tribool) String() string {
    var ret string
    switch a.value {
        case tribool_true: ret = "true"
        case tribool_indeterminate: ret = "indeterminate"
        case tribool_false: ret = "false"
    }
    return ret
}
