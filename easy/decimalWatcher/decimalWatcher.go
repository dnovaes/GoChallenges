package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * DecimalWatcher

 input (integer):

 0
 03
 003
 0030
 030
 3

 output (always in decimal):

 0.00
 0.03
 0.30
 0.30
 3.00

*/

const COMMA rune = ','
const DOT rune = ','

func decimalWatcher(content string, numDigits int) string {

	count := len(content)
	if count == 0 {
		return ""
	} else if count <= numDigits {
		content = fillZeroes(content, (numDigits-count)+1)
	}

	content = applyMask(content, numDigits)

	var str strings.Builder
	str.WriteString(content)

	fmt.Println("Result of decimalWatcher: ", str.String())
	return str.String()
}

func fillZeroes(content string, numZeroes int) string {
	var str strings.Builder
	str.WriteString(content)
	for i := 0; i < numZeroes; i++ {
		str.WriteString("0")
	}
	return str.String()
}

func applyMask(content string, numDigits int) string {
	var count = len(content)
	var startPosLessSignificative = count - numDigits

	content = content[:startPosLessSignificative] + "." + content[startPosLessSignificative:]
	countSignificativeLetters := float64(len(content[:startPosLessSignificative]))

	for i := int(countSignificativeLetters); i-3 > 0; i = i - 3 {
		var posNearestComma = indexOf(COMMA, content)
		//fmt.Println(fmt.Sprintf("pos next comma: %d, %s", posNearestComma, content))
		if posNearestComma != -1 {
			content = content[:i-3] + "," + content[i-3:]
		} else {
			content = content[:i-3] + "," + content[i-3:i] + content[startPosLessSignificative:]
		}
	}
	return content
}

func indexOf(element rune, data string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	valueTyped := " "
	var result string

	valueTyped = readLine(reader)
	for valueTyped != "q" {
		result = decimalWatcher(valueTyped, 2)
		fmt.Fprintf(writer, "%s\n", result)
		valueTyped = readLine(reader)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
