package network

import (
	"testing"
	"fmt"
	"time"
)

func TestMessage_RoundTrip(t *testing.T) {
	createTime := time.Now()
	m := Message{
		created: createTime,
		content: "Burning all the bridges now",
		kind: MT_COMBAT,
	}
	b := m.ToBytes()
	if len(b) < 45 {
		t.Errorf("Message should have been 90 bytes, was %d", len(b))
		t.FailNow()
	}
	m1 := MessageFromBytes(b)
	expectedContent := fmt.Sprintf("%s:%d:%s", MT_COMBAT, createTime.Unix(), m.content)
	if m1.content != expectedContent {
		t.Errorf("Decoded content mismatch; expected=%s decoded=%s", expectedContent, m1.content)
		t.FailNow()
	}
	if m1.kind != MT_FROM_CLIENT {
		t.Errorf("Decoded content type wrong; expected=%s got=%s", MT_FROM_CLIENT, m1.kind)
		t.FailNow()
	}
}
