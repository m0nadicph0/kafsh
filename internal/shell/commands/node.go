package commands

import (
	"errors"
	"fmt"
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/printer"
	"os"
)

func nodeCmd(cli client.KafkaClient, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("sub command required")
	}

	handler, err := LookUpNodeSubCmd(args[0])
	if err != nil {
		return err
	}
	return handler(cli, args[1:])
}

func nodesCmd(cli client.KafkaClient, args []string) error {
	nodes, err := cli.ListNodes()
	if err != nil {
		return err
	}

	printer.NewPrinter(printer.TabPrinter, os.Stdout).PrintNodes(nodes)
	return nil
}

var nodeCmdsHandlerMap = map[string]CmdHandler{
	"ls": nodesCmd,
}

func LookUpNodeSubCmd(cmd string) (CmdHandler, error) {
	h, ok := nodeCmdsHandlerMap[cmd]
	if ok {
		return h, nil
	}
	return nil, errors.New("unsupported command")
}
