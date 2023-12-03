package matcher

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sei40kr/zsh-fast-alias-tips/model"
)

func Match(defs []model.AliasDef, command string) string {
	sort.Slice(defs, func(i, j int) bool {
		return len(defs[j].Abbr) <= len(defs[i].Abbr)
	})

	for true {
		var match model.AliasDef
		for _, def := range defs {

			if command == def.Abbr {
				match = def
				break
			} else if strings.HasPrefix(command, def.Abbr+" ") {
				match = def
				break
			}
		}

		if match != (model.AliasDef{}) {
			command = fmt.Sprintf("%s%s", match.Name, command[len(match.Abbr):])
		} else {
			break
		}
	}
	return command
}
