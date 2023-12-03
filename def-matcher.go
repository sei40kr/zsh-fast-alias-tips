package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sei40kr/zsh-fast-alias-tips/matcher"
	"github.com/sei40kr/zsh-fast-alias-tips/model"
	"github.com/sei40kr/zsh-fast-alias-tips/parser"
)

// def-matcher.go
// author: Seong Yong-ju <sei40kr@gmail.com>

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		os.Exit(1)
	}

	defs := make([]model.AliasDef, 0, 512)

	scanner := bufio.NewScanner(bufio.NewReaderSize(os.Stdin, 1024))
	for scanner.Scan() {
		line := scanner.Text()
		defs = append(defs, parser.Parse(line))
	}

	command := os.Args[1]
	fmt.Println(matcher.Match(defs, command))
}
