package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func writeFile(content string, path string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// write some text to file
	_, err = file.WriteString(content)
	checkError(err)

	// save changes
	err = file.Sync()
	checkError(err)
}

func createFile(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}

func readFile(path string) string {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// read file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			checkError(err)
		}
		if n == 0 {
			break
		}
	}
	checkError(err)
	return string(text)
}

func division(valueTotal float64, nextNumber float64, position int) float64 {
	var ret float64
	if position == 0 {
		ret = nextNumber
	} else {
		ret = valueTotal / nextNumber
	}
	return ret
}

func multiplication(valueTotal float64, nextNumber float64, position int) float64 {
	var ret float64
	if position == 0 {
		ret = nextNumber
	} else {
		ret = valueTotal * nextNumber
	}
	return ret
}

func addition(valueTotal float64, nextNumber float64, position int) float64 {
	return valueTotal + nextNumber
}

func subtraction(valueTotal float64, nextNumber float64, position int) float64 {
	var ret float64
	if position == 0 {
		ret = nextNumber
	} else {
		ret = valueTotal - nextNumber
	}
	return ret
}

func doTheOperations(fileContent string, path string) {
	var auxStr = strings.NewReplacer("'", "", "[", "", "]", "").Replace(fileContent)
	var aux2Str = strings.Trim(auxStr, "\x00")
	values := strings.Split(aux2Str, ",")

	var val float64
	var fn func(float64, float64, int) float64

	switch path {
	case "multiplication/multiplication.txt":
		fn = multiplication
		break
	case "subtraction/subtraction.txt":
		fn = subtraction
		break
	case "addition/addition.txt":
		fn = addition
		break
	case "division/division.txt":
		fn = division

		break
	}

	for i, v := range values {
		var number = parsingFloat(v)
		val = fn(val, number, i)
	}
	fmt.Println(strconv.FormatFloat(val, 'f', 6, 64))

}

func parsingFloat(value string) float64 {
	aux, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		fmt.Println(err)
	}
	return aux
}

func main() {
	directories := [4]string{"multiplication", "subtraction", "addition", "division"}
	path := []string{}
	for _, str := range directories {
		var concatPath = str + "/" + str + ".txt"
		os.Mkdir(str, 0777)
		path = append(path, concatPath)
		createFile(concatPath)
		switch str {
		case "multiplication":
			writeFile("[10,20,30,40]", concatPath)
			break
		case "subtraction":
			writeFile("1.1,0.5,0.001", concatPath)
			break
		case "addition":
			writeFile("'11','12','13','14'", concatPath)
			break
		case "division":
			writeFile("80,70,60", concatPath)
			break
		}
	}

	fmt.Println(string("Finished writing files\n"))

	for _, str := range path {
		fmt.Println(string("Path: " + str))
		fileContent := readFile(str)
		fmt.Println(string(fileContent))
		doTheOperations(fileContent, str)
	}

}
