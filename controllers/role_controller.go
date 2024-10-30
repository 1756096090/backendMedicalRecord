package controllers

import (
    "encoding/json"
    "net/http"
    "backendMedicalRecord/models"
    "backendMedicalRecord/repository"
	"github.com/gorilla/mux"
)

func CreateRole(w http.ResponseWriter, r *http.Request) {
    var role models.Role
    if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    
    result, err := repository.CreateRole(role)
    if err != nil {
        http.Error(w, "Error creating role", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}

func GetRole(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    roleID := vars["id"]

    role, err := repository.GetRoleByID(roleID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if role == nil {
        http.Error(w, "Role not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(role)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    roleID := vars["id"]

    var role models.Role
    err := json.NewDecoder(r.Body).Decode(&role)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    result, err := repository.UpdateRole(roleID, role)
    if err != nil {
        http.Error(w, "Error updating role", http.StatusInternalServerError)
        return
    }

    if result.ModifiedCount == 0 {
        http.Error(w, "Role not found or no changes made", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(result)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    roleID := vars["id"]
    
    result, err := repository.DeleteRole(roleID)
    if err != nil {
        http.Error(w, "Error deleting role", http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "Role not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) 
}

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
    roles, err := repository.GetAllRoles()
    if err != nil {
        http.Error(w, "Error getting roles", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(roles)
}
