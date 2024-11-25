// routes/user_routes.go
package routes

import (
    "backendMedicalRecord/controllers"
    "github.com/gorilla/mux"
)


func ProcedureRoutes(router *mux.Router) {
    router.HandleFunc("", controllers.CreateProcedure).Methods("POST")       
    router.HandleFunc("", controllers.GetAllProcedures).Methods("GET")       
    router.HandleFunc("/{id}", controllers.GetProcedure).Methods("GET")      
    router.HandleFunc("/{id}", controllers.UpdateProcedure).Methods("PUT")  
    router.HandleFunc("/{id}", controllers.DeleteProcedure).Methods("DELETE")  
}
