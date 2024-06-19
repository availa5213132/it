package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb/gvb_server/config"
	"gvb/gvb_server/global"
	"io/ioutil"
	"log"
)

// initConf 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal: %v", err)
	}
	log.Println("config yamlFile load init success.")
	global.Config = c
}
