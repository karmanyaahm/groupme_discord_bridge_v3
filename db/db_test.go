package db

import (
	"bytes"
	"log"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/spf13/viper"
)

//func TestParse(t *testing.T) {
//	Parse()
//}

//func TestSave(t *testing.T) {
//	Parse()
//	AddWebhook("60941041", "aaaa", "bbbb")
//}

//TODO: check for condition where some ids are the same in different structs
var testConnections = []Connection{
	Connection{
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
	Connection{
		BotID:            "botid2",
		GroupID:          "groupid2",
		PrimaryChannelID: "chanid2.1",
		ChannelID: map[string]string{
			"chanid2.1": "",
			"chanid2.2": "name2.2",
		},
	},
	Connection{
		BotID:            "botid3",
		GroupID:          "groupid3",
		PrimaryChannelID: "chanid3.1",
		ChannelID: map[string]string{
			"chanid3.1": "",
		},
	},
	Connection{
		BotID:            "botid4",
		GroupID:          "groupid4",
		PrimaryChannelID: "chanid4.1",
		ChannelID: map[string]string{
			"chanid4.1": "name4.1",
		},
	},
	Connection{
		BotID:            "botid5",
		GroupID:          "groupid5",
		PrimaryChannelID: "chanid5.1",
		ChannelID:        nil,
	},
}
var testStdCases = []struct {
	inp      string
	isErr    bool
	out      string
	function func(string) (string, error)
}{
	{"chanid1.1", false, "botid1", BotIDFromChannelID},
	{"chanid2.2", false, "botid2", BotIDFromChannelID},
	{"chanid5.1", false, "botid5", BotIDFromChannelID},
	{"chanidNone", true, "Not Found", BotIDFromChannelID},

	{"chanid1.1", false, "name1.1", ChannelNameFromChannelID},
	{"chanid2.2", false, "name2.2", ChannelNameFromChannelID},
	{"chanid2.1", true, "No Name", ChannelNameFromChannelID},
	{"chanid3.1", true, "Only Name", ChannelNameFromChannelID},
	{"chanid4.1", true, "Only Name", ChannelNameFromChannelID},
	{"chanid5.1", true, "Only Name", ChannelNameFromChannelID},
	{"chanidNone", true, "Not Found", ChannelNameFromChannelID},

	{"groupid1", false, "chanid1.1", ChannelIDFromGroupID},
	{"groupid5", false, "chanid5.1", ChannelIDFromGroupID},
	{"groupNone", true, "Not Found", ChannelIDFromGroupID},
}

var testConfigStr = `
[[Connection]]
groupme_group_id = "gid"
primary_channel = "cid"
`
var testAddWebhookStr = `Address = ""
DiscordBotToken = "Bot "

[[Connection]]
  groupme_bot_id = ""
  groupme_group_id = "gid"
  primary_channel = "cid"
  webhook_id = "wi"
  webhook_token = "wt"

  [Connection.discord_channels]
    cid = ""
`

var testAddChannelNameStr = `Address = ""
DiscordBotToken = "Bot "

[[Connection]]
  groupme_bot_id = ""
  groupme_group_id = "gid"
  primary_channel = "cid"
  webhook_id = ""
  webhook_token = ""

  [Connection.discord_channels]
    cid = "name"
`

func parse() {
	//Parse()
	conf = Config{Connection: testConnections}
	postParseInit()
}

func TestStringFuncs(t *testing.T) {
	parse()
	for _, i := range testStdCases {
		a, err := i.function(i.inp)
		if i.isErr {
			if err.Error() != i.out {
				t.Log("Error Not Matched: ", i.inp, i.out, a, err)
				t.Fail()
			}
		} else {
			if i.out != a || err != nil {
				t.Log("Fail: ", i.inp, i.out, a, err)
				t.Fail()
			}
		}
	}
}

func TestWebhookGroupID(t *testing.T) {
	parse()
	a, b, err := WebhookFromGroupID("groupid1")
	if err != nil || a != "whid1" || b != "wht1" {
		t.Log("Webhook: ", "groupid1", a, b, err)
	}
	a, b, err = WebhookFromGroupID("groupid2")
	if err.Error() != "No Webhook" {
		t.Log("Webhook: ", "groupid1", a, b, err)
	}

	a, b, err = WebhookFromGroupID("groupidNone")
	if err.Error() != "Not Found" {
		t.Log("Webhook: ", "groupid1", a, b, err)
	}

}
func TestParse(t *testing.T) {
	ogdbRead := dbRead
	ogdbWrite := dbWrite
	defer func() { dbRead = ogdbRead }()
	defer func() { dbWrite = ogdbWrite }()

	buf := []byte(testConfigStr)

	dbRead = func() {
		viper.SetConfigType("toml")
		viper.ReadConfig(bytes.NewReader(buf))
	}

	dbWrite = func(b []byte) error {
		buf = b
		return nil
	}

	Parse()

	AddWebhook("gid", "wi", "wt")

	if b := string(buf); b != testAddWebhookStr {
		log.Println(diff.CharacterDiff(testAddWebhookStr, b))
		t.Fail()
	}

	buf = []byte(testConfigStr)
	Parse()

	AddChannelName("cid", "name")
	if b := string(buf); b != testAddChannelNameStr {
		log.Println(diff.CharacterDiff(testAddChannelNameStr, b))
		t.Fail()
	}

}

//var testConfig = Config{Address: "localhost:5000", Discord
