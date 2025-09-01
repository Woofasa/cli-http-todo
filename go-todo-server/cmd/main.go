package main

import (
	"context"
	"fmt"
	"log"
	"main/internal/app"
	"main/internal/cli"
	httpclient "main/internal/http_client"
)

func main() {
	dbPath := "./internal/repo/sqlite/tasks.db"
	app, err := app.NewApp(context.Background(), dbPath)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("choose the client version (cli | http): ")
	var version string
	fmt.Scan(&version)

	switch version{
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
