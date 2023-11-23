package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
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
	reader := csv.NewReader(readFile)
	reader.Comma = ';'
	reader.LazyQuotes = true

	csvFile, err := os.Create("AP13-2.csv")
	writer := csv.NewWriter(csvFile)

	//var lines []string

	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return fmt.Errorf("failed to read file %s", err)
		}

		for key, line := range record {
			if key == 12 {
				split := strings.Split(line, "|")

				for _, splitLine := range split {
					var lines []string
					newLine := record[0] + splitLine
					lines = append(lines, newLine)

					if err := writer.Write(lines); err != nil {
						log.Fatalln("error writing record to csv:", err)
					}
				}

			}
		}
	}

	writer.Flush()

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
