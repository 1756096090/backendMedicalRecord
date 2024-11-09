package controllers

import (
	"backendMedicalRecord/models"
	"backendMedicalRecord/repository"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey = []byte("ddc30a48c376a93d918e30d69822ecc263e108cb1a0a7e783b2172427132243af1a2c29b7e2a3e384b7e702e7abae9f33247180e964445ed66433ec21557a802017cd5d5e0cbe7e1db0504251175ec3ac6fb4a6532a42f8fcdb96e5f3708ea9c13fa8b277dbd9c9e0cd3c4458273d7b39203d43d7df84fa665233864aa6a92f8f31fbf8fe871ba2986a6948b7cfcae134bbb569dc9473b0db646609fe82bbb680d68eaddae427391731f4508eea88bc8c58745b66f09da77a6d791d0e709a82fb19a6bee5385dccb9a0f648ef7062b2b27b4de3ec44f0929a916ea4fbfb5f734729b63d4d79db8e7a3d2e82f0a8bb9c26f060e460d32ba5d195961d829c95bd8")

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    existingUser, err := repository.GetUserByDNIOrEmail(user.DNI, user.Email)
    if err != nil {
        http.Error(w, "Error checking DNI", http.StatusInternalServerError)
        return
    }
    if existingUser != nil {
        http.Error(w, "El DNI o el mail ya existen", http.StatusConflict)
        return
    }

    result, err := repository.CreateUser(user)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}


func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]

    user, err := repository.GetUserByID(userID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if user == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]

    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    existingUser, err := repository.GetUserByDNIOrEmail(user.DNI, user.Email)
    if err != nil {
        http.Error(w, "Error checkin", http.StatusInternalServerError)
        return
    }
    if existingUser != nil && existingUser.ID.Hex() != userID {
        http.Error(w, "El DNI o el Mail ya existe", http.StatusConflict)
        return
    }

    result, err := repository.UpdateUser(userID, user)
    if err != nil {
        http.Error(w, "Error updating user", http.StatusInternalServerError)
        return
    }

    if result.ModifiedCount == 0 {
        http.Error(w, "No ha realizado cambios", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(result)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]
    
    result, err := repository.DeleteUser(userID)
    if err != nil {
        http.Error(w, "Error deleting user", http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) 
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := repository.GetAllUsers()
    if err != nil {
        http.Error(w, "Error getting users", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("User login attempt: %s", user.Email)

	storedUser, err := repository.Login(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID: storedUser.ID, 
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token":   tokenString,
		"message": "Login successful",
		"user":    storedUser, 
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


type Claims struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	jwt.RegisteredClaims
}

