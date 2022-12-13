package config

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"log"
	"path"
	"strings"
)

type Config struct {
	Name string
}

func init() {
	// config instance
	c := Config{Name: ""}

	// init config file
	if err := c.Init(); err != nil {
		log.Print("unable to init config")
		return
	}
}

func (c *Config) Init() (err error) {
	println("0000000==========================")
	println(HomeDirPath)
	println("0000000==========================")
	// config
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // if config file is set, load it accordingly
	} else {
		//path.Join(HomeDirPath, configDirName, ConfigName)
		viper.AddConfigPath(path.Join(HomeDirPath, "workdir", ".crawlab")) // if no config file is set, load by default
		viper.SetConfigName("config")
	}

	// config type as yaml
	viper.SetConfigType("yaml") // default yaml

	// auto env
	viper.AutomaticEnv() // load matched environment variables

	// env prefix
	viper.SetEnvPrefix("CRAWLAB") // environment variable prefix as CRAWLAB

	// replacer
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	//###
	println("==========================")
	var cfg YamlSchema
	if err := viper.ReadInConfig(); err != nil {
		println("eee==========================")
		return err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	println(cfg.Grpc.Server.Address)
	println("==========================")
	//###

	// read default config
	defaultConfBuf := bytes.NewBufferString(DefaultConfigYaml)
	if err := viper.ReadConfig(defaultConfBuf); err != nil {
		return err
	}
	var defaults map[string]interface{}
	err = yaml.Unmarshal([]byte(DefaultConfigYaml), &defaults)
	if err != nil {
		println(fmt.Sprintf("yaml eee==========================,(%v)", err))
		return err
	}
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	// merge config
	if err := viper.MergeInConfig(); err != nil { // viper parsing config file
		return err
	}

	var cfg2 YamlSchema
	if err := viper.ReadInConfig(); err != nil {
		println("eee22==========================")
		return err
	}
	if err := viper.Unmarshal(&cfg2); err != nil {
		return err
	}
	oo, err := yaml.Marshal(&cfg2)
	if err != nil {
		println("eee33==========================")
		return err
	}
	println(string(oo))

	return nil
}
