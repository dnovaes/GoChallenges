package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type UniqChar struct {
	char  byte
	count int
}

func (uchar UniqChar) String() string {
	return fmt.Sprintf("{%c %d}", uchar.char, uchar.count)
}

// Complete the alternate function below.
// 1) remove consecutive chars
// 2) remove min count chars
func alternate(s string) int32 {
	uniqueChars := UniqChars(s)
	fmt.Println("unique chars found: ", uniqueChars)

	minChar, uniqueChars := removeMinCharCount(uniqueChars)
	s = strings.Replace(s, string(minChar), "", -1)
	fmt.Println(fmt.Sprintf("New string after trim: %c, %s", minChar, s))

	minChar, s = removeConsecutiveChar(s)
	fmt.Println(fmt.Sprintf("minChar from consectutive func: '%c'", minChar))
	if minChar != ' ' {
		s = strings.Replace(s, string(minChar), "", -1)
		fmt.Println(fmt.Sprintf("New string after removing a consecutive char: %c, %s", minChar, s))
	}

	minChar, uniqueChars = removeMinCharCount(uniqueChars)
	s = strings.Replace(s, string(minChar), "", -1)
	fmt.Println(fmt.Sprintf("New string after trim: %c, %s", minChar, s))

	return 0
}

func removeMinCharCount(uniqueChars []UniqChar) (byte, []UniqChar) {
	minVal := 1001
	minPos := -1
	var minItem UniqChar
	for i, item := range uniqueChars {
		if item.count < minVal {
			minItem = item
			minVal = item.count
			minPos = i
		}
	}
	uniqueChars = removeUniqueCharByPos(uniqueChars, minPos)
	return minItem.char, uniqueChars
}

func removeUniqueChar(uniqueChars []UniqChar, char byte) []UniqChar {
	var posToRemove []int
	for i, item := range uniqueChars {
		if char == item.char {
			posToRemove = append(posToRemove, i)
		}
	}

	for i, _ := range posToRemove {
		uniqueChars = removeUniqueCharByPos(uniqueChars, i)
	}
	return uniqueChars
}

func removeUniqueCharByPos(uniqueChars []UniqChar, pos int) []UniqChar {
	uniqueChars[pos] = uniqueChars[len(uniqueChars)-1]
	return uniqueChars[:len(uniqueChars)-1]
}

func removeConsecutiveChar(s string) (byte, string) {
	var prevChar rune = ' '

	for i, item := range s {
		if i != 0 {
			fmt.Println(fmt.Sprintf("%c == %c ?", prevChar, item))
			if prevChar == item {
				s = strings.Replace(s, string(prevChar), "", -1)
				return byte(prevChar), s
			}
		}
		prevChar = item
	}
	return ' ', s
}

func UniqChars(s string) []UniqChar {
	var slice []UniqChar
	len := len(s)
	for i := 0; i < len; i++ {
		pos, findResult := Find(slice, s[i])
		if !findResult {
			slice = append(slice, UniqChar{s[i], 1})
		} else {
			slice[pos].count++
		}
	}
	return slice
}

func Find(slice []UniqChar, char byte) (int, bool) {
	for i, item := range slice {
		if item.char == char {
			return i, true
		}
	}
	return -1, false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	lTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	l := int32(lTemp)

	s := readLine(reader)

	result := alternate(s)

	fmt.Fprintf(writer, "%d %d\n", result, l)

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
