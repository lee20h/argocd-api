package main

import (
	"argocd-api/argocd"
	"argocd-api/config"
)

func main() {
	cfg := config.Init()

	connection := argocd.Connection{
		Address: cfg.Address,
		Token:   cfg.Token,
	}

	client, err := argocd.NewClient(&connection)
	if err != nil {
		panic(err)
	}

	err = client.SyncApplication("default", true)
	if err != nil {
		panic(err)
	}
}
