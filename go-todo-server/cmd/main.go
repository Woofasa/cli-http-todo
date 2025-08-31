package main

import (
	"context"
	"log"
	"main/internal/app"
	"main/internal/cli"
)

func main() {
	dbPath := "./internal/repo/sqlite/tasks.db"
	app, err := app.NewApp(context.Background(), dbPath)
	if err != nil{
		log.Fatal(err)
	}
			
	if err := cli.Run(app); err != nil {
		log.Fatal("fatal error: ", err)
	}
}
