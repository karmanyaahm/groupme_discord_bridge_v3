package config

import "github.com/bwmarrin/discordgo"

var Addr string
var Discord_Session *discordgo.Session

func init() {
	Addr = "localhost:5000"
}
