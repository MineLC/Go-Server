package cmds

import (
	"github.com/minelc/go-server-api/cmd"
)

type CommandManager struct {
	cmds map[string]*cmd.StructCommand
}

func (c *CommandManager) ReplaceManager(newManager cmd.CommandManager) {

}

func NewCommandManager() CommandManager {
	return CommandManager{
		cmds: make(map[string]*cmd.StructCommand),
	}
}

func (c *CommandManager) Add(command cmd.Command, values ...string) {
	for _, prefix := range values {
		tab, ok := command.(cmd.TabCommand)
		if ok {
			c.cmds[prefix] = &cmd.StructCommand{
				Execute: command.Execute,
				Tab:     tab.Tab,
			}
			continue
		}
		c.cmds[prefix] = &cmd.StructCommand{
			Execute: command.Execute,
		}
	}
}

func (c *CommandManager) AddStruct(command cmd.StructCommand, values ...string) {
	for _, prefix := range values {
		c.cmds[prefix] = &command
	}
}

func (c *CommandManager) Delete(values ...string) {
	for _, prefix := range values {
		delete(c.cmds, prefix)
	}
}

func (c *CommandManager) Get(prefix string) *cmd.StructCommand {
	if prefix[0] == '/' {
		prefix = prefix[1:]
	}
	return c.cmds[prefix]
}
