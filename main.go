package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := flag.String("file", "", "path to log file")
	keyword := flag.String("keyword", "", "optional keyword filter")
	level := flag.String("level", "", "optional severity level (INFO, WARN, ERROR)")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("usage: log-line-filter -file app.log [-level ERROR] [-keyword timeout]")
		return
	}

	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("could not open file:", err)
		return
	}
	defer f.Close()

	counts := map[string]int{
		"INFO":  0,
		"WARN":  0,
		"ERROR": 0,
		"OTHER": 0,
	}

	total := 0
