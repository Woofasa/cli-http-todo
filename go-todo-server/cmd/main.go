package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/cli"
	httpclient "main/internal/http_client"
	app "main/internal/usecase"
)

func main() {
	app, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("choose the client version (cli | http): ")
	var version string
	fmt.Scan(&version)

	switch version {
	case "cli":
		if err := cli.Run(app); err != nil {
			log.Fatal("fatal error: ", err)
		}
	case "http":
		handler := &httpclient.Handler{App: app}
		httpclient.RunServer(handler, ":3001")
	default:
		fmt.Println("Unknown client name.")
	}
}
