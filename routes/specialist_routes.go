// routes/user_routes.go
package routes

import (
    "backendMedicalRecord/controllers"
    "github.com/gorilla/mux"
)


func SpecialistRoutes(router *mux.Router) {
    router.HandleFunc("", controllers.CreateSpecialist).Methods("POST")       
    router.HandleFunc("", controllers.GetAllSpecialists).Methods("GET")       
    router.HandleFunc("/{id}", controllers.GetSpecialist).Methods("GET")      
    router.HandleFunc("/{id}", controllers.UpdateSpecialist).Methods("PUT")  
    router.HandleFunc("/{id}", controllers.DeleteSpecialist).Methods("DELETE")  
}
