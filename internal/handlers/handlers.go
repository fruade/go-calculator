package handlers

import (
	"calc_service/internal/constants"
	"calc_service/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func handleError(w http.ResponseWriter, err error) {
	errorMapping := map[string]struct {
		status int
		msg    string
	}{
		constants.InvalidExpression:    {http.StatusUnprocessableEntity, constants.ErrInvalidExpression},
		constants.ErrUnknownOperator:   {http.StatusBadRequest, constants.ErrUnknownOperator},
		constants.ErrDivisionByZero:    {http.StatusUnprocessableEntity, constants.ErrDivisionByZero},
		constants.ErrNotEnoughOperands: {http.StatusUnprocessableEntity, constants.ErrNotEnoughOperands},
	}

	if mappedError, exists := errorMapping[err.Error()]; exists {
		w.WriteHeader(mappedError.status)
		json.NewEncoder(w).Encode(Response{Error: mappedError.msg})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response{Error: constants.ErrInternalError})
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: constants.ErrInvalidJSON})
		return
	}

	expression := strings.TrimSpace(req.Expression)
	result, err := service.Calc(expression)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: result})
}
