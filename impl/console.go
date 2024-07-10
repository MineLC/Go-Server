package impl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/minelc/go-server/api"
	"github.com/minelc/go-server/api/cmd"
	"github.com/minelc/go-server/api/data/chat"
)

type Console struct {
	pendientCommand *cmd.StructCommand
	cmdArgs         []string
}

func (c *Console) SendMsg(messages ...string) {
	for _, msg := range messages {
		fmt.Println(msg)
	}
}

func (c *Console) SendMsgColor(messages ...string) {
	for _, msg := range messages {
		fmt.Println(chat.TranslateConsole(msg))
	}
}

func (c *Console) start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')

		if c.pendientCommand != nil {
			c.SendMsgColor("&cWait to execute next command")
			continue
		}

		if len(input) <= 1 {
			continue
		}

		input = strings.Replace(input, "\n", "", -1)
		split := strings.Split(input, " ")
		command := api.GetServer().GetCommandManager().Get(split[0])
		if command == nil {
			c.SendMsgColor("&cCommand inexistent | &6" + split[0])
			continue
		}
		c.pendientCommand = command

		if len(split) > 1 {
			c.cmdArgs = split[1:]
			continue
		}
		c.cmdArgs = nil
	}
}

func (c *Console) executePendient() {
	if c.pendientCommand == nil {
		return
	}
	c.pendientCommand.Execute(c, c.cmdArgs)
	c.pendientCommand = nil
}
