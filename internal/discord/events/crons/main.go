package crons

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func StartCron(session *discordgo.Session) {
	s := gocron.NewScheduler(time.UTC)

	Challenges(s, session)
	Reactions(s, session)

	s.StartAsync()
}
