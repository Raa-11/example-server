package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("[INFO] : Shutting Down")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("[ERROR] : Timeout %d ms has been Elapsed, Force Exit", timeout.Microseconds())
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("[INFO] : Cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("[ERROR : %s: Clean up Failed: %s", innerKey, err.Error())
					return
				}
				log.Printf("[INFO] : %s was Shutdown Gracefully", innerKey)
			}()
		}
		wg.Wait()
		close(wait)
	}()
	return wait
}
