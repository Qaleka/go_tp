package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"errors"
)

func twoNumOp(num1 float64, num2 float64, op byte) (float64, error) { // Вычисление
	switch op {
	case '+':
		return num1 + num2, nil
	case '-':
		return num1 - num2, nil
	case '*':
		return num1 * num2, nil
	case '/':
		if num2 == 0{
			return 0, errors.New("division by zero")
		}
		return num1 / num2, nil
	}
	return 0, nil
}

func compareOp(op1 byte, op2 byte) bool { // Сравнение приоритета
	switch {
	case op1 == op2:
		return true
	case (op1 == '*' || op1 == '/') && (op2 == '+' || op2 == '-'|| op2 == '*'|| op2 == '/'):
		return true
	case (op1 == '+' || op1 == '-') && (op2 == '+' || op2 == '-'):
		return true
	case (op1 == '+' || op1 == '-') && (op2 == '*' || op2 == '/'):
		return false
	}
	return false
}

func operationWithNumbers(numbers []float64, operations []byte) ([]float64, []byte, error) {
	newNumbers := append([]float64{}, numbers...)
	newOperations := append([]byte{}, operations...)

	num1 := newNumbers[len(newNumbers)-1]
	newNumbers = newNumbers[:len(newNumbers)-1]
	num2 := newNumbers[len(newNumbers)-1]
	newNumbers = newNumbers[:len(newNumbers)-1]

	oneOperation := newOperations[len(newOperations)-1]
	newOperations = newOperations[:len(newOperations)-1]

	opResult, err := twoNumOp(num2, num1, oneOperation)
	if err != nil {
		return nil, nil, err
	}

	newNumbers = append(newNumbers, opResult)

	return newNumbers, newOperations, nil
}

func calculate(expression string) (float64, error) {
	var (
		operations []byte
		numbers []float64
		one_number string
		minus_braket int
		err error
	)
	for i := 0; i < len(expression); i++ {
		switch {
		case expression[i] == '(':
			if len(numbers) == 0 && len(operations) >= 1 && operations[len(operations) - 1] == '-' {
				minus_braket++
				operations = operations[:len(operations) - 1]
			}
			operations = append(operations, expression[i])
		case expression[i] == ')':
			for len(operations) > 0 && operations[len(operations) - 1] != '(' {
				numbers, operations, err = operationWithNumbers(numbers, operations)
				if err != nil {
					return 0, err
				}
			}
			if minus_braket !=0 && len(numbers) == 1 {
				numbers[len(numbers) - 1] = -numbers[len(numbers) - 1]
				minus_braket--
			}
			operations = operations[:len(operations) - 1]
		case expression[i] >= '0' && expression[i] <= '9':
			one_number = ""
			if len(operations) > 0 && operations[len(operations) - 1] == '-' {
				if len(numbers) == 0 {
					operations = operations[:len(operations) - 1]
					one_number = one_number + "-"
				} else if len(operations) > 1 && expression[i-1] == '(' {
					operations = operations[:len(operations) - 1]
					one_number = one_number + "-"
				}
			}
			for len(expression) > i && expression[i] >= '0' && expression[i] <= '9' {
				one_number = one_number + string(expression[i])
				i++
			}
			i--
			num, err := strconv.ParseFloat(one_number,64)
			if err != nil {
				return 0, err
			}
			numbers = append(numbers, num)
		default:
			for len(operations) > 0 { // Производим операции
				if compareOp(operations[len(operations) - 1], expression[i]) {
					numbers, operations, err = operationWithNumbers(numbers, operations)
					if err != nil {
						return 0, err
					}
				} else {
					break
				}
			}
			operations = append(operations, expression[i])
		}
	}
	for len(operations) > 0 { // Добивае финальные операции
		numbers, operations, err = operationWithNumbers(numbers, operations)
		if err != nil {
			return 0, err
		}
	}
	return numbers[0], nil
}

func calc() error{
	data := bufio.NewScanner(os.Stdin)
	data.Scan()
	expression := data.Text()
	result, err := calculate(expression); 
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}