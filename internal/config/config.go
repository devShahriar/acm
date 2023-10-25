package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	logger "magic.pathao.com/data/de-logger"
)

const provider = "consul"

var (
	m       sync.Mutex
	appConf *AppConfig
)

type AppConfig struct {
	AppKey string   `json:"app_key" mapstructure:"app_key"`
	DbConf DbConfig `json:"db_conf" mapstructure:"db_conf"`
}

type DbConfig struct {
	Host               string `json:"host" mapstructure:"host"`
	Password           string `json:"password" mapstructure:"password"`
	User               string `json:"user" mapstructure:"user"`
	DbName             string `json:"db_name" mapstructure:"db_name"`
	Port               string `json:"port" mapstructure:"port"`
	SlowQueryThreshold int    `json:"slow_query_threshold" mapstructure:"slow_query_threshold"`
}

func GetViper(confType string) *viper.Viper {

	var viperInstance *viper.Viper
	if viperInstance == nil {
		viperInstance = viper.New()

		err := viperInstance.BindEnv("CONSUL_URL")
		if err != nil {
			fmt.Println(err)
			return nil
		}
		err = viperInstance.BindEnv("CONSUL_PATH")
		if err != nil {
			fmt.Println(err)
			return nil
		}
		consulUrl := viperInstance.GetString("CONSUL_URL")
		consulPath := viperInstance.GetString("CONSUL_PATH")
		if consulUrl == "" {
			log.Println("ACM_CONSUL_URL is not provided")
		}
		if consulPath == "" {
			log.Println("ACM_CONSUL_PATH is not provided")
		}

		viperInstance.SetConfigType("json")
		err = viperInstance.AddRemoteProvider(provider, consulUrl, fmt.Sprintf("%v/%v", consulPath, confType))

		if err != nil {
			log.Println(err)
		}
		err = viperInstance.ReadRemoteConfig()
		if err != nil {
			log.Panicf("Error during reading remote configuration: %v", err)
		}
	}
	return viperInstance
}

func GetAppConfig() *AppConfig {

	if appConf == nil {
		m.Lock()
		defer m.Unlock()
		viperConf := GetViper("appconf")

		var err error
		appConf, err = LoadAppConf(viperConf)
		if err != nil {
			log.Panicf("Error during loading configuration: %v", err)
		}
	}
	return appConf
}

func LoadAppConf(viperInstance *viper.Viper) (*AppConfig, error) {
	appConfig := AppConfig{}
	err := viperInstance.Unmarshal(&appConfig)
	if err != nil {
		log.Panicf("Couldn't read the application configuration: %v", err)
		return nil, err
	}
	appConfJson, _ := json.Marshal(appConfig)
	logger.GetSimpleLogger("Service").Infof("App Config: %v", string(appConfJson))
	return &appConfig, err
}
