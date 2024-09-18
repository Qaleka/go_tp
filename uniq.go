package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"flag"
	"strings"
)

type arguments struct{ //список аргументов ком.строки
	countLine bool
	duplicates bool
	unique bool
	numFields int
	numChars int
	allSmall bool
	inputFile string
	outputFile string
}

type linesInfo struct {//информация для 1 строчки
	line string
	count int
}

// func uniq(input io.Reader, output io.Writer) error {
// 	in := bufio.NewScanner(input)
// 	text, err := io.ReadAll(input)
// 	if err != nil {
		
// 	}

// 	var prev string
// 	for in.Scan() {
// 		txt := in.Text()
// 		if txt == prev{
// 			continue
// 		}
// 		prev = txt
// 		fmt.Fprintln(output, txt)
// 	}
// 	return nil
// }

func parseArguments() arguments {//парсинг аргументов при помощи flag
	var args arguments
	//привязка к переменным
	flag.BoolVar(&args.countLine, "c", false, "Подсчет количества встречаний строки")
	flag.BoolVar(&args.duplicates, "d", false, "Вывод только повторяющихся строк")
	flag.BoolVar(&args.unique, "u", false, "Вывод только уникальных строк")
	flag.IntVar(&args.numFields, "f", 0, "Не учитывать введенное количество полей в начале строки")
	flag.IntVar(&args.numChars, "s", 0, "Не учитывать введенное количество символов в начале строки")
	flag.BoolVar(&args.allSmall, "i", false, "Не учитывать регистр букв")

	flag.Parse()
	//проверка на совместимость
	if (args.countLine && args.duplicates) || (args.countLine && args.unique) || (args.duplicates && args.unique) {
		fmt.Println("Нельзя использовать -c -d -u вместе")
		os.Exit(1)
	}
	//проверка на входной файд
	if len(flag.Args()) > 0{
		args.inputFile = flag.Arg(0)
	}
	//проверка на выходной файл
	if len(flag.Args()) > 1 {
		args.outputFile = flag.Arg(1)
	}

	return args
}

func readInput(input string) io.ReadCloser {
	if input == "" {
		return os.Stdin
	}
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Ошибка при открытии файла")
		os.Exit(1)
	}
	return file
}

func readFromInput(input io.Reader) []string {
	data := bufio.NewScanner(input)
	var lines []string
	for data.Scan() {
		lines = append(lines, data.Text())
	}
	return lines
}

//чтение из ввода, фильтрация согласно параметрам
func filterLines(allLines []string, args arguments) []linesInfo {
	// var readFrom io.Reader
	// if args.inputFile == "" {//если нет файла
	// 	readFrom = os.Stdin
	// } else {
	// 	var err error
	// 	readFrom, err = os.Open(args.inputFile)
	// 	if err != nil {
	// 		fmt.Println("Ошибка при чтении файла")
	// 		os.Exit(1)
	// 	}
	// }
	// inData := bufio.NewScanner(readFrom)
	var lines []linesInfo//список строк
	var prev string//предыдущая строка
	var count int//количество идущих подряд строк
	var flagOneLine bool = false//флаг для случая 1 строки
	var original string//оригинал дубликатов до преобразований строки
	var flagOriginal bool = false//флаг для того, что использовать оригинал или нет
	var originalPrev string //предыдущая строка без преобразований
	for _, txt:= range allLines {
		buff := txt
		if args.numFields > 0 {//параметр f
			fields := strings.Fields(buff)
			if len(fields) > args.numFields {
				buff = strings.Join(fields[args.numFields:], " ")
			} else
			{
				buff = ""
			}
		}
		if args.numChars > 0 && len(buff) >= args.numChars { //параметр s
			buff = buff[args.numChars:]
		}
		if args.allSmall { //параметр i
			buff = strings.ToLower(buff)
		}
		if buff == prev {
			if !flagOriginal {
				original = originalPrev
				flagOriginal = true
			}
			count++
			continue
		} 
		if flagOriginal {
			lines = append(lines, linesInfo{original, count})
			flagOriginal = false
		} else if flagOneLine {
			lines = append(lines, linesInfo{originalPrev, count})
		}
		count = 1
		prev = buff
		originalPrev = txt
		flagOneLine = true
	}
	if flagOriginal {
		lines = append(lines, linesInfo{original, count})
		flagOriginal = false
	} else if flagOneLine {
		lines = append(lines, linesInfo{originalPrev, count})
	}
	
	return filterDU(lines, args)
}

func filterDU (lines []linesInfo, args arguments) []linesInfo {
	var filterWithDU []linesInfo
	for _,line:= range lines {
		if args.countLine {//флаг c
			filterWithDU = append(filterWithDU, line)
		} else if args.duplicates && line.count > 1 { //параметр d
			filterWithDU = append(filterWithDU, line)
		} else if args.unique && line.count == 1 { //параметр u
			filterWithDU = append(filterWithDU, line)
		} else if !args.countLine && !args.duplicates && !args.unique {
			filterWithDU = append(filterWithDU, line)
		}
	}
	return filterWithDU
}

func writeIntoOutput(lines []linesInfo, args arguments) {
	var writer io.Writer = os.Stdout
	if args.outputFile != "" {
		outputFile, err := os.Create(args.outputFile)
		if err != nil {
			fmt.Println("Ошибка при создании файла")
			os.Exit(1)
		}
		defer outputFile.Close()
		writer = outputFile
	}

	for _,line:= range lines {
		if args.countLine {//флаг c
			fmt.Fprintf(writer, "%d %s\n", line.count, line.line)
		} else if args.duplicates && line.count > 1 { //параметр d
			fmt.Fprintln(writer, line.line)
		} else if args.unique && line.count == 1 { //параметр u
			fmt.Fprintln(writer, line.line)
		} else if !args.countLine && !args.duplicates && !args.unique {
			fmt.Fprintln(writer, line.line)
		}
	}
}
// func filterLine(line string, args arguments) string {
// 	if args.numFields > 0 {

// 	}

// 	if args.numChars > 0 && len(line) > args.numChars{
// 		line = line[args.numChars:]
// 	}
// }

func main() {
	args := parseArguments()
	// fmt.Println(args)
	input := readInput(args.inputFile)

	lines := readFromInput(input)

	filteredLines := filterLines(lines, args)
	writeIntoOutput(filteredLines,args)
	// err := uniq(os.Stdin, os.Stdout)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}