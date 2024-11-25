// routes/user_routes.go
package routes

import (
    "backendMedicalRecord/controllers"
    "github.com/gorilla/mux"
)


func DiagnosisRoutes(router *mux.Router) {
    router.HandleFunc("", controllers.CreateDiagnosis).Methods("POST")       
    router.HandleFunc("", controllers.GetAllDiagnosiss).Methods("GET")       
    router.HandleFunc("/{id}", controllers.GetDiagnosis).Methods("GET")      
    router.HandleFunc("/{id}", controllers.UpdateDiagnosis).Methods("PUT")  
    router.HandleFunc("/{id}", controllers.DeleteDiagnosis).Methods("DELETE")  
}
