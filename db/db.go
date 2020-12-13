package db

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/pelletier/go-toml"
	"github.com/spf13/viper"
)

var conf Config
var DiscordToken string
var Addr string

type Connection struct {
	BotID   string `mapstructure:"groupme_bot_id" toml:"groupme_bot_id"`
	GroupID string `mapstructure:"groupme_group_id" toml:"groupme_group_id"`

	ChannelID        map[string]string `mapstructure:"discord_channels" toml:"discord_channels"`
	PrimaryChannelID string            `mapstructure:"primary_channel" toml:"primary_channel"`
	WebhookID        string            `mapstructure:"webhook_id,omitempty" toml:"webhook_id"`
	WebhookToken     string            `mapstructure:"webhook_token,omitempty" toml:"webhook_token"`
}

type Config struct {
	Address         string
	DiscordBotToken string
	Connection      []Connection
}

func postParseInit() {
	//Primary Channel Mod
	for i, con := range conf.Connection { //for every channelid
		if len(con.ChannelID) == 0 { //create if doesn't exist
			conf.Connection[i].ChannelID = map[string]string{}
		}

		if conf.Connection[i].ChannelID[con.PrimaryChannelID] == "" { //if primarychannel isn't a value
			conf.Connection[i].ChannelID[con.PrimaryChannelID] = "" //make it one
		}
		//	for j, k := range conf.Connection[i].ChannelID {
		//		log.Println(i, j, k, len(k))
		//	}
		//	log.Println(con.ChannelID["chanid1."])
	}

	//Bot Identifier Discord
	if !strings.HasPrefix(conf.DiscordBotToken, "Bot ") {
		conf.DiscordBotToken = "Bot " + conf.DiscordBotToken
	}

}

func Parse() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	dbRead()
	conf = Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("unable to decode into config struct, %v", err)
	}
	postParseInit()

	DiscordToken = conf.DiscordBotToken
	Addr = conf.Address
}

var dbRead = func() {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
var dbWrite = func(b []byte) error {
	err := ioutil.WriteFile(viper.ConfigFileUsed(), b, 0644)
	return err
}

func updateDB(conf Config) error {

	b, err := toml.Marshal(conf)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := dbWrite(b); err != nil {
		log.Println(err)
		return err
	}
	Parse()
	return nil
}
func BotIDFromChannelID(s string) (string, error) {
	for _, i := range conf.Connection {
		for j := range i.ChannelID {
			if j == s {
				return i.BotID, nil
			}
		}

	}
	return "", errors.New("Not Found")
}

func ChannelNameFromChannelID(s string) (string, error) {
	for _, i := range conf.Connection {
		for j, k := range i.ChannelID {
			if j == s {

				if len(i.ChannelID) == 1 {
					return "", errors.New("Only Name")
				} else if len(k) == 0 {
					return "", errors.New("No Name")

				} else {
					return k, nil
				}
			}
		}
	}
	return "", errors.New("Not Found")
}

func ChannelIDFromGroupID(s string) (string, error) {
	for _, i := range conf.Connection {
		if i.GroupID == s {
			return i.PrimaryChannelID, nil
		}
	}
	return "", errors.New("Not Found")
}

func WebhookFromGroupID(s string) (string, string, error) {
	for _, i := range conf.Connection {
		if i.GroupID == s {
			if (len(i.WebhookID) == 0) || (len(i.WebhookToken) == 0) {
				return "", "", errors.New("No Webhook")
			}
			return i.WebhookID, i.WebhookToken, nil
		}
	}
	return "", "", errors.New("Not Found")
}

//TODO: Lock something during both Add methods to avoid race condition
func AddWebhook(gid, wi, wt string) error {
	myConf := Config{}
	copier.Copy(&myConf, &conf)
	var index int
	for j, i := range myConf.Connection {
		if i.GroupID == gid {
			index = j
			break
		}
	}
	myConf.Connection[index].WebhookID = wi
	myConf.Connection[index].WebhookToken = wt
	return updateDB(myConf)
}

func AddChannelName(channelID, channelName string) error {
	myConf := Config{}
	copier.Copy(&myConf, &conf)

	var index int
	for k, i := range myConf.Connection {
		for j := range i.ChannelID {
			if j == channelID {
				index = k
			}
		}
	}
	myConf.Connection[index].ChannelID[channelID] = channelName
	return updateDB(myConf)

}
