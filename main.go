package main

import (
	"context"
	"os"

	"github.com/ProxySafe/site-backend/src/app"
	"github.com/ProxySafe/site-backend/src/app/config"
	"github.com/cosiner/flag"
	"gopkg.in/yaml.v2"
)

type Params struct {
	ConfigFile    string `names:"--config_dir, -c" usage:"configuration directory" default:"./config/config.yaml"`
	LogsDir       string `names:"--logs_dir, -l" usage:"log directory" default:"./var/log/"`
	WebServerPort int    `names:"--web_server_port, -wp" usage:"port if you need to run web server"`
}

func getParams() *Params {
	params := &Params{}
	err := flag.Commandline.ParseStruct(params)
	if err != nil {
		panic(err)
	}
	return params
}

func getConfig(params *Params) *config.Config {
	content, err := os.ReadFile(params.ConfigFile)
	if err != nil {
		panic(err)
	}

	c := &config.Config{}
	if err := yaml.Unmarshal(content, c); err != nil {
		panic(err)
	}
	return c
}

func main() {
	ctx := context.Background()
	p := getParams()
	c := getConfig(p)
	a := app.NewApp(c)
	a.Init(ctx)
	a.Run(p.WebServerPort)
}
