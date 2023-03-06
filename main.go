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
