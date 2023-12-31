package printer

import (
	"fmt"
	"github.com/m0nadicph0/kafsh/internal/models"
	"text/tabwriter"
)

type tabPrinter struct {
	w *tabwriter.Writer
}

func (t *tabPrinter) PrintGroupDescription(group *models.GroupDesc) {
	fmt.Fprintf(t.w, "Group ID:\t%v\n", group.GroupID)
	fmt.Fprintf(t.w, "State:\t%v\n", group.State)
	fmt.Fprintf(t.w, "Protocol:\t%v\n", group.Protocol)
	fmt.Fprintf(t.w, "Protocol Type:\t%v\n", group.ProtocolType)
	t.w.Flush()
}

func (t *tabPrinter) PrintGroups(groups []*models.Group) {
	_, _ = fmt.Fprintln(t.w, "NAME\tSTATE\tCONSUMERS\t")

	for _, group := range groups {
		_, _ = fmt.Fprintln(t.w, group.ToTSRec())
	}
	_ = t.w.Flush()
}

func (t *tabPrinter) PrintNodes(nodes []*models.Node) {
	_, _ = fmt.Fprintln(t.w, "ID\tADDRESS\tCONTROLLER\t")

	for _, node := range nodes {
		_, _ = fmt.Fprintln(t.w, node.ToTSRec())
	}
	_ = t.w.Flush()
}

func (t *tabPrinter) PrintTopicDesc(desc *models.TopicDesc) {
	fmt.Fprintf(t.w, "Name:\t%v\t\n", desc.Name)
	fmt.Fprintf(t.w, "Internal:\t%v\t\n", desc.Internal)
	fmt.Fprintf(t.w, "Compacted:\t%v\t\n", desc.Compacted)
	fmt.Fprintf(t.w, "Partitions:\n")
	fmt.Fprintln(t.w, "Partition\tLeader\tReplicas\tISR")
	fmt.Fprintln(t.w, "---------\t------\t--------\t---")
	for _, partition := range desc.Partitions {
		fmt.Fprintln(t.w, partition.ToTSRec())
	}
	fmt.Fprintf(t.w, "Config:\n")
	fmt.Fprintln(t.w, "Name\tValue\tReadOnly\tSensitive")
	fmt.Fprintln(t.w, "----\t-----\t--------\t---------")
	for _, config := range desc.Config {
		fmt.Fprintln(t.w, config.ToTSRec())
	}
	t.w.Flush()
}

func (t *tabPrinter) PrintTopics(topics []*models.Topic) {
	_, _ = fmt.Fprintln(t.w, "NAME\tPARTITIONS\tREPLICAS\t")

	for _, topic := range topics {
		_, _ = fmt.Fprintln(t.w, topic.ToTSRec())
	}
	_ = t.w.Flush()
}
