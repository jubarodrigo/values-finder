package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	err := ReadCurrentDir()
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(file string) error {
	readFile, err := os.Open("./files/" + file)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for _, line := range lines {
		if res := strings.Contains(line, "d7acf78b"); res == true {
			fmt.Println("The file is:", file)
			fmt.Println("result", res)
		}
	}

	return nil
}

func ReadCurrentDir() error {
	files, err := os.Open("./files")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
		return err
	}
	defer files.Close()

	list, _ := files.Readdirnames(0) // 0 to read all files and folders
	for _, file := range list {
		err = readFile(file)
		if err != nil {
			return err
		}
	}

	return nil
}
