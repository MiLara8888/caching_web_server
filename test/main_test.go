package test

import (
	doc "github.com/MiLara8888/caching_web_server/internal/document_service"
	"log"
	"runtime"
	"testing"

	"github.com/MiLara8888/caching_web_server/pkg/settings"
)

var (
	err error
	// настройка подключения
	config *settings.Config
)

func TestMain(m *testing.M) {
	config, err = settings.InitEnv()
	if err != nil {
		log.Fatal(err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	m.Run()
}


func TestDocRun(t *testing.T) {

	service, err := doc.New(config)
	if err != nil {
		t.Fatal(err)
	}
	err = service.Start()
	if err != nil {
		t.Fatal(err)
	}
}
