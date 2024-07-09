package motd

import (
	"encoding/json"

	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/impl/conf"
)

type Response struct {
	Version     Version `json:"version"`
	Players     Players `json:"players"`
	Description Message `json:"description"`
	Favicon     string  `json:"favicon"`
}

type Version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type Players struct {
	Max    int            `json:"max"`
	Online int            `json:"online"`
	Sample []SamplePlayer `json:"sample"`
}

type SamplePlayer struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Message struct {
	Text string `json:"text"`
}

var response string

func SetResponse(motd conf.Motd) {
	res := Response{
		Version: Version{
			Name:     "GoLang Server",
			Protocol: 47,
		},
		Players: Players{
			Max:    motd.Max,
			Online: motd.Online,
			Sample: []SamplePlayer{},
		},
		Description: Message{
			Text: chat.Translate(motd.Line),
		},
		Favicon: motd.Favicon,
	}
	text, err := json.Marshal(res)
	if err != nil {
		return
	}
	response = string(text)
}

func GetResponse() string {
	return response
}
