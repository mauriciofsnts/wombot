package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"code.db.cafe/wombot/internal/config"
	"code.db.cafe/wombot/internal/discord/slash"
	"code.db.cafe/wombot/internal/discord/slash/categories/utils"
	"github.com/bwmarrin/discordgo"
)

func Start() error {

	dg, err := discordgo.New("Bot " + config.Wombot.Token)

	if err != nil {
		return err
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()

	if err != nil {
		return err
	}

	slash.RegisterSlashCommand(utils.Ping)
	slash.RegisterSlashCommand(utils.Setup)

	err = slash.Start(dg)

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
