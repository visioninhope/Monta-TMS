package utils

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"trenova/app/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/wader/gormstore/v2"
)

func ParseBody(r *http.Request, body interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(body); err != nil {
		return err
	}
	return nil
}

func ParseBodyAndValidate(w http.ResponseWriter, r *http.Request, body interface{}) error {
	if err := ParseBody(r, body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	if err := Validate(body); err != nil {
		var validationErr models.ValidationErrorResponse
		if jsonErr := json.Unmarshal([]byte(err.Error()), &validationErr); jsonErr == nil {
			errorBytes, _ := json.Marshal(validationErr)
			http.Error(w, string(errorBytes), http.StatusBadRequest)
		} else {
			// Fallback in case the error is not a validation error
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return err
	}

	return nil
}

// GetMuxVar retrieves a variable from the route pattern match and writes an error if it's not found.
func GetMuxVar(w http.ResponseWriter, r *http.Request, key string) (value string) {
	vars := mux.Vars(r)
	value, ok := vars[key]
	if !ok {
		ResponseWithError(w, http.StatusBadRequest, models.ValidationErrorDetail{
			Code:   "invalid",
			Detail: "The required parameter is missing.",
			Attr:   key,
		})
		value = "" // Return an empty string if the value is not found
	}
	return value
}

// RegisterGob registers the UUID type with gob, so it can be used in sessions.
func RegisterGob() {
	gob.Register(uuid.UUID{})
	gob.Register(models.User{})
}

func GetSystemSessionID() string {
	key := os.Getenv("SESSION_ID")
	if key == "" {
		log.Fatal("SESSION_ID not found in environment")
	}

	return key
}

// GetSessionDetails retrieves user ID, organization ID, and business unit ID from the session.
func GetSessionDetails(r *http.Request, store *gormstore.Store) (userID uuid.UUID, orgID uuid.UUID, buID uuid.UUID, ok bool) {
	if store == nil {
		log.Println("Session store is not initialized")
		return uuid.Nil, uuid.Nil, uuid.Nil, false
	}

	sessionID := GetSystemSessionID()
	session, err := store.Get(r, sessionID)
	if err != nil {
		log.Printf("Error retrieving session: %v", err)
		return uuid.Nil, uuid.Nil, uuid.Nil, false
	}

	userID, userOk := session.Values["userID"].(uuid.UUID)
	orgID, orgOk := session.Values["organizationID"].(uuid.UUID)
	buID, buOk := session.Values["businessUnitID"].(uuid.UUID)

	return userID, orgID, buID, userOk && orgOk && buOk
}