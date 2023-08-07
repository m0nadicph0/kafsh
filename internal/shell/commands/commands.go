package commands

import (
	"errors"
	"github.com/m0nadicph0/kafsh/internal/client"
	"os"
)

type CmdHandler func(kCli client.KafkaClient, args []string) error

var handlerMap = map[string]CmdHandler{
	"exit": func(kCli client.KafkaClient, args []string) error {
		os.Exit(1)
		return nil
	},
	"topics": topicsCmd,
	"topic":  topicCmd,
}

func LookUp(cmd string) (CmdHandler, error) {
	h, ok := handlerMap[cmd]
	if ok {
		return h, nil
	}
	return nil, errors.New("unsupported command")
}
