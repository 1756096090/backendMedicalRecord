// routes/user_routes.go
package routes

import (
    "backendMedicalRecord/controllers"
    "github.com/gorilla/mux"
)


func RoleRoutes(router *mux.Router) {
    router.HandleFunc("", controllers.CreateRole).Methods("POST")       
    router.HandleFunc("", controllers.GetAllRoles).Methods("GET")       
    router.HandleFunc("/{id}", controllers.GetRole).Methods("GET")      
    router.HandleFunc("/{id}", controllers.UpdateRole).Methods("PUT")  
    router.HandleFunc("/{id}", controllers.DeleteRole).Methods("DELETE")  
}
