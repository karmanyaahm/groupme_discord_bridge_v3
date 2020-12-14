package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/karmanyaahm/groupme_discord_bridge_v3/db"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/discord"
	"github.com/karmanyaahm/groupme_discord_bridge_v3/groupme"
)

func main() {
	db.Parse()
	//discord.Main()
	//groupme.Listen()

	log.Printf(`Now running. Press CTRL-C to exit.`)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		discord.Close()
		groupme.Close()

		os.Exit(0)
	}()

	discord.Main()
	groupme.Listen()
	for {
		log.Println("Sleeping")
		time.Sleep(1 * time.Hour)
	}
}
