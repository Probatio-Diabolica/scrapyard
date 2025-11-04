package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"scrapyard/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Missing port")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalln("Missing port")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln("can't connect to the db")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	//not oke because it handles every req
	// v1Router.HandleFunc("/ready", handleReadiness)
	/////////////////////////////////////////////////////
	v1Router.Get("/ready", handleReadiness)
	v1Router.Get("/err", handleErr)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.handleGetUserByAPIKey)
	router.Mount("/v1", v1Router)
	////////////////////////////////////////////////////
	httpServer := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("ok this seems to be working")

	err = httpServer.ListenAndServe()

	if err != nil {
		log.Fatalln("can't connect to the port")
	}
	fmt.Println(" port => ", port)
}
