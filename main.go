package main

import "log"

func main() {
	store, err := newPGStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil{
		log.Fatal(err)
	}
	server := NewAPIServer(":3010",store)
	server.Run()
}