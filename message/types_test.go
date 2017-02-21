package message

import (
	"testing"
	"fmt"
	"time"
)

func TestMessage_RoundTrip(t *testing.T) {
	createTime := time.Now()
	m := Message{
		Created: createTime,
		Content: []byte("Burning all the bridges now"),
		Kind: MT_COMBAT,
	}
	b := m.ToBytes()
	if len(b) < 45 {
		t.Errorf("Message should have been 90 bytes, was %d", len(b))
		t.FailNow()
	}
	m1 := FromBytes(b)
	expectedContent := fmt.Sprintf("%s:%d:%s", MT_COMBAT, createTime.Unix(), string(m.Content))
	if string(m1.Content) != expectedContent {
		t.Errorf("Decoded content mismatch; expected=%s decoded=%s", expectedContent, m1.Content)
		t.FailNow()
	}
	if m1.Kind != MT_FROM_CLIENT {
		t.Errorf("Decoded content type wrong; expected=%s got=%s", MT_FROM_CLIENT, m1.Kind)
		t.FailNow()
	}
}
