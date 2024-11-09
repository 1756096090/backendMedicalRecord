package controllers

import (
	"backendMedicalRecord/models"
	"backendMedicalRecord/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePatient(w http.ResponseWriter, r *http.Request) {
    var patient models.Patient
    log.Printf(patient.DNI, "Creating Patient")
    if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    existingPatient, err := repository.GetPatientByDNIOrEmail(patient.DNI, patient.Mail)
    if err != nil {
        http.Error(w, "El dni o el mailse encuentra duplicado", http.StatusInternalServerError)
        return
    }
    if existingPatient != nil {
        http.Error(w, "DNI already exists", http.StatusConflict)
        return
    }

    result, err := repository.CreatePatient(patient)
    if err != nil {
        http.Error(w, "Error creating patient", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}
func GetPatient(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    patientID := vars["id"]

    patient, err := repository.GetPatientByID(patientID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if patient == nil {
        http.Error(w, "Patient not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(patient)
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    patientID := vars["id"]

    var patient models.Patient

    err := json.NewDecoder(r.Body).Decode(&patient)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    existingPatient, err := repository.GetPatientByDNIOrEmail(patient.DNI, patient.Mail)
    if err != nil {
        http.Error(w, "El dni o le mail se encuentra duplicado", http.StatusInternalServerError)
        return
    }
    if existingPatient != nil && existingPatient.ID.Hex() != patientID {
        http.Error(w, "No se encuentra el paciente ", http.StatusConflict)
        return
    }

    result, err := repository.UpdatePatient(patientID, patient)
    if err != nil {
        http.Error(w, "Error updating patient", http.StatusInternalServerError)
        return
    }

    if result.ModifiedCount == 0 {
        http.Error(w, "No hay cambios", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(result)
}


func DeletePatient(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    patientID := vars["id"]
    
    result, err := repository.DeletePatient(patientID)
    if err != nil {
        http.Error(w, "Error deleting patient", http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "Patient not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) 
}

func GetAllPatients(w http.ResponseWriter, r *http.Request) {
    patients, err := repository.GetAllPatients()
    if err != nil {
        http.Error(w, "Error getting patients", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(patients)
}
