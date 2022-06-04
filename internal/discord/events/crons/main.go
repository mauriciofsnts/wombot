package crons

import (
	"time"

	"code.db.cafe/wombot/internal/utils/challenges"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

func StartCron(session *discordgo.Session) {
	challenges.LoadChallenges()

	s := gocron.NewScheduler(time.UTC)

	Challenges(s, session)
	Reactions(s, session)

	s.StartAsync()
}
