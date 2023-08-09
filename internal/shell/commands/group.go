package commands

import (
	"errors"
	"fmt"
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/printer"
	"os"
)

func groupCmd(cli client.KafkaClient, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("sub command required")
	}

	handler, err := LookUpGroupSubCmd(args[0])
	if err != nil {
		return err
	}
	return handler(cli, args[1:])
}

func groupsCmd(cli client.KafkaClient, args []string) error {
	groups, err := cli.ListGroups()
	if err != nil {
		return err
	}

	printer.NewPrinter(printer.TabPrinter, os.Stdout).PrintGroups(groups)
	return nil
}

var groupCmdsHandlerMap = map[string]CmdHandler{
	"ls": groupsCmd,
}

func LookUpGroupSubCmd(cmd string) (CmdHandler, error) {
	h, ok := groupCmdsHandlerMap[cmd]
	if ok {
		return h, nil
	}
	return nil, errors.New("unsupported command")
}
