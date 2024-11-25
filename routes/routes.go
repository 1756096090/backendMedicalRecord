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

    scheduleRouter :=  router.PathPrefix("/schedule").Subrouter()
    ScheduleRoutes(scheduleRouter)

    diagnosisRouter :=  router.PathPrefix("/diagnosis").Subrouter()
    DiagnosisRoutes(diagnosisRouter)

    patientRoutes := router.PathPrefix("/patient").Subrouter()
    PatientRoutes(patientRoutes)

    procedureRoutes := router.PathPrefix("/procedure").Subrouter()
    ProcedureRoutes(procedureRoutes)

    diagnosisProcedureRoutes := router.PathPrefix("/diagnosisProcedure").Subrouter()
    DiagnosisProcedureRoutes(diagnosisProcedureRoutes)

    return router
}
