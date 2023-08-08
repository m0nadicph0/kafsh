package models

import "fmt"

type Node struct {
	ID           int32
	Address      string
	IsController bool
}

func (n *Node) ToTSRec() string {
	return fmt.Sprintf("%d\t%s\t%t", n.ID, n.Address, n.IsController)
}
