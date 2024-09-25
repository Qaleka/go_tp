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

func readFromInput(input io.ReadCloser) ([]string, error) {
	defer input.Close()
	data := bufio.NewScanner(input)
	var lines []string
	for data.Scan() {
		err := data.Err()
		if err != nil {
			return nil, err
		}
		line := data.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

//чтение из ввода, фильтрация согласно параметрам
func filterLines(allLines []string, args arguments) []linesInfo {
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

func filterDU(lines []linesInfo, args arguments) []linesInfo {
	var filterWithDU []linesInfo
	for _,line:= range lines {
		switch {
		case args.countLine:
			filterWithDU = append(filterWithDU, line)
		case args.duplicates && line.count > 1:
			filterWithDU = append(filterWithDU, line)
		case args.unique && line.count == 1:
			filterWithDU = append(filterWithDU, line)
		case !args.countLine && !args.duplicates && !args.unique:
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
		switch {
		case args.countLine:
			fmt.Fprintf(writer, "%d %s\n", line.count, line.line)
		case args.duplicates && line.count > 1:
			fmt.Fprintln(writer, line.line)
		case args.unique && line.count == 1:
			fmt.Fprintln(writer, line.line)
		case !args.countLine && !args.duplicates && !args.unique:
			fmt.Fprintln(writer, line.line)
		}
	}
}

func uniq() {
	args := parseArguments()
	input := readInput(args.inputFile)

	lines, err := readFromInput(input)
	if err != nil {
		fmt.Printf("Ошибка при чтении: %s", err)
	}

	filteredLines := filterLines(lines, args)
	writeIntoOutput(filteredLines,args)
}