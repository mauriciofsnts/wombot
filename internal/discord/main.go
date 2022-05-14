package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"code.db.cafe/wombot/internal/config"
	"github.com/bwmarrin/discordgo"
)

func Start() error {

	dg, err := discordgo.New("Bot " + config.Wombot.Token)

	if err != nil {
		return err
	}

	dg.Identify.Intents = discordgo.PermissionAdministrator

	err = dg.Open()
	if err != nil {
		return err
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

	return nil
}
