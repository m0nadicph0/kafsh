package models

import "fmt"

type Group struct {
	Name      string
	State     string
	Consumers int
}

func (g *Group) ToTSRec() string {
	return fmt.Sprintf("%s\t%s\t%d", g.Name, g.State, g.Consumers)
}

type GroupDesc struct {
	GroupID      string
	State        string
	Protocol     string
	ProtocolType string
}
