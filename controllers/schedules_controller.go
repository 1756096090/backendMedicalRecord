package controllers

import (
	"backendMedicalRecord/models"
	"backendMedicalRecord/repository"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
    var schedule models.Schedule
    if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
        log.Printf("Error decoding schedule", err )
        http.Error(w, "Invalid request payload schedule", http.StatusBadRequest)
        return
    }
    
    result, err := repository.CreateSchedule(schedule)
    if err != nil {
        http.Error(w, "Error creating schedule", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    scheduleID := vars["id"]

    schedule, err := repository.GetScheduleByID(scheduleID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if schedule == nil {
        http.Error(w, "Schedule not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(schedule)
}

func UpdateSchedule(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    scheduleID := vars["id"]

    var schedule models.Schedule
    err := json.NewDecoder(r.Body).Decode(&schedule)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    result, err := repository.UpdateSchedule(scheduleID, schedule)
    if err != nil {
        
        http.Error(w, "Error updating schedule", http.StatusInternalServerError)
        return
    }

    if result.ModifiedCount == 0 {
        http.Error(w, "No hay cambios", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(result)
}

func DeleteSchedule(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    scheduleID := vars["id"]
    
    result, err := repository.DeleteSchedule(scheduleID)
    if err != nil {
        http.Error(w, "Error deleting schedule", http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "Schedule not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) 
}

func GetAllSchedules(w http.ResponseWriter, r *http.Request) {
    schedules, err := repository.GetAllSchedules()
    if err != nil {
        http.Error(w, "Error getting schedules", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(schedules)
}


func GetShedulesByMonthYear(w http.ResponseWriter, r *http.Request){

    var requestData struct {
        Month int `json:"month"`
        Year int `json:"year"`
    }

    if err:= json.NewDecoder(r.Body).Decode(&requestData); err != nil{
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if requestData.Month < 1 || requestData.Month > 12 {
        http.Error(w, "Invalid month. Must be between 1 and 12.", http.StatusBadRequest)
        return
    }

    if requestData.Year < 0 {
        http.Error(w, "Invalid year. Must be a positive integer.", http.StatusBadRequest)
        return
    }
    

    schedules, err := repository.GetShedulesByMonthYear(requestData.Month, requestData.Year)

    if err!= nil{
        http.Error(w, "Error getting schedules", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(schedules)
}

func GetSchedulesByUser(w http.ResponseWriter, r *http.Request){
    var requestData struct {
        IDUser string `json:"IDUser"`
    }

    if err:= json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Error decoding JSON ", http.StatusInternalServerError)	
    }

    schedules, err := repository.GetSchedulesByUserAndDate(requestData.IDUser)
    if err!= nil {
        http.Error(w, "Error getting schedules", http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(schedules)
}
