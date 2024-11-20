// routes/user_routes.go
package routes

import (
    "backendMedicalRecord/controllers"
    "github.com/gorilla/mux"
)


func UserRoutes(router *mux.Router) {
    router.HandleFunc("", controllers.CreateUser).Methods("POST")       
    router.HandleFunc("", controllers.GetAllUsers).Methods("GET")       
    router.HandleFunc("/doctors", controllers.GetDoctors).Methods("GET")       
    router.HandleFunc("/{id}", controllers.GetUser).Methods("GET")      
    router.HandleFunc("/all-info/{id}", controllers.GetUserWithRoleAndSpecialty).Methods("GET")      
    router.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")  
    router.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")  
    router.HandleFunc("/login", controllers.Login).Methods("POST") 
}
