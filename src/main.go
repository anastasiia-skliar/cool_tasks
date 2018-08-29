package main

import (
	"flag"
	"github.com/Nastya-Kruglikova/cool_tasks/src/config"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/auth"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service"
)

func main() {
	configFile := flag.String("config", "/opt/cool_tasks/config/config.json", "Configuration file in JSON-format")
	flag.Parse()

	if len(*configFile) > 0 {
		config.FilePath = *configFile
	}

	err := config.Load()
	if err != nil {
		log.Fatalf("error while reading config: %s", err)
	}

	f, err := os.OpenFile(config.Config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	database.DB, err = database.SetupPostgres(config.Config.Database)
	if err != nil {
		log.Fatalf("error while loading postgreSQL: %s:", err)
	}

	database.Cache, err = database.SetupRedis(config.Config.Database)

	if err != nil {
		log.Fatalf("error while loading redis: %s:", err)
	}

	defer f.Close()

	log.SetOutput(f)

	// setting up web server middlewares
	middlewareManager := negroni.New(
		negroni.HandlerFunc(auth.IsAuthorized),
		negroni.HandlerFunc(auth.AccessPermission),
	)
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.UseHandler(service.NewRouter())

	log.Println("Starting HTTP listener...")
	err = http.ListenAndServe(getPort(), middlewareManager)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Stop running application: %s", err)
}

//get heroku port
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = config.Config.ListenURL
		log.Printf("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}