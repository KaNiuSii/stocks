package stocks

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var Path string = "./stocks/names/"

func readTwoFiles() ([]string, []string) {
	files, err := os.ReadDir(Path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil, nil
	}

	if len(files) < 2 {
		fmt.Println("Not enough files in the directory")
		return nil, nil
	}

	firstIndex := rand.Intn(len(files))
	secondIndex := firstIndex
	for secondIndex == firstIndex {
		secondIndex = rand.Intn(len(files))
	}

	firstFileName := files[firstIndex].Name()
	firstFileContent, err := os.ReadFile(Path + firstFileName)
	if err != nil {
		fmt.Println("Error reading file:", firstFileName, err)
		return nil, nil
	}

	secondFileName := files[secondIndex].Name()
	secondFileContent, err := os.ReadFile(Path + secondFileName)
	if err != nil {
		fmt.Println("Error reading file:", secondFileName, err)
		return nil, nil
	}

	return strings.Split(string(firstFileContent), "\n"), strings.Split(string(secondFileContent), "\n")
}

func generateNewName() string {
	first, second := readTwoFiles()

	if len(first) == 0 || len(second) == 0 {
		return "" // Handle the case where either file is empty or error in reading files
	}

	firstWord := first[rand.Intn(len(first))]
	secondWord := second[rand.Intn(len(second))]

	return strings.ReplaceAll(firstWord, "\x00", "") + " " + strings.ReplaceAll(secondWord, "\x00", "")
}
