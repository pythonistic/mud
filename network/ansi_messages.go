package network

import (
	"github.com/mgutz/ansi"
	"mud/message"
)

var tell = ansi.ColorFunc("magenta:white")
var system = ansi.ColorFunc("blue:white")
var connect = ansi.ColorFunc("white:black")
var disconnect = ansi.ColorFunc("white:black")
var combat = ansi.ColorFunc("red:white")
var description = ansi.ColorFunc("black:white")
var emote = ansi.ColorFunc("green:white")
var other = ansi.ColorFunc("black:cyan")
var say = ansi.ColorFunc("black:white")

func AnsiAdapter(msg *message.Message) []byte {
	content := string(msg.Content)

	switch msg.Kind {
	case message.MT_TELL:
		content = tell(content)
	case message.MT_SYSTEM:
		content = system(content)
	case message.MT_CONNECT:
		content = connect(content)
	case message.MT_DISCONNECT:
		content = disconnect(content)
	case message.MT_COMBAT:
		content = combat(content)
	case message.MT_DESCRIPTION:
		content = description(content)
	case message.MT_EMOTE:
		content = emote(content)
	case message.MT_OTHER:
		content = other(content)
	case message.MT_SAY:
		content = say(content)
	default:
		content = other(content)
	}

	return []byte(content)
}
