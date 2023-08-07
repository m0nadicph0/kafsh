package printer

import (
	"fmt"
	"github.com/m0nadicph0/kafsh/internal/models"
	"text/tabwriter"
)

type tabPrinter struct {
	w *tabwriter.Writer
}

func (t *tabPrinter) PrintTopics(topics []*models.Topic) {
	_, _ = fmt.Fprintln(t.w, "NAME\tPARTITIONS\tREPLICAS\t")

	for _, topic := range topics {
		_, _ = fmt.Fprintln(t.w, topic.ToTSRec())
	}
	_ = t.w.Flush()
}
