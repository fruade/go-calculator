package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
    "calc_service/internal/constants"
)

type Request struct {
	Expression string `json:"expression"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func validateExpression(expression string) bool {
	validExpression := regexp.MustCompile(`^[0-9+\-*/().\s]+$`)
	return validExpression.MatchString(expression)
}

func ValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ErrorResponse{Error: constants.ErrMethodNotAllowed})
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: constants.ErrReadRequestBody})
			return
		}

		log.Printf("Middleware received body: %s", string(body))

		var req Request
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: constants.ErrInvalidRequest})
			return
		}

		if req.Expression == "" || !validateExpression(req.Expression) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ErrorResponse{Error: constants.ErrInvalidExpression})
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next(w, r)
	}
}
