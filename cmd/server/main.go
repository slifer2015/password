package main

import (
	"github.com/sirupsen/logrus"
	"test.com/test/services/config"
	"test.com/test/services/initializer"
	"test.com/test/services/signal"
)

const appName = "server"

func main() {
	config.Initialize(appName)
	defer initializer.Initialize()()
	sig := signal.WaitExitSignal()
	logrus.Infof("terminated using %s signal", sig.String())
}
