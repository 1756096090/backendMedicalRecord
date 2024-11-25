package controllers

import (
    "encoding/json"
    "net/http"
    "backendMedicalRecord/models"
    "backendMedicalRecord/repository"
	"github.com/gorilla/mux"
)

func CreateDiagnosis(w http.ResponseWriter, r *http.Request) {
    var diagnosis models.Diagnosis
    if err := json.NewDecoder(r.Body).Decode(&diagnosis); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    
    result, err := repository.CreateDiagnosis(diagnosis)
    if err != nil {
        http.Error(w, "Error creating diagnosis", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}

func GetDiagnosis(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    diagnosisID := vars["id"]

    diagnosis, err := repository.GetDiagnosisByID(diagnosisID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if diagnosis == nil {
        http.Error(w, "Diagnosis not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(diagnosis)
}

func UpdateDiagnosis(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    diagnosisID := vars["id"]

    var diagnosis models.Diagnosis
    err := json.NewDecoder(r.Body).Decode(&diagnosis)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    result, err := repository.UpdateDiagnosis(diagnosisID, diagnosis)
    if err != nil {
        http.Error(w, "Error updating diagnosis", http.StatusInternalServerError)
        return
    }

    if result.ModifiedCount == 0 {
        http.Error(w, "No hay cambios", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(result)
}

func DeleteDiagnosis(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    diagnosisID := vars["id"]
    
    result, err := repository.DeleteDiagnosis(diagnosisID)
    if err != nil {
        http.Error(w, "Error deleting diagnosis", http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "Diagnosis not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) 
}

func GetAllDiagnosiss(w http.ResponseWriter, r *http.Request) {
    diagnosiss, err := repository.GetAllDiagnosiss()
    if err != nil {
        http.Error(w, "Error getting diagnosiss", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(diagnosiss)
}


