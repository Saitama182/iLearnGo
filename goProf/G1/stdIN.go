package main

import (
	"bufio"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	/* 	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
		if scanner.Text() == "stop" {
			break

		}

	} */
}
