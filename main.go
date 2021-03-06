package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/crackeer/goweb/container"
	"github.com/crackeer/goweb/server"
)

var configPath string

func main() {
	fmt.Println(os.Args)

	flag.StringVar(&configPath, "c", "./config/app.yaml", "config file")
	flag.Parse()

	if len(configPath) < 1 {
		panic("please set config path")
	}

	container.Init(configPath)

	errChan := make(chan error)

	go func() {
		err := server.Run(container.GetAppConfig())
		errChan <- err
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, os.Interrupt)

	select {
	case err := <-errChan:
		fmt.Printf("encounter error when starting server with [%s]\n", err.Error())
	case <-signalChan:
		fmt.Printf("received signal to shutdown, process will exit\n")
	}
}
