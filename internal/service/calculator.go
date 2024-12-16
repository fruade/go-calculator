package service

import (
	"calc_service/internal/constants"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (string, error) {
	expressionList := tokenize(expression)
	if len(expressionList) == 0 {
		return "", fmt.Errorf(constants.InvalidExpression)
	}

	result, err := evaluate(expressionList)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", result), nil
}

func tokenize(expression string) []string {
	for _, operator := range []string{"+", "-", "*", "/", "(", ")"} {
		pattern := " " + operator + " "
		expression = strings.ReplaceAll(expression, operator, pattern)
	}

	return strings.Fields(expression)
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func priority(operator string) int {
	if operator == "/" || operator == "*" {
		return 2
	} else if operator == "+" || operator == "-" {
		return 1
	}
	return 0
}

func evaluate(expressionList []string) (float64, error) {
	var numbers []float64
	var operators []string

	for _, token := range expressionList {
		if isNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf(constants.InvalidNumber, token)
			}
			numbers = append(numbers, num)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				if err := calculateOperator(&numbers, &operators); err != nil {
					return 0, err
				}
			}
			operators = operators[:len(operators)-1]
		} else {
			for len(operators) > 0 && priority(operators[len(operators)-1]) >= priority(token) {
				if err := calculateOperator(&numbers, &operators); err != nil {
					return 0, err
				}
			}
			operators = append(operators, token)
		}
	}

	for len(operators) > 0 {
		if err := calculateOperator(&numbers, &operators); err != nil {
			return 0, err
		}
	}

	if len(numbers) != 1 {
		return 0, fmt.Errorf(constants.InvalidExpressionResult)
	}

	return numbers[0], nil
}

func calculateOperator(numbers *[]float64, operators *[]string) error {
	if len(*numbers) < 2 {
		return fmt.Errorf(constants.InvalidExpression)
	}

	right := (*numbers)[len(*numbers)-1]
	*numbers = (*numbers)[:len(*numbers)-1]
	left := (*numbers)[len(*numbers)-1]
	*numbers = (*numbers)[:len(*numbers)-1]

	op := (*operators)[len(*operators)-1]
	*operators = (*operators)[:len(*operators)-1]

	var result float64
	switch op {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		if right == 0 {
			return fmt.Errorf(constants.ErrDivisionByZero)
		}
		result = left / right
	default:
		return fmt.Errorf(constants.ErrUnknownOperator, op)
	}

	*numbers = append(*numbers, result)
	return nil
}
