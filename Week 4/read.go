package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	names := make([]Name, 0, 3)

	fmt.Println("Enter text file's name :")
	filename, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	filename = strings.Trim(filename, "\n")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var name Name
		name.fname, name.lname = s[0], s[1]
		names = append(names, name)

	}

	for _, v := range names {
		fmt.Println(v)
	}

}
