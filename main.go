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

	//actions, err := client.GetCustomActionList("gops-swagger-ui-blue", "Deployment", "swagger-ui", "gops")
	//if err != nil {
	//	panic(err)
	//}

	//err = client.RunCustomAction("gops-swagger-ui-blue", "Deployment", "swagger-ui", "gops", "restart")
	//if err != nil {
	//	panic(err)
	//}

	err = client.SyncApplication("gops-swagger-ui-blue", true)
	if err != nil {
		panic(err)
	}
}
