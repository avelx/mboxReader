package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var path string = "/Users/pavel/devcore/DataStore/all_mail.mbox"

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Fix: Set fixed read buffer: to resolve error => " bufio.Scanner: token too long"
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	var emailsNumber int = 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "From:") {
			emailsNumber += 1
			fmt.Printf("%s\n", line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("Total emails found: %d\n", emailsNumber)
}
