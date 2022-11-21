package config

var DefaultConfigYaml = `
server:
  host: 0.0.0.0
  port: 8000
spider:
  fs: "/spiders"
  workspace: "/workspace"
task:
  workers: 16
  cancelWaitSeconds: 30
grpc:
  address: localhost:9666
  server:
    address: 0.0.0.0:9666
  authKey: Crawlab2021!
`

type YamlSchema struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`

	Spider struct {
		Fs        string `mapstructure:"fs"`
		Workspace string `mapstructure:"workspace"`
	} `mapstructure:"spider"`

	Task struct {
		Workers           int `mapstructure:"workers"`
		CancelWaitSeconds int `mapstructure:"cancelWaitSeconds"`
	} `mapstructure:"task"`

	Grpc struct {
		Address string     `mapstructure:"address"`
		Server  GrpcServer `mapstructure:"server"`
		AuthKey string     `mapstructure:"authKey"`
	} `mapstructure:"grpc"`
}

type GrpcServer struct {
	Address string `mapstructure:"address"`
}
