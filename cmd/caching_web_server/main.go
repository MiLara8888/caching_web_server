package main


import (
	"errors"
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"

	doc "github.com/MiLara8888/caching_web_server/internal/document_service"
	settings "github.com/MiLara8888/caching_web_server/pkg/settings"
)

var (
	config *settings.Config
)

func init() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	config, err = settings.InitEnv()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := &cli.App{
		Name:  "client service",
		Usage: "",
		Commands: []cli.Command{
			{
				Name:  "serve",
				Usage: "Start service",
				Action: func(*cli.Context) error {
					return ItemsApiRun()
				},
			},
		},
	}
	app.Run(os.Args)
}

func ItemsApiRun() error {
	service, err := doc.New(config)
	if err != nil {
		log.Fatal(err)
	}
	err = service.Start()
	if err != nil {
		log.Fatal(errors.Join(errors.New("err on start service"), err))
	}
	return nil
}
