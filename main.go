package main

import (
	"go_by_example/application"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &application.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":27017", server))
}
