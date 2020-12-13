package mvc

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/google/go-cmp/cmp"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
)

var testConnections = []db.Connection{
	db.Connection{
		BotID:            "botid1",
		GroupID:          "groupid1",
		PrimaryChannelID: "chanid1.1",
		WebhookID:        "whid1",
		WebhookToken:     "wht1",
		ChannelID: map[string]string{
			"chanid1.1": "name1.1",
			"chanid1.2": "name1.2",
		},
	},
	db.Connection{
		BotID:            "botid2",
		GroupID:          "groupid2",
		PrimaryChannelID: "chanid2.1",
		ChannelID: map[string]string{
			"chanid2.1": "",
			"chanid2.2": "name2.2",
		},
	},
}

func parse() {
	db.SetConfig(db.Config{Connection: testConnections})
}

var testData = &discordgo.WebhookParams{
	Content:   "mycontent",
	Username:  "myname",
	AvatarURL: "myavatarurl",
	Embeds: []*discordgo.MessageEmbed{
		&discordgo.MessageEmbed{
			URL:   "myimageurl",
			Image: &discordgo.MessageEmbedImage{URL: "myimageurl"},
		},
	},
}

func TestGroupmeReceive(t *testing.T) {
	success := false

	ogCallwebhook := CallWebhook
	defer func() { CallWebhook = ogCallwebhook }()

	CallWebhook = func(webhookID, webhookToken string, data *discordgo.WebhookParams) {
		if d := cmp.Diff(testData, data); webhookID != "whid1" || webhookToken != "wht1" || d != "" {
			t.Log(webhookID, webhookToken, d)
		} else {
			success = true
		}
	}
	parse()
	GroupmeReceive("myname", "myavatarurl", "mycontent", "groupid1", []map[string]interface{}{
		{"type": "image", "url": "myimageurl"},
	})
	if !success {
		t.Fail()
	}
}
