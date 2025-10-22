package handlers

import(
	"jwt-server/auth"
	"net/http"
	"log"
	"encoding/json"
)

type Handler struct {
	JWTSecret []byte
}

func NewHandler (secret []byte) *Handler {
	return &Handler{
		JWTSecret: secret,
	}
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method", http.StatusMethodNotAllowed)
		return
	}

	var creds struct { // Simple struct to hold login credentials
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request Payload", http.StatusBadRequest)
		return
	}

	if creds.Username != "admin" || creds.Password != "adminpass" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}


	tokenString, err := auth.GenerateToken(42, "admin", h.JWTSecret)

	if err != nil {
		log.Printf("Error generating token: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(struct { 
        Token string `json:"token"`
    }{Token: tokenString})
}