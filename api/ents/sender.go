package ents

type Sender interface {
	SendMsg(message ...string)
	SendMsgColor(message ...string)
}

type Console interface {
	Sender
}
