// main.go
package main

import (
    "log"
    "net/http"
    "backendMedicalRecord/config"
    "backendMedicalRecord/routes"
    "github.com/gorilla/handlers"
)

func main() {
    config.ConnectDB()

    router := routes.SetupRoutes()

    
    corsOptions := handlers.AllowedOrigins([]string{"*"}) 
    corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}) 
    corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}) 

    log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsMethods, corsHeaders)(router)))
}
