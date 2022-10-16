package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/drone/envsubst/v2"

	"shopping-cart/common/global"
)

const (
	defaultLogDir = "log"

	devConfigFile = "config/app_dev.json"
	configFileFmt = "config/app_%s.json"
)

var (
	// AppConf ...
	AppConf = &RuntimeConfig{
		Log: LogConfig{
			LogDir: defaultLogDir,
		},
	}
)

type AppConfig struct {
	Port      uint16   `json:"port"`
	Debug     bool     `json:"debug"`
	DBDebug   bool     `json:"db_debug"`
	Pprof     bool     `json:"pprof"`
	WhiteList []string `json:"white_list"`
}

// LogConfig 日志配置
type LogConfig struct {
	LogLevel   int    `json:"log_lvl,omitempty"`
	Env        string `json:"env"`
	ShowCaller bool   `json:"show_caller"`
	LogDir     string `json:"log_dir"`
	Stdout     bool   `json:"stdout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	ConnectUrls DBConnUrls `json:"connect_urls"`
	MaxIdle     int        `json:"max_idle"`
	MaxOpen     int        `json:"max_open"`
	LifeTime    int        `json:"life_time"`
}

type DBConnUrls struct {
	Master string   `json:"master"`
	Slaves []string `json:"slaves"`
}

// RuntimeConfig ...
type RuntimeConfig struct {
	App AppConfig `json:"app"`

	Log LogConfig      `json:"log"`
	DB  DatabaseConfig `json:"db"`
}

func (conf *RuntimeConfig) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}

func (conf *RuntimeConfig) JsonString() ([]byte, error) {
	b, err := json.Marshal(*conf)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func LoadConfig(config string) error {
	var filePath string
	switch {
	case os.Getenv(global.APPEnv) == global.DevEnv:
		filePath = devConfigFile
	case config != "":
		filePath = config
	default:
		filePath = fmt.Sprintf(configFileFmt, os.Getenv(global.APPEnv))
	}
	configJson, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	newJson, err := envsubst.EvalEnv(string(configJson))
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(newJson), &AppConf)
	if err != nil {
		return err
	}
	return nil
}
