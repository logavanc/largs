package largs

import (
	_ "fmt"
	_ "os"
)

//                                                                      PRIVATE
///////////////////////////////////////////////////////////////////////////////

type group struct {
	name   string
	groups []*group
	//flags  []*flag

}

func newGroup() (g *group) {
	g = new(group)
	return
}

//                                                                       PUBLIC
///////////////////////////////////////////////////////////////////////////////

func (g *group) SetName(n string) (err error) {
	g.name = n
	return
}

func (g *group) AddGroup() (gr *group) {
	gr = newGroup()
	g.groups = append(g.groups, gr)

	return
}

func (g *group) AddFlag() {

}
