package demo2

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"reflect"
)

func MergeConfigIntoDefault() {
	var viperIn = viper.New()
	viperIn.SetConfigType("yaml")

	var defaults map[string]interface{}
	err := yaml.Unmarshal([]byte(DefaultConfigYaml), &defaults)
	if err != nil {
		println(fmt.Sprintf("yaml eee==========================,(%v)", err))
		return
	}
	for key, value := range defaults {
		viperIn.SetDefault(key, value)
	}

	//
	name := viperIn.GetString("server.port")
	fmt.Println(reflect.TypeOf(name))
	println(name)

	println("=========11========")
	var cfg1 YamlSchema
	if err := viperIn.Unmarshal(&cfg1); err != nil {
		return
	}
	oo1, err := yaml.Marshal(&cfg1)
	if err != nil {
		println("eee33==========================")
		return
	}
	println(string(oo1))

	// oneConfigYaml
	err = viperIn.ReadConfig(bytes.NewBufferString(oneConfigYaml))
	if err != nil {
		println(fmt.Sprintf("=================,err:(%v)", err))
		return
	}
	println("=========22========")
	name = viperIn.GetString("server.port")
	fmt.Println(reflect.TypeOf(name))
	println(name)

	println("=========33========")
	var cfg2 YamlSchema
	if err := viperIn.Unmarshal(&cfg2); err != nil {
		return
	}
	oo2, err := yaml.Marshal(&cfg2)
	if err != nil {
		println("eee33==========================")
		return
	}
	println(string(oo2))

	println("=========end!!!========")
	// merge config
	if err := viperIn.MergeInConfig(); err != nil { // viper parsing config file
		println(fmt.Sprintf("eee00000001==========================(%v)", err))
		return
	}
	var cfg3 YamlSchema
	if err := viperIn.Unmarshal(&cfg3); err != nil {
		return
	}
	oo3, err := yaml.Marshal(&cfg3)
	if err != nil {
		println("eee33==========================")
		return
	}
	println(string(oo3))
}
