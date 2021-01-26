package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-sfomuseum-mapshaper"
	"github.com/sfomuseum/go-sfomuseum-mapshaper/api"
	"log"
	"net/http"
)

func main() {

	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")
	mapshaper_path := flag.String("mapshaper-path", "/usr/local/bin/mapshaper", "...")

	ctx := context.Background()

	ms, err := mapshaper.NewMapshaper(ctx, *mapshaper_path)

	if err != nil {
		log.Fatalf("Failed to create new mapshaper for '%s', %v", *mapshaper_path, err)
	}

	mux := http.NewServeMux()

	handler, err := api.InnerPointHandler(ms)

	if err != nil {
		log.Fatalf("Failed to create inner point handler, %v", err)
	}

	mux.Handle("/api/innerpoint", handler)

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create server for '%s', %v", *server_uri, err)
	}

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server for '%s', %v", *server_uri, err)
	}

}
