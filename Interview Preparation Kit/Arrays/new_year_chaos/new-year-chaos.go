// https://www.hackerrank.com/challenges/new-year-chaos/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Solution is referenced from "Isming"'s solution, translated from C++ into GoLang
// https://www.hackerrank.com/challenges/new-year-chaos/forum/comments/143969?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
//
// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	var bribes = 0

	// Iterate from the last element
	for i := int32(len(q) - 1); i >= 0; i-- {

		// Since the original array was a sorted number array of 1 to n with interval of 1,
		// For example:
		// In Test Case of {2, 5, 1, 3, 4}
		// The element 5 (with original position of 5, index of 4) can only be placed up to position 3, index of 2
		// On this case, element 5 is placed on index 1 or position 2, so we need to find whether the position is valid or not
		// KEY: The difference must not be greater than 2
		// So, we substract the current position (which is array index + 1) from the current value, which is 5
		// and the difference = 5 - (1 + 1) = 5 - 2 = 3, indicated that there was an invalid move
		if q[i]-int32(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}

		// Here, we need to find out how many briberies were happened, knowing that the bribery limit is 2.
		// So, we need to ensure that the two queue number in front of the current number is always lower than the current number.
		// For example:
		//
		// Under the Test Case of {2, 1, 5, 3, 4}
		// Let's say we are checking the first element from the back, which is 4.
		// We expect that the two people in front of number 4 has lower number
		var max int32
		// i: 4 -> Is 0 >= 4 - 2? false, then max = 2
		// i: 3 -> Is 0 >= 3 - 2? false, then max = 1
		// i: 2 -> Is 0 >= 5 - 2? false, then max = 3
		// i: 1 -> Is 0 >= 1 - 2? true, then max = 0
		// i: 0 -> Is 0 >= 0 - 2? true, then max = 0

		// To prevent looping from negative index
		if 0 >= q[i]-2 {
			max = 0
		} else {
			max = q[i] - 2
		}

		// Loop through the ranged specification
		// The first element from the back is 4, and the two in front of it is 2.
		// We start from the 2nd index element.
		// The element at 2nd index is 5, and 5 is greater than the element at 4th index which is 4.
		// This indicated that number 5 had bribed number 4.
		// On the next step, the 3rd index contains number 3, and 3 is lower than 4. That means no bribery happened.
		//
		// Then we check the second element from the back, which is 3.
		// The second element from the back is placed at 3rd index position, so we loop through 1st index position.
		// The element at 1st index is 1, and 1 is not greater than 3. No bribery happened.
		// Then we check the element at 2nd index, which is 5. 5 is greater than 3, that means 5 had bribed 3.
		for j := max; j < i; j++ {
			// max: 2; i: 4
			// 5 > 4? true, bribes++
			// 3 > 4? false
			//
			// max: 1; i: 3
			// 1 > 3? false
			// 5 > 3? true, bribes++
			//
			// max: 3; i: 2
			// no checking
			//
			// max: 0; i: 1
			// 2 > 1? true, bribes++
			//
			// max: 0; i: 0
			// no checking

			if q[j] > q[i] {
				bribes++
			}
		}
	}
	// Expected result: 3
	fmt.Println(bribes)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
