package config

type Config struct {
	Token string
	Pg    struct {
		Host     string
		Port     int
		Username string
		Password string
		// snif snif
		DbName string `yaml:"db_name"`
	}
}
