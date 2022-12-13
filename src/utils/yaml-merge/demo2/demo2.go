package demo2

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes: brown
beard: true
`)

func readAndMergeConfig() {
	var viperIn = viper.New()
	viperIn.SetConfigType("yaml")
	/*err := viperIn.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		println(fmt.Sprintf("=================,err:(%v)", err))
		return
	}*/
	err := viperIn.ReadConfig(bytes.NewBufferString(DefaultConfigYaml))
	if err != nil {
		println(fmt.Sprintf("=================,err:(%v)", err))
		return
	}
	println("=========22========")
	name := viperIn.Get("server.port")
	name = viperIn.Get("spider.fs")
	/*switch name.(type) {
	case string:
		println(cast.ToString(name))
		return
	case int:
		println(cast.ToInt(name))
		return
	}*/
	fmt.Println(reflect.TypeOf(name))
	//name = viperIn.GetString("name")
	//println(name)
	println("=========22========")
}
