package main

import (
	"embed"
	"errors"
	"fmt"
	"github.com/coreos/go-systemd/v22/sdjournal"
	"github.com/taranovegor/jurnalo/internal/config"
	"github.com/taranovegor/jurnalo/internal/container"
	"github.com/taranovegor/jurnalo/internal/handler/journal"
	"io/fs"
	"log"
	"net/http"
	"time"
)

var version = "unknown"

//go:embed web/dist
var app embed.FS

func main() {
	fmt.Println(fmt.Sprintf("ĵurnalo! version: %s", version))

	ctn, err := container.Init()
	if err != nil {
		log.Fatalf("Error during container initialization: %s", err)
	}

	reader := ctn.Get(container.JournalReader).(*sdjournal.JournalReader)
	journalHandler := ctn.Get(container.HandlerJournal).(*journal.Handler)

	go func() {
		err := reader.Follow(time.After(1<<63-1), journalHandler)
		if !errors.Is(err, sdjournal.ErrExpired) {
			log.Fatalf("Error during follow: %s", err)
		}
	}()

	dist, err := fs.Sub(app, "web/dist")
	if err != nil {
		log.Fatalf("Error during sub: %s", err)
	}

	router := ctn.Get(container.HttpRouter).(*http.ServeMux)
	router.Handle("/", http.FileServer(http.FS(dist)))

	log.Fatal(http.ListenAndServe(config.Get().Http.Addr(), router))
}
