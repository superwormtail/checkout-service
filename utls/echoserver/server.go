package echoserver

import (
	"checkout-service/config"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	server *echo.Echo
	mutex  sync.Once
)

func GetServer() *echo.Echo {
	mutex.Do(func() {
		server = newServer()
	})

	return server
}

func newServer() *echo.Echo {
	return echo.New()
}

func InitServer() {

	// Hide banner
	GetServer().HideBanner = true

	// Set debug status parameter
	//GetServer().Debug = config.GetConfig().AppDebug

	// healthCheck endpoint
	GetServer().GET("/infrastructure/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	// custom error handler
	GetServer().HTTPErrorHandler = Handler

	// Custom binder
	//GetServer().Binder = &CustomBinder{bind: &echo.DefaultBinder{}}
}

func StartServer() {
	if err := GetServer().StartServer(&http.Server{
		Addr: fmt.Sprintf(":%v", config.Env.App.HTTPPort),
	}); err != nil {
		log.Fatal(err.Error())
	}
}

func Shutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := GetServer().Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
