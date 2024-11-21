package controllers

import (
    "encoding/json"
    "net/http"
    "backendMedicalRecord/models"
    "backendMedicalRecord/repository"
	"github.com/gorilla/mux"
)

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
    var schedule models.Schedule
    if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
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
