package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(fileToParse string) string {
	var separator string = "Content-Type:"
	var sepatorFound bool = false

	file, err := os.Open(fileToParse) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	var fileContent strings.Builder

	data := make([]byte, 100)

	count, err := file.Read(data)
	//var myString = string(data[:])
	for count > 0 {
		actualBytes := data[1:count]
		line := string(actualBytes[:])
		if sepatorFound {
			fileContent.Write(actualBytes)
		}
		if strings.Contains(line, separator) && sepatorFound == false {
			sepatorFound = true
		}

		if err != nil {
			log.Fatal(err)
		}
		count, err = file.Read(data)

		//n, err :=
	}

	// Attempt to extract pure text from MIME message
	doc, err2 := html.Parse(strings.NewReader(fileContent.String()))
	if err2 != nil {
		log.Fatal(err)
	}
	for n := range doc.Descendants() {
		if n.Type == html.DocumentNode {
			fmt.Println(n.Data)
		}
	}

	//for n := range doc.Descendants() {
	//	if n.Type == html.ElementNode && n.DataAtom == atom.A {
	//		for _, a := range n.Attr {
	//			//if a.Key == "href" {
	//			fmt.Println(a.Val)
	//			break
	//			//}
	//		}
	//	}
	//}

	//fmt.Printf("FileContent %s\n", fileContent.String())

	// Brutal approach:: just parse what we have
	//doc, err := html.Parse
	return fileContent.String()
}

func main() {
	var path string = "/Users/pavel/devcore/DataStore/all_mail.mbox"
	var processedEmails string = "/Users/pavel/devcore/DataStore/mails/"

	var fileToParse = processedEmails + "email-1000.txt"
	readFile(fileToParse)
	//fmt.Printf("Here is result: %s", result)
	return

	// Wait for a bit

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
