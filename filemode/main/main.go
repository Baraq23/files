package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	file, err := os.OpenFile("go.txt", os.O_RDWR, 0o644)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	line := ""
	newLine := ""

	for scanner.Scan() {
		line = scanner.Text()

		if strings.Contains(line, "package") {
			newLine = strings.Replace(line, "package", "package main", -1)
			lines = append(lines, newLine)
		} else if strings.Contains(line, "import") {
			newLine = strings.Replace(line, "import", "import \"fmt\"", -1)
			lines = append(lines, newLine)
		} else {
			lines = append(lines, line)
		}

		
	}
	// Clear file content
	if err := file.Truncate(0); err != nil {
		checkErr(err)
	}

	// Move cursor to the star of file
	if _, err := file.Seek(0, 0); err != nil {
		checkErr(err)
	}

	// Write content line by line 
	for _, line := range lines {
		_, err := fmt.Fprintln(file, line)
		checkErr(err)
	}

	

	fmt.Println("File modified successfully")
	err2 := os.WriteFile("gofile.go", []byte(file), 0644)
	checkErr(err2)


	RunTXT()
}


func RunTXT() {
	cmd := exec.Command("go", "run", "/home/barraotieno/Desktop/filemode/go.go")

	output, err := cmd.Output()
	checkErr(err)
	


	fmt.Println("The ouput is:", string(output))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
