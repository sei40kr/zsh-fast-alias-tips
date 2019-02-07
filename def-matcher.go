package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// def-matcher.go
// author: Seong Yong-ju <sei40kr@gmail.com>

type Def struct {
	alias string
	abbr  string
}

func ParseDef(line string) Def {
	alias := make([]rune, 0, 1024)
	abbr := make([]rune, 0, 1024)

	afterEscape := false
	inQuote := false
	inRightExp := false
	for _, aRune := range line {
		if aRune == '\\' {
			afterEscape = !afterEscape

			if afterEscape {
				continue
			}
		}

		if aRune == '\'' && !afterEscape {
			inQuote = !inQuote
		} else if aRune == '=' && !inQuote {
			inRightExp = true
		} else if !inRightExp {
			alias = append(alias, aRune)
		} else {
			abbr = append(abbr, aRune)
		}

		afterEscape = false
	}

	return Def{alias: string(alias), abbr: string(abbr)}
}

func MatchDef(defs []Def, command string) string {
	sort.Slice(defs, func(i, j int) bool {
		return strings.Compare(defs[i].abbr, defs[j].abbr) <= 0
	})

	ok := 0
	ng := len(defs)
	for 1 < ng-ok {
		mid := (ok + ng) / 2

		if 0 <= strings.Compare(command, defs[mid].abbr) {
			ok = mid
		} else {
			ng = mid
		}
	}

	if strings.HasPrefix(command, defs[ok].abbr) {
		return defs[ok].alias
	} else {
		return ""
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		os.Exit(1)
	}

	defs := make([]Def, 0, 512)

	scanner := bufio.NewScanner(bufio.NewReaderSize(os.Stdin, 1024))
	for scanner.Scan() {
		line := scanner.Text()
		defs = append(defs, ParseDef(line))
	}

	fmt.Printf("%s\n", MatchDef(defs, os.Args[1]))
}
