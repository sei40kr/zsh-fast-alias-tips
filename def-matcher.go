package main

import (
	"sort"
	"strings"
)

// def-matcher.go
// author: Seong Yong-ju <sei40kr@gmail.com>

type Def struct {
	alias string
	abbr  string
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
	// TODO Implement this
}
