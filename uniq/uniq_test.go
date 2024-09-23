package main

import (
  "reflect"
  "testing"
  "bufio"
)

func TestProcessLines_NoFlags(t *testing.T) {
  lines := []string{
    "I love music.",
	  "I love music.",
	  "I love music.",
	  "",
    "I love music of Kartik.",
    "I love music of Kartik.",
    "Thanks.",
    "I love music of Kartik.",
    "I love music of Kartik.",
  }
  args := &arguments{}

  expected := []linesInfo{
    {"I love music.", 3},
    {"", 1},
    {"I love music of Kartik.", 2},
    {"Thanks.", 1},
    {"I love music of Kartik.", 2},
  }

  result := filterLines(lines, *args)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestProcessLines_IgnoreCase(t *testing.T) {
  lines := []string{
    "I LOVE MUSIC.",
    "I love music.",
    "I LoVe MuSiC.",
    "",
    "I love MuSIC of Kartik.",
    "I love music of kartik.",
    "Thanks.",
    "I love music of kartik.",
    "I love MuSIC of Kartik.",
  }
  args := &arguments{
    allSmall: true,
  }

  expected := []linesInfo{
    {"I LOVE MUSIC.", 3},
    {"", 1},
    {"I love MuSIC of Kartik.", 2},
    {"Thanks.", 1},
    {"I love music of kartik.", 2},
  }

  result := filterLines(lines, *args)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestProcessLines_CountFlag(t *testing.T) {
  lines := []string{
    "I love music.",
    "I love music.",
    "I love music.",
    "",
    "I love music of Kartik.",
    "I love music of Kartik.",
    "Thanks.",
    "I love music of Kartik.",
    "I love music of Kartik.",
  }
  args := &arguments{
    countLine: true,
  }

  expected := []linesInfo{
    {"I love music.", 3},
    {"", 1},
    {"I love music of Kartik.", 2},
    {"Thanks.", 1},
    {"I love music of Kartik.", 2},
  }

  result := filterLines(lines, *args)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestProcessLines_DuplicatesFlag(t *testing.T) {
  lines := []string{
    "I love music.",
    "I love music.",
    "I love music.",
    "",
    "I love music of Kartik.",
    "I love music of Kartik.",
    "Thanks.",
    "I love music of Kartik.",
    "I love music of Kartik.",
  }
  args := &arguments{
    duplicates: true,
  }

  expected := []linesInfo{
    {"I love music.", 3},
    {"I love music of Kartik.", 2},
    {"I love music of Kartik.", 2},
  }

  result := filterLines(lines, *args)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestProcessLines_UniqueFlag(t *testing.T) {
  lines := []string{
    "I love music.",
    "I love music.",
    "I love music.",
    "",
    "I love music of Kartik.",
    "I love music of Kartik.",
    "Thanks.",
    "I love music of Kartik.",
    "I love music of Kartik.",
  }
  args := &arguments{
    unique: true,
  }

  expected := []linesInfo{
    {"", 1},
    {"Thanks.", 1},
  }

  result := filterLines(lines, *args)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestProcessLines_InputFile(t *testing.T) {
  lines := []string{}
  args := &arguments{
    inputFile: "test.txt",
  }
  input := readInput(args.inputFile)
  lines = readFromInput(input)
  expected := []linesInfo{
    {"I love music.", 3},
    {"", 1},
    {"I love music of Kartik.", 2},
    {"Thanks.", 1},
    {"I love music of Kartik.", 2},
  }

  result := filterLines(lines, *args)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestProcessLines_OutputFile(t *testing.T) {
  lines := []string{
    "I love music.",
    "I love music.",
    "I love music.",
    "",
    "I love music of Kartik.",
    "I love music of Kartik.",
    "Thanks.",
    "I love music of Kartik.",
    "I love music of Kartik.",
  }
  args := &arguments{
    outputFile: "testOutput.txt",
  }
  expected := []linesInfo{
    {"I love music.", 3},
    {"", 1},
    {"I love music of Kartik.", 2},
    {"Thanks.", 1},
    {"I love music of Kartik.", 2},
  }

  
  result := filterLines(lines, *args)
  writeIntoOutput(result,*args)

  output := readInput(args.outputFile)
  data := bufio.NewScanner(output)
	var outputLines []string
	for data.Scan() {
		lines = append(outputLines, data.Text())
	}

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, outputLines)
  }
}

func TestProcessLines_InputAndOutputFile(t *testing.T) {
  lines := []string{}
  args := &arguments{
    inputFile: "test.txt",
    outputFile: "testOutput.txt",
  }
  expected := []linesInfo{
    {"I love music.", 3},
    {"", 1},
    {"I love music of Kartik.", 2},
    {"Thanks.", 1},
    {"I love music of Kartik.", 2},
  }

  input := readInput(args.inputFile)
  lines = readFromInput(input)

  result := filterLines(lines, *args)
  writeIntoOutput(result,*args)

  output := readInput(args.outputFile)
  data := bufio.NewScanner(output)
	var outputLines []string
	for data.Scan() {
		lines = append(outputLines, data.Text())
	}

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, outputLines)
  }
}

func TestProcessLines_numFields(t *testing.T) {
  lines := []string{
    "We love music.",
    "I love music.",
    "They love music.",
    "",
    "I love music of Kartik.",
    "We love music of Kartik.",
    "Thanks.",
  }
  args := &arguments{
    numFields: 1,
  }

  expected := []linesInfo{
    {"We love music.", 3},
    {"", 1},
    {"I love music of Kartik.", 2},
    {"Thanks.", 1},
  }

  result := filterLines(lines, *args)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}

func TestProcessLines_numChars(t *testing.T) {
  lines := []string{
    "I love music.",
    "A love music.",
    "C love music.",
    "",
    "I love music of Kartik.",
    "We love music of Kartik.",
    "Thanks.",
  }
  args := &arguments{
    numChars: 1,
  }

  expected := []linesInfo{
    {"I love music.", 3},
    {"", 1},
    {"I love music of Kartik.", 1},
    {"We love music of Kartik.", 1},
    {"Thanks.", 1},
  }

  result := filterLines(lines, *args)

  if !reflect.DeepEqual(result, expected) {
    t.Errorf("Expected %v, got %v", expected, result)
  }
}