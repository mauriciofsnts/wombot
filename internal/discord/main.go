package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"code.db.cafe/wombot/internal/config"
	"github.com/bwmarrin/discordgo"
)

func Start() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Wombot.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.PermissionAdministrator

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
