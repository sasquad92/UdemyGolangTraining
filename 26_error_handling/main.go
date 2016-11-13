package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt") // nf - pointer to  file
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no_file.txt") // only catching error while opening non existing file
	if err != nil {
		// fmt.Println("error happend", err)
		log.Println("error happend", err) // the same but with time stamp
		// log.Fatalln(err) // writing error with exit status 1
		// panic(err) // msg with some stack stuff
	}
}
