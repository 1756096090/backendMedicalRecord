// routes/user_routes.go
package routes

import (
    "backendMedicalRecord/controllers"
    "github.com/gorilla/mux"
)


func DiagnosisProcedureRoutes(router *mux.Router) {
    router.HandleFunc("", controllers.CreateDiagnosisProcedure).Methods("POST")       
    router.HandleFunc("", controllers.GetAllDiagnosisProcedures).Methods("GET")       
    router.HandleFunc("/{id}", controllers.GetDiagnosisProcedure).Methods("GET")      
    router.HandleFunc("/{id}", controllers.UpdateDiagnosisProcedure).Methods("PUT")  
    router.HandleFunc("/{id}", controllers.DeleteDiagnosisProcedure).Methods("DELETE")  
    router.HandleFunc("/by-patient", controllers.GetAllDiagnosisProceduresByID).Methods("POST") 	  
    router.HandleFunc("/report", controllers.GenerateReportsOfProcedures).Methods("POST") 	  
}
