package parser

import "github.com/sei40kr/zsh-fast-alias-tips/model"

func Parse(line string) model.AliasDef {
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

	return model.AliasDef{Name: string(alias), Abbr: string(abbr)}
}
