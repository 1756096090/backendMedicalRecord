// routes/user_routes.go
package routes

import (
    "backendMedicalRecord/controllers"
    "github.com/gorilla/mux"
)


func ScheduleRoutes(router *mux.Router) {
    router.HandleFunc("", controllers.CreateSchedule).Methods("POST")       
    router.HandleFunc("", controllers.GetAllSchedules).Methods("GET")       
    router.HandleFunc("/{id}", controllers.GetSchedule).Methods("GET")      
    router.HandleFunc("/{id}", controllers.UpdateSchedule).Methods("PUT")  
    router.HandleFunc("/{id}", controllers.DeleteSchedule).Methods("DELETE")  
    router.HandleFunc("/by-month-year", controllers.GetShedulesByMonthYear).Methods("POST")
    router.HandleFunc("/by-user", controllers.GetSchedulesByUser).Methods("POST")

}