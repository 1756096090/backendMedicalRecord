package routes

import (
    "github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    
    userRouter := router.PathPrefix("/user").Subrouter() 
    UserRoutes(userRouter)

    roleRouter := router.PathPrefix("/role").Subrouter()
    RoleRoutes(roleRouter)

    specialistRouter :=  router.PathPrefix("/specialist").Subrouter()
    SpecialistRoutes(specialistRouter)

    patientRoutes := router.PathPrefix("/patient").Subrouter()
    PatientRoutes(patientRoutes)

    return router
}
