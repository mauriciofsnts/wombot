package config

type Config struct {
	Token string
	Pg    struct {
		Host     string
		Port     int
		Username string
		Password string
		DbName   string
	}
}
