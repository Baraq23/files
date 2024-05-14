package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	content, err := os.ReadFile("file.txt")
	checkErr("Error reading file", err)
	strContent := string(content)

	// Replace string
	strContent = strings.ReplaceAll(strContent, "package", "package main")
	code := strings.ReplaceAll(strContent, "import", "import \"fmt\"")

	//creating a temporary go file to store the code
	temp := "goFile.go"
	er := os.WriteFile(temp, []byte(code), 0644)
	checkErr("Error writing to goFile.go", er)
	// Delete the go file after it has been run
	defer os.Remove(temp)

	// running the code
	cmd := exec.Command("go", "run", "goFile.go")

	//Writting output in ouput.txt file
	outputFile, err := cmd.CombinedOutput()
	checkErr("Error obtaining output", err)
	errr := os.WriteFile("output.txt", []byte(outputFile), 0644)
	checkErr("Error writing output information to file", errr)
}

func checkErr(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		return
	}
}
