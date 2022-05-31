package entities

type User struct {
	UserID        string `gorm:"primary_key"`
	GuildID       string
	Guild         Guild
	Streak        int64 `gorm:"not null"`
	HighestStreak int64 `gorm:"not null"`
	CurrentDay    int64 `gorm:"not null"`
}
