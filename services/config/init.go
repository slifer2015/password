package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Initialize init config
func Initialize(prefix string) {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	viper.SetConfigName(prefix + "_config")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Debug("Config file changed:", e.Name)
	})
}
