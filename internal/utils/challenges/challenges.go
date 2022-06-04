package challenges

import (
	"os"

	"github.com/Pauloo27/logger"
	"github.com/ghodss/yaml"
)

type ChallengesModel struct {
	Name        string
	Description string
	Image       string
}

var (
	ChallengesData []ChallengesModel
)

func LoadChallenges() {
	f, err := os.ReadFile("./asset/challenges.yml")

	if err != nil {
		logger.Error("Failed to read challenge file", err)
		return
	}

	err = yaml.Unmarshal(f, &ChallengesData)

	if err != nil {
		logger.Error("Failed to parse challenges", err)
		return
	}
}
