package discord

import (
	"os"
	"os/signal"
	"syscall"

	"code.db.cafe/wombot/internal/config"
	"code.db.cafe/wombot/internal/discord/events"
	"code.db.cafe/wombot/internal/discord/events/crons"
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"

	// register commands
	_ "code.db.cafe/wombot/internal/discord/events/commands"
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

	err = events.Start(dg)

	if err != nil {
		return err
	}

	crons.StartCron(dg)

	logger.Success("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

	return nil
}
