package discord

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

//func TestConnection(t *testing.T) {
//	db.DiscordToken = "Bot "
//	Main()
//	time.Sleep(time.Second)
//	Close()
//}

func TestNameProccessor(t *testing.T) {
	objs := map[string]discordgo.MessageCreate{
		//	"myname": discordgo.MessageCreate{
		//		&discordgo.Message{Author: &discordgo.User{Username: "myname"}}},
		"mynick":      discordgo.MessageCreate{Message: &discordgo.Message{Author: &discordgo.User{Username: "myname"}, Member: &discordgo.Member{Nick: "mynick"}}},
		"mynameagain": discordgo.MessageCreate{Message: &discordgo.Message{Author: &discordgo.User{Username: "mynameagain"}, Member: &discordgo.Member{Nick: ""}}},
	}

	for ans, msg := range objs {
		t.Log("running: " + ans)
		if ans != nameFromMessage(msg.Message) {
			t.Log(ans)
			t.Fail()
		}
	}
}
