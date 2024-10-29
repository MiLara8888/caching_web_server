package documentservice

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MiLara8888/caching_web_server/pkg/settings"
	"github.com/MiLara8888/caching_web_server/pkg/storage"
	"github.com/MiLara8888/caching_web_server/pkg/storage/document_db/postgres"
)

var wait time.Duration

type Rest struct {
	Routes *gin.Engine

	Config *settings.Config

	DB storage.IDocumentDB

	ctx context.Context

	errChan chan error

	// токен администратора Фиксированный, задается в конфиге приложения
	TokenAdmin string
}

func New(c *settings.Config) (*Rest, error) {

	db, err := postgres.New(c)
	if err != nil {
		return nil, err
	}

	rest := &Rest{
		Routes:     gin.Default(),
		Config:     c,
		DB:         db,
		TokenAdmin: c.TokenAdmin,
	}

	rest.initializeRoutes()
	return rest, err
}

// общий слушатель ошибок для потоков
func (ms *Rest) errorListener(ctx context.Context) chan error {
	var (
		wg  sync.WaitGroup
		out = make(chan error)
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-out:
				if !ok {
					return
				}
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func (s *Rest) Start() error {

	connWs := net.JoinHostPort(s.Config.Host, s.Config.Port)
	log.Printf(`merch server start : %s`, connWs)

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	s.ctx = ctx

	go s.errorListener(ctx)

	srv := &http.Server{
		Addr:           connWs,
		WriteTimeout:   time.Second * 15,
		ReadTimeout:    time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
		Handler:        s.Routes,
		// BaseContext: func(l net.Listener) context.Context {
		// 	return s.ctx
		// },
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// https://ru.wikipedia.org/wiki/%D0%A1%D0%B8%D0%B3%D0%BD%D0%B0%D0%BB_(Unix)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)

	<-c

	s.DB.Close(ctx)
	srv.Shutdown(ctx)

	log.Println("shutting down")

	return nil
}
