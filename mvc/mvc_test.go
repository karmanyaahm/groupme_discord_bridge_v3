package mvc

import (
	"testing"
	"time"

	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/discord"
)

func TestGroupmeReceive(t *testing.T) {
	db.Parse()
	go discord.Main()
	time.Sleep(1 * time.Second)
	GroupmeReceive("name", "", "msgcontent", "60941041")

}
