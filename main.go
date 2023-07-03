package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var (
	listenAddr string
	healthy    int32
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] : .env file found")
	}
}

func main() {
	value, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		log.Printf("[INFO] : Value not found : %v", ok)
	} else {
		log.Printf("[INFO] : Value found : %s", value)
	}

	flag.StringVar(&listenAddr, "port", value, "server listen address")
	flag.Parse()

	addr := fmt.Sprintf("0.0.0.0:%s", listenAddr)

	server := &http.Server{
		Addr:    addr,
		Handler: service(),
	}

	log.Println("[INFO] : Server is Starting...")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("[ERROR] : starting server: %s", err.Error())
		}
	}()

	log.Printf("[INFO] : Server is Ready to Handle Request at %s", listenAddr)

	atomic.StoreInt32(&healthy, 1)

	wait := gracefulShutdown(context.Background(), 5*time.Second, map[string]operation{
		"server": func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	<-wait
}
