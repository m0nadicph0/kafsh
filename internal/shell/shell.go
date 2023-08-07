package shell

import (
	"fmt"
	"github.com/buildkite/shellwords"
	"github.com/chzyer/readline"
	"github.com/m0nadicph0/kafsh/internal/client"
	"github.com/m0nadicph0/kafsh/internal/shell/commands"
	"io"
	"log"
	"strings"
)

const Prompt = "\033[31m»\033[0m "

var completer = readline.NewPrefixCompleter(
	readline.PcItem("topics"),
	readline.PcItem("topic",
		readline.PcItem("create"),
		readline.PcItem("delete"),
		readline.PcItem("describe"),
		readline.PcItem("ls"),
	),
	readline.PcItem("exit"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}
func Start(kCli client.KafkaClient) error {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          Prompt,
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		return err
	}
	defer l.Close()
	l.CaptureExitSignal()

	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		words, err := shellwords.Split(line)

		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}

		if len(words) > 0 {
			handler, err := commands.LookUp(words[0])
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
				continue
			}
			err = handler(kCli, words[1:])
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
				continue
			}
		}
	}

	return nil
}
