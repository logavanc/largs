PACKAGE DOCUMENTATION

package largs
    import "largs"


TYPES

type Arg struct {
    Long  string
    Short string
}

type Largs struct {
    Args []Arg
    // contains filtered or unexported fields
}


