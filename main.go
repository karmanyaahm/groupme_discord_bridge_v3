package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/discord"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/groupme"
)

func main() {
	db.Parse()
	discord.Main()
	groupme.Listen()

	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		oscall := <-sc
		log.Printf("system call:%+v", oscall)

		discord.Close()
		groupme.Close()
	}()
}
