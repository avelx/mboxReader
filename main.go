package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Simple MBOX reader: prints "From " line and Subject for each message
func main() {
	var path string = "/Users/pavel/devcore/DataStore/all_mail.mbox"
	//if len(os.Args) != 2 {
	//  fmt.Println("Usage: go run mbox_reader.go <mboxfile>")
	//  return
	//}

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

	//inMessage := false
	//var fromLine string
	//var subject string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 100 && strings.HasPrefix(line, "From ") {
			fmt.Printf("%s Line: %s\n", line, len(line))
		}
		//	// Print previous message info, if any
		//	if inMessage {
		//		fmt.Printf("%s | Subject: %s\n", fromLine, subject)
		//	}
		//	fromLine = line
		//	subject = ""
		//	inMessage = true
		//} else if strings.HasPrefix(line, "Subject: ") && subject == "" {
		//	subject = strings.TrimPrefix(line, "Subject: ")
		//}
	}
	// Print last message
	//if inMessage {
	//fmt.Printf("%s | Subject: %s\n", fromLine, subject)

	//}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
