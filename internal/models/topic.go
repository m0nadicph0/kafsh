package models

import "fmt"

type Topic struct {
	Name       string
	Partitions int32
	Replicas   int16
}

func (t Topic) ToTSRec() string {
	return fmt.Sprintf("%s\t%d\t%d\t", t.Name, t.Partitions, t.Replicas)
}
