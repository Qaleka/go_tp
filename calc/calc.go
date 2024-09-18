package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func twoNumOp(num1 int, num2 int, op byte) int {
	switch op {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	case '*':
		return num1 * num2
	case '/':
		return num1 / num2
	}
	return 0
}

func compareOp(op1 byte, op2 byte) bool {
	fmt.Println(string(op1), string(op2))
	if op1 == op2 {
		return true
	} else if (op1 == '*' || op1 == '/') && (op2 == '+' || op2 == '-'|| op2 == '*'|| op2 == '/') {
		return true
	} else if (op1 == '+' || op1 == '-') && (op2 == '+' || op2 == '-') {
		return true
	} else if (op1 == '+' || op1 == '-') && (op2 == '*' || op2 == '/') {
		return false
	}
	return false
}

func calculate (expression string) int {
	var operations []byte
	var numbers []int
	var one_number string
	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			operations = append(operations, expression[i])
		} else if expression[i] == ')' {
			for len(operations) > 0 && operations[len(operations) - 1] != '(' {
				num1 := numbers[len(numbers) - 1]
				numbers = numbers[:len(numbers) - 1]
				num2 := numbers[len(numbers) - 1]
				numbers = numbers[:len(numbers) - 1]
				oneOperaion := operations[len(operations) - 1]
				operations = operations[:len(operations) - 1]
				numbers = append(numbers, twoNumOp(num2, num1, oneOperaion))
			}
			operations = operations[:len(operations) - 1]
		} else if expression[i] >= '0' && expression[i] <= '9' {
			one_number = ""
			for len(expression) > i && expression[i] >= '0' && expression[i] <= '9' {
				one_number = one_number + string(expression[i])
				i++
			}
			i--
			num, _ := strconv.Atoi(one_number)
			// fmt.Println(num)
			numbers = append(numbers, num)
		} else {
			for len(operations) > 0 {
				if compareOp(operations[len(operations) - 1], expression[i]) {
					num1 := numbers[len(numbers) - 1]
					numbers = numbers[:len(numbers) - 1]
					num2 := numbers[len(numbers) - 1]
					numbers = numbers[:len(numbers) - 1]
					oneOperation := operations[len(operations) - 1]
					operations = operations[:len(operations) - 1]
					// fmt.Println(num2,num1,twoNumOp(num2, num1, oneOperation))
					numbers = append(numbers, twoNumOp(num2, num1, oneOperation))
				} else {
					break
				}
			}
			operations = append(operations, expression[i])
		}
	}

	for len(operations) > 0 {
		// fmt.Println(numbers)
		num1 := numbers[len(numbers) - 1]
		numbers = numbers[:len(numbers) - 1]
		num2 := numbers[len(numbers) - 1]
		numbers = numbers[:len(numbers) - 1]
		oneOperaion := operations[len(operations) - 1]
		operations = operations[:len(operations) - 1]
		numbers = append(numbers, twoNumOp(num2, num1, oneOperaion))
	}
	return numbers[0]
}

func main() {
	data := bufio.NewScanner(os.Stdin)
	data.Scan()
	exspression := data.Text()
	fmt.Println(calculate(exspression))
}