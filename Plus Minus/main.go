package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	positive := int32(0)
	negative := int32(0)
	zeroes := int32(0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroes++
		} else if arr[i] < 0 {
			negative++
		} else if arr[i] > 0 {
			positive++
		}
	}
	positiveRatio := fmt.Sprintf("%.6f", float32(positive)/float32(len(arr)))
	negativeRatio := fmt.Sprintf("%.6f", float32(negative)/float32(len(arr)))
	zeroRatio := fmt.Sprintf("%.6f", float32(zeroes)/float32(len(arr)))

	fmt.Println(positiveRatio + "\n" + negativeRatio + "\n" + zeroRatio)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
