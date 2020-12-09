package db

import (
	"testing"
)

func TestAAA(t *testing.T) {
	Parse()
}

func TestChanIDConversion(t *testing.T) {
	Parse()
	a, err := BotIDFromChannelID("662789281453965315")
	if err != nil {
		t.Fail()
	}
	t.Log(a)
	a, err = ChannelIDFromBotID(a)
	if err != nil {
		t.Fail()
	}
	t.Log(a)
}
