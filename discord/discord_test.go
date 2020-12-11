package discord

import (
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
)

func TestConnection(t *testing.T) {
	db.Parse()
	Main()
	time.Sleep(time.Second)
	Close()
}

func TestNameProccessor(t *testing.T) {
	objs := map[string]discordgo.MessageCreate{
		//	"myname": discordgo.MessageCreate{
		//		&discordgo.Message{Author: &discordgo.User{Username: "myname"}}},
		"mynick": discordgo.MessageCreate{
			&discordgo.Message{Author: &discordgo.User{Username: "myname"}, Member: &discordgo.Member{Nick: "mynick"}}},
		"mynameagain": discordgo.MessageCreate{
			&discordgo.Message{Author: &discordgo.User{Username: "mynameagain"}, Member: &discordgo.Member{Nick: ""}}},
	}

	for ans, msg := range objs {
		t.Log("running: " + ans)
		if ans != nameFromMessage(&msg) {
			t.Log(ans)
			t.Fail()
		}
	}
}
