package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {

	start := time.Now()

	filename := "test.dat"
	newfile := "testGo.txt"
	e := os.Remove(newfile)
	if e != nil {
		log.Println(e)
	}

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("The file already exists and has been deleted successfully")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024*1024)
	scanner.Buffer(buf, 1024*1024*1024)

	lineCount := 0
	for scanner.Scan() {
		lineCount += 1
		if lineCount > 1 {
			newLine := scanner.Text()[438:]
			charCount := len(newLine)

			file, err := os.OpenFile(newfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("failed creating file: %s", err)
			}
			datawriter := bufio.NewWriter(file)
			for i := 0; i < charCount/90; i++ {
				data := newLine[i*90 : (i+1)*90]
				_, _ = datawriter.WriteString(data + "\n")
			}
			datawriter.Flush()
			file.Close()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)

}
