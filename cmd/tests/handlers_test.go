package tests

import (
	"bytes"
	"calc_service/internal/constants"
	"calc_service/internal/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
    tests := []struct {
        name           string
        input          string
        expectedCode   int
        expectedOutput string
    }{
        {
            name:           "Valid expression",
            input:          `{"expression": "2+2*2"}`,
            expectedCode:   http.StatusOK,
            expectedOutput: `{"result":"6.00"}`,
        },
        {
            name:           "Invalid expression",
            input:          `{"expression": "2+2w*2"}`,
            expectedCode:   http.StatusUnprocessableEntity,
            expectedOutput: `{"error":"` + constants.ErrInvalidExpression + `"}`,
        },
        {
            name:           "Invalid expression",
            input:          `{"expression": "2++2"}`,
            expectedCode:   http.StatusUnprocessableEntity,
            expectedOutput: `{"error":"` + constants.ErrInvalidExpression + `"}`,
        },
        {
            name:           "Empty json",
            input:          `{}`,
            expectedCode:   http.StatusUnprocessableEntity,
            expectedOutput: `{"error":"` + constants.ErrInvalidExpression + `"}`,
        },
        {
            name:           "Division zero",
            input:          `{"expression":"2/0"}`,
            expectedCode:   http.StatusUnprocessableEntity,
            expectedOutput: `{"error":"` + constants.ErrDivisionByZero + `"}`,
        },
        {
            name:           "Invalid JSON payload",
            input:          `{"expression":"2+2*2"`,
            expectedCode:   http.StatusBadRequest,
            expectedOutput: `{"error":"` + constants.ErrInvalidJSON + `"}`,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(tt.input)))
            req.Header.Set("Content-Type", "application/json")

            rec := httptest.NewRecorder()
			handlers.CalculateHandler(rec, req)

            if rec.Code != tt.expectedCode {
                t.Errorf("expected status code %d, got %d", tt.expectedCode, rec.Code)
            }

            if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(tt.expectedOutput) {
                t.Errorf("expected response body %s, got %s", tt.expectedOutput, rec.Body.String())
            }
        })
    }
}
