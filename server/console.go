package server

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/cmd"
	"github.com/minelc/go-server-api/data/chat"
)

var codeToForm = map[chat.ChatColor]color.Attribute{
	chat.DarkRed:    color.FgHiRed,
	chat.Red:        color.FgRed,
	chat.Gold:       color.FgYellow,
	chat.Yellow:     color.FgHiYellow,
	chat.DarkGreen:  color.FgGreen,
	chat.Green:      color.FgHiGreen,
	chat.DarkAqua:   color.FgCyan,
	chat.Aqua:       color.FgHiCyan,
	chat.DarkBlue:   color.FgBlue,
	chat.Blue:       color.FgHiBlue,
	chat.DarkPurple: color.FgMagenta,
	chat.Purple:     color.FgHiMagenta,
	chat.White:      color.FgHiWhite,
	chat.Black:      color.FgBlack,
	chat.DarkGray:   color.FgHiBlack,
	chat.Gray:       color.FgWhite,

	chat.Reset:         color.Reset,
	chat.Obfuscated:    color.BlinkRapid,
	chat.Bold:          color.Bold,
	chat.Strikethrough: color.CrossedOut,
	chat.Underline:     color.Underline,
	chat.Italic:        color.Italic,
}

var charToCode = map[rune]chat.ChatColor{
	'4': chat.DarkRed,
	'c': chat.Red,
	'6': chat.Gold,
	'e': chat.Yellow,

	'2': chat.DarkGreen,
	'a': chat.Green,

	'3': chat.DarkAqua,
	'b': chat.Aqua,

	'1': chat.DarkBlue,
	'9': chat.Blue,

	'5': chat.DarkPurple,
	'd': chat.Purple,

	'f': chat.White,
	'0': chat.Black,

	'8': chat.DarkGray,
	'7': chat.Gray,

	'k': chat.Obfuscated,
	'l': chat.Bold,
	'm': chat.Strikethrough,
	'n': chat.Underline,
	'o': chat.Italic,
	'r': chat.Reset,
}

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
		fmt.Println(translateConsole(msg))
	}
}

func (c *Console) start() {
	reader := bufio.NewReader(os.Stdin)
	cmdManager := api.GetServer().GetPluginManager().GetCommandManager()
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
		command := cmdManager.Get(split[0])
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

func translateConsole(text string) string {
	text = chat.Translate(text)

	build := strings.Builder{}
	temps := strings.Builder{}

	chars := []rune(text)
	forms := make([]color.Attribute, 0)

	for i := 0; i < len(chars); i++ {
		r := chars[i]
		if r != chat.ColorCChar || i+1 >= len(chars) {
			temps.WriteRune(r)
			continue
		}

		f, con := codeToForm[charToCode[chars[i+1]]]
		if !con {
			temps.WriteRune(r)
			continue
		}

		if temps.Len() > 0 {
			build.WriteString(color.New(forms...).Sprint(temps.String()))
			temps.Reset()
		}

		i++
		if f <= color.CrossedOut {
			forms = append(forms, f)
		} else {
			forms = make([]color.Attribute, 0)
			forms = append(forms, f)
		}
	}

	if temps.Len() > 0 {
		build.WriteString(color.New(forms...).Sprint(temps.String()))
	}

	return build.String()
}
