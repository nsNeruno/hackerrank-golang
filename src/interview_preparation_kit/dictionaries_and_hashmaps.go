package interview_preparation_kit

import (
	"fmt"
	"strings"
)

// https://www.hackerrank.com/challenges/ctci-ransom-note/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
func checkMagazine(magazine []string, note []string) {
	// Turn both arrays into maps
	// and get count of each words
	mMap, nMap := make(map[string]int), make(map[string]int)
	for i := 0; i < len(magazine); i++ {
		mMap[magazine[i]]++
	}
	for i := 0; i < len(note); i++ {
		nMap[note[i]]++
	}
	// If the word count from the magazine is less than the requirement of the note then it's a NO NO
	for word, count := range nMap {
		if mMap[word] < count {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

// https://www.hackerrank.com/challenges/two-strings/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
func twoStrings(s1 string, s2 string) string {
	// Iterate through the shorter string (Optional)
	var shorter, other string
	if len(s1) <= len(s2) {
		shorter, other = s1, s2
	} else {
		shorter, other = s2, s1
	}
	// This is considered as cheating though, using method Contains from "strings" package.
	for _, c := range shorter {
		if strings.Contains(other, string(c)) {
			return "YES"
		}
	}
	return "NO"
}
