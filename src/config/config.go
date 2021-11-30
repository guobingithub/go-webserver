package config

type Config struct {
	WebServer struct {
		Address string       `mapstructure:"addr"`
		Logger  LoggerConfig `mapstructure:"logger"`
	} `mapstructure:"webserver"`

	DataBase struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"database"`

	EMS struct {
		Address string `mapstructure:"addr"`
		User    string `mapstructure:"user"`
		Passwd  string `mapstructure:"passwd"`
	} `mapstructure:"ems"`
}

type LoggerConfig struct {
	Level         string   `mapstructure:"level"`
	EnableContext bool     `mapstructure:"enable_context"`
	MustSetFields []string `mapstructure:"must_set_fields"`
}
