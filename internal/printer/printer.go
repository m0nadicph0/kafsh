package printer

import (
	"github.com/m0nadicph0/kafsh/internal/models"
	"io"
	"os"
	"text/tabwriter"
)

type Type int

const (
	TabPrinter   = 0
	TablePrinter = 1
)

type Printer interface {
	PrintTopics([]*models.Topic)
	PrintTopicDesc(*models.TopicDesc)
	PrintNodes([]*models.Node)
	PrintGroups([]*models.Group)
	PrintGroupDescription(group *models.GroupDesc)
}

func NewPrinter(t Type, out io.Writer) Printer {
	switch t {
	case TablePrinter:
		return &tablePrinter{}
	case TabPrinter:
		return &tabPrinter{w: tabwriter.NewWriter(os.Stdout, 6, 4, 2, ' ', 0)}
	default:
		return &tabPrinter{w: tabwriter.NewWriter(os.Stdout, 6, 4, 2, ' ', 0)}
	}

}
