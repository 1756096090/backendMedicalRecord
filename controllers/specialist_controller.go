package controllers

import (
    "encoding/json"
    "net/http"
    "backendMedicalRecord/models"
    "backendMedicalRecord/repository"
	"github.com/gorilla/mux"
)

func CreateSpecialist(w http.ResponseWriter, r *http.Request) {
    var specialist models.Specialist
    if err := json.NewDecoder(r.Body).Decode(&specialist); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    
    result, err := repository.CreateSpecialist(specialist)
    if err != nil {
        http.Error(w, "Error creating specialist", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}

func GetSpecialist(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    specialistID := vars["id"]

    specialist, err := repository.GetSpecialistByID(specialistID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if specialist == nil {
        http.Error(w, "Specialist not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(specialist)
}

func UpdateSpecialist(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    specialistID := vars["id"]

    var specialist models.Specialist
    err := json.NewDecoder(r.Body).Decode(&specialist)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    result, err := repository.UpdateSpecialist(specialistID, specialist)
    if err != nil {
        http.Error(w, "Error updating specialist", http.StatusInternalServerError)
        return
    }

    if result.ModifiedCount == 0 {
        http.Error(w, "Specialist not found or no changes made", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(result)
}

func DeleteSpecialist(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    specialistID := vars["id"]
    
    result, err := repository.DeleteSpecialist(specialistID)
    if err != nil {
        http.Error(w, "Error deleting specialist", http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "Specialist not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) 
}

func GetAllSpecialists(w http.ResponseWriter, r *http.Request) {
    specialists, err := repository.GetAllSpecialists()
    if err != nil {
        http.Error(w, "Error getting specialists", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(specialists)
}
