package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/shomuMatch/goCoverTest/api/path1"
	"github.com/shomuMatch/goCoverTest/api/path2"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	})
	http.HandleFunc("/api1", path1.Api1)
	http.HandleFunc("/api2", path2.Api2)

	server := &http.Server{
		Addr:    ":8888",
		Handler: nil,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	fmt.Println("start receiving at :8888")
	log.Fatal(server.ListenAndServe())
}
