package message

import (
	"time"
	"fmt"
)

type MessageType uint8
const (
	MT_UNKNOWN MessageType = iota
	MT_FROM_CLIENT
	MT_SYSTEM
	MT_SAY
	MT_EMOTE
	MT_TELL
	MT_DESCRIPTION
	MT_COMBAT
	MT_DISCONNECT
	MT_CONNECT
	MT_OTHER
)

func (mt MessageType) String() string {
	switch mt {
	case MT_UNKNOWN:
		return "UNKNOWN"
	case MT_FROM_CLIENT:
		return "FROM_CLIENT"
	case MT_SYSTEM:
		return "SYSTEM"
	case MT_SAY:
		return "SAY"
	case MT_EMOTE:
		return "EMOTE"
	case MT_TELL:
		return "TELL"
	case MT_DESCRIPTION:
		return "DESCRIPTION"
	case MT_COMBAT:
		return "COMBAT"
	case MT_DISCONNECT:
		return "DISCONNECT"
	case MT_CONNECT:
		return "CONNECT"
	case MT_OTHER:
		return "OTHER"
	}

	return "UNDEFINED"
}

type Message struct {
	Kind    MessageType
	Created time.Time
	Content []byte
}

func (m *Message) ToBytes() []byte {
	// convert the content to double-byte runes
	return []byte(fmt.Sprintf("%s:%d:%s", m.Kind, m.Created.Unix(), m.Content))
}

func (m *Message) String() string {
	return fmt.Sprintf("Message{kind=%s,created=%s,content=%s}", m.Kind, m.Created, m.Content)
}

func FromBytes(b []byte) *Message {
	m := &Message{
		Kind: MT_FROM_CLIENT,
		Created: time.Now(),
		Content: b,
	}
	return m
}

func New(content []byte, kind MessageType) *Message {
	return &Message{
		Content: content,
		Created: time.Now(),
		Kind: kind,
	}
}