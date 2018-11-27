// https://www.hackerrank.com/challenges/crush/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Solution is based from the discussion forum on the Challenge page itself

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	max, arr, sum := int64(0), make([]int64, n), int64(0)

	// Instead of observing increases on each individual value...
	// We observe each beginning of range increase, and where it ends.
	//
	// For example:
	// n = 5, m = 3
	// a b  k
	// 1 2 100 -> this will add 100 to n[1 - 1] and -100 to n[2]
	// 2 5 100 -> this will add 100 to n[2 - 1] and should add -100 to n[5] or none at all since it is beyond index range
	// 3 4 100
	// Expected output: 200
	//
	// Begin with: [5]int64 {0, 0, 0, 0, 0}
	// Then we iterate through the queries
	// m[0]: {100, 0, -100, 0, 0}
	// m[1]: {100, 100, -100, 0, 0}
	// m[2]: {100, 100, 0, 0, -100}
	//
	// Then we'll get sum of the whole array
	// while observing the peak sum as max value
	// (0)+100 100(+100) 200(+0) 200(+0) 100(-100)

	for i := 0; i < len(queries); i++ {
		query := queries[i]
		a := query[0] - 1
		b := query[1] - 1
		k := int64(query[2])
		arr[a] += k
		if b+1 < n {
			arr[b+1] -= k
		}
	}
	for i := int32(0); i < n; i++ {
		if arr[i] != 0 {
			sum += arr[i]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

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
