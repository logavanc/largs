package largs

import (
	_ "fmt"
	_ "os"
)

//                                                                       PUBLIC
///////////////////////////////////////////////////////////////////////////////

type Group struct {
	name   string
	groups []*Group
	//flags  []*flag

}

func (g *Group) SetName(n string) (err error) {
	g.name = n
	return
}

func (g *Group) AddGroup() (gr *Group) {
	gr = newGroup()
	g.groups = append(g.groups, gr)

	return
}

func (g *Group) AddFlag() {
}

//                                                                      PRIVATE
///////////////////////////////////////////////////////////////////////////////

func newGroup() (g *Group) {
	g = &Group{}
	return
}
