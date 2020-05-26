package main

import (
	"awesomeProject/gokit/users"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
	httpServer := flag.String("http", ":8080", "http addr")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
	}
	level.Info(logger).Log("msg", "Service Started")
	defer level.Info(logger).Log("msg", "Service stoped")

	ctx := context.Background()

	userService := users.NewService(users.NewRepo(logger), logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := users.CreateEndPoint(userService)

	go func() {
		fmt.Println("listening", *httpServer)
		handler := users.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpServer, handler)
	}()
	level.Error(logger).Log("exit", <-errs)
}
