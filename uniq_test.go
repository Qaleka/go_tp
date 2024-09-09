package main

import (
  "reflect"
  "testing"
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

// func TestProcessLines_IgnoreCase(t *testing.T) {
//   lines := []string{
//     "I LOVE MUSIC.",
//     "I love music.",
//     "I LoVe MuSiC.",
//   }
//   opts := &arguments{
//     ignoreCase: true,
//   }

//   expected := []linesInfo{
//     {"I LOVE MUSIC.", "i love music.", 3},
//   }

//   result := processLines(lines, opts)

//   if !reflect.DeepEqual(result, expected) {
//     t.Errorf("Expected %v, got %v", expected, result)
//   }
// }

// func TestProcessLines_CountFlag(t *testing.T) {
//   lines := []string{
//     "I love music.",
//     "I love music.",
//     "I love music of Kartik.",
//     "Thanks.",
//     "I love music of Kartik.",
//   }
//   opts := &arguments{
//     count: true,
//   }

//   expected := []linesInfo{
//     {"I love music.", "I love music.", 2},
//     {"I love music of Kartik.", "I love music of Kartik.", 2},
//     {"Thanks.", "Thanks.", 1},
//   }

//   result := processLines(lines, opts)

//   if !reflect.DeepEqual(result, expected) {
//     t.Errorf("Expected %v, got %v", expected, result)
//   }
// }

// func TestProcessLines_DuplicatesFlag(t *testing.T) {
//   lines := []string{
//     "I love music.",
//     "I love music.",
//     "I love music of Kartik.",
//     "Thanks.",
//     "I love music of Kartik.",
//   }
//   opts := &arguments{
//     duplicates: true,
//   }

//   expected := []linesInfo{
//     {"I love music.", "I love music.", 2},
//     {"I love music of Kartik.", "I love music of Kartik.", 2},
//   }

//   result := processLines(lines, opts)

//   if !reflect.DeepEqual(result, expected) {
//     t.Errorf("Expected %v, got %v", expected, result)
//   }
// }

// func TestProcessLines_UniqueFlag(t *testing.T) {
//   lines := []string{
//     "I love music.",
//     "I love music.",
//     "I love music of Kartik.",
//     "Thanks.",
//     "I love music of Kartik.",
//   }
//   opts := &arguments{
//     unique: true,
//   }

//   expected := []linesInfo{
//     {"Thanks.", "Thanks.", 1},
//   }

//   result := processLines(lines, opts)

//   if !reflect.DeepEqual(result, expected) {
//     t.Errorf("Expected %v, got %v", expected, result)
//   }
// }
