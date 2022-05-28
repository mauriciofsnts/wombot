package slash

import (
	"code.db.cafe/wombot/internal/i18n"
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
)

type DiscordContext struct {
	S *discordgo.Session
	I *discordgo.InteractionCreate
}

type SlashCommand struct {
	*discordgo.ApplicationCommand
	Handler func(ctx *DiscordContext, t *i18n.Language)
}

var commands = make(map[string]*SlashCommand)

func RegisterSlashCommand(cmds ...*SlashCommand) {
	for _, command := range cmds {
		commands[command.Name] = command
	}
}

func Start(s *discordgo.Session) error {

	applicationCommands := make([]*discordgo.ApplicationCommand, len(commands))

	i := 0
	for _, c := range commands {
		logger.Debugf("Registering command: %s", c.Name)

		applicationCommands[i] = c.ApplicationCommand
		i++
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commandName := i.ApplicationCommandData().Name

		if command, ok := commands[commandName]; ok {
			command.Handler(
				&DiscordContext{
					S: s,
					I: i,
				},
				i18n.GetLanguage(i18n.LanguageCria),
			)
		}

	})

	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", applicationCommands)

	return err
}
