package main

import (
	"fmt"
	// "log"
	"net/http"
	"os"

	"github.com/zeemzo/jigsaw-gateway/api/routes"
	"github.com/zeemzo/jigsaw-gateway/services"
	"github.com/gorilla/handlers"
	// "github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8000"
}


func main() {

	// getEnvironment()

	port := getPort()
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	c := cron.New()
	c.AddFunc("@every 30m", func() {
		services.CheckCOCStatus()
	})
	c.Start()

	router := routes.NewRouter()
	fmt.Println("Jigsaw Gateway Started @port " + port)
	http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router))

}