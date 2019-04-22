package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sankalpjonn/go-timeloop"
)

func main() {
	tl := timeloop.New()

	tl.Job(func() {
		log.Println("printing Test2SecTimer 1")
	}, time.Second*1)

	tl.Job(func() {
		log.Println("printing Test5SecTimer 2 ")
	}, time.Second*1)

	tl.Job(func() {
		log.Println("printing Test10SecTimer 3")
	}, time.Second*1)

	tl.Start()
	defer tl.Stop()

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	signal.Notify(done, os.Interrupt, syscall.SIGINT)
	<-done
}
