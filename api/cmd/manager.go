package cmd

type CommandManager struct {
	cmds map[string]*StructCommand
}

func NewCommandManager() CommandManager {
	return CommandManager{
		cmds: make(map[string]*StructCommand),
	}
}

func (c *CommandManager) Add(cmd Command, values ...string) {
	for _, prefix := range values {
		tab, ok := cmd.(TabCommand)
		if ok {
			c.cmds[prefix] = &StructCommand{
				Execute: cmd.Execute,
				Tab:     tab.Tab,
			}
			continue
		}
		c.cmds[prefix] = &StructCommand{
			Execute: cmd.Execute,
		}
	}
}

func (c *CommandManager) AddStruct(cmd StructCommand, values ...string) {
	for _, prefix := range values {
		c.cmds[prefix] = &cmd
	}
}

func (c *CommandManager) Delete(values ...string) {
	for _, prefix := range values {
		delete(c.cmds, prefix)
	}
}

func (c *CommandManager) Get(prefix string) *StructCommand {
	if prefix[0] == '/' {
		prefix = prefix[1:]
	}
	return c.cmds[prefix]
}
