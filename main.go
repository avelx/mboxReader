package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var path string = "/Users/pavel/devcore/DataStore/all_mail.mbox"
	var processedEmails string = "/Users/pavel/devcore/DataStore/mails/"

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
	var b strings.Builder
	var emailSeparator string = "From "
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, emailSeparator) {
			emailsNumber += 1
			var fileName string = processedEmails + "/email-" + strconv.Itoa(emailsNumber) + ".txt"
			//fmt.Printf("%s\n\n", b.String())

			err := os.WriteFile(fileName, []byte(b.String()), 0644)
			if err != nil {
				panic(err)
			}
			// Start new email string accumulation
			b.Reset()
			b.Write([]byte(line + "\n"))
		} else {
			// buffer all lines scanned
			b.Write([]byte(line + "\n"))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Printf("Total emails found: %d\n", emailsNumber)
}
