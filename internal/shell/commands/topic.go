package commands

import (
	"errors"
	"fmt"
	"github.com/m0nadicph0/kafsh/internal/cache"
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/constants"
	"github.com/m0nadicph0/kafsh/internal/printer"
	"github.com/m0nadicph0/kafsh/internal/udf"
	"github.com/m0nadicph0/kafsh/internal/util"
	"github.com/spf13/pflag"
	"os"
)

func topicCmd(cli client.KafkaClient, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("sub command required")
	}

	handler, err := LookUpTopicSubCmd(args[0])
	if err != nil {
		return err
	}
	return handler(cli, args[1:])
}

func topicsCmd(cli client.KafkaClient, args []string) error {
	topics, err := cli.ListTopics()
	if err != nil {
		return err
	}

	cache.Set(constants.CacheKeyTopicNames, util.GetTopicNames(topics))
	printer.NewPrinter(printer.TabPrinter, os.Stdout).PrintTopics(topics)
	return nil
}

var topicCmdsHandlerMap = map[string]CmdHandler{
	"create":   createTopicCmd,
	"delete":   deleteTopicCmd,
	"describe": describeTopicCmd,
	"ls":       topicsCmd,
}

func LookUpTopicSubCmd(cmd string) (CmdHandler, error) {
	h, ok := topicCmdsHandlerMap[cmd]
	if ok {
		return h, nil
	}
	return nil, errors.New("unsupported command")
}

func createTopicCmd(cli client.KafkaClient, args []string) error {
	fs := pflag.NewFlagSet("createTopic", pflag.ContinueOnError)

	var partitions int32
	var replication int16

	fs.Int32VarP(&partitions, "partitions", "p", 1, "Number of partitions")
	fs.Int16VarP(&replication, "replicas", "r", 1, "Number of replicas")

	err := fs.Parse(args)
	if err != nil {
		fmt.Println("USAGE:")
		fmt.Println(fs.FlagUsages())
		return err
	}

	topicName := fs.Arg(0)
	name, err := udf.Expand(topicName)
	if err != nil {
		return err
	}

	err = cli.CreateTopic(name, partitions, replication)
	if err != nil {
		return err
	}

	util.Success("created topic %s with partitions=%d and replication=%d\n", name, partitions, replication)
	return nil
}

func deleteTopicCmd(cli client.KafkaClient, args []string) error {
	for _, arg := range args {
		err := cli.DeleteTopic(arg)
		if err != nil {
			return err
		}
		util.Success("deleted topic %s\n", arg)
	}
	return nil
}

func describeTopicCmd(cli client.KafkaClient, args []string) error {
	topicDesc, err := cli.DescribeTopic(args[0])
	if err != nil {
		return err
	}
	printer.NewPrinter(printer.TabPrinter, os.Stdout).PrintTopicDesc(topicDesc)

	return nil
}
