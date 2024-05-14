package main

import (
	"bufio"
	"os"
    "fmt"
)


func main() {
	asciiMap := readBanner("standard.txt")
	input := os.Args[1]
	storeArt2(asciiMap, input)
}

// This function stores the ascii art to then prints the ascii art
func storeArt2(asciiMap map[rune][]string, input string) {
	//slice of a slice to store the Art
	store := [][]string{}
	for _, chr := range input {
		store = append(store, asciiMap[chr])
	}
	//Display what is in the store
	for i := 0; i < 8; i++ {
	 for j := 0; j < len(store); j++ {
			fmt.Print(store[j][i])
		}
		println()
	}
} 
//this function maps the ascii charracters to their corresponding art
func readBanner(file string) map[rune][]string {
	bannerFile, err := os.Open(file)

	if err != nil {
		panic(err)
	}
	defer bannerFile.Close()
	scanner := bufio.NewScanner(bannerFile)
	asciiMap := make(map[rune][]string)
	
	for i := 32; i <= 126; i++{
		runeValue := i
		scanner.Scan() 
		var art []string
		for i := 0; i < 8; i++ {
			scanner.Scan()
			line := scanner.Text()
			art = append(art, line)
		}
		asciiMap[rune(runeValue)] = art
		
	}
	return asciiMap
}
