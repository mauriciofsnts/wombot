package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"code.db.cafe/wombot/internal/config"
	"code.db.cafe/wombot/internal/discord/commands"
	"code.db.cafe/wombot/internal/discord/commands/categories/utils"
	"github.com/bwmarrin/discordgo"
)

func Start() error {

	dg, err := discordgo.New("Bot " + config.Wombot.Token)

	if err != nil {
		return err
	}

	registerCommands(dg, config.Wombot)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

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

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := commands.NewCommandHandler(cfg.Prefix)

	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command Execution failed: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&utils.CmdPing{})

	s.AddHandler(cmdHandler.HandleMessage)
}
