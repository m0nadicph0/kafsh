package models

import (
	"fmt"
)

type Topic struct {
	Name       string
	Partitions int32
	Replicas   int16
}

func (t Topic) ToTSRec() string {
	return fmt.Sprintf("%s\t%d\t%d\t", t.Name, t.Partitions, t.Replicas)
}

type TopicDesc struct {
	Name                string
	Internal            bool
	Compacted           bool
	SummedHighWatermark int
	Partitions          []*PartitionDetail
	Config              []*PartitionConfig
}

type PartitionDetail struct {
	Partition int32
	Leader    int32
	Replicas  []int32
	ISR       []int32
}

func (d *PartitionDetail) ToTSRec() string {
	return fmt.Sprintf("%d\t%d\t%v\t%v\t", d.Partition, d.Leader, d.Replicas, d.ISR)
}

type PartitionConfig struct {
	Name      string
	Value     string
	ReadOnly  bool
	Sensitive bool
}

func (c *PartitionConfig) ToTSRec() string {
	return fmt.Sprintf("%s\t%s\t%t\t%t\t", c.Name, c.Value, c.ReadOnly, c.Sensitive)
}
