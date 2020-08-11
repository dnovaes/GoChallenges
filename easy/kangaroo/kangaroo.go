package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// Cases
/*
no start valid positions

start valid position
0 3 2 2 YES
0 3 3 2 YES
0 5 6 2 YES
0 2 5 3 NO
*/

// Complete the kangaroo function below.
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	const YES string = "YES"
	const NO string = "NO"

	type Kango struct {
		pos  int32
		jump int32
	}
	var arr []Kango

	// append ordered kangaroo initial positions
	kango1 := Kango{x1, v1}
	kango2 := Kango{x2, v2}
	if x1 < x2 {
		arr = append(arr, kango1)
		arr = append(arr, kango2)
	} else {
		arr = append(arr, kango2)
		arr = append(arr, kango1)
	}
	fmt.Println("arr: ", arr)

	// check for valid start position + valid jumps
	if kango1.jump < kango2.jump {
		return NO
	}

	// check if they meet at same position until kango2 passes kango1
	var numJump int32
	for numJump = 1; kango1.pos < kango2.pos; numJump++ {
		kango1.pos = kango1.pos + kango1.jump
		kango2.pos = kango2.pos + kango2.jump
		fmt.Println("kango1 pos: ", kango1.pos, "kango2 pos: ", kango2.pos, "jumps: ", numJump)
		if kango1.pos == kango2.pos {
			fmt.Println(YES + " by checking jumps")
			return YES
		}
		time.Sleep(time.Second / 3)
	}

	fmt.Println(NO + " end of file")
	return NO
}

func kangaroo2(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// numJumps
	// question:
	// x1 + v1*numJumps == x2 + v2*numJumps
	// v1*numJumps - v2*numJumps = x2 - x1
	// numJumps * (v1 - v2) = x2 - x1
	// numJumps = (x2 - x1) / (v1 - v2)  must be > 0 and integer  (== rest zero)
	// numJumps = (x2 - x1) % (v1 - v2) == 0

	const YES string = "YES"
	const NO string = "NO"
	type Kango struct {
		pos  int32
		jump int32
	}
	var arr []Kango

	kango1 := Kango{x1, v1}
	kango2 := Kango{x2, v2}
	if x1 < x2 {
		arr = append(arr, kango1)
		arr = append(arr, kango2)
	} else {
		arr = append(arr, kango2)
		arr = append(arr, kango1)
	}

	if arr[0].jump < arr[1].jump {
		return NO
	}

	if arr[0].jump-arr[1].jump == 0 {
		return NO
	}

	numJumps := (arr[1].pos-arr[0].pos)%(arr[0].jump-arr[1].jump) == 0
	if numJumps {
		fmt.Println("returned YES on mod")
		return YES
	}
	return NO
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	x1V1X2V2 := strings.Split(readLine(reader), " ")

	x1Temp, err := strconv.ParseInt(x1V1X2V2[0], 10, 64)
	checkError(err)
	x1 := int32(x1Temp)

	v1Temp, err := strconv.ParseInt(x1V1X2V2[1], 10, 64)
	checkError(err)
	v1 := int32(v1Temp)

	x2Temp, err := strconv.ParseInt(x1V1X2V2[2], 10, 64)
	checkError(err)
	x2 := int32(x2Temp)

	v2Temp, err := strconv.ParseInt(x1V1X2V2[3], 10, 64)
	checkError(err)
	v2 := int32(v2Temp)

	//result := kangaroo(x1, v1, x2, v2)
	result := kangaroo2(x1, v1, x2, v2)

	fmt.Fprintf(writer, "%s\n", result)

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
