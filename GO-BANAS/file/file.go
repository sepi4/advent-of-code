package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	f, err := os.Create("kissa.txt")
	if err != nil {
		log.Fatal(err)
	}

	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, v := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(v))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open("kissa.txt")
	if err != nil {
		log.Fatal(err)
	}
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("prime:", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// =============
	_, err = os.Stat("kissa.txt")
	if errors.Is(err, os.ErrNotExist) {
		// panic("File doesn't exist")
		pl("File doesn't exist")
	} else {

		f, err := os.OpenFile(
			"kissa.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = f.WriteString("kana\n")
		if err != nil {
			log.Fatal(err)
		}

	}
}
