package status

import (
	"github.com/golangmc/minecraft-server/apis/data"
	"github.com/golangmc/minecraft-server/apis/data/chat"
	"github.com/golangmc/minecraft-server/impl/conf"
)

type Response struct {
	Version     Version `json:"version,string"`
	Players     Players `json:"players,string"`
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

func DefaultResponse() Response {
	return Response{
		Version: Version{
			Name:     "GoLang Server",
			Protocol: data.CurrentProtocol.Protocol(),
		},
		Players: Players{
			Max:    conf.Config.Motd.Max,
			Online: conf.Config.Motd.Online,
			Sample: []SamplePlayer{},
		},
		Description: Message{
			Text: chat.Translate(conf.Config.Motd.Line),
		},
		Favicon: conf.Config.Motd.Favicon,
	}
}
