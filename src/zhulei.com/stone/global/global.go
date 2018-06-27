package global

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func init() {
	loadConfigs()
}

var Configs GlobalConfigs

type GlobalConfigs struct {
	Db *DBConfig `json:"db"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func loadConfigs() error {
	data, err := ioutil.ReadFile("./config/configs.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	jerr := json.Unmarshal(data, &Configs)
	if jerr != nil {
		fmt.Println(err.Error())
		return jerr
	}
	return nil
}
