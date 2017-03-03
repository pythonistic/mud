package network

import (
	"github.com/mgutz/ansi"
)

var colorizeTell = ansi.ColorFunc("magenta:white")
var colorizeSystem = ansi.ColorFunc("blue:white")
var colorizeConnect = ansi.ColorFunc("white:black")
var colorizeDisconnect = ansi.ColorFunc("white:black")
var colorizeCombat = ansi.ColorFunc("red:white")
var colorizeDescription = ansi.ColorFunc("black:white")
var colorizeEmote = ansi.ColorFunc("green:white")
var colorizeOther = ansi.ColorFunc("black:cyan")
var colorizeSay = ansi.ColorFunc("black:white")
var colorizeError = ansi.ColorFunc("red+h:white+h")

func AnsiAdapter(msg *Message) []byte {
	content := string(msg.Content)

	switch msg.Kind {
	case MT_TELL:
		content = colorizeTell(content)
	case MT_SYSTEM:
		content = colorizeSystem(content)
	case MT_CONNECT:
		content = colorizeConnect(content)
	case MT_DISCONNECT:
		content = colorizeDisconnect(content)
	case MT_COMBAT:
		content = colorizeCombat(content)
	case MT_DESCRIPTION:
		content = colorizeDescription(content)
	case MT_EMOTE:
		content = colorizeEmote(content)
	case MT_OTHER:
		content = colorizeOther(content)
	case MT_SAY:
		content = colorizeSay(content)
	case MT_ERROR:
		content = colorizeError(content)
	default:
		content = colorizeOther(content)
	}

	return []byte(content)
}
