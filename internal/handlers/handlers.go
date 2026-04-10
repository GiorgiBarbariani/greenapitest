package handlers

import (
	"encoding/json"
	"net/http"

	"greenapitest/internal/api"
	"greenapitest/internal/models"
)

type Handler struct {
	client *api.Client
}

func NewHandler(client *api.Client) *Handler {
	return &Handler{client: client}
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, models.APIResponse{
		Success: false,
		Error:   message,
	})
}

func (h *Handler) GetSettings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if creds.IDInstance == "" || creds.APITokenInstance == "" {
		respondError(w, http.StatusBadRequest, "idInstance and apiTokenInstance are required")
		return
	}

	result, err := h.client.GetSettings(creds.IDInstance, creds.APITokenInstance)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    result,
	})
}

func (h *Handler) GetStateInstance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if creds.IDInstance == "" || creds.APITokenInstance == "" {
		respondError(w, http.StatusBadRequest, "idInstance and apiTokenInstance are required")
		return
	}

	result, err := h.client.GetStateInstance(creds.IDInstance, creds.APITokenInstance)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    result,
	})
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.IDInstance == "" || req.APITokenInstance == "" {
		respondError(w, http.StatusBadRequest, "idInstance and apiTokenInstance are required")
		return
	}

	if req.PhoneNumber == "" || req.Message == "" {
		respondError(w, http.StatusBadRequest, "phoneNumber and message are required")
		return
	}

	chatID := req.PhoneNumber + "@c.us"
	result, err := h.client.SendMessage(req.IDInstance, req.APITokenInstance, chatID, req.Message)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    result,
	})
}

func (h *Handler) SendFileByURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.SendFileByURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.IDInstance == "" || req.APITokenInstance == "" {
		respondError(w, http.StatusBadRequest, "idInstance and apiTokenInstance are required")
		return
	}

	if req.PhoneNumber == "" || req.FileURL == "" || req.FileName == "" {
		respondError(w, http.StatusBadRequest, "phoneNumber, fileUrl, and fileName are required")
		return
	}

	chatID := req.PhoneNumber + "@c.us"
	result, err := h.client.SendFileByURL(req.IDInstance, req.APITokenInstance, chatID, req.FileURL, req.FileName, req.Caption)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    result,
	})
}
