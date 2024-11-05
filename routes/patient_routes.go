// routes/user_routes.go
package routes

import (
    "backendMedicalRecord/controllers"
    "github.com/gorilla/mux"
)


func PatientRoutes(router *mux.Router) {
    router.HandleFunc("", controllers.CreatePatient).Methods("POST")       
    router.HandleFunc("", controllers.GetAllPatients).Methods("GET")       
    router.HandleFunc("/{id}", controllers.GetPatient).Methods("GET")      
    router.HandleFunc("/{id}", controllers.UpdatePatient).Methods("PUT")  
    router.HandleFunc("/{id}", controllers.DeletePatient).Methods("DELETE")  
}
