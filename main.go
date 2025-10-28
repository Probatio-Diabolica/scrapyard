package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalln("Missing port")
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
	router.Mount("/v1", v1Router)
	////////////////////////////////////////////////////
	httpServer := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("ok this seems to be working")

	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatalln("can't connect to the port")
	}
	fmt.Println(" port => ", port)
}
