package main

import (
	"reflect"
	"testing"
)

func TestCalculate_Plus(t *testing.T) {
	expression := "41+3"
	expected := float64(44)

	result,err := calculate(expression)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculate_Minus(t *testing.T) {
	expression := "41-37"
	expected :=  float64(4)

	result,err := calculate(expression)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculate_Multiplication(t *testing.T) {
	expression := "41*37"
	expected :=  float64(1517)

	result,err := calculate(expression)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculate_Divide(t *testing.T) {
	expression := "256/16"
	expected :=  float64(16)

	result, err := calculate(expression)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculate_Brackets(t *testing.T) {
	expression := "35+4*(17-7)-41*30"
	expected :=  float64(-1155)

	result, err := calculate(expression)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestCalculate_MinusBrackets(t *testing.T) {
	expression := "-(-11-(1*20/2)-11/2*3)"
	expected := 37.5

	result,err := calculate(expression)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}