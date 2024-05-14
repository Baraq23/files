package main

import "fmt"

func main() {
	testMap := make(map[int][]string)
	// slice to be appended to the map
	mapSlice := []string{}
	// Char to append
	char := "a"
	// The key are 1,2 & 3
	key := 1
	// Adding key and values
	for i := 0; i < 12; i++ {
		if i > 2 && i < 6 {
			char = "b"
		}
		if i >= 6 && i < 9 {
			char = "c"
		}
		if i >= 9 {
			char = "d"
		}
		// Append the right character
		mapSlice = append(mapSlice, char)
		// Appends the slice to the map wit the correct key
		if i == 2 || i == 5 || i == 8 || i == 11 {
			testMap[key] = mapSlice
			key++
			mapSlice = []string{}
		}
	}
	fmt.Println(testMap)
	mapLooper(testMap)
}

func mapLooper(myMap map[int][]string) {
	// looping through the map
	for i := 0; i < len(myMap[1]); i++ {
		for j := 1; j <= len(myMap); j++ {
			print(myMap[j][i])
		}
		println()
	}
}
