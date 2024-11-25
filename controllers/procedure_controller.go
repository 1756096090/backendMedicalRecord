package controllers

import (
    "encoding/json"
    "net/http"
    "backendMedicalRecord/models"
    "backendMedicalRecord/repository"
	"github.com/gorilla/mux"
)

func CreateProcedure(w http.ResponseWriter, r *http.Request) {
    var procedure models.Procedure
    if err := json.NewDecoder(r.Body).Decode(&procedure); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    
    result, err := repository.CreateProcedure(procedure)
    if err != nil {
        http.Error(w, "Error creating procedure", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}

func GetProcedure(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    procedureID := vars["id"]

    procedure, err := repository.GetProcedureByID(procedureID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if procedure == nil {
        http.Error(w, "Procedure not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(procedure)
}

func UpdateProcedure(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    procedureID := vars["id"]

    var procedure models.Procedure
    err := json.NewDecoder(r.Body).Decode(&procedure)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    result, err := repository.UpdateProcedure(procedureID, procedure)
    if err != nil {
        http.Error(w, "Error updating procedure", http.StatusInternalServerError)
        return
    }

    if result.ModifiedCount == 0 {
        http.Error(w, "No hay cambios", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(result)
}

func DeleteProcedure(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    procedureID := vars["id"]
    
    result, err := repository.DeleteProcedure(procedureID)
    if err != nil {
        http.Error(w, "Error deleting procedure", http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "Procedure not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) 
}

func GetAllProcedures(w http.ResponseWriter, r *http.Request) {
    procedures, err := repository.GetAllProcedures()
    if err != nil {
        http.Error(w, "Error getting procedures", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(procedures)
}


