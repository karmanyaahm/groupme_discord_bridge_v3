package main

import (
	"testing"
	"time"

	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/discord"
)

func TestDiscordConnection(t *testing.T) {
	db.Parse()
	discord.Main()
	time.Sleep(500 * time.Millisecond)
	discord.Close()
}
