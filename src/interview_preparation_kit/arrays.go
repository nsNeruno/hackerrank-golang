package interview_preparation_kit

import "fmt"

// https://www.hackerrank.com/challenges/crush/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
// Solution is based from the discussion forum on the Challenge page itself
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

// https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func rotLeft(a []int32, d int32) []int32 {
	var l = len(a)
	rotated := make([]int32, l)
	for i := 0; i < l; i++ {
		var s = i + int(d)
		for s >= l {
			s -= l
		}
		rotated[i] = a[s]
	}
	return rotated
}

// https://www.hackerrank.com/challenges/minimum-swaps-2/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func minimumSwaps(arr []int32) int32 {
	swaps := int32(0)
	l := len(arr)

	// Simply iterate through each element from the start, find the correct one somewhere, and swap it!
	for i := 0; i < l-1; i++ {
		k := int32(i + 1)
		if arr[i] != k {
			for j := k; j < int32(l); j++ {
				if arr[j] == k {
					swap := arr[i]
					arr[i] = k
					arr[j] = swap
					swaps++
					break
				}
			}
		}
	}
	return swaps
}

// https://www.hackerrank.com/challenges/new-year-chaos/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
//
// Solution is referenced from "Isming"'s solution, translated from C++ into GoLang
// https://www.hackerrank.com/challenges/new-year-chaos/forum/comments/143969?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
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
