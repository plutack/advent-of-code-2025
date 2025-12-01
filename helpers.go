package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readLineFromFile(p string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		file, err := os.Open(p)
		check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out
}
