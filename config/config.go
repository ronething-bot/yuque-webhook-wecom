package config

import (
	"log"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func SetConfig(filePath string) {
	log.Printf("[config] run the env with:%s", filePath)

	Config = viper.New()
	Config.SetConfigFile(filePath)
	if err := Config.ReadInConfig(); err != nil {
		log.Fatalf("[config] read config err: %v", err)
	}

	log.Printf("config is %v\n", Config)

	watchFileConfig(filePath)
}

func watchFileConfig(filepath string) {
	Config.WatchConfig()
	Config.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file change: %v %v\n", e.Name, e.Op)
		if e.Op == fsnotify.Write {
			// 重新读取
			if err := Config.ReadInConfig(); err != nil {
				log.Fatalf("[config] read config err: %v", err)
			}
			log.Printf("after change, config is %v\n", Config)
		}
	})
}
