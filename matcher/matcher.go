package matcher

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sei40kr/zsh-fast-alias-tips/model"
)

func Match(defs []model.AliasDef, command string) (*model.AliasDef, bool) {
	sort.Slice(defs, func(i, j int) bool {
		return len(defs[j].Abbr) <= len(defs[i].Abbr)
	})

	var candidate model.AliasDef
	isFullMatch := false

	for true {
		var match model.AliasDef
		for _, def := range defs {

			if command == def.Abbr {
				match = def
				isFullMatch = true
				break
			} else if strings.HasPrefix(command, def.Abbr) {
				match = def
				break
			}
		}

		if match != (model.AliasDef{}) {
			command = fmt.Sprintf("%s%s", match.Name, command[len(match.Abbr):])
			candidate = match
		} else {
			break
		}
	}

	if candidate != (model.AliasDef{}) {
		return &candidate, isFullMatch
	} else {
		return nil, false
	}
}
