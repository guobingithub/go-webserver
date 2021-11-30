package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func ReadYamlFile(filePath string, out interface{}) (err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(file, out); err != nil {
		return
	}
	return
}

func ReadJSONFile(filePath string, out interface{}) (err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	if err = json.Unmarshal(file, out); err != nil {
		return
	}
	return
}

func ReadConfig(filePath string, out interface{}, defaults map[string]interface{}) error {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigFile(filePath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&out); err != nil {
		return err
	}

	return nil
}
