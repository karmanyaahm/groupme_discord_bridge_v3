package db

import (
	"testing"
)

func TestAAA(t *testing.T) {
	Parse()
}

func TestSave(t *testing.T) {
	Parse()
	AddWebhook("60941041", "aaaa", "bbbb")
}

func TestChanIDConversion(t *testing.T) {
	Parse()
	a, err := BotIDFromChannelID("662789281453965315")
	if err != nil {
		t.Fail()
	}
	t.Log(a)

	a, err = ChannelIDFromGroupID("60941041")
	if err != nil {
		t.Fail()
	}
	t.Log(a)

	a, b, err := WebhookFromGroupID("60941041")
	if err != nil {
		t.Fail()
	}
	t.Log(a, b)
}
