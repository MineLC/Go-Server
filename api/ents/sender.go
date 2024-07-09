package ents

type Sender interface {
	SendMessage(message ...interface{})
}
